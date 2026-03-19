# MaestroNvim

MaestroNvim is a Go library for managing Neovim configurations in the DevOpsMaestro ecosystem. It provides a complete toolkit for initializing configs from templates, managing plugins and packages declaratively via YAML, generating Lua configuration trees, and syncing plugins from external sources.

## Packages

| Package | Import Path | Description |
|---------|-------------|-------------|
| `nvim` | `github.com/rmkohlman/MaestroNvim/nvim` | Nvim config lifecycle management |
| `nvimops` | `github.com/rmkohlman/MaestroNvim/nvimops` | Plugin apply, generate, store, and list |
| `nvimops/plugin` | `github.com/rmkohlman/MaestroNvim/nvimops/plugin` | Plugin types, YAML parsing, Lua generation |
| `nvimops/store` | `github.com/rmkohlman/MaestroNvim/nvimops/store` | Plugin persistence (file, memory, read-only) |
| `nvimops/config` | `github.com/rmkohlman/MaestroNvim/nvimops/config` | Core Neovim config generation from YAML |
| `nvimops/library` | `github.com/rmkohlman/MaestroNvim/nvimops/library` | Embedded library of 51 curated plugins |
| `nvimops/package` | `github.com/rmkohlman/MaestroNvim/nvimops/package` | Plugin package types and YAML parsing |
| `nvimops/package/library` | `github.com/rmkohlman/MaestroNvim/nvimops/package/library` | Embedded library of 12 curated packages |
| `nvimops/sync` | `github.com/rmkohlman/MaestroNvim/nvimops/sync` | Sync framework interfaces and registry |

## Installation

```
go get github.com/rmkohlman/MaestroNvim
```

## Quick Examples

**Initialize a config from a template:**

```go
mgr := nvim.NewManager()
err := mgr.Init(nvim.InitOptions{
    Template: "lazyvim",
})
```

**Parse a plugin YAML file and generate Lua:**

```go
p, err := plugin.ParseYAMLFile("telescope.yaml")
gen := plugin.NewGenerator()
lua, err := gen.GenerateLua(p)
```

**Look up a plugin in the embedded library:**

```go
lib, _ := library.NewLibrary()
p, ok := lib.Get("telescope")
```

**Generate a complete Neovim config directory:**

```go
cfg, _ := config.ParseYAMLFile("core.yaml")
gen := config.NewGenerator()
gen.WriteToDirectory(cfg, plugins, "/home/dev/.config/nvim")
```

## Source

[https://github.com/rmkohlman/MaestroNvim](https://github.com/rmkohlman/MaestroNvim)

Part of the [DevOpsMaestro ecosystem](https://github.com/rmkohlman/devopsmaestro).
