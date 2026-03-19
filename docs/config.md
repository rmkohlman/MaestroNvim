# Core Config

Import path: `github.com/rmkohlman/MaestroNvim/nvimops/config`

The `config` package generates a complete, structured Neovim Lua configuration directory from a single YAML definition. It produces all necessary files for a namespace-scoped lazy.nvim setup.

## CoreConfig

```go
type CoreConfig struct {
    APIVersion  string
    Kind        string
    Namespace   string
    Leader      string
    Options     map[string]interface{}
    Globals     map[string]interface{}
    Keymaps     []Keymap
    Autocmds    []Autocmd
    BasePlugins []string
}
```

| Field | Default | Description |
|-------|---------|-------------|
| `Namespace` | `"workspace"` | Lua module namespace used for directory structure |
| `Leader` | `" "` (space) | Neovim leader key |
| `Options` | 16 defaults | `vim.opt` settings |
| `Globals` | map | `vim.g` global variables |
| `Keymaps` | 14 defaults | Key mappings registered at startup |
| `Autocmds` | 2 defaults | Autocommand groups |
| `BasePlugins` | 2 defaults | Plugins always included (`nvim-lua/plenary.nvim`, `christoomey/vim-tmux-navigator`) |

### Keymap (config package)

```go
type Keymap struct {
    Mode    string
    Key     string
    Action  string
    Desc    string
    Silent  bool
    Noremap *bool
}
```

Note: this `Keymap` type is distinct from `plugin.Keymap`.

### Autocmd

```go
type Autocmd struct {
    Group    string
    Events   []string
    Pattern  string
    Callback string
    Command  string
    Desc     string
}
```

## Defaults

```go
func DefaultCoreConfig() *CoreConfig
```

Returns a fully populated config with:
- 16 options (e.g. `number: true`, `relativenumber: true`, `expandtab: true`)
- 14 keymaps (e.g. `"i"/"kj"/<ESC>` through `"n"/"Y"/yy`)
- 2 autocmds (both in the `HighlightYank` group)
- 2 base plugins

## Parsing

```go
func ParseYAML(data []byte) (*CoreConfig, error)
func ParseYAMLFile(path string) (*CoreConfig, error)
```

If `Namespace` or `Leader` are empty in the parsed YAML, defaults (`"workspace"` and `" "`) are applied automatically.

## Serialization

```go
func (c *CoreConfig) ToYAML() ([]byte, error)
func (c *CoreConfig) WriteYAMLFile(path string) error
```

## Lua Generation

```go
gen := config.NewGenerator()
generated, err := gen.Generate(cfg)
```

### Generator

```go
type Generator struct {
    IndentSize int
    UseTabs    bool
}

func NewGenerator() *Generator
```

Default: `IndentSize = 2`, `UseTabs = true`.

### GeneratedConfig

`Generate` returns a `GeneratedConfig` containing all seven Lua files as strings:

```go
type GeneratedConfig struct {
    InitLua       string
    LazyLua       string
    CoreInitLua   string
    OptionsLua    string
    KeymapsLua    string
    AutocmdsLua   string
    PluginsInitLua string
}
```

## Writing to Disk

```go
err := gen.WriteToDirectory(cfg, plugins, "/home/dev/.config/nvim")
```

`WriteToDirectory` generates and writes the full directory structure:

| File | Contents |
|------|----------|
| `init.lua` | Entry point, loads lazy and namespace |
| `lua/{ns}/lazy.lua` | lazy.nvim bootstrap |
| `lua/{ns}/core/init.lua` | Core module loader |
| `lua/{ns}/core/options.lua` | `vim.opt` and `vim.g` settings |
| `lua/{ns}/core/keymaps.lua` | `vim.keymap.set` calls |
| `lua/{ns}/core/autocmds.lua` | `nvim_create_augroup` / `nvim_create_autocmd` blocks |
| `lua/{ns}/plugins/init.lua` | Plugin list loader |
| `lua/{ns}/plugins/{name}.lua` | One file per enabled plugin |

`{ns}` is the value of `CoreConfig.Namespace`.

## Options Format

Options with an `"append:value"` format in the YAML are emitted as `opt.key:append("value")` in Lua rather than `opt.key = "value"`. This is useful for `opt.rtp`, `opt.packpath`, and similar list-type options.

## Autocmds

Autocmds are grouped by `Group` name. Each group becomes an `nvim_create_augroup` call followed by one `nvim_create_autocmd` call per autocmd in that group.

```lua
local highlight_group = vim.api.nvim_create_augroup("HighlightYank", { clear = true })
vim.api.nvim_create_autocmd("TextYankPost", {
  group = highlight_group,
  pattern = "*",
  callback = function() vim.highlight.on_yank() end,
})
```

## Example YAML

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimCoreConfig
namespace: myapp
leader: " "
options:
  number: true
  relativenumber: true
  tabstop: 2
  shiftwidth: 2
  expandtab: true
keymaps:
  - mode: "i"
    key: "kj"
    action: "<ESC>"
    desc: "Exit insert mode"
    silent: true
basePlugins:
  - nvim-lua/plenary.nvim
```
