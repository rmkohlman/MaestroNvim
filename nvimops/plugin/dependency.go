// Package plugin provides types and utilities for Neovim plugin management.
package plugin

import (
	"fmt"
	"strings"
)

// DependencyResolver handles plugin dependency graph analysis including
// transitive dependency resolution, circular dependency detection,
// topological sorting, and diamond dependency deduplication.
type DependencyResolver struct {
	// plugins maps plugin repo (or name) to its Plugin definition
	plugins map[string]*Plugin
}

// NewDependencyResolver creates a new DependencyResolver from a list of plugins.
// Plugins are indexed by both their Name and Repo for flexible lookup.
func NewDependencyResolver(plugins []*Plugin) *DependencyResolver {
	m := make(map[string]*Plugin, len(plugins)*2)
	for _, p := range plugins {
		if p.Name != "" {
			m[p.Name] = p
		}
		if p.Repo != "" {
			m[p.Repo] = p
		}
	}
	return &DependencyResolver{plugins: m}
}

// ResolveError represents an error during dependency resolution.
type ResolveError struct {
	Plugin  string
	Message string
	Chain   []string // dependency chain for context
}

func (e *ResolveError) Error() string {
	if len(e.Chain) > 0 {
		return fmt.Sprintf("dependency error for %q: %s (chain: %s)",
			e.Plugin, e.Message, strings.Join(e.Chain, " → "))
	}
	return fmt.Sprintf("dependency error for %q: %s", e.Plugin, e.Message)
}

// CircularDependencyError is returned when a circular dependency is detected.
type CircularDependencyError struct {
	Cycle []string // the cycle path, e.g. [A, B, C, A]
}

func (e *CircularDependencyError) Error() string {
	return fmt.Sprintf("circular dependency detected: %s", strings.Join(e.Cycle, " → "))
}

// MissingDependencyError is returned when a required dependency is not found.
type MissingDependencyError struct {
	Plugin     string
	Dependency string
}

func (e *MissingDependencyError) Error() string {
	return fmt.Sprintf("plugin %q depends on %q which is not in the plugin set",
		e.Plugin, e.Dependency)
}

// Resolve performs full dependency resolution for a given plugin.
// It returns all transitive dependencies in topological order (dependencies first),
// with diamond dependencies deduplicated.
func (r *DependencyResolver) Resolve(pluginKey string) ([]*Plugin, error) {
	root, ok := r.plugins[pluginKey]
	if !ok {
		return nil, &ResolveError{Plugin: pluginKey, Message: "plugin not found"}
	}

	// Track visited nodes and the current path for cycle detection
	visited := make(map[string]bool)
	inStack := make(map[string]bool)
	var order []*Plugin

	if err := r.dfs(root, visited, inStack, &order, nil); err != nil {
		return nil, err
	}

	return order, nil
}

// ResolveAll resolves dependencies for all plugins and returns a global
// topological ordering. All plugins must be resolvable.
func (r *DependencyResolver) ResolveAll() ([]*Plugin, error) {
	visited := make(map[string]bool)
	inStack := make(map[string]bool)
	var order []*Plugin

	// Process all plugins (use a deduplicated set)
	seen := make(map[string]bool)
	for _, p := range r.plugins {
		if seen[p.Repo] {
			continue
		}
		seen[p.Repo] = true

		if err := r.dfs(p, visited, inStack, &order, nil); err != nil {
			return nil, err
		}
	}

	return order, nil
}

// dfs performs a depth-first traversal for topological sort with cycle detection.
// The path parameter tracks the current traversal chain for error reporting.
func (r *DependencyResolver) dfs(
	p *Plugin,
	visited map[string]bool,
	inStack map[string]bool,
	order *[]*Plugin,
	path []string,
) error {
	key := p.Repo
	if key == "" {
		key = p.Name
	}

	// Cycle detection: if this node is already in the current DFS stack
	if inStack[key] {
		cycle := append(path, key)
		return &CircularDependencyError{Cycle: cycle}
	}

	// Already fully processed — skip (handles diamond dependencies)
	if visited[key] {
		return nil
	}

	inStack[key] = true
	currentPath := append(path, key)

	// Recurse into dependencies
	for _, dep := range p.Dependencies {
		depPlugin, ok := r.plugins[dep.Repo]
		if !ok {
			return &MissingDependencyError{
				Plugin:     key,
				Dependency: dep.Repo,
			}
		}
		if err := r.dfs(depPlugin, visited, inStack, order, currentPath); err != nil {
			return err
		}
	}

	inStack[key] = false
	visited[key] = true
	*order = append(*order, p)
	return nil
}

// DependencyTree represents a plugin and its dependency tree for display.
type DependencyTree struct {
	Plugin   *Plugin
	Children []*DependencyTree
}

// BuildTree builds a dependency tree for display purposes.
// Unlike Resolve, this does not fail on missing deps — it just omits them.
func (r *DependencyResolver) BuildTree(pluginKey string) *DependencyTree {
	p, ok := r.plugins[pluginKey]
	if !ok {
		return nil
	}
	visited := make(map[string]bool)
	return r.buildTreeNode(p, visited)
}

func (r *DependencyResolver) buildTreeNode(p *Plugin, visited map[string]bool) *DependencyTree {
	key := p.Repo
	if key == "" {
		key = p.Name
	}

	node := &DependencyTree{Plugin: p}

	// Prevent infinite recursion on cycles in tree display
	if visited[key] {
		return node
	}
	visited[key] = true

	for _, dep := range p.Dependencies {
		child, ok := r.plugins[dep.Repo]
		if !ok {
			continue
		}
		node.Children = append(node.Children, r.buildTreeNode(child, visited))
	}

	return node
}

// FormatTree returns a human-readable tree string for a DependencyTree.
func FormatTree(tree *DependencyTree) string {
	if tree == nil {
		return ""
	}
	var sb strings.Builder
	name := tree.Plugin.Name
	if name == "" {
		name = tree.Plugin.Repo
	}
	sb.WriteString(name + "\n")
	formatTreeLines(&sb, tree.Children, "")
	return sb.String()
}

func formatTreeLines(sb *strings.Builder, children []*DependencyTree, prefix string) {
	for i, child := range children {
		isLast := i == len(children)-1
		connector := "├── "
		childPrefix := "│   "
		if isLast {
			connector = "└── "
			childPrefix = "    "
		}

		name := child.Plugin.Name
		if name == "" {
			name = child.Plugin.Repo
		}
		sb.WriteString(prefix + connector + name + "\n")

		if len(child.Children) > 0 {
			formatTreeLines(sb, child.Children, prefix+childPrefix)
		}
	}
}
