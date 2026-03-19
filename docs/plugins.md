# Plugins

Import paths:
- `github.com/rmkohlman/MaestroNvim/nvimops/plugin` — plugin types, YAML parsing, Lua generation
- `github.com/rmkohlman/MaestroNvim/nvimops` — high-level plugin manager

## Plugin Type

```go
type Plugin struct {
    Name         string
    Description  string
    Repo         string
    Branch       string
    Version      string
    Priority     int
    Lazy         bool
    Event        []string
    Ft           []string
    Cmd          []string
    Keys         []Keymap
    Dependencies []Dependency
    Build        string
    Config       string
    Init         string
    Opts         interface{}
    Keymaps      []Keymap
    Category     string
    Tags         []string
    Enabled      bool
    CreatedAt    *time.Time
    UpdatedAt    *time.Time
}
```

### Dependency

```go
type Dependency struct {
    Repo    string
    Build   string
    Version string
    Branch  string
    Config  bool
}
```

### Keymap

```go
type Keymap struct {
    Key    string
    Mode   []string
    Action string
    Desc   string
}
```

## Constructors

```go
func NewPlugin(name, repo string) *Plugin
func NewPluginYAML(name, repo string) *PluginYAML
```

## YAML Format

Plugin YAML files use the `devopsmaestro.io/v1` API version and `NvimPlugin` kind.

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  description: Fuzzy finder over lists
  category: editor
  tags:
    - search
    - files
spec:
  repo: nvim-telescope/telescope.nvim
  lazy: true
  event:
    - VimEnter
  dependencies:
    - nvim-lua/plenary.nvim
    - repo: nvim-tree/nvim-web-devicons
      config: false
  config: |
    require("telescope").setup({})
  keymaps:
    - key: "<leader>ff"
      mode: [n]
      action: "<cmd>Telescope find_files<cr>"
      desc: Find files
```

### Metadata Fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Plugin identifier (required) |
| `description` | string | Human-readable description |
| `category` | string | Plugin category |
| `tags` | []string | Searchable tags |
| `labels` | map[string]string | Key-value labels |
| `annotations` | map[string]string | Key-value annotations |

### Spec Fields

| Field | Type | Description |
|-------|------|-------------|
| `repo` | string | GitHub repo path, e.g. `user/repo` (required) |
| `branch` | string | Git branch to pin to |
| `version` | string | Git tag or version to pin to |
| `build` | string | Shell command to run after install |
| `config` | string | Lua config block |
| `init` | string | Lua init block (runs before config) |
| `priority` | int | Load priority (higher loads first) |
| `lazy` | bool | Defer loading until triggered |
| `enabled` | bool | Whether the plugin is active |
| `event` | string or []string | Events that trigger loading |
| `ft` | string or []string | Filetypes that trigger loading |
| `cmd` | string or []string | Commands that trigger loading |
| `keys` | []KeymapYAML | Keys that trigger loading |
| `dependencies` | []DependencyYAML | Required plugins |
| `opts` | any | Options table passed to `setup()` |
| `keymaps` | []KeymapYAML | Keymaps to register via `vim.keymap.set` |

### Dependency YAML

Dependencies can be written as a plain string (repo path only) or a mapping:

```yaml
dependencies:
  - nvim-lua/plenary.nvim
  - repo: nvim-tree/nvim-web-devicons
    config: false
    version: v1.0.0
```

## Parsing

```go
func ParseYAMLFile(path string) (*Plugin, error)
func ParseYAML(data []byte) (*Plugin, error)
func ParseYAMLMultiple(data []byte) ([]*Plugin, error)
```

`ParseYAMLMultiple` handles YAML files with multiple documents separated by `---`.

Validation rules:
- `kind` must be empty or `"NvimPlugin"`
- `metadata.name` is required
- `spec.repo` is required

## Conversion

```go
func (y *PluginYAML) ToPlugin() *Plugin
func (p *Plugin) ToYAML() *PluginYAML
func (p *Plugin) ToYAMLBytes() ([]byte, error)
```

## Lua Generation

The `Generator` converts a `Plugin` to lazy.nvim-compatible Lua.

```go
gen := plugin.NewGenerator()        // IndentSize defaults to 2
lua, err := gen.GenerateLua(p)      // returns Lua string
path, err := gen.GenerateLuaFile(p) // writes file, returns path
```

### Generator

```go
type Generator struct {
    IndentSize int
}

func NewGenerator() *Generator
```

### LuaGenerator Interface

```go
type LuaGenerator interface {
    GenerateLua(*Plugin) (string, error)
    GenerateLuaFile(*Plugin) (string, error)
}
```

### Lua Output Example

Given a plugin with `Repo: "nvim-telescope/telescope.nvim"`, `Lazy: true`, and a config string, the generator produces:

```lua
return {
  "nvim-telescope/telescope.nvim",
  lazy = true,
  config = function()
    require("telescope").setup({})
  end,
}
```

**Lua expression detection:** If `Config` or `Init` begins with `require(`, `require'`, `require "`, or `function`, it is emitted as a raw Lua expression rather than a string. The string `"true"` in `Config` with no keymaps emits the Lua boolean `true`.

**Keymaps:** Each `Keymap` entry generates a `vim.keymap.set()` call inside the config function.

## Plugin Manifest

```go
type PluginManifest struct {
    InstalledPlugins []string
    Features         PluginFeatures
}

type PluginFeatures struct {
    HasMason      bool
    HasTreesitter bool
    HasTelescope  bool
    HasLSPConfig  bool
}
```

```go
func ResolveManifest(plugins []*Plugin) *PluginManifest
func ResolveManifestFromNames(names []string) *PluginManifest
```

`ResolveManifest` inspects the plugin list and sets feature flags based on well-known repos. `ResolveManifestFromNames` does the same from a list of repo name strings.

## High-Level Manager (nvimops)

The `nvimops` package provides a `Manager` that combines plugin storage, URL fetching, and Lua generation into a single interface.

```go
mgr, err := nvimops.New()

// or with options
mgr, err := nvimops.NewWithOptions(nvimops.Options{
    StoreDir:  "~/.nvim-manager/plugins",
    Generator: plugin.NewGenerator(),
})
```

### Manager Interface

```go
type Manager interface {
    ApplyFile(path string) error
    ApplyURL(url string) error
    Apply(*plugin.Plugin) error
    Get(name string) (*plugin.Plugin, error)
    List() ([]*plugin.Plugin, error)
    Delete(name string) error
    GenerateLua(name string) error
    GenerateLuaFor(name string) (string, error)
    Store() store.PluginStore
    Generator() plugin.LuaGenerator
    Close() error
}
```

### Options

```go
type Options struct {
    Store     store.PluginStore
    StoreDir  string
    Generator plugin.LuaGenerator
}
```

### FetchURL

```go
func FetchURL(url string) ([]byte, string, error)
```

Downloads a plugin YAML from a URL. The `github:user/repo/path/file.yaml` shorthand is expanded to:

```
https://raw.githubusercontent.com/{user}/{repo}/main/{path}
```

The function uses a 30-second HTTP timeout and returns the raw bytes, the resolved URL, and any error.

## Testing

### MockManager (nvimops)

```go
mock := nvimops.NewMockManager()
mock.GetResult = &plugin.Plugin{Name: "telescope"}
mock.ListResult = []*plugin.Plugin{...}
mock.InjectApplyError(errors.New("write failed"))
```

Available fields: `GetResult`, `ListResult`, `GenerateLuaForResult`, `MockStore`, `MockGenerator`, per-method error fields (`ApplyFileError`, `ApplyURLError`, `ApplyError`, `GetError`, `ListError`, `DeleteError`, `GenerateLuaError`, `GenerateLuaForError`, `CloseError`), `Calls []MockManagerCall`, and per-method argument slices.

| Helper | Description |
|--------|-------------|
| `Reset()` | Clear all call history |
| `CallCount(method string) int` | Count calls to a method |
| `GetCalls(method string) []MockManagerCall` | All calls to a method |
| `LastCall() *MockManagerCall` | Most recent call |
