# Nvim Manager

Import path: `github.com/rmkohlman/MaestroNvim/nvim`

The `nvim` package manages the lifecycle of a Neovim configuration directory. It supports initializing configs from curated templates, tracking sync status, and enumerating workspaces.

## Manager Interface

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
```

`NewManager` targets `~/.config/nvim` and reads sync status from the standard MaestroSDK paths location. `NewManagerWithPath` accepts any absolute config path.

## Initialization

```go
err := mgr.Init(nvim.InitOptions{
    Template:  "lazyvim",
    Overwrite: false,
})
```

### InitOptions

| Field | Type | Description |
|-------|------|-------------|
| `ConfigPath` | `string` | Override the config directory path |
| `Template` | `string` | Template name to initialize from |
| `Overwrite` | `bool` | Overwrite an existing config if present |
| `GitClone` | `bool` | Clone a git repository instead of writing files |
| `GitURL` | `string` | Git URL to clone (required when `Template` is `"custom"`) |
| `Subdir` | `string` | Subdirectory within the cloned repo to use |

### Supported Templates

| Template | Source |
|----------|--------|
| `kickstart` | `https://github.com/nvim-lua/kickstart.nvim.git` |
| `lazyvim` | `https://github.com/LazyVim/starter.git` |
| `astronvim` | `https://github.com/AstroNvim/template.git` |
| `minimal` | Embedded `init.lua` written directly |
| `custom` | Requires `GitURL` field to be set |

## Sync and Push

```go
err := mgr.Sync(remote, nvim.SyncPull)
err := mgr.Push(remote)
```

!!! note
    `Sync` and `Push` are stub implementations in the current release. Both return a `"sync not yet implemented"` error. They are reserved for a future release that will add remote sync capability.

### SyncDirection

```go
type SyncDirection int

const (
    SyncPull          SyncDirection = 0
    SyncPush          SyncDirection = 1
    SyncBidirectional SyncDirection = 2
)
```

`SyncDirection.String()` returns `"pull"`, `"push"`, `"bidirectional"`, or `"unknown"`.

## Status

```go
status, err := mgr.Status()
```

### Status Struct

| Field | Type | Description |
|-------|------|-------------|
| `ConfigPath` | `string` | Absolute path to the Neovim config directory |
| `Exists` | `bool` | Whether the config directory exists |
| `LastSync` | `time.Time` | Timestamp of the last sync operation |
| `SyncedWith` | `string` | Remote location last synced with |
| `LocalChanges` | `bool` | Uncommitted local changes detected |
| `RemoteChanges` | `bool` | Remote changes not yet pulled |
| `Template` | `string` | Template used during initialization |

Status is persisted as JSON to the MaestroSDK `NvimSyncStatus` path.

## ListWorkspaces

```go
workspaces, err := mgr.ListWorkspaces()
```

Returns a slice of `Workspace` values. In the current release this always returns an empty slice — it is reserved for future integration with the DevOpsMaestro workspace database.

### Workspace Struct

| Field | Type | Description |
|-------|------|-------------|
| `ID` | `string` | Workspace identifier |
| `Name` | `string` | Display name |
| `Active` | `bool` | Whether this is the currently active workspace |
| `NvimPath` | `string` | Path to the workspace-specific Neovim config |

## URL Utilities

The `nvim` package includes helpers for working with git URLs in template initialization.

```go
func IsGitURL(s string) bool
func NormalizeGitURL(s string) string
func ParseGitURL(s string) GitURLInfo
```

### GitURLInfo

| Field | Type | Description |
|-------|------|-------------|
| `FullURL` | `string` | Normalized full HTTPS URL |
| `Platform` | `string` | `"github"`, `"gitlab"`, `"bitbucket"`, or `"git"` |
| `RepoName` | `string` | Repository name extracted from the URL |

Shorthand URLs are supported:

```
github:user/repo   -> https://github.com/user/repo.git
gitlab:user/repo   -> https://gitlab.com/user/repo.git
bitbucket:user/repo -> https://bitbucket.org/user/repo.git
```

## Testing

A `MockManager` is provided for use in tests.

```go
mock := nvim.NewMockManager()
mock.SetStatus(&nvim.Status{Exists: true, Template: "lazyvim"})
mock.SimulateLocalChanges()
mock.InjectError("Sync", errors.New("network unavailable"))
```

### MockManager Fields

| Field | Type |
|-------|------|
| `Initialized` | `bool` |
| `StatusData` | `*Status` |
| `Workspaces` | `[]Workspace` |
| `Calls` | `[]MockManagerCall` |
| `InitError` | `error` |
| `SyncError` | `error` |
| `PushError` | `error` |
| `StatusError` | `error` |
| `ListWorkspacesError` | `error` |

### MockManager Methods

| Method | Description |
|--------|-------------|
| `Reset()` | Clear all call history and errors |
| `CallCount(method string) int` | Count calls to a named method |
| `GetCalls(method string) []MockManagerCall` | All calls to a named method |
| `LastCall() *MockManagerCall` | Most recent call, or nil |
| `SetStatus(*Status)` | Set the status returned by `Status()` |
| `AddWorkspace(Workspace)` | Add a workspace to the list |
| `SetWorkspaces([]Workspace)` | Replace the workspace list |
| `SimulateLocalChanges()` | Set `LocalChanges = true` on the stored status |
| `SimulateRemoteChanges()` | Set `RemoteChanges = true` on the stored status |
| `SetInitialized(bool)` | Set the `Initialized` flag |
| `InjectError(method string, err error)` | Inject an error for a named method |

### MockManagerCall

```go
type MockManagerCall struct {
    Method string
    Args   []interface{}
}
```
