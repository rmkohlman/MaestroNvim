# MaestroNvim

Go library for managing Neovim configurations in the DevOpsMaestro ecosystem.

## Overview

MaestroNvim provides a programmatic interface for initializing, syncing, and managing Neovim configurations across workspaces. It includes:

- **Nvim Manager** — initialize configs from templates, sync status tracking, workspace-aware config paths
- **Plugin management** — parse, store, generate, and apply Neovim plugin definitions from YAML
- **Plugin store** — file-backed, in-memory, and read-only store implementations
- **Core config generation** — generate complete Lua config directory trees from a single YAML definition
- **Plugin library** — 51 curated, embedded plugin definitions ready to use
- **Package system** — group plugins into named packages with single-inheritance composition
- **Package library** — 12 curated, embedded package definitions (language-specific dev environments)
- **Sync framework** — extensible source handler system for syncing plugins from external sources (LazyVim, AstroNvim, NvChad, etc.)

## Installation

```
go get github.com/rmkohlman/MaestroNvim
```

Requires Go 1.25.6 or later.

## Quick Usage

### Initialize a Neovim config from a template

```go
import "github.com/rmkohlman/MaestroNvim/nvim"

mgr := nvim.NewManager()
err := mgr.Init(nvim.InitOptions{
    Template: "lazyvim",
    Overwrite: false,
})
```

### Parse a plugin YAML and generate Lua

```go
import (
    "github.com/rmkohlman/MaestroNvim/nvimops/plugin"
)

p, err := plugin.ParseYAMLFile("telescope.yaml")
if err != nil {
    log.Fatal(err)
}

gen := plugin.NewGenerator()
lua, err := gen.GenerateLua(p)
```

### Load a plugin from the embedded library

```go
import "github.com/rmkohlman/MaestroNvim/nvimops/library"

lib, err := library.NewLibrary()
if err != nil {
    log.Fatal(err)
}

p, ok := lib.Get("telescope")
if ok {
    fmt.Println(p.Repo) // nvim-telescope/telescope.nvim
}
```

## Documentation

Full API reference and package documentation: [https://rmkohlman.github.io/MaestroNvim/](https://rmkohlman.github.io/MaestroNvim/)

## Ecosystem

MaestroNvim is part of the DevOpsMaestro ecosystem: [https://github.com/rmkohlman/devopsmaestro](https://github.com/rmkohlman/devopsmaestro)

## License

GPL-3.0 with commercial dual-license. See LICENSE for details.
