// Package plugin provides types and utilities for Neovim plugin management.
package plugin

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// LockEntry represents a single plugin's pinned state in a lock file.
type LockEntry struct {
	Branch string `json:"branch"`
	Commit string `json:"commit"`
}

// LockFile represents the lazy-lock.json format used by lazy.nvim.
// It maps plugin short names to their pinned branch and commit.
type LockFile struct {
	Entries map[string]LockEntry
}

// NewLockFile creates an empty LockFile.
func NewLockFile() *LockFile {
	return &LockFile{Entries: make(map[string]LockEntry)}
}

// ParseLockFile reads and parses a lazy-lock.json file.
func ParseLockFile(path string) (*LockFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read lock file: %w", err)
	}
	return ParseLockFileData(data)
}

// ParseLockFileData parses lazy-lock.json content from bytes.
func ParseLockFileData(data []byte) (*LockFile, error) {
	var entries map[string]LockEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("failed to parse lock file: %w", err)
	}
	if entries == nil {
		entries = make(map[string]LockEntry)
	}
	return &LockFile{Entries: entries}, nil
}

// GenerateLockFile creates a LockFile from a set of plugins.
// Each plugin's current branch and version (as commit) are recorded.
// If commit is empty, the plugin is skipped.
func GenerateLockFile(plugins []*Plugin) *LockFile {
	lf := NewLockFile()
	for _, p := range plugins {
		if !p.Enabled {
			continue
		}
		// Use the short name (last segment of repo) as the key, matching lazy.nvim convention
		shortName := repoShortName(p.Repo)
		if shortName == "" {
			continue
		}

		entry := LockEntry{
			Branch: p.Branch,
			Commit: p.Version, // version field holds the commit/tag
		}
		// Only include if there's a commit to pin
		if entry.Commit != "" || entry.Branch != "" {
			lf.Entries[shortName] = entry
		}
	}
	return lf
}

// Marshal serializes the lock file to JSON bytes matching lazy-lock.json format.
// Output is sorted by plugin name for deterministic output.
func (lf *LockFile) Marshal() ([]byte, error) {
	// Use sorted keys for deterministic output
	sorted := make(map[string]LockEntry, len(lf.Entries))
	for k, v := range lf.Entries {
		sorted[k] = v
	}
	return json.MarshalIndent(sorted, "", "  ")
}

// WriteTo writes the lock file to the specified path.
func (lf *LockFile) WriteTo(path string) error {
	data, err := lf.Marshal()
	if err != nil {
		return fmt.Errorf("failed to marshal lock file: %w", err)
	}
	// Append newline for POSIX compliance
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

// repoShortName extracts the short name from a GitHub repo path.
// e.g. "nvim-telescope/telescope.nvim" → "telescope.nvim"
func repoShortName(repo string) string {
	parts := strings.Split(repo, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return repo
}

// LockMismatch describes a difference between the current config and lock file.
type LockMismatch struct {
	Plugin   string
	Field    string // "commit", "branch", or "missing"
	Expected string
	Actual   string
}

func (m *LockMismatch) String() string {
	if m.Field == "missing_in_lock" {
		return fmt.Sprintf("%s: not in lock file", m.Plugin)
	}
	if m.Field == "missing_in_config" {
		return fmt.Sprintf("%s: in lock file but not in config", m.Plugin)
	}
	return fmt.Sprintf("%s: %s mismatch (lock=%q, config=%q)",
		m.Plugin, m.Field, m.Expected, m.Actual)
}

// Verify checks if the given plugins match the lock file entries.
// Returns a list of mismatches; empty list means everything matches.
func (lf *LockFile) Verify(plugins []*Plugin) []LockMismatch {
	var mismatches []LockMismatch

	// Build map of current plugins by short name
	current := make(map[string]*Plugin)
	for _, p := range plugins {
		if !p.Enabled {
			continue
		}
		shortName := repoShortName(p.Repo)
		current[shortName] = p
	}

	// Check lock entries against current config
	for name, entry := range lf.Entries {
		p, ok := current[name]
		if !ok {
			mismatches = append(mismatches, LockMismatch{
				Plugin: name,
				Field:  "missing_in_config",
			})
			continue
		}
		if entry.Commit != "" && p.Version != entry.Commit {
			mismatches = append(mismatches, LockMismatch{
				Plugin:   name,
				Field:    "commit",
				Expected: entry.Commit,
				Actual:   p.Version,
			})
		}
		if entry.Branch != "" && p.Branch != entry.Branch {
			mismatches = append(mismatches, LockMismatch{
				Plugin:   name,
				Field:    "branch",
				Expected: entry.Branch,
				Actual:   p.Branch,
			})
		}
	}

	// Check for plugins not in lock file
	for name := range current {
		if _, ok := lf.Entries[name]; !ok {
			mismatches = append(mismatches, LockMismatch{
				Plugin: name,
				Field:  "missing_in_lock",
			})
		}
	}

	return mismatches
}
