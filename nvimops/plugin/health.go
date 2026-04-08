// Package plugin provides types and utilities for Neovim plugin management.
package plugin

import (
	"fmt"
	"strings"
)

// DefaultHealthChecks generates default health checks for a plugin based on
// its configuration when no explicit health checks are defined.
// It infers a lua_module check from the plugin name/repo.
func DefaultHealthChecks(p *Plugin) []HealthCheck {
	if len(p.HealthChecks) > 0 {
		return p.HealthChecks
	}

	var checks []HealthCheck

	// Infer lua_module check from the plugin name
	// Convention: plugin "foo-bar" -> try require("foo-bar") and require("foo_bar")
	moduleName := inferModuleName(p)
	if moduleName != "" {
		checks = append(checks, HealthCheck{
			Type:        HealthCheckLuaModule,
			Value:       moduleName,
			Description: fmt.Sprintf("Module '%s' is loadable", moduleName),
		})
	}

	return checks
}

// inferModuleName tries to determine the Lua module name from the plugin config.
// It checks config code for require() calls first, then falls back to repo name.
func inferModuleName(p *Plugin) string {
	// Try to extract from config code — find the first require() call
	if p.Config != "" {
		if mod := extractRequireModule(p.Config); mod != "" {
			return mod
		}
	}

	// Fall back to extracting from repo name
	// e.g. "nvim-telescope/telescope.nvim" -> "telescope"
	if p.Repo != "" {
		parts := strings.Split(p.Repo, "/")
		if len(parts) >= 2 {
			name := parts[len(parts)-1]
			// Strip common suffixes
			name = strings.TrimSuffix(name, ".nvim")
			name = strings.TrimSuffix(name, ".lua")
			name = strings.TrimPrefix(name, "nvim-")
			return name
		}
	}

	return ""
}

// extractRequireModule extracts the first module name from require() calls
// in Lua code. Returns empty string if no require() is found.
func extractRequireModule(luaCode string) string {
	// Look for require("module") or require('module')
	for _, prefix := range []string{`require("`, `require('`, `require "`} {
		idx := strings.Index(luaCode, prefix)
		if idx < 0 {
			continue
		}
		start := idx + len(prefix)
		// Find closing delimiter
		var closer byte
		switch prefix[len(prefix)-1] {
		case '"':
			closer = '"'
		case '\'':
			closer = '\''
		}
		end := strings.IndexByte(luaCode[start:], closer)
		if end > 0 {
			mod := luaCode[start : start+end]
			// Skip sub-modules (e.g. "telescope.actions") - use root
			if dotIdx := strings.IndexByte(mod, '.'); dotIdx > 0 {
				mod = mod[:dotIdx]
			}
			return mod
		}
	}
	return ""
}

// ValidateHealthCheckType checks if a health check type string is valid.
func ValidateHealthCheckType(t string) error {
	switch HealthCheckType(t) {
	case HealthCheckLuaModule, HealthCheckCommand,
		HealthCheckTreesitter, HealthCheckLSP:
		return nil
	default:
		return fmt.Errorf("invalid health check type: %q (valid: lua_module, command, treesitter, lsp)", t)
	}
}
