package plugin

import (
	"strings"
	"testing"
)

func TestParseYAML(t *testing.T) {
	yaml := `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  description: "Fuzzy finder"
  category: fuzzy-finder
  tags: ["finder", "search"]
spec:
  repo: nvim-telescope/telescope.nvim
  branch: 0.1.x
  dependencies:
    - nvim-lua/plenary.nvim
    - repo: nvim-telescope/telescope-fzf-native.nvim
      build: make
  config: |
    local telescope = require("telescope")
    telescope.setup({})
  keymaps:
    - key: "<leader>ff"
      mode: n
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files"
`
	p, err := ParseYAML([]byte(yaml))
	if err != nil {
		t.Fatalf("ParseYAML failed: %v", err)
	}

	// Check basic fields
	if p.Name != "telescope" {
		t.Errorf("Name = %q, want %q", p.Name, "telescope")
	}
	if p.Description != "Fuzzy finder" {
		t.Errorf("Description = %q, want %q", p.Description, "Fuzzy finder")
	}
	if p.Repo != "nvim-telescope/telescope.nvim" {
		t.Errorf("Repo = %q, want %q", p.Repo, "nvim-telescope/telescope.nvim")
	}
	if p.Branch != "0.1.x" {
		t.Errorf("Branch = %q, want %q", p.Branch, "0.1.x")
	}

	// Check tags
	if len(p.Tags) != 2 || p.Tags[0] != "finder" || p.Tags[1] != "search" {
		t.Errorf("Tags = %v, want [finder, search]", p.Tags)
	}

	// Check dependencies
	if len(p.Dependencies) != 2 {
		t.Fatalf("Dependencies count = %d, want 2", len(p.Dependencies))
	}
	if p.Dependencies[0].Repo != "nvim-lua/plenary.nvim" {
		t.Errorf("Dep[0].Repo = %q, want %q", p.Dependencies[0].Repo, "nvim-lua/plenary.nvim")
	}
	if p.Dependencies[1].Build != "make" {
		t.Errorf("Dep[1].Build = %q, want %q", p.Dependencies[1].Build, "make")
	}

	// Check keymaps
	if len(p.Keymaps) != 1 {
		t.Fatalf("Keymaps count = %d, want 1", len(p.Keymaps))
	}
	if p.Keymaps[0].Key != "<leader>ff" {
		t.Errorf("Keymap.Key = %q, want %q", p.Keymaps[0].Key, "<leader>ff")
	}
}

func TestStringOrSlice(t *testing.T) {
	tests := []struct {
		name     string
		yaml     string
		expected []string
	}{
		{
			name: "single string",
			yaml: `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: test
spec:
  repo: test/test
  event: VeryLazy
`,
			expected: []string{"VeryLazy"},
		},
		{
			name: "array of strings",
			yaml: `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: test
spec:
  repo: test/test
  event:
    - BufRead
    - BufNewFile
`,
			expected: []string{"BufRead", "BufNewFile"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := ParseYAML([]byte(tt.yaml))
			if err != nil {
				t.Fatalf("ParseYAML failed: %v", err)
			}
			if len(p.Event) != len(tt.expected) {
				t.Errorf("Event count = %d, want %d", len(p.Event), len(tt.expected))
			}
			for i, e := range tt.expected {
				if i >= len(p.Event) || p.Event[i] != e {
					t.Errorf("Event[%d] = %q, want %q", i, p.Event[i], e)
				}
			}
		})
	}
}

func TestGenerateLua(t *testing.T) {
	p := &Plugin{
		Name:        "telescope",
		Description: "Fuzzy finder",
		Repo:        "nvim-telescope/telescope.nvim",
		Branch:      "0.1.x",
		Dependencies: []Dependency{
			{Repo: "nvim-lua/plenary.nvim"},
			{Repo: "nvim-telescope/telescope-fzf-native.nvim", Build: "make"},
		},
		Keys: []Keymap{
			{Key: "<leader>ff", Action: "<cmd>Telescope find_files<cr>", Desc: "Find files", Mode: []string{"n"}},
		},
		Config: `local telescope = require("telescope")
telescope.setup({})`,
	}

	g := NewGenerator()
	lua, err := g.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	// Check that output contains expected elements
	checks := []string{
		`"nvim-telescope/telescope.nvim"`,
		`branch = "0.1.x"`,
		`dependencies = {`,
		`"nvim-lua/plenary.nvim"`,
		`build = "make"`,
		`keys = {`,
		`"<leader>ff"`,
		`desc = "Find files"`,
		`config = function()`,
		`local telescope = require("telescope")`,
	}

	for _, check := range checks {
		if !strings.Contains(lua, check) {
			t.Errorf("Generated Lua missing: %q\n\nGenerated:\n%s", check, lua)
		}
	}
}

func TestPluginRoundTrip(t *testing.T) {
	// Create a plugin
	original := &Plugin{
		Name:        "test-plugin",
		Description: "A test plugin",
		Repo:        "test/plugin",
		Branch:      "main",
		Version:     "v1.0.0",
		Priority:    1000,
		Lazy:        true,
		Event:       []string{"VeryLazy"},
		Ft:          []string{"go", "lua"},
		Cmd:         []string{"TestCmd"},
		Dependencies: []Dependency{
			{Repo: "dep/one"},
			{Repo: "dep/two", Build: "make", Version: "v2.0.0"},
		},
		Keys: []Keymap{
			{Key: "<leader>t", Mode: []string{"n"}, Action: ":Test<cr>", Desc: "Run test"},
		},
		Build:    "make",
		Config:   "require('test').setup({})",
		Init:     "vim.g.test = true",
		Category: "testing",
		Tags:     []string{"test", "utility"},
		Enabled:  true,
	}

	// Convert to YAML
	py := original.ToYAML()

	// Convert back to Plugin
	converted := py.ToPlugin()

	// Check fields match
	if converted.Name != original.Name {
		t.Errorf("Name = %q, want %q", converted.Name, original.Name)
	}
	if converted.Repo != original.Repo {
		t.Errorf("Repo = %q, want %q", converted.Repo, original.Repo)
	}
	if len(converted.Dependencies) != len(original.Dependencies) {
		t.Errorf("Dependencies count = %d, want %d", len(converted.Dependencies), len(original.Dependencies))
	}
	if len(converted.Keys) != len(original.Keys) {
		t.Errorf("Keys count = %d, want %d", len(converted.Keys), len(original.Keys))
	}
}

func TestValidation(t *testing.T) {
	tests := []struct {
		name    string
		yaml    string
		wantErr bool
	}{
		{
			name: "valid plugin",
			yaml: `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: test
spec:
  repo: test/test
`,
			wantErr: false,
		},
		{
			name: "missing name",
			yaml: `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  description: no name
spec:
  repo: test/test
`,
			wantErr: true,
		},
		{
			name: "missing repo",
			yaml: `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: test
spec:
  branch: main
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseYAML([]byte(tt.yaml))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseYAML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestGenerateLuaBranchField_Regression254 is a regression test for GitHub issue #254.
// The Lua generator must emit branch = "..." when spec.branch is set.
// Without this, lazy.nvim clones from main instead of the specified branch,
// which breaks plugins like nvim-treesitter that require the master branch.
func TestGenerateLuaBranchField_Regression254(t *testing.T) {
	tests := []struct {
		name           string
		branch         string
		expectBranch   bool
		expectedOutput string
	}{
		{
			name:           "treesitter master branch",
			branch:         "master",
			expectBranch:   true,
			expectedOutput: `branch = "master"`,
		},
		{
			name:           "telescope 0.1.x branch",
			branch:         "0.1.x",
			expectBranch:   true,
			expectedOutput: `branch = "0.1.x"`,
		},
		{
			name:           "harpoon2 branch",
			branch:         "harpoon2",
			expectBranch:   true,
			expectedOutput: `branch = "harpoon2"`,
		},
		{
			name:         "empty branch omitted",
			branch:       "",
			expectBranch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Plugin{
				Name:   "test-plugin",
				Repo:   "test/test-plugin",
				Branch: tt.branch,
			}

			g := NewGenerator()
			lua, err := g.GenerateLua(p)
			if err != nil {
				t.Fatalf("GenerateLua failed: %v", err)
			}

			hasBranch := strings.Contains(lua, "branch = ")
			if tt.expectBranch && !hasBranch {
				t.Errorf("Generated Lua missing branch field\n\nGenerated:\n%s", lua)
			}
			if !tt.expectBranch && hasBranch {
				t.Errorf("Generated Lua has unexpected branch field\n\nGenerated:\n%s", lua)
			}
			if tt.expectBranch && !strings.Contains(lua, tt.expectedOutput) {
				t.Errorf("Generated Lua missing %q\n\nGenerated:\n%s", tt.expectedOutput, lua)
			}
		})
	}
}

// TestGenerateLuaTreesitterFromYAML_Regression254 verifies the full YAML → Plugin → Lua
// pipeline preserves the branch field, matching the exact scenario from issue #254.
func TestGenerateLuaTreesitterFromYAML_Regression254(t *testing.T) {
	yamlData := `
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: treesitter
  description: "Syntax highlighting"
  category: syntax
spec:
  repo: nvim-treesitter/nvim-treesitter
  branch: master
  build: ":TSUpdate"
  config: |
    -- treesitter config
    vim.treesitter.start()
`
	p, err := ParseYAML([]byte(yamlData))
	if err != nil {
		t.Fatalf("ParseYAML failed: %v", err)
	}

	// Verify branch survives YAML parsing
	if p.Branch != "master" {
		t.Fatalf("ParseYAML lost branch: got %q, want %q", p.Branch, "master")
	}

	// Verify branch appears in generated Lua
	g := NewGenerator()
	lua, err := g.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	if !strings.Contains(lua, `branch = "master"`) {
		t.Errorf("Generated Lua missing branch = \"master\" (issue #254 regression)\n\nGenerated:\n%s", lua)
	}

	// Also verify other fields survived
	checks := []string{
		`"nvim-treesitter/nvim-treesitter"`,
		`build = ":TSUpdate"`,
		`config = function()`,
	}
	for _, check := range checks {
		if !strings.Contains(lua, check) {
			t.Errorf("Generated Lua missing: %q\n\nGenerated:\n%s", check, lua)
		}
	}
}

// TestGenerateLuaVersionField verifies the version field is emitted correctly.
func TestGenerateLuaVersionField(t *testing.T) {
	p := &Plugin{
		Name:    "test-plugin",
		Repo:    "test/test-plugin",
		Version: "*",
	}

	g := NewGenerator()
	lua, err := g.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	if !strings.Contains(lua, `version = "*"`) {
		t.Errorf("Generated Lua missing version field\n\nGenerated:\n%s", lua)
	}
}

// TestGenerateLuaBranchAndVersionMutualExclusion verifies both branch and version
// can be set (lazy.nvim resolves precedence, we just emit both).
func TestGenerateLuaBranchAndVersionMutualExclusion(t *testing.T) {
	p := &Plugin{
		Name:    "test-plugin",
		Repo:    "test/test-plugin",
		Branch:  "develop",
		Version: "v1.0.0",
	}

	g := NewGenerator()
	lua, err := g.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	if !strings.Contains(lua, `branch = "develop"`) {
		t.Errorf("Generated Lua missing branch\n\nGenerated:\n%s", lua)
	}
	if !strings.Contains(lua, `version = "v1.0.0"`) {
		t.Errorf("Generated Lua missing version\n\nGenerated:\n%s", lua)
	}
}

func TestGenerateLuaWithKeymaps(t *testing.T) {
	// Test that the 'keymaps' field generates vim.keymap.set() calls in config
	p := &Plugin{
		Name: "test-plugin",
		Repo: "test/test-plugin",
		Keymaps: []Keymap{
			{Key: "<leader>tf", Action: "<cmd>TestFile<cr>", Desc: "Test current file", Mode: []string{"n"}},
			{Key: "<leader>tn", Action: "<cmd>TestNearest<cr>", Desc: "Test nearest", Mode: []string{"n", "v"}},
		},
	}

	g := NewGenerator()
	lua, err := g.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	// Check that output contains vim.keymap.set calls
	checks := []string{
		`config = function()`,
		`vim.keymap.set("n", "<leader>tf"`,
		`"<cmd>TestFile<cr>"`,
		`desc = "Test current file"`,
		`vim.keymap.set({ "n", "v" }, "<leader>tn"`,
		`"<cmd>TestNearest<cr>"`,
		`desc = "Test nearest"`,
	}

	for _, check := range checks {
		if !strings.Contains(lua, check) {
			t.Errorf("Generated Lua missing: %q\n\nGenerated:\n%s", check, lua)
		}
	}
}

func TestGenerateLuaWithConfigAndKeymaps(t *testing.T) {
	// Test that keymaps are appended to existing config
	p := &Plugin{
		Name:   "test-plugin",
		Repo:   "test/test-plugin",
		Config: `require("test").setup({})`,
		Keymaps: []Keymap{
			{Key: "<leader>t", Action: "<cmd>Test<cr>", Desc: "Run test", Mode: []string{"n"}},
		},
	}

	g := NewGenerator()
	lua, err := g.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	// Check that both config and keymaps are present
	checks := []string{
		`config = function()`,
		`require("test").setup({})`,
		`-- Keymaps`,
		`vim.keymap.set("n", "<leader>t"`,
	}

	for _, check := range checks {
		if !strings.Contains(lua, check) {
			t.Errorf("Generated Lua missing: %q\n\nGenerated:\n%s", check, lua)
		}
	}
}
