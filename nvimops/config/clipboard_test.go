package config

// =============================================================================
// Tests for clipboard conditional fix (issue #381)
// Verifies that clipboard settings are wrapped in provider-availability checks
// so that yank does not fail with 'clipboard: no provider' in containers.
// =============================================================================

import (
	"strings"
	"testing"
)

// TestWriteOption_Clipboard_UsesConditional verifies that setting clipboard via
// append: syntax generates a conditional guard rather than an unconditional set.
func TestWriteOption_Clipboard_UsesConditional(t *testing.T) {
	gen := NewGenerator()
	var lua strings.Builder

	gen.writeOption(&lua, "clipboard", "append:unnamedplus")
	output := lua.String()

	if !strings.Contains(output, "vim.fn.has('clipboard')") {
		t.Error("clipboard option missing has('clipboard') guard")
	}
	if !strings.Contains(output, "vim.fn.executable('xclip')") {
		t.Error("clipboard option missing xclip executable check")
	}
	if !strings.Contains(output, "vim.fn.executable('pbcopy')") {
		t.Error("clipboard option missing pbcopy executable check")
	}
	if !strings.Contains(output, "vim.fn.executable('wl-copy')") {
		t.Error("clipboard option missing wl-copy executable check")
	}
	if !strings.Contains(output, "vim.env.SSH_TTY") {
		t.Error("clipboard option missing SSH_TTY check")
	}
	if !strings.Contains(output, `opt.clipboard:append("unnamedplus")`) {
		t.Error("clipboard option missing append call")
	}
	if !strings.Contains(output, "end") {
		t.Error("clipboard conditional missing closing 'end'")
	}
}

// TestWriteOption_Clipboard_DoesNotSetUnconditionally verifies the generated Lua
// does NOT contain a bare unconditional clipboard assignment.
func TestWriteOption_Clipboard_DoesNotSetUnconditionally(t *testing.T) {
	gen := NewGenerator()
	var lua strings.Builder

	gen.writeOption(&lua, "clipboard", "append:unnamedplus")
	output := lua.String()

	// The assignment must not appear outside a conditional block.
	// A simple heuristic: the line must not start without indentation.
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "opt.clipboard") && !strings.HasPrefix(line, " ") {
			t.Errorf("clipboard assignment appears unconditionally (no indentation): %q", line)
		}
	}
}

// TestWriteOption_NonClipboard_NotWrapped verifies non-clipboard append options
// are NOT wrapped in a provider conditional.
func TestWriteOption_NonClipboard_NotWrapped(t *testing.T) {
	gen := NewGenerator()
	var lua strings.Builder

	gen.writeOption(&lua, "shortmess", "append:c")
	output := lua.String()

	if strings.Contains(output, "vim.fn.has('clipboard')") {
		t.Error("non-clipboard option should not contain clipboard conditional")
	}
	if !strings.Contains(output, `opt.shortmess:append("c")`) {
		t.Error("non-clipboard append option not written correctly")
	}
}

// TestGenerateCoreLua_ClipboardConditional verifies end-to-end that when a
// config includes clipboard=append:unnamedplus, the generated init.lua wraps
// the clipboard setting in provider checks.
func TestGenerateCoreLua_ClipboardConditional(t *testing.T) {
	cfg := DefaultCoreConfig()

	// Ensure clipboard option is present (it should be by default)
	found := false
	for k, v := range cfg.Options {
		if k == "clipboard" {
			found = true
			_ = v
		}
	}
	if !found {
		t.Skip("default config does not set clipboard option — skipping end-to-end check")
	}

	gen := NewGenerator()
	lua := gen.generateOptionsLua(cfg)

	if !strings.Contains(lua, "vim.fn.has('clipboard')") {
		t.Error("generated options lua missing clipboard provider conditional")
	}
}
