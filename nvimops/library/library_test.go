package library

import (
	"strings"
	"testing"

	"github.com/rmkohlman/MaestroNvim/nvimops/plugin"
)

func TestLibrary(t *testing.T) {
	lib, err := NewLibrary()
	if err != nil {
		t.Fatalf("NewLibrary failed: %v", err)
	}

	// Should have plugins loaded
	count := lib.Count()
	if count == 0 {
		t.Error("Library should have plugins loaded")
	}
	t.Logf("Loaded %d plugins from library", count)

	// Test Get
	telescope, ok := lib.Get("telescope")
	if !ok {
		t.Error("Library should have telescope plugin")
	}
	if telescope != nil && telescope.Repo != "nvim-telescope/telescope.nvim" {
		t.Errorf("Telescope repo = %q, want nvim-telescope/telescope.nvim", telescope.Repo)
	}

	// Test List
	plugins := lib.List()
	if len(plugins) != count {
		t.Errorf("List() returned %d plugins, want %d", len(plugins), count)
	}

	// Test Categories
	categories := lib.Categories()
	if len(categories) == 0 {
		t.Error("Library should have categories")
	}
	t.Logf("Categories: %v", categories)

	// Test Tags
	tags := lib.Tags()
	if len(tags) == 0 {
		t.Error("Library should have tags")
	}
	t.Logf("Tags: %v", tags)

	// Test ListByCategory
	lspPlugins := lib.ListByCategory("lsp")
	t.Logf("LSP plugins: %d", len(lspPlugins))

	// Test ListByTag
	finderPlugins := lib.ListByTag("finder")
	t.Logf("Finder plugins: %d", len(finderPlugins))

	// Test Info
	info := lib.Info()
	if len(info) != count {
		t.Errorf("Info() returned %d items, want %d", len(info), count)
	}
}

func TestLibraryPluginContent(t *testing.T) {
	lib, err := NewLibrary()
	if err != nil {
		t.Fatalf("NewLibrary failed: %v", err)
	}

	// Check some expected plugins
	expectedPlugins := []string{"telescope", "treesitter", "lspconfig"}

	for _, name := range expectedPlugins {
		p, ok := lib.Get(name)
		if !ok {
			t.Errorf("Expected plugin %q not found", name)
			continue
		}

		// Basic validation
		if p.Name == "" {
			t.Errorf("Plugin %q has empty name", name)
		}
		if p.Repo == "" {
			t.Errorf("Plugin %q has empty repo", name)
		}
	}
}

// TestLibraryBranchPreservation_Regression254 verifies that plugins with branch fields
// in the embedded YAML library retain those fields after parsing.
// Regression guard for GitHub issue #254.
func TestLibraryBranchPreservation_Regression254(t *testing.T) {
	lib, err := NewLibrary()
	if err != nil {
		t.Fatalf("NewLibrary failed: %v", err)
	}

	// Plugins that MUST have branch set in the library
	branchExpectations := map[string]string{
		"treesitter": "master",
		"telescope":  "0.1.x",
		"harpoon":    "harpoon2",
	}

	for name, expectedBranch := range branchExpectations {
		p, ok := lib.Get(name)
		if !ok {
			t.Errorf("Library plugin %q not found", name)
			continue
		}

		if p.Branch != expectedBranch {
			t.Errorf("Plugin %q: Branch = %q, want %q (issue #254 regression)",
				name, p.Branch, expectedBranch)
		}
	}
}

// TestLibraryTreesitterLuaGeneration_Regression254 is an end-to-end test that loads
// treesitter from the embedded library and verifies the generated Lua contains branch.
func TestLibraryTreesitterLuaGeneration_Regression254(t *testing.T) {
	lib, err := NewLibrary()
	if err != nil {
		t.Fatalf("NewLibrary failed: %v", err)
	}

	ts, ok := lib.Get("treesitter")
	if !ok {
		t.Fatal("treesitter plugin not found in library")
	}

	if ts.Branch != "master" {
		t.Fatalf("treesitter Branch = %q, want %q", ts.Branch, "master")
	}

	// Generate Lua and verify branch is present
	gen := plugin.NewGenerator()
	lua, err := gen.GenerateLua(ts)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	if !strings.Contains(lua, `branch = "master"`) {
		t.Errorf("treesitter Lua missing branch = \"master\" (issue #254 regression)\n\nGenerated:\n%s", lua)
	}

	// Also verify other expected fields
	if !strings.Contains(lua, `"nvim-treesitter/nvim-treesitter"`) {
		t.Errorf("treesitter Lua missing repo\n\nGenerated:\n%s", lua)
	}
	if !strings.Contains(lua, `build = ":TSUpdate"`) {
		t.Errorf("treesitter Lua missing build\n\nGenerated:\n%s", lua)
	}
}
