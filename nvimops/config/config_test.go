package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDefaultCoreConfig(t *testing.T) {
	cfg := DefaultCoreConfig()

	if cfg.Namespace != "workspace" {
		t.Errorf("expected namespace 'workspace', got %q", cfg.Namespace)
	}

	if cfg.Leader != " " {
		t.Errorf("expected leader ' ', got %q", cfg.Leader)
	}

	// Check options
	if v, ok := cfg.Options["relativenumber"]; !ok || v != true {
		t.Error("expected relativenumber = true")
	}

	if v, ok := cfg.Options["tabstop"]; !ok || v != 2 {
		t.Error("expected tabstop = 2")
	}

	// Check keymaps
	if len(cfg.Keymaps) == 0 {
		t.Error("expected default keymaps")
	}

	// Check autocmds
	if len(cfg.Autocmds) == 0 {
		t.Error("expected default autocmds")
	}

	// Check base plugins
	if len(cfg.BasePlugins) == 0 {
		t.Error("expected default base plugins")
	}
}

func TestParseYAML(t *testing.T) {
	yaml := `
apiVersion: nvp/v1
kind: CoreConfig
namespace: myproject
leader: " "
options:
  relativenumber: true
  number: true
  tabstop: 4
keymaps:
  - mode: n
    key: "<leader>ff"
    action: "<cmd>Telescope find_files<cr>"
    desc: Find files
autocmds:
  - group: MyGroup
    events: [BufEnter]
    callback: print("hello")
basePlugins:
  - nvim-lua/plenary.nvim
`

	cfg, err := ParseYAML([]byte(yaml))
	if err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	if cfg.Namespace != "myproject" {
		t.Errorf("expected namespace 'myproject', got %q", cfg.Namespace)
	}

	if v, ok := cfg.Options["tabstop"]; !ok || v != 4 {
		t.Errorf("expected tabstop = 4, got %v", v)
	}

	if len(cfg.Keymaps) != 1 {
		t.Errorf("expected 1 keymap, got %d", len(cfg.Keymaps))
	}

	if len(cfg.Autocmds) != 1 {
		t.Errorf("expected 1 autocmd, got %d", len(cfg.Autocmds))
	}
}

func TestParseYAML_Defaults(t *testing.T) {
	yaml := `
options:
  number: true
`
	cfg, err := ParseYAML([]byte(yaml))
	if err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	// Should have defaults applied
	if cfg.Namespace != "workspace" {
		t.Errorf("expected default namespace 'workspace', got %q", cfg.Namespace)
	}

	if cfg.Leader != " " {
		t.Errorf("expected default leader ' ', got %q", cfg.Leader)
	}
}

func TestGenerator_GenerateInitLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{Namespace: "workspace", Leader: " "}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	expected := `-- Set leader key BEFORE loading lazy.nvim (required for mappings)
vim.g.mapleader = " "
vim.g.maplocalleader = " "

require("workspace.core")
require("workspace.lazy")
`
	if generated.InitLua != expected {
		t.Errorf("InitLua mismatch:\ngot:\n%s\nexpected:\n%s", generated.InitLua, expected)
	}
}

func TestGenerator_GenerateLazyLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{Namespace: "workspace"}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	// Check key parts of lazy.lua
	if !strings.Contains(generated.LazyLua, "lazy/lazy.nvim") {
		t.Error("LazyLua should contain lazy.nvim path")
	}
	if !strings.Contains(generated.LazyLua, `import = "workspace.plugins"`) {
		t.Error("LazyLua should import workspace.plugins")
	}
}

func TestGenerator_GenerateCoreInitLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{Namespace: "workspace"}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	expected := `require("workspace.core.options")
require("workspace.core.keymaps")
require("workspace.core.autocmds")
`
	if generated.CoreInitLua != expected {
		t.Errorf("CoreInitLua mismatch:\ngot:\n%s\nexpected:\n%s", generated.CoreInitLua, expected)
	}
}

func TestGenerator_GenerateOptionsLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		Globals: map[string]interface{}{
			"netrw_liststyle": 3,
		},
		Options: map[string]interface{}{
			"relativenumber": true,
			"number":         true,
			"tabstop":        2,
			"shiftwidth":     2,
			"background":     "dark",
		},
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	// Check options.lua content
	if !strings.Contains(generated.OptionsLua, "let g:netrw_liststyle = 3") {
		t.Error("OptionsLua should contain netrw_liststyle")
	}
	if !strings.Contains(generated.OptionsLua, "opt.relativenumber = true") {
		t.Error("OptionsLua should contain relativenumber")
	}
	if !strings.Contains(generated.OptionsLua, "opt.tabstop = 2") {
		t.Error("OptionsLua should contain tabstop")
	}
	if !strings.Contains(generated.OptionsLua, `opt.background = "dark"`) {
		t.Error("OptionsLua should contain background")
	}
}

func TestGenerator_GenerateKeymapsLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		Leader:    " ",
		Keymaps: []Keymap{
			{Mode: "i", Key: "kj", Action: "<ESC>", Desc: "Exit insert mode"},
			{Mode: "n", Key: "<leader>sv", Action: "<C-w>v", Desc: "Split vertically"},
			{Mode: "v", Key: "J", Action: ":m '>+1<CR>gv=gv", Silent: true},
		},
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	// Check keymaps.lua content
	if !strings.Contains(generated.KeymapsLua, `vim.g.mapleader = " "`) {
		t.Error("KeymapsLua should set mapleader")
	}
	if !strings.Contains(generated.KeymapsLua, `keymap.set("i", "kj"`) {
		t.Error("KeymapsLua should contain kj keymap")
	}
	if !strings.Contains(generated.KeymapsLua, `desc = "Exit insert mode"`) {
		t.Error("KeymapsLua should contain description")
	}
	if !strings.Contains(generated.KeymapsLua, `silent = true`) {
		t.Error("KeymapsLua should contain silent option")
	}
}

func TestGenerator_GenerateAutocmdsLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		Autocmds: []Autocmd{
			{
				Group:    "HighlightYank",
				Events:   []string{"TextYankPost"},
				Callback: `vim.highlight.on_yank({ timeout = 300 })`,
			},
		},
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	// Check autocmds.lua content
	if !strings.Contains(generated.AutocmdsLua, `nvim_create_augroup("HighlightYank"`) {
		t.Error("AutocmdsLua should create augroup")
	}
	if !strings.Contains(generated.AutocmdsLua, `"TextYankPost"`) {
		t.Error("AutocmdsLua should contain event")
	}
	if !strings.Contains(generated.AutocmdsLua, `vim.highlight.on_yank`) {
		t.Error("AutocmdsLua should contain callback")
	}
}

func TestGenerator_GeneratePluginsInitLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		BasePlugins: []string{
			"nvim-lua/plenary.nvim",
			"christoomey/vim-tmux-navigator",
		},
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	expected := `return {
	"nvim-lua/plenary.nvim",
	"christoomey/vim-tmux-navigator",
}
`
	if generated.PluginsInitLua != expected {
		t.Errorf("PluginsInitLua mismatch:\ngot:\n%s\nexpected:\n%s", generated.PluginsInitLua, expected)
	}
}

func TestGenerator_WriteToDirectory(t *testing.T) {
	// Create temp directory
	tmpDir, err := os.MkdirTemp("", "nvp-config-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	gen := NewGenerator()
	cfg := DefaultCoreConfig()

	err = gen.WriteToDirectory(cfg, nil, tmpDir)
	if err != nil {
		t.Fatalf("failed to write to directory: %v", err)
	}

	// Check that all expected files exist
	expectedFiles := []string{
		"init.lua",
		"lua/workspace/lazy.lua",
		"lua/workspace/core/init.lua",
		"lua/workspace/core/options.lua",
		"lua/workspace/core/keymaps.lua",
		"lua/workspace/core/autocmds.lua",
		"lua/workspace/plugins/init.lua",
	}

	for _, file := range expectedFiles {
		path := filepath.Join(tmpDir, file)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected file %s does not exist", file)
		}
	}

	// Verify init.lua content
	initContent, err := os.ReadFile(filepath.Join(tmpDir, "init.lua"))
	if err != nil {
		t.Fatalf("failed to read init.lua: %v", err)
	}
	if !strings.Contains(string(initContent), `require("workspace.core")`) {
		t.Error("init.lua should require workspace.core")
	}
}

func TestToYAML(t *testing.T) {
	cfg := DefaultCoreConfig()
	data, err := cfg.ToYAML()
	if err != nil {
		t.Fatalf("failed to serialize to YAML: %v", err)
	}

	// Parse it back
	cfg2, err := ParseYAML(data)
	if err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	if cfg2.Namespace != cfg.Namespace {
		t.Errorf("namespace mismatch after round-trip")
	}

	if cfg2.Leader != cfg.Leader {
		t.Errorf("leader mismatch after round-trip")
	}
}

func TestParseYAML_CustomHighlights(t *testing.T) {
	yaml := `
namespace: workspace
customHighlights:
  MyCustomGroup:
    fg: "#ff0000"
    bg: "#000000"
    bold: true
  AnotherGroup:
    link: Comment
  UnderlinedGroup:
    fg: "#00ff00"
    sp: "#ff0000"
    underline: true
    italic: true
`
	cfg, err := ParseYAML([]byte(yaml))
	if err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}

	if len(cfg.CustomHighlights) != 3 {
		t.Fatalf("expected 3 custom highlights, got %d", len(cfg.CustomHighlights))
	}

	// Check MyCustomGroup
	hl, ok := cfg.CustomHighlights["MyCustomGroup"]
	if !ok {
		t.Fatal("expected MyCustomGroup highlight")
	}
	if hl.Fg != "#ff0000" {
		t.Errorf("expected fg '#ff0000', got %q", hl.Fg)
	}
	if hl.Bg != "#000000" {
		t.Errorf("expected bg '#000000', got %q", hl.Bg)
	}
	if !hl.Bold {
		t.Error("expected bold = true")
	}

	// Check AnotherGroup (link)
	hl2, ok := cfg.CustomHighlights["AnotherGroup"]
	if !ok {
		t.Fatal("expected AnotherGroup highlight")
	}
	if hl2.Link != "Comment" {
		t.Errorf("expected link 'Comment', got %q", hl2.Link)
	}

	// Check UnderlinedGroup
	hl3, ok := cfg.CustomHighlights["UnderlinedGroup"]
	if !ok {
		t.Fatal("expected UnderlinedGroup highlight")
	}
	if !hl3.Underline {
		t.Error("expected underline = true")
	}
	if !hl3.Italic {
		t.Error("expected italic = true")
	}
	if hl3.Sp != "#ff0000" {
		t.Errorf("expected sp '#ff0000', got %q", hl3.Sp)
	}
}

func TestGenerator_GenerateHighlightsLua(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		CustomHighlights: map[string]HighlightGroup{
			"MyCustomGroup": {
				Fg:   "#ff0000",
				Bg:   "#000000",
				Bold: true,
			},
			"LinkedGroup": {
				Link: "Comment",
			},
			"FancyGroup": {
				Fg:        "#00ff00",
				Sp:        "#0000ff",
				Italic:    true,
				Undercurl: true,
			},
		},
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	lua := generated.HighlightsLua

	// Check header
	if !strings.Contains(lua, "-- Custom highlight groups") {
		t.Error("HighlightsLua should contain header comment")
	}

	// Check FancyGroup attributes
	if !strings.Contains(lua, `nvim_set_hl(0, "FancyGroup", { fg = "#00ff00", sp = "#0000ff", italic = true, undercurl = true })`) {
		t.Errorf("HighlightsLua should contain FancyGroup with correct attrs, got:\n%s", lua)
	}

	// Check LinkedGroup (link syntax)
	if !strings.Contains(lua, `nvim_set_hl(0, "LinkedGroup", { link = "Comment" })`) {
		t.Errorf("HighlightsLua should contain LinkedGroup link, got:\n%s", lua)
	}

	// Check MyCustomGroup
	if !strings.Contains(lua, `nvim_set_hl(0, "MyCustomGroup", { fg = "#ff0000", bg = "#000000", bold = true })`) {
		t.Errorf("HighlightsLua should contain MyCustomGroup, got:\n%s", lua)
	}
}

func TestGenerator_HighlightsLua_Empty(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	if generated.HighlightsLua != "" {
		t.Errorf("expected empty HighlightsLua when no custom highlights, got:\n%s", generated.HighlightsLua)
	}
}

func TestGenerator_CoreInitLua_IncludesHighlights(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		CustomHighlights: map[string]HighlightGroup{
			"MyGroup": {Fg: "#ff0000"},
		},
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	if !strings.Contains(generated.CoreInitLua, `require("workspace.core.highlights")`) {
		t.Errorf("CoreInitLua should require highlights when custom highlights present, got:\n%s", generated.CoreInitLua)
	}
}

func TestGenerator_CoreInitLua_ExcludesHighlightsWhenEmpty(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	if strings.Contains(generated.CoreInitLua, "highlights") {
		t.Errorf("CoreInitLua should NOT require highlights when none defined, got:\n%s", generated.CoreInitLua)
	}
}

func TestGenerator_WriteToDirectory_WithHighlights(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "nvp-highlights-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		CustomHighlights: map[string]HighlightGroup{
			"TestHighlight": {Fg: "#ff0000", Bold: true},
		},
	}

	err = gen.WriteToDirectory(cfg, nil, tmpDir)
	if err != nil {
		t.Fatalf("failed to write to directory: %v", err)
	}

	// Check highlights.lua was written
	hlPath := filepath.Join(tmpDir, "lua", "workspace", "core", "highlights.lua")
	content, err := os.ReadFile(hlPath)
	if err != nil {
		t.Fatalf("highlights.lua not found: %v", err)
	}

	if !strings.Contains(string(content), `nvim_set_hl(0, "TestHighlight"`) {
		t.Errorf("highlights.lua should contain TestHighlight, got:\n%s", string(content))
	}
}

func TestGenerator_WriteToDirectory_NoHighlightsFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "nvp-no-highlights-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
	}

	err = gen.WriteToDirectory(cfg, nil, tmpDir)
	if err != nil {
		t.Fatalf("failed to write to directory: %v", err)
	}

	// Check highlights.lua was NOT written
	hlPath := filepath.Join(tmpDir, "lua", "workspace", "core", "highlights.lua")
	if _, err := os.Stat(hlPath); !os.IsNotExist(err) {
		t.Error("highlights.lua should NOT exist when no custom highlights defined")
	}
}

func TestCustomHighlights_AllAttributes(t *testing.T) {
	gen := NewGenerator()
	cfg := &CoreConfig{
		Namespace: "workspace",
		CustomHighlights: map[string]HighlightGroup{
			"FullGroup": {
				Fg:            "#aabbcc",
				Bg:            "#112233",
				Sp:            "#445566",
				Bold:          true,
				Italic:        true,
				Underline:     true,
				Undercurl:     true,
				Strikethrough: true,
				Reverse:       true,
			},
		},
	}

	generated, err := gen.Generate(cfg)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	lua := generated.HighlightsLua

	// Verify all attributes are present
	for _, attr := range []string{"fg = \"#aabbcc\"", "bg = \"#112233\"", "sp = \"#445566\"",
		"bold = true", "italic = true", "underline = true", "undercurl = true",
		"strikethrough = true", "reverse = true"} {
		if !strings.Contains(lua, attr) {
			t.Errorf("HighlightsLua missing attribute %q, got:\n%s", attr, lua)
		}
	}
}

func TestCustomHighlights_YAMLRoundTrip(t *testing.T) {
	cfg := &CoreConfig{
		Namespace: "workspace",
		Leader:    " ",
		CustomHighlights: map[string]HighlightGroup{
			"RoundTrip": {
				Fg:     "#ff0000",
				Bold:   true,
				Italic: true,
			},
			"LinkedHL": {
				Link: "Normal",
			},
		},
	}

	data, err := cfg.ToYAML()
	if err != nil {
		t.Fatalf("failed to serialize: %v", err)
	}

	cfg2, err := ParseYAML(data)
	if err != nil {
		t.Fatalf("failed to parse: %v", err)
	}

	if len(cfg2.CustomHighlights) != 2 {
		t.Fatalf("expected 2 custom highlights after round-trip, got %d", len(cfg2.CustomHighlights))
	}

	rt, ok := cfg2.CustomHighlights["RoundTrip"]
	if !ok {
		t.Fatal("expected RoundTrip highlight after round-trip")
	}
	if rt.Fg != "#ff0000" || !rt.Bold || !rt.Italic {
		t.Errorf("RoundTrip attrs mismatch: fg=%q bold=%v italic=%v", rt.Fg, rt.Bold, rt.Italic)
	}

	linked, ok := cfg2.CustomHighlights["LinkedHL"]
	if !ok {
		t.Fatal("expected LinkedHL highlight after round-trip")
	}
	if linked.Link != "Normal" {
		t.Errorf("LinkedHL link mismatch: got %q", linked.Link)
	}
}
