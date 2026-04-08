package plugin

import (
	"testing"
)

func TestHealthCheckYAMLParsing(t *testing.T) {
	yamlData := `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  description: "Fuzzy finder"
spec:
  repo: nvim-telescope/telescope.nvim
  health_checks:
    - type: lua_module
      value: telescope
      description: "Telescope module is loadable"
    - type: command
      value: Telescope
      description: "Telescope command exists"
`
	p, err := ParseYAML([]byte(yamlData))
	if err != nil {
		t.Fatalf("ParseYAML failed: %v", err)
	}

	if len(p.HealthChecks) != 2 {
		t.Fatalf("HealthChecks count = %d, want 2", len(p.HealthChecks))
	}

	// Check first health check
	hc0 := p.HealthChecks[0]
	if hc0.Type != HealthCheckLuaModule {
		t.Errorf("HealthChecks[0].Type = %q, want %q", hc0.Type, HealthCheckLuaModule)
	}
	if hc0.Value != "telescope" {
		t.Errorf("HealthChecks[0].Value = %q, want %q", hc0.Value, "telescope")
	}
	if hc0.Description != "Telescope module is loadable" {
		t.Errorf("HealthChecks[0].Description = %q, want %q",
			hc0.Description, "Telescope module is loadable")
	}

	// Check second health check
	hc1 := p.HealthChecks[1]
	if hc1.Type != HealthCheckCommand {
		t.Errorf("HealthChecks[1].Type = %q, want %q", hc1.Type, HealthCheckCommand)
	}
	if hc1.Value != "Telescope" {
		t.Errorf("HealthChecks[1].Value = %q, want %q", hc1.Value, "Telescope")
	}
}

func TestHealthCheckRoundTrip(t *testing.T) {
	original := &Plugin{
		Name:    "test-plugin",
		Repo:    "test/plugin",
		Enabled: true,
		HealthChecks: []HealthCheck{
			{
				Type:        HealthCheckLuaModule,
				Value:       "test",
				Description: "Test module loadable",
			},
			{
				Type:        HealthCheckCommand,
				Value:       "TestCmd",
				Description: "Test command exists",
			},
		},
	}

	// Convert to YAML and back
	py := original.ToYAML()
	converted := py.ToPlugin()

	if len(converted.HealthChecks) != 2 {
		t.Fatalf("HealthChecks count = %d, want 2", len(converted.HealthChecks))
	}

	for i, hc := range converted.HealthChecks {
		orig := original.HealthChecks[i]
		if hc.Type != orig.Type {
			t.Errorf("HealthChecks[%d].Type = %q, want %q", i, hc.Type, orig.Type)
		}
		if hc.Value != orig.Value {
			t.Errorf("HealthChecks[%d].Value = %q, want %q", i, hc.Value, orig.Value)
		}
		if hc.Description != orig.Description {
			t.Errorf("HealthChecks[%d].Description = %q, want %q",
				i, hc.Description, orig.Description)
		}
	}
}

func TestValidateHealthCheckType(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"lua_module valid", "lua_module", false},
		{"command valid", "command", false},
		{"treesitter valid", "treesitter", false},
		{"lsp valid", "lsp", false},
		{"invalid type", "foobar", true},
		{"empty type", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateHealthCheckType(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateHealthCheckType(%q) error = %v, wantErr %v",
					tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestInferModuleName(t *testing.T) {
	tests := []struct {
		name     string
		plugin   *Plugin
		expected string
	}{
		{
			name: "from config require",
			plugin: &Plugin{
				Name:   "telescope",
				Repo:   "nvim-telescope/telescope.nvim",
				Config: `local telescope = require("telescope")`,
			},
			expected: "telescope",
		},
		{
			name: "from config require sub-module",
			plugin: &Plugin{
				Name:   "telescope",
				Repo:   "nvim-telescope/telescope.nvim",
				Config: `local actions = require("telescope.actions")`,
			},
			expected: "telescope",
		},
		{
			name: "from repo name",
			plugin: &Plugin{
				Name: "gitsigns",
				Repo: "lewis6991/gitsigns.nvim",
			},
			expected: "gitsigns",
		},
		{
			name: "strip nvim- prefix from repo",
			plugin: &Plugin{
				Name: "tree",
				Repo: "nvim-tree/nvim-tree.lua",
			},
			expected: "tree",
		},
		{
			name: "from single-quote require",
			plugin: &Plugin{
				Name:   "alpha",
				Repo:   "goolord/alpha-nvim",
				Config: `local alpha = require('alpha')`,
			},
			expected: "alpha",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := inferModuleName(tt.plugin)
			if result != tt.expected {
				t.Errorf("inferModuleName() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestDefaultHealthChecks(t *testing.T) {
	tests := []struct {
		name         string
		plugin       *Plugin
		expectCount  int
		expectModule string
	}{
		{
			name: "explicit checks returned as-is",
			plugin: &Plugin{
				Name: "test",
				Repo: "test/test",
				HealthChecks: []HealthCheck{
					{Type: HealthCheckCommand, Value: "TestCmd"},
				},
			},
			expectCount: 1,
		},
		{
			name: "inferred from config",
			plugin: &Plugin{
				Name:   "telescope",
				Repo:   "nvim-telescope/telescope.nvim",
				Config: `local telescope = require("telescope")`,
			},
			expectCount:  1,
			expectModule: "telescope",
		},
		{
			name: "inferred from repo",
			plugin: &Plugin{
				Name: "gitsigns",
				Repo: "lewis6991/gitsigns.nvim",
			},
			expectCount:  1,
			expectModule: "gitsigns",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checks := DefaultHealthChecks(tt.plugin)
			if len(checks) != tt.expectCount {
				t.Fatalf("DefaultHealthChecks() count = %d, want %d",
					len(checks), tt.expectCount)
			}
			if tt.expectModule != "" && len(checks) > 0 {
				if checks[0].Value != tt.expectModule {
					t.Errorf("check value = %q, want %q",
						checks[0].Value, tt.expectModule)
				}
			}
		})
	}
}
