# Plugin Store

Import path: `github.com/rmkohlman/MaestroNvim/nvimops/store`

The `store` package provides three implementations of `PluginStore` for persisting and retrieving plugin definitions.

## PluginStore Interface

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

## FileStore

`FileStore` persists each plugin as a `<name>.yaml` file in a directory on disk. Plugins are lazy-loaded on first access.

```go
func NewFileStore(baseDir string) (*FileStore, error)
func DefaultFileStore() (*FileStore, error)
```

`DefaultFileStore` uses `~/.nvim-manager/plugins`. Tilde expansion is supported in `baseDir`.

Timestamps (`CreatedAt`, `UpdatedAt`) are set automatically on write.

### Extra Methods

```go
func (s *FileStore) Reload() error
func (s *FileStore) BaseDir() string
```

`Reload` clears the in-memory cache and re-reads all YAML files from disk. `BaseDir` returns the resolved base directory path.

## MemoryStore

`MemoryStore` is a thread-safe in-memory store. It returns copies of plugins on `Get` and `List` to prevent external mutation.

```go
func NewMemoryStore() *MemoryStore
```

Useful for testing or for building transient in-process plugin sets.

## ReadOnlyStore

`ReadOnlyStore` wraps a `ReadOnlySource` and allows reads but returns `ErrReadOnly` for all write operations.

```go
func NewReadOnlyStore(src ReadOnlySource) *ReadOnlyStore
```

### ReadOnlySource Interface

```go
type ReadOnlySource interface {
    Get(name string) (*plugin.Plugin, bool)
    List() []*plugin.Plugin
    ListByCategory(category string) []*plugin.Plugin
    ListByTag(tag string) []*plugin.Plugin
}
```

The embedded `library.Library` type satisfies `ReadOnlySource`, so the plugin library can be wrapped as a read-only store:

```go
lib, _ := library.NewLibrary()
store := store.NewReadOnlyStore(lib)
p, err := store.Get("telescope")
```

## Error Types

### ErrNotFound

```go
type ErrNotFound struct {
    Name string
}
// Error() returns "plugin not found: <name>"

func IsNotFound(err error) bool
```

### ErrAlreadyExists

```go
type ErrAlreadyExists struct {
    Name string
}
// Error() returns "plugin already exists: <name>"

func IsAlreadyExists(err error) bool
```

### ErrReadOnly

```go
type ErrReadOnly struct {
    Operation string
}
// Error() returns "operation not permitted on read-only store: <operation>"

func IsReadOnly(err error) bool
```

## Usage Examples

**File-backed store:**

```go
s, err := store.NewFileStore("~/.config/my-app/plugins")
if err != nil {
    log.Fatal(err)
}
defer s.Close()

err = s.Upsert(p)
plugins, err := s.List()
```

**Check existence before creating:**

```go
exists, err := s.Exists("telescope")
if !exists {
    err = s.Create(p)
}
```

**Handle not-found gracefully:**

```go
p, err := s.Get("unknown-plugin")
if store.IsNotFound(err) {
    fmt.Println("plugin not in store")
}
```
