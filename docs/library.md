# Plugin Library

Import path: `github.com/rmkohlman/MaestroNvim/nvimops/library`

The `library` package provides an embedded collection of 51 curated Neovim plugin definitions. No filesystem access is required — the plugins are compiled into the binary.

## Constructors

```go
func NewLibrary() (*Library, error)
func NewLibraryFromDir(dir string) (*Library, error)
```

`NewLibrary` loads from the embedded filesystem. `NewLibraryFromDir` loads YAML files from a directory on disk, useful for testing or custom plugin sets.

## Library Methods

```go
func (l *Library) Get(name string) (*plugin.Plugin, bool)
func (l *Library) List() []*plugin.Plugin
func (l *Library) ListByCategory(category string) []*plugin.Plugin
func (l *Library) ListByTag(tag string) []*plugin.Plugin
func (l *Library) Categories() []string
func (l *Library) Tags() []string
func (l *Library) Count() int
func (l *Library) Info() []PluginInfo
```

`List` returns plugins sorted by name. `Categories` and `Tags` return sorted, deduplicated slices.

### PluginInfo

```go
type PluginInfo struct {
    Name        string
    Description string
    Category    string
    Tags        []string
    Repo        string
}
```

`Info()` returns a lightweight summary for each plugin without the full `Plugin` struct.

## Embedded Plugins

The library contains 51 plugins across categories including colorscheme, editor, lsp, syntax, ui, formatting, linting, and language-specific tooling.

| # | Name | Category |
|---|------|----------|
| 01 | colorscheme | theme |
| 02 | telescope | editor |
| 03 | treesitter | syntax |
| 04 | treesitter-textobjects | syntax |
| 05 | nvim-cmp | lsp |
| 06 | mason | lsp |
| 07 | lspconfig | lsp |
| 08 | gitsigns | editor |
| 09 | lazygit | editor |
| 10 | which-key | ui |
| 11 | lualine | ui |
| 12 | autopairs | editor |
| 13 | comment | editor |
| 14 | surround | editor |
| 15 | alpha | ui |
| 16 | copilot | coding |
| 17 | dressing | ui |
| 18 | indent-blankline | ui |
| 19 | bufferline | ui |
| 20 | nvim-tree | editor |
| 21 | todo-comments | editor |
| 22 | trouble | editor |
| 23 | formatting | formatting |
| 24 | linting | linting |
| 25 | auto-session | editor |
| 26 | vim-maximizer | editor |
| 27 | substitute | editor |
| 28 | harpoon | editor |
| 29 | toggleterm | editor |
| 30 | copilot-cmp | coding |
| 31 | copilot-chat | coding |
| 32 | dadbod | utility |
| 33 | dadbod-ui | utility |
| 34 | dadbod-completion | utility |
| 35 | dbee | utility |
| 36 | render-markdown | utility |
| 37 | markdown-preview | utility |
| 38 | obsidian | utility |
| 39 | nvim-dap | editor |
| 40 | neotest | editor |
| 41 | nvim-dap-go | editor |
| 42 | neotest-go | editor |
| 43 | gopher-nvim | editor |
| 44 | nvim-dap-python | editor |
| 45 | neotest-python | editor |
| 46 | venv-selector | editor |
| 47 | rustaceanvim | editor |
| 48 | crates-nvim | editor |
| 49 | neotest-rust | editor |
| 50 | neotest-jest | editor |
| 51 | nvim-jdtls | editor |

## Usage Examples

**Look up a single plugin:**

```go
lib, err := library.NewLibrary()
if err != nil {
    log.Fatal(err)
}

p, ok := lib.Get("telescope")
if ok {
    fmt.Printf("%s: %s\n", p.Name, p.Repo)
}
```

**List all plugins in a category:**

```go
lspPlugins := lib.ListByCategory("lsp")
for _, p := range lspPlugins {
    fmt.Println(p.Name)
}
```

**Get a count and summary:**

```go
fmt.Printf("%d plugins available\n", lib.Count())
for _, info := range lib.Info() {
    fmt.Printf("%-20s %s\n", info.Name, info.Description)
}
```

**Use as a read-only store:**

```go
s := store.NewReadOnlyStore(lib)
p, err := s.Get("mason")
```

The `Library` type satisfies the `store.ReadOnlySource` interface.
