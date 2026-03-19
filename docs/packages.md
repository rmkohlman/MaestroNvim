# Packages

Import paths:
- `github.com/rmkohlman/MaestroNvim/nvimops/package` — package types, parsing, creator
- `github.com/rmkohlman/MaestroNvim/nvimops/package/library` — embedded library of 12 curated packages

A Package groups a set of plugins under a single name and supports single-inheritance via `Extends`. This allows language-specific environments to be composed on top of a common base.

## Package Type

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

`Extends` names another package this one inherits from. A package cannot extend itself.

## Constructors

```go
func NewPackage(name string) *Package
func NewPackageYAML(name string) *PackageYAML
```

## YAML Format

Package YAML files use the `devopsmaestro.io/v1` API version and `NvimPackage` kind.

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: maestro-go
  description: Go development environment
  category: language
  tags:
    - go
    - development
spec:
  extends: maestro
  plugins:
    - nvim-dap-go
    - neotest-go
    - gopher-nvim
  enabled: true
```

### Metadata Fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Package identifier (required) |
| `description` | string | Human-readable description |
| `category` | string | Package category |
| `tags` | []string | Searchable tags |
| `labels` | map[string]string | Key-value labels |
| `annotations` | map[string]string | Key-value annotations |

### Spec Fields

| Field | Type | Description |
|-------|------|-------------|
| `extends` | string | Parent package to inherit from |
| `plugins` | string or []string | Plugin names included in this package |
| `enabled` | bool | Whether the package is active |

## Parsing

```go
func ParseYAMLFile(path string) (*Package, error)
func ParseYAML(data []byte) (*Package, error)
func ParseYAMLMultiple(data []byte) ([]*Package, error)
```

Validation rules:
- `apiVersion` must be empty or `"devopsmaestro.io/v1"`
- `kind` must be empty or `"NvimPackage"`
- `metadata.name` is required
- Plugin names must not be empty strings
- A package cannot extend itself

## Conversion

```go
func (y *PackageYAML) ToPackage() *Package
func (p *Package) ToYAML() *PackageYAML
func (p *Package) ToYAMLBytes() ([]byte, error)
```

## FilePackageCreator

`FilePackageCreator` creates package YAML files on disk. It is the primary implementation of the `sync.PackageCreator` interface and is used by the sync framework to persist packages received from source handlers.

```go
type FilePackageCreator struct {
    PackagesDir string
}

func NewFilePackageCreator(packagesDir string) *FilePackageCreator
```

```go
err := creator.CreatePackage("lazyvim", []string{"telescope", "treesitter", "lspconfig"})
```

The created YAML file is written to `<PackagesDir>/<sourceName>.yaml` and includes the following auto-generated labels:

| Label | Value |
|-------|-------|
| `source` | The `sourceName` argument |
| `auto-generated` | `"true"` |
| `sync-time` | RFC3339 timestamp of creation |

## Package Library

Import path: `github.com/rmkohlman/MaestroNvim/nvimops/package/library`

The package library contains 12 embedded packages for common development environments.

```go
func NewLibrary() (*Library, error)
func NewLibraryFromDir(dir string) (*Library, error)
```

### Library Methods

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

### PackageInfo

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

### Embedded Packages

| Name | Description |
|------|-------------|
| `core` | Core plugins for all environments |
| `maestro` | Full Maestro base environment |
| `full` | All available plugins |
| `maestro-go` | Go development environment |
| `maestro-python` | Python development environment |
| `maestro-rust` | Rust development environment |
| `maestro-node` | Node.js development environment |
| `maestro-java` | Java development environment |
| `maestro-dotnet` | .NET development environment |
| `maestro-gleam` | Gleam development environment |
| `go-dev` | Lightweight Go development |
| `python-dev` | Lightweight Python development |

## Usage Examples

**Load and inspect the library:**

```go
lib, err := pkglibrary.NewLibrary()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%d packages\n", lib.Count())

pkg, ok := lib.Get("maestro-go")
if ok {
    fmt.Printf("%s extends %s, %d plugins\n", pkg.Name, pkg.Extends, len(pkg.Plugins))
}
```

**List language-specific packages:**

```go
for _, info := range lib.Info() {
    if info.Extends != "" {
        fmt.Printf("%s (extends %s): %d plugins\n", info.Name, info.Extends, info.PluginCount)
    }
}
```
