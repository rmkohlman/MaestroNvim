# API Reference

Complete reference for all exported types, functions, and interfaces in MaestroNvim.

---

## Package `nvim`

Import: `github.com/rmkohlman/MaestroNvim/nvim`

### Interfaces

#### Manager

```go
type Manager interface {
    Init(InitOptions) error
    Sync(string, SyncDirection) error
    Push(string) error
    Status() (*Status, error)
    ListWorkspaces() ([]Workspace, error)
}
```

### Constructors

```go
func NewManager() Manager
func NewManagerWithPath(configPath string) Manager
func NewMockManager() *MockManager
```

### Types

#### InitOptions

```go
type InitOptions struct {
    ConfigPath string
    Template   string
    Overwrite  bool
    GitClone   bool
    GitURL     string
    Subdir     string
}
```

#### SyncDirection

```go
type SyncDirection int

const (
    SyncPull          SyncDirection = 0
    SyncPush          SyncDirection = 1
    SyncBidirectional SyncDirection = 2
)

func (d SyncDirection) String() string
```

#### Status

```go
type Status struct {
    ConfigPath    string
    Exists        bool
    LastSync      time.Time
    SyncedWith    string
    LocalChanges  bool
    RemoteChanges bool
    Template      string
}
```

#### Workspace

```go
type Workspace struct {
    ID       string
    Name     string
    Active   bool
    NvimPath string
}
```

#### GitURLInfo

```go
type GitURLInfo struct {
    FullURL  string
    Platform string
    RepoName string
}
```

#### MockManager

```go
type MockManager struct {
    Initialized           bool
    StatusData            *Status
    Workspaces            []Workspace
    Calls                 []MockManagerCall
    InitError             error
    SyncError             error
    PushError             error
    StatusError           error
    ListWorkspacesError   error
}
```

Methods: `Reset()`, `CallCount(string) int`, `GetCalls(string) []MockManagerCall`, `LastCall() *MockManagerCall`, `SetStatus(*Status)`, `AddWorkspace(Workspace)`, `SetWorkspaces([]Workspace)`, `SimulateLocalChanges()`, `SimulateRemoteChanges()`, `SetInitialized(bool)`, `InjectError(string, error)`

#### MockManagerCall

```go
type MockManagerCall struct {
    Method string
    Args   []interface{}
}
```

### Functions

```go
func IsGitURL(s string) bool
func NormalizeGitURL(s string) string
func ParseGitURL(s string) GitURLInfo
```

---

## Package `nvimops`

Import: `github.com/rmkohlman/MaestroNvim/nvimops`

### Interfaces

#### Manager

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

### Constructors

```go
func New() (Manager, error)
func NewWithOptions(opts Options) (Manager, error)
func NewMockManager() *MockManager
```

### Types

#### Options

```go
type Options struct {
    Store     store.PluginStore
    StoreDir  string
    Generator plugin.LuaGenerator
}
```

### Functions

```go
func FetchURL(url string) ([]byte, string, error)
```

---

## Package `nvimops/plugin`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/plugin`

### Interfaces

#### LuaGenerator

```go
type LuaGenerator interface {
    GenerateLua(*Plugin) (string, error)
    GenerateLuaFile(*Plugin) (string, error)
}
```

### Types

#### Plugin

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

#### Dependency

```go
type Dependency struct {
    Repo    string
    Build   string
    Version string
    Branch  string
    Config  bool
}
```

#### Keymap

```go
type Keymap struct {
    Key    string
    Mode   []string
    Action string
    Desc   string
}
```

#### Generator

```go
type Generator struct {
    IndentSize int
}

func NewGenerator() *Generator
func (g *Generator) GenerateLua(*Plugin) (string, error)
func (g *Generator) GenerateLuaFile(*Plugin) (string, error)
```

#### PluginManifest

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

#### StringOrSlice

```go
type StringOrSlice []string
```

Marshals as a scalar string when it contains exactly one element.

### Constructors

```go
func NewPlugin(name, repo string) *Plugin
func NewPluginYAML(name, repo string) *PluginYAML
```

### Functions

```go
func ParseYAML(data []byte) (*Plugin, error)
func ParseYAMLFile(path string) (*Plugin, error)
func ParseYAMLMultiple(data []byte) ([]*Plugin, error)
func ResolveManifest(plugins []*Plugin) *PluginManifest
func ResolveManifestFromNames(names []string) *PluginManifest
```

### Methods

```go
func (y *PluginYAML) ToPlugin() *Plugin
func (p *Plugin) ToYAML() *PluginYAML
func (p *Plugin) ToYAMLBytes() ([]byte, error)
```

---

## Package `nvimops/store`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/store`

### Interfaces

#### PluginStore

```go
type PluginStore interface {
    Create(*plugin.Plugin) error
    Update(*plugin.Plugin) error
    Upsert(*plugin.Plugin) error
    Delete(name string) error
    Get(name string) (*plugin.Plugin, error)
    List() ([]*plugin.Plugin, error)
    ListByCategory(category string) ([]*plugin.Plugin, error)
    ListByTag(tag string) ([]*plugin.Plugin, error)
    Exists(name string) (bool, error)
    Close() error
}
```

#### ReadOnlySource

```go
type ReadOnlySource interface {
    Get(name string) (*plugin.Plugin, bool)
    List() []*plugin.Plugin
    ListByCategory(category string) []*plugin.Plugin
    ListByTag(tag string) []*plugin.Plugin
}
```

### Constructors

```go
func NewFileStore(baseDir string) (*FileStore, error)
func DefaultFileStore() (*FileStore, error)
func NewMemoryStore() *MemoryStore
func NewReadOnlyStore(src ReadOnlySource) *ReadOnlyStore
```

### Extra FileStore Methods

```go
func (s *FileStore) Reload() error
func (s *FileStore) BaseDir() string
```

### Error Types

```go
type ErrNotFound struct{ Name string }
type ErrAlreadyExists struct{ Name string }
type ErrReadOnly struct{ Operation string }

func IsNotFound(err error) bool
func IsAlreadyExists(err error) bool
func IsReadOnly(err error) bool
```

---

## Package `nvimops/config`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/config`

### Types

#### CoreConfig

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

#### Keymap

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

#### Autocmd

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

#### Generator

```go
type Generator struct {
    IndentSize int
    UseTabs    bool
}

func NewGenerator() *Generator
func (g *Generator) Generate(cfg *CoreConfig) (*GeneratedConfig, error)
func (g *Generator) WriteToDirectory(cfg *CoreConfig, plugins []*plugin.Plugin, dir string) error
```

#### GeneratedConfig

```go
type GeneratedConfig struct {
    InitLua        string
    LazyLua        string
    CoreInitLua    string
    OptionsLua     string
    KeymapsLua     string
    AutocmdsLua    string
    PluginsInitLua string
}
```

### Functions

```go
func DefaultCoreConfig() *CoreConfig
func ParseYAML(data []byte) (*CoreConfig, error)
func ParseYAMLFile(path string) (*CoreConfig, error)
```

### Methods

```go
func (c *CoreConfig) ToYAML() ([]byte, error)
func (c *CoreConfig) WriteYAMLFile(path string) error
```

---

## Package `nvimops/library`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/library`

### Types

#### Library

```go
type Library struct { /* embedded FS */ }
```

#### PluginInfo

```go
type PluginInfo struct {
    Name        string
    Description string
    Category    string
    Tags        []string
    Repo        string
}
```

### Constructors

```go
func NewLibrary() (*Library, error)
func NewLibraryFromDir(dir string) (*Library, error)
```

### Methods

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

`Library` satisfies `store.ReadOnlySource`.

---

## Package `nvimops/package`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/package`

Go package name: `pkg`

### Types

#### Package

```go
type Package struct {
    Name        string
    Description string
    Category    string
    Tags        []string
    Extends     string
    Plugins     []string
    Enabled     bool
    CreatedAt   *time.Time
    UpdatedAt   *time.Time
}
```

#### FilePackageCreator

```go
type FilePackageCreator struct {
    PackagesDir string
}

func NewFilePackageCreator(packagesDir string) *FilePackageCreator
func (c *FilePackageCreator) CreatePackage(sourceName string, plugins []string) error
```

### Constructors

```go
func NewPackage(name string) *Package
func NewPackageYAML(name string) *PackageYAML
```

### Functions

```go
func ParseYAML(data []byte) (*Package, error)
func ParseYAMLFile(path string) (*Package, error)
func ParseYAMLMultiple(data []byte) ([]*Package, error)
```

### Methods

```go
func (y *PackageYAML) ToPackage() *Package
func (p *Package) ToYAML() *PackageYAML
func (p *Package) ToYAMLBytes() ([]byte, error)
```

---

## Package `nvimops/package/library`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/package/library`

Go package name: `library`

### Types

#### Library

```go
type Library struct { /* embedded FS */ }
```

#### PackageInfo

```go
type PackageInfo struct {
    Name        string
    Description string
    Category    string
    Tags        []string
    Extends     string
    PluginCount int
}
```

### Constructors

```go
func NewLibrary() (*Library, error)
func NewLibraryFromDir(dir string) (*Library, error)
```

### Methods

```go
func (l *Library) Get(name string) (*pkg.Package, bool)
func (l *Library) List() []*pkg.Package
func (l *Library) ListByCategory(category string) []*pkg.Package
func (l *Library) ListByTag(tag string) []*pkg.Package
func (l *Library) Categories() []string
func (l *Library) Tags() []string
func (l *Library) Count() int
func (l *Library) Has(name string) bool
func (l *Library) Info() []PackageInfo
```

---

## Package `nvimops/sync`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/sync`

### Interfaces

#### SourceHandler

```go
type SourceHandler interface {
    Name() string
    Description() string
    Sync(ctx context.Context, opts SyncOptions) (*SyncResult, error)
    ListAvailable(ctx context.Context) ([]AvailablePlugin, error)
    Validate(ctx context.Context) error
}
```

#### SourceHandlerFactory

```go
type SourceHandlerFactory interface {
    CreateHandler(source string) (SourceHandler, error)
    ListSources() []string
    IsSupported(source string) bool
    GetHandlerInfo(source string) (*SourceInfo, error)
}
```

#### PackageCreator

```go
type PackageCreator interface {
    CreatePackage(sourceName string, plugins []string) error
}
```

### Types

#### SourceType

```go
type SourceType string

const (
    SourceTypeGitHub   SourceType = "github"
    SourceTypeLocal    SourceType = "local"
    SourceTypeRemote   SourceType = "remote"
    SourceTypeRegistry SourceType = "registry"
)
```

#### SourceInfo

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

#### HandlerRegistration

```go
type HandlerRegistration struct {
    Name       string
    Info       SourceInfo
    CreateFunc func() SourceHandler
}
```

#### SyncOptions

```go
type SyncOptions struct {
    DryRun         bool
    Filters        map[string]string
    TargetDir      string
    Overwrite      bool
    PackageCreator PackageCreator
}

func (o SyncOptions) HasFilter(key string) bool
func (o SyncOptions) GetFilter(key string) string
func (o SyncOptions) MatchesFilter(key, value string) bool
func (o SyncOptions) MatchesAvailablePlugin(p AvailablePlugin) bool
```

#### SyncOptionsBuilder

```go
func NewSyncOptions() *SyncOptionsBuilder
func (b *SyncOptionsBuilder) DryRun(v bool) *SyncOptionsBuilder
func (b *SyncOptionsBuilder) WithFilter(key, value string) *SyncOptionsBuilder
func (b *SyncOptionsBuilder) WithFilters(f map[string]string) *SyncOptionsBuilder
func (b *SyncOptionsBuilder) WithTargetDir(dir string) *SyncOptionsBuilder
func (b *SyncOptionsBuilder) Overwrite(v bool) *SyncOptionsBuilder
func (b *SyncOptionsBuilder) WithPackageCreator(c PackageCreator) *SyncOptionsBuilder
func (b *SyncOptionsBuilder) Build() SyncOptions
```

#### SyncResult

```go
type SyncResult struct {
    PluginsCreated  []string
    PluginsUpdated  []string
    PackagesCreated []string
    PackagesUpdated []string
    Errors          []error
    SourceName      string
    TotalAvailable  int
    TotalSynced     int
}

func (r *SyncResult) AddError(err error)
func (r *SyncResult) HasErrors() bool
func (r *SyncResult) AddPluginCreated(name string)
func (r *SyncResult) AddPluginUpdated(name string)
func (r *SyncResult) AddPackageCreated(name string)
func (r *SyncResult) AddPackageUpdated(name string)
func (r *SyncResult) Summary() string
```

#### AvailablePlugin

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

#### SourceStatus

```go
type SourceStatus struct {
    Name          string
    IsImplemented bool
    IsRegistered  bool
    HandlerType   string
    SourceInfo    *SourceInfo
}
```

### Constructors

```go
func NewSourceRegistry() *SourceRegistry
func NewSourceHandlerFactory() SourceHandlerFactory
func NewSourceHandlerFactoryWithRegistry(r *SourceRegistry) SourceHandlerFactory
```

### Functions

```go
func RegisterBuiltinSources(r *SourceRegistry) error
func GetGlobalRegistry() *SourceRegistry
func RegisterGlobalSource(reg HandlerRegistration) error
func InitializeGlobalRegistry() error
func GetSourceStatus(name string) (*SourceStatus, error)
func ListAllSourceStatus() ([]*SourceStatus, error)

var BuiltinSources []SourceInfo
```

### Error Types

```go
type ErrSourceNotFound struct{ Source string }
type ErrSourceAlreadyRegistered struct{ Source string }
type ErrSyncFailed struct {
    Source string
    Err    error
}
func (e ErrSyncFailed) Unwrap() error
```

---

## Package `nvimops/sync/sources`

Import: `github.com/rmkohlman/MaestroNvim/nvimops/sync/sources`

### Functions

```go
func NewLazyVimHandler() sync.SourceHandler
func RegisterLazyVimHandler(registry *sync.SourceRegistry) error
func RegisterAllHandlers(registry *sync.SourceRegistry) error
func RegisterAllGlobalHandlers() error
```

### Types

#### GitHubContent

```go
type GitHubContent struct {
    Name        string
    Path        string
    Type        string
    SHA         string
    URL         string
    GitURL      string
    HTMLURL     string
    DownloadURL string
    Size        int
}
```

#### GitHubRelease

```go
type GitHubRelease struct {
    TagName string
    Name    string
    Draft   bool
    Created time.Time
}
```
