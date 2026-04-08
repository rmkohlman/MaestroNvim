package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseLockFileData(t *testing.T) {
	data := []byte(`{
  "telescope.nvim": {
    "branch": "master",
    "commit": "abc123def456"
  },
  "plenary.nvim": {
    "branch": "main",
    "commit": "789xyz000111"
  }
}`)

	lf, err := ParseLockFileData(data)
	if err != nil {
		t.Fatalf("ParseLockFileData failed: %v", err)
	}

	if len(lf.Entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(lf.Entries))
	}

	entry, ok := lf.Entries["telescope.nvim"]
	if !ok {
		t.Fatal("missing telescope.nvim entry")
	}
	if entry.Branch != "master" {
		t.Errorf("branch = %q, want master", entry.Branch)
	}
	if entry.Commit != "abc123def456" {
		t.Errorf("commit = %q, want abc123def456", entry.Commit)
	}
}

func TestGenerateLockFile(t *testing.T) {
	plugins := []*Plugin{
		{Name: "telescope", Repo: "nvim-telescope/telescope.nvim", Branch: "master", Version: "abc123", Enabled: true},
		{Name: "plenary", Repo: "nvim-lua/plenary.nvim", Branch: "main", Version: "def456", Enabled: true},
		{Name: "disabled", Repo: "user/disabled", Branch: "main", Version: "xxx", Enabled: false},
	}

	lf := GenerateLockFile(plugins)
	if len(lf.Entries) != 2 {
		t.Fatalf("expected 2 entries (disabled skipped), got %d", len(lf.Entries))
	}

	entry, ok := lf.Entries["telescope.nvim"]
	if !ok {
		t.Fatal("missing telescope.nvim entry")
	}
	if entry.Commit != "abc123" {
		t.Errorf("commit = %q, want abc123", entry.Commit)
	}
}

func TestLockFileRoundTrip(t *testing.T) {
	original := NewLockFile()
	original.Entries["telescope.nvim"] = LockEntry{Branch: "master", Commit: "abc123"}
	original.Entries["plenary.nvim"] = LockEntry{Branch: "main", Commit: "def456"}

	data, err := original.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	parsed, err := ParseLockFileData(data)
	if err != nil {
		t.Fatalf("ParseLockFileData failed: %v", err)
	}

	if len(parsed.Entries) != len(original.Entries) {
		t.Fatalf("entry count mismatch: got %d, want %d", len(parsed.Entries), len(original.Entries))
	}

	for name, origEntry := range original.Entries {
		parsedEntry, ok := parsed.Entries[name]
		if !ok {
			t.Errorf("missing entry: %s", name)
			continue
		}
		if parsedEntry.Branch != origEntry.Branch {
			t.Errorf("%s branch = %q, want %q", name, parsedEntry.Branch, origEntry.Branch)
		}
		if parsedEntry.Commit != origEntry.Commit {
			t.Errorf("%s commit = %q, want %q", name, parsedEntry.Commit, origEntry.Commit)
		}
	}
}

func TestLockFileWriteAndRead(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "lazy-lock.json")

	lf := NewLockFile()
	lf.Entries["test.nvim"] = LockEntry{Branch: "main", Commit: "abc123"}

	if err := lf.WriteTo(path); err != nil {
		t.Fatalf("WriteTo failed: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("lock file not created: %v", err)
	}

	// Read it back
	parsed, err := ParseLockFile(path)
	if err != nil {
		t.Fatalf("ParseLockFile failed: %v", err)
	}

	if len(parsed.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(parsed.Entries))
	}
	if parsed.Entries["test.nvim"].Commit != "abc123" {
		t.Errorf("commit mismatch after round-trip")
	}
}

func TestLockFileVerify_Match(t *testing.T) {
	lf := NewLockFile()
	lf.Entries["telescope.nvim"] = LockEntry{Branch: "master", Commit: "abc123"}

	plugins := []*Plugin{
		{Name: "telescope", Repo: "nvim-telescope/telescope.nvim", Branch: "master", Version: "abc123", Enabled: true},
	}

	mismatches := lf.Verify(plugins)
	if len(mismatches) != 0 {
		t.Errorf("expected no mismatches, got %d: %v", len(mismatches), mismatches)
	}
}

func TestLockFileVerify_CommitMismatch(t *testing.T) {
	lf := NewLockFile()
	lf.Entries["telescope.nvim"] = LockEntry{Branch: "master", Commit: "abc123"}

	plugins := []*Plugin{
		{Name: "telescope", Repo: "nvim-telescope/telescope.nvim", Branch: "master", Version: "different", Enabled: true},
	}

	mismatches := lf.Verify(plugins)
	if len(mismatches) != 1 {
		t.Fatalf("expected 1 mismatch, got %d", len(mismatches))
	}
	if mismatches[0].Field != "commit" {
		t.Errorf("mismatch field = %q, want commit", mismatches[0].Field)
	}
}

func TestLockFileVerify_MissingInLock(t *testing.T) {
	lf := NewLockFile()
	// Empty lock file

	plugins := []*Plugin{
		{Name: "telescope", Repo: "nvim-telescope/telescope.nvim", Enabled: true},
	}

	mismatches := lf.Verify(plugins)
	found := false
	for _, m := range mismatches {
		if m.Field == "missing_in_lock" {
			found = true
		}
	}
	if !found {
		t.Error("expected missing_in_lock mismatch")
	}
}

func TestLockFileVerify_MissingInConfig(t *testing.T) {
	lf := NewLockFile()
	lf.Entries["removed.nvim"] = LockEntry{Branch: "main", Commit: "xxx"}

	plugins := []*Plugin{} // empty config

	mismatches := lf.Verify(plugins)
	found := false
	for _, m := range mismatches {
		if m.Field == "missing_in_config" && m.Plugin == "removed.nvim" {
			found = true
		}
	}
	if !found {
		t.Error("expected missing_in_config mismatch for removed.nvim")
	}
}

func TestParseLockFile_MissingFile(t *testing.T) {
	_, err := ParseLockFile("/nonexistent/path/lazy-lock.json")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestRepoShortName(t *testing.T) {
	tests := []struct {
		repo     string
		expected string
	}{
		{"nvim-telescope/telescope.nvim", "telescope.nvim"},
		{"nvim-lua/plenary.nvim", "plenary.nvim"},
		{"simple", "simple"},
		{"a/b/c", "c"},
	}
	for _, tt := range tests {
		got := repoShortName(tt.repo)
		if got != tt.expected {
			t.Errorf("repoShortName(%q) = %q, want %q", tt.repo, got, tt.expected)
		}
	}
}

func TestGenerateLuaWithLockFile(t *testing.T) {
	lf := NewLockFile()
	lf.Entries["telescope.nvim"] = LockEntry{Branch: "master", Commit: "abc123def456"}

	gen := NewGeneratorWithLock(lf)

	p := &Plugin{
		Name:    "telescope",
		Repo:    "nvim-telescope/telescope.nvim",
		Branch:  "master",
		Enabled: true,
	}

	lua, err := gen.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	if !strings.Contains(lua, `commit = "abc123def456"`) {
		t.Errorf("generated Lua missing commit pin\n\nGenerated:\n%s", lua)
	}
}

func TestGenerateLuaWithoutLockFile(t *testing.T) {
	gen := NewGenerator()

	p := &Plugin{
		Name:    "telescope",
		Repo:    "nvim-telescope/telescope.nvim",
		Enabled: true,
	}

	lua, err := gen.GenerateLua(p)
	if err != nil {
		t.Fatalf("GenerateLua failed: %v", err)
	}

	if strings.Contains(lua, "commit =") {
		t.Errorf("generated Lua should NOT contain commit without lock file\n\nGenerated:\n%s", lua)
	}
}
