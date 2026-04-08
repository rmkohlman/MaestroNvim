package plugin

import (
	"strings"
	"testing"
)

func TestGenerateHealthCheckLua(t *testing.T) {
	plugins := []*Plugin{
		{
			Name:    "telescope",
			Repo:    "nvim-telescope/telescope.nvim",
			Enabled: true,
			HealthChecks: []HealthCheck{
				{Type: HealthCheckLuaModule, Value: "telescope"},
				{Type: HealthCheckCommand, Value: "Telescope"},
			},
		},
		{
			Name:    "disabled-plugin",
			Repo:    "test/disabled",
			Enabled: false,
			HealthChecks: []HealthCheck{
				{Type: HealthCheckLuaModule, Value: "disabled"},
			},
		},
		{
			Name:    "treesitter",
			Repo:    "nvim-treesitter/nvim-treesitter",
			Enabled: true,
			HealthChecks: []HealthCheck{
				{Type: HealthCheckTreesitter, Value: "lua"},
			},
		},
	}

	lua := GenerateHealthCheckLua(plugins)

	// Disabled plugin should NOT be included
	if strings.Contains(lua, "disabled-plugin") {
		t.Error("Generated Lua should not contain disabled plugin checks")
	}

	// Enabled plugins should have checks
	checks := []string{
		`pcall(require, "telescope")`,
		`plugin = "telescope"`,
		`check_type = "lua_module"`,
		`check_type = "command"`,
		`check_value = "Telescope"`,
		`check_type = "treesitter"`,
		`check_value = "lua"`,
		`vim.json.encode(results)`,
		`vim.cmd('qa!')`,
	}

	for _, check := range checks {
		if !strings.Contains(lua, check) {
			t.Errorf("Generated Lua missing: %q\n\nGenerated:\n%s", check, lua)
		}
	}
}

func TestGenerateHealthCheckLuaEmpty(t *testing.T) {
	// Plugins with no checks (but can infer) should still generate
	plugins := []*Plugin{
		{
			Name:    "gitsigns",
			Repo:    "lewis6991/gitsigns.nvim",
			Enabled: true,
			// No explicit health checks — should infer lua_module
		},
	}

	lua := GenerateHealthCheckLua(plugins)

	if !strings.Contains(lua, `pcall(require, "gitsigns")`) {
		t.Errorf("Expected inferred lua_module check for gitsigns\n\nGenerated:\n%s", lua)
	}
}

func TestGenerateHealthCheckLuaLSP(t *testing.T) {
	plugins := []*Plugin{
		{
			Name:    "lspconfig",
			Repo:    "neovim/nvim-lspconfig",
			Enabled: true,
			HealthChecks: []HealthCheck{
				{Type: HealthCheckLSP, Value: "lua_ls"},
			},
		},
	}

	lua := GenerateHealthCheckLua(plugins)

	checks := []string{
		`require, "lspconfig"`,
		`check_type = "lsp"`,
		`check_value = "lua_ls"`,
	}

	for _, check := range checks {
		if !strings.Contains(lua, check) {
			t.Errorf("Generated Lua missing: %q\n\nGenerated:\n%s", check, lua)
		}
	}
}

func TestHealthCheckerStaticCheck(t *testing.T) {
	checker := NewHealthChecker()

	plugins := []*Plugin{
		{
			Name:    "telescope",
			Repo:    "nvim-telescope/telescope.nvim",
			Enabled: true,
			HealthChecks: []HealthCheck{
				{Type: HealthCheckLuaModule, Value: "telescope"},
			},
		},
		{
			Name:    "disabled",
			Repo:    "test/disabled",
			Enabled: false,
		},
		{
			Name: "no-checks",
			Repo: "some/empty-repo",
			// No repo name that can be inferred easily
			Enabled: true,
		},
	}

	reports := checker.StaticCheck(plugins)

	if len(reports) != 3 {
		t.Fatalf("StaticCheck returned %d reports, want 3", len(reports))
	}

	// Reports are sorted by name
	// "disabled" comes first
	if reports[0].PluginName != "disabled" {
		t.Errorf("reports[0].PluginName = %q, want %q", reports[0].PluginName, "disabled")
	}
	if reports[0].Status != HealthStatusSkipped {
		t.Errorf("disabled plugin status = %q, want %q", reports[0].Status, HealthStatusSkipped)
	}

	// "no-checks" — will actually get inferred check from repo name
	if reports[1].PluginName != "no-checks" {
		t.Errorf("reports[1].PluginName = %q, want %q", reports[1].PluginName, "no-checks")
	}

	// "telescope" has valid checks
	if reports[2].PluginName != "telescope" {
		t.Errorf("reports[2].PluginName = %q, want %q", reports[2].PluginName, "telescope")
	}
	if reports[2].Status != HealthStatusUnknown {
		t.Errorf("telescope status = %q, want %q", reports[2].Status, HealthStatusUnknown)
	}
}

func TestHealthCheckerParseNvimResults(t *testing.T) {
	checker := NewHealthChecker()

	jsonOutput := `[
		{"plugin":"telescope","check_type":"lua_module","check_value":"telescope","status":"healthy","message":"module loadable"},
		{"plugin":"telescope","check_type":"command","check_value":"Telescope","status":"healthy","message":"command available"},
		{"plugin":"gitsigns","check_type":"lua_module","check_value":"gitsigns","status":"unhealthy","message":"module not found"}
	]`

	reports, err := checker.ParseNvimResults([]byte(jsonOutput))
	if err != nil {
		t.Fatalf("ParseNvimResults failed: %v", err)
	}

	if len(reports) != 2 {
		t.Fatalf("ParseNvimResults returned %d reports, want 2", len(reports))
	}

	// Sorted by name: gitsigns first
	if reports[0].PluginName != "gitsigns" {
		t.Errorf("reports[0].PluginName = %q, want %q", reports[0].PluginName, "gitsigns")
	}
	if reports[0].Status != HealthStatusUnhealthy {
		t.Errorf("gitsigns status = %q, want %q", reports[0].Status, HealthStatusUnhealthy)
	}

	if reports[1].PluginName != "telescope" {
		t.Errorf("reports[1].PluginName = %q, want %q", reports[1].PluginName, "telescope")
	}
	if reports[1].Status != HealthStatusHealthy {
		t.Errorf("telescope status = %q, want %q", reports[1].Status, HealthStatusHealthy)
	}
	if len(reports[1].Results) != 2 {
		t.Errorf("telescope results count = %d, want 2", len(reports[1].Results))
	}
}

func TestHealthCheckerParseInvalidJSON(t *testing.T) {
	checker := NewHealthChecker()

	_, err := checker.ParseNvimResults([]byte("not json"))
	if err == nil {
		t.Error("ParseNvimResults should fail on invalid JSON")
	}
}
