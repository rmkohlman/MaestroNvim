# Sync Framework

Import path: `github.com/rmkohlman/MaestroNvim/nvimops/sync`

The `sync` package provides an extensible framework for syncing Neovim plugin definitions from external sources. Sources are registered by name and invoked to produce plugin and package definitions that are persisted locally.

## SourceHandler Interface

```go
type SourceHandler interface {
    Name() string
    Description() string
    Sync(ctx context.Context, opts SyncOptions) (*SyncResult, error)
    ListAvailable(ctx context.Context) ([]AvailablePlugin, error)
    Validate(ctx context.Context) error
}
```

## AvailablePlugin

```go
type AvailablePlugin struct {
    Name         string
    Description  string
    Category     string
    Repo         string
    Config       string
    SourceName   string
    Labels       map[string]string
    Dependencies []string
}
```

## SyncOptions

```go
type SyncOptions struct {
    DryRun         bool
    Filters        map[string]string
    TargetDir      string
    Overwrite      bool
    PackageCreator PackageCreator
}
```

### SyncOptionsBuilder

```go
opts := sync.NewSyncOptions().
    DryRun(true).
    WithFilter("category", "lsp").
    WithTargetDir("/tmp/plugins").
    Overwrite(true).
    Build()
```

| Method | Description |
|--------|-------------|
| `DryRun(bool)` | Enable or disable dry-run mode |
| `WithFilter(key, value string)` | Add a single filter |
| `WithFilters(map[string]string)` | Set all filters |
| `WithTargetDir(string)` | Set the output directory |
| `Overwrite(bool)` | Allow overwriting existing plugins |
| `WithPackageCreator(PackageCreator)` | Set the package creator |
| `Build() SyncOptions` | Return the final options value |

### SyncOptions Methods

```go
func (o SyncOptions) HasFilter(key string) bool
func (o SyncOptions) GetFilter(key string) string
func (o SyncOptions) MatchesFilter(key, value string) bool
func (o SyncOptions) MatchesAvailablePlugin(p AvailablePlugin) bool
```

## SyncResult

```go
type SyncResult struct {
    PluginsCreated   []string
    PluginsUpdated   []string
    PackagesCreated  []string
    PackagesUpdated  []string
    Errors           []error
    SourceName       string
    TotalAvailable   int
    TotalSynced      int
}
```

### SyncResult Methods

```go
func (r *SyncResult) AddError(err error)
func (r *SyncResult) HasErrors() bool
func (r *SyncResult) AddPluginCreated(name string)
func (r *SyncResult) AddPluginUpdated(name string)
func (r *SyncResult) AddPackageCreated(name string)
func (r *SyncResult) AddPackageUpdated(name string)
func (r *SyncResult) Summary() string
```

## PackageCreator Interface

```go
type PackageCreator interface {
    CreatePackage(sourceName string, plugins []string) error
}
```

`FilePackageCreator` from `nvimops/package` implements this interface.

## Source Registry

The registry stores named `HandlerRegistration` entries and is used by the factory to resolve handlers by name.

```go
registry := sync.NewSourceRegistry()
```

### SourceRegistry Methods

```go
func (r *SourceRegistry) Register(reg HandlerRegistration) error
func (r *SourceRegistry) Unregister(name string) error
func (r *SourceRegistry) GetRegistration(name string) (HandlerRegistration, bool)
func (r *SourceRegistry) ListSources() []string
func (r *SourceRegistry) ListRegistrations() []HandlerRegistration
func (r *SourceRegistry) Clear()
func (r *SourceRegistry) Size() int
func (r *SourceRegistry) IsRegistered(name string) bool
func (r *SourceRegistry) GetSourceInfo(name string) (*SourceInfo, error)
func (r *SourceRegistry) ListSourcesByType(t SourceType) []string
func (r *SourceRegistry) SearchSources(query string) []HandlerRegistration
```

### HandlerRegistration

```go
type HandlerRegistration struct {
    Name       string
    Info       SourceInfo
    CreateFunc func() SourceHandler
}
```

### SourceInfo

```go
type SourceInfo struct {
    Name         string
    Description  string
    URL          string
    Type         string
    RequiresAuth bool
    ConfigKeys   []string
}
```

### SourceType Constants

```go
const (
    SourceTypeGitHub   SourceType = "github"
    SourceTypeLocal    SourceType = "local"
    SourceTypeRemote   SourceType = "remote"
    SourceTypeRegistry SourceType = "registry"
)
```

## Global Registry

```go
func GetGlobalRegistry() *SourceRegistry
func RegisterGlobalSource(reg HandlerRegistration) error
func InitializeGlobalRegistry() error
```

`InitializeGlobalRegistry` registers all builtin sources as `NotImplementedHandler` placeholders into the global registry.

## SourceHandlerFactory

```go
func NewSourceHandlerFactory() SourceHandlerFactory
func NewSourceHandlerFactoryWithRegistry(r *SourceRegistry) SourceHandlerFactory
```

```go
type SourceHandlerFactory interface {
    CreateHandler(source string) (SourceHandler, error)
    ListSources() []string
    IsSupported(source string) bool
    GetHandlerInfo(source string) (*SourceInfo, error)
}
```

## Builtin Sources

Six builtin sources are registered as `NotImplementedHandler` placeholders. The LazyVim source has a full implementation in the `nvimops/sync/sources` package.

| Source | Type | Implemented |
|--------|------|-------------|
| `lazyvim` | github | Yes (via `sources.RegisterLazyVimHandler`) |
| `astronvim` | github | No (placeholder) |
| `nvchad` | github | No (placeholder) |
| `kickstart` | github | No (placeholder) |
| `lunarvim` | github | No (placeholder) |
| `local` | local | No (placeholder) |

```go
var BuiltinSources []SourceInfo // exported slice of 6 builtin source info records

func RegisterBuiltinSources(r *SourceRegistry) error
```

## Error Types

```go
type ErrSourceNotFound struct{ Source string }
type ErrSourceAlreadyRegistered struct{ Source string }
type ErrSyncFailed struct {
    Source string
    Err    error
}
// ErrSyncFailed implements Unwrap() error
```

## Source Status

```go
type SourceStatus struct {
    Name          string
    IsImplemented bool
    IsRegistered  bool
    HandlerType   string
    SourceInfo    *SourceInfo
}

func GetSourceStatus(name string) (*SourceStatus, error)
func ListAllSourceStatus() ([]*SourceStatus, error)
```

## LazyVim Source Handler

Import path: `github.com/rmkohlman/MaestroNvim/nvimops/sync/sources`

`LazyVimHandler` is the only fully implemented source handler. It fetches plugin definitions from the LazyVim GitHub repository by reading Lua files in `lua/lazyvim/plugins/` and converting them to `NvimPlugin` YAML.

```go
func NewLazyVimHandler() sync.SourceHandler
func RegisterLazyVimHandler(registry *sync.SourceRegistry) error
func RegisterAllHandlers(registry *sync.SourceRegistry) error
func RegisterAllGlobalHandlers() error
```

`RegisterLazyVimHandler` unregisters the placeholder entry and registers the real handler. `RegisterAllHandlers` and `RegisterAllGlobalHandlers` are entry points to register all available real handlers at once.

### Category Mapping

LazyVim plugin files are mapped to categories by filename:

| Filename contains | Category |
|-------------------|----------|
| `coding` | `coding` |
| `colorscheme` | `theme` |
| `editor` | `editor` |
| `formatting` | `formatting` |
| `linting` | `linting` |
| `treesitter` | `syntax` |
| `ui` | `ui` |
| `util` | `utility` |
| `*lsp*` | `lsp` |
| (default) | `misc` |

### Usage

```go
registry := sync.NewSourceRegistry()
sync.RegisterBuiltinSources(registry)
sources.RegisterLazyVimHandler(registry)

factory := sync.NewSourceHandlerFactoryWithRegistry(registry)
handler, err := factory.CreateHandler("lazyvim")

result, err := handler.Sync(ctx, sync.NewSyncOptions().Build())
fmt.Println(result.Summary())
```
