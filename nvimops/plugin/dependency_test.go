package plugin

import (
	"errors"
	"strings"
	"testing"
)

// helper to create a minimal plugin with dependencies
func depPlugin(name, repo string, deps ...string) *Plugin {
	var depList []Dependency
	for _, d := range deps {
		depList = append(depList, Dependency{Repo: d})
	}
	return &Plugin{
		Name:         name,
		Repo:         repo,
		Enabled:      true,
		Dependencies: depList,
	}
}

func TestDependencyResolver_SimpleChain(t *testing.T) {
	// A → B → C (transitive)
	plugins := []*Plugin{
		depPlugin("A", "user/A", "user/B"),
		depPlugin("B", "user/B", "user/C"),
		depPlugin("C", "user/C"),
	}

	resolver := NewDependencyResolver(plugins)
	order, err := resolver.Resolve("user/A")
	if err != nil {
		t.Fatalf("Resolve failed: %v", err)
	}

	// Order should be: C, B, A (dependencies first)
	if len(order) != 3 {
		t.Fatalf("expected 3 plugins, got %d", len(order))
	}
	if order[0].Name != "C" {
		t.Errorf("order[0] = %q, want C", order[0].Name)
	}
	if order[1].Name != "B" {
		t.Errorf("order[1] = %q, want B", order[1].Name)
	}
	if order[2].Name != "A" {
		t.Errorf("order[2] = %q, want A", order[2].Name)
	}
}

func TestDependencyResolver_CircularDetection(t *testing.T) {
	// A → B → C → A
	plugins := []*Plugin{
		depPlugin("A", "user/A", "user/B"),
		depPlugin("B", "user/B", "user/C"),
		depPlugin("C", "user/C", "user/A"),
	}

	resolver := NewDependencyResolver(plugins)
	_, err := resolver.Resolve("user/A")
	if err == nil {
		t.Fatal("expected circular dependency error, got nil")
	}

	var circErr *CircularDependencyError
	if !errors.As(err, &circErr) {
		t.Fatalf("expected CircularDependencyError, got %T: %v", err, err)
	}

	// Verify the cycle contains A
	errMsg := circErr.Error()
	if !strings.Contains(errMsg, "user/A") {
		t.Errorf("cycle error should mention user/A: %s", errMsg)
	}
}

func TestDependencyResolver_DiamondDependency(t *testing.T) {
	// A → B, A → C, B → D, C → D  (D should appear once)
	plugins := []*Plugin{
		depPlugin("A", "user/A", "user/B", "user/C"),
		depPlugin("B", "user/B", "user/D"),
		depPlugin("C", "user/C", "user/D"),
		depPlugin("D", "user/D"),
	}

	resolver := NewDependencyResolver(plugins)
	order, err := resolver.Resolve("user/A")
	if err != nil {
		t.Fatalf("Resolve failed: %v", err)
	}

	// D should appear exactly once
	if len(order) != 4 {
		t.Fatalf("expected 4 plugins (D deduped), got %d", len(order))
	}

	// D must come before B and C, which come before A
	indexOf := func(name string) int {
		for i, p := range order {
			if p.Name == name {
				return i
			}
		}
		return -1
	}

	dIdx := indexOf("D")
	bIdx := indexOf("B")
	cIdx := indexOf("C")
	aIdx := indexOf("A")

	if dIdx == -1 || bIdx == -1 || cIdx == -1 || aIdx == -1 {
		t.Fatalf("missing plugin in order: D=%d B=%d C=%d A=%d", dIdx, bIdx, cIdx, aIdx)
	}
	if dIdx >= bIdx || dIdx >= cIdx {
		t.Errorf("D (idx %d) should come before B (%d) and C (%d)", dIdx, bIdx, cIdx)
	}
	if bIdx >= aIdx || cIdx >= aIdx {
		t.Errorf("B (%d) and C (%d) should come before A (%d)", bIdx, cIdx, aIdx)
	}
}

func TestDependencyResolver_MissingDependency(t *testing.T) {
	plugins := []*Plugin{
		depPlugin("A", "user/A", "user/missing"),
	}

	resolver := NewDependencyResolver(plugins)
	_, err := resolver.Resolve("user/A")
	if err == nil {
		t.Fatal("expected missing dependency error, got nil")
	}

	var missingErr *MissingDependencyError
	if !errors.As(err, &missingErr) {
		t.Fatalf("expected MissingDependencyError, got %T: %v", err, err)
	}
	if missingErr.Dependency != "user/missing" {
		t.Errorf("missing dep = %q, want user/missing", missingErr.Dependency)
	}
}

func TestDependencyResolver_NoDeps(t *testing.T) {
	plugins := []*Plugin{
		depPlugin("A", "user/A"),
	}

	resolver := NewDependencyResolver(plugins)
	order, err := resolver.Resolve("user/A")
	if err != nil {
		t.Fatalf("Resolve failed: %v", err)
	}
	if len(order) != 1 || order[0].Name != "A" {
		t.Errorf("expected [A], got %v", order)
	}
}

func TestDependencyResolver_ResolveAll(t *testing.T) {
	plugins := []*Plugin{
		depPlugin("A", "user/A", "user/B"),
		depPlugin("B", "user/B"),
		depPlugin("C", "user/C"),
	}

	resolver := NewDependencyResolver(plugins)
	order, err := resolver.ResolveAll()
	if err != nil {
		t.Fatalf("ResolveAll failed: %v", err)
	}

	if len(order) != 3 {
		t.Fatalf("expected 3 plugins, got %d", len(order))
	}

	// B must come before A
	indexOf := func(name string) int {
		for i, p := range order {
			if p.Name == name {
				return i
			}
		}
		return -1
	}
	if indexOf("B") >= indexOf("A") {
		t.Errorf("B should come before A in topological order")
	}
}

func TestDependencyResolver_NotFound(t *testing.T) {
	resolver := NewDependencyResolver(nil)
	_, err := resolver.Resolve("nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent plugin")
	}
}

func TestFormatTree(t *testing.T) {
	plugins := []*Plugin{
		depPlugin("telescope", "nvim-telescope/telescope.nvim", "nvim-lua/plenary.nvim", "nvim-telescope/telescope-fzf-native.nvim"),
		depPlugin("plenary", "nvim-lua/plenary.nvim"),
		depPlugin("fzf-native", "nvim-telescope/telescope-fzf-native.nvim"),
	}

	resolver := NewDependencyResolver(plugins)
	tree := resolver.BuildTree("nvim-telescope/telescope.nvim")
	if tree == nil {
		t.Fatal("BuildTree returned nil")
	}

	output := FormatTree(tree)
	if !strings.Contains(output, "telescope") {
		t.Errorf("tree output missing telescope: %s", output)
	}
	if !strings.Contains(output, "├── ") || !strings.Contains(output, "└── ") {
		t.Errorf("tree output missing tree connectors: %s", output)
	}
}

func TestDependencyResolver_LookupByName(t *testing.T) {
	// Verify that plugins can be looked up by name as well as repo
	plugins := []*Plugin{
		depPlugin("telescope", "nvim-telescope/telescope.nvim"),
	}

	resolver := NewDependencyResolver(plugins)
	order, err := resolver.Resolve("telescope")
	if err != nil {
		t.Fatalf("Resolve by name failed: %v", err)
	}
	if len(order) != 1 || order[0].Name != "telescope" {
		t.Errorf("expected [telescope], got %v", order)
	}
}
