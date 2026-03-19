# nvp (NvimOps) -- Neovim Plugin and Theme Manager

[![Release](https://img.shields.io/github/v/release/rmkohlman/devopsmaestro)](https://github.com/rmkohlman/devopsmaestro/releases/latest)
[![CI](https://github.com/rmkohlman/devopsmaestro/actions/workflows/ci.yml/badge.svg)](https://github.com/rmkohlman/devopsmaestro/actions/workflows/ci.yml)
[![License](https://img.shields.io/badge/license-GPL--3.0-blue)](LICENSE)

**kubectl-style Neovim plugin and theme manager using YAML.**

`nvp` lets you define Neovim plugins and themes as YAML files, apply them from files, URLs, or GitHub repositories, and generate Lua configuration for lazy.nvim. No more hand-writing plugin specs.

---

## Key Features

- **YAML-based plugins** - Define plugins in YAML, generate Lua for lazy.nvim
- **38+ curated library** - Browse and install pre-configured plugins instantly
- **Plugin packages** - Install grouped plugin sets with a single command
- **34+ embedded themes** - All themes available instantly without installation
  - **21 CoolNight variants** - Blue, purple, green, warm, red/pink, monochrome, and special families
  - **13+ additional themes** - Catppuccin, Dracula, Everforest, Gruvbox, Tokyo Night, Nord, and more
- **Parametric theme generator** - Create custom CoolNight variants from a hue angle, hex color, or preset name
- **kubectl-style IaC** - `nvp apply -f theme.yaml`, URL support, GitHub shorthand
- **External source sync** - Import plugins from LazyVim, AstroNvim, NvChad, and other sources
- **Standalone** - Works completely independently, no containers required

---

## Installation

### Homebrew (Recommended)

```bash
# Add the tap
brew tap rmkohlman/tap

# Install DevOpsMaestro toolkit (includes dvm, nvp, and dvt)
brew install devopsmaestro

# Or install nvp standalone (no containers needed)
brew install nvimops

# Verify installation
nvp version
```

### From Source

```bash
git clone https://github.com/rmkohlman/devopsmaestro.git
cd devopsmaestro

# Build nvp
go build -o nvp ./cmd/nvp/

# Install to PATH
sudo mv nvp /usr/local/bin/
```

Requires Go 1.25+ to build from source.

---

## Quick Start

```bash
# Initialize nvp
nvp init

# Browse the plugin library (38+ curated plugins)
nvp library list
nvp library list --category lsp

# Install plugins from library
nvp library install telescope treesitter lspconfig

# Browse the theme library (34+ embedded themes)
nvp theme library list

# Install a theme and set it as active
nvp theme library install coolnight-ocean --use

# Or use any library theme directly (no installation needed)
nvp theme use catppuccin-mocha

# Generate Lua files for Neovim
nvp generate
# Creates ~/.config/nvim/lua/plugins/nvp/*.lua
```

### Apply from YAML (Infrastructure as Code)

```bash
# Apply a plugin from a file
nvp apply -f plugin.yaml

# Apply from a URL
nvp apply -f https://example.com/plugin.yaml

# Apply from GitHub shorthand
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml

# Apply from stdin
cat plugin.yaml | nvp apply -f -
```

### Create a Custom Theme

```bash
# From a hue angle (0-360)
nvp theme create --from "210" --name my-blue-theme --use

# From a hex color
nvp theme create --from "#8B00FF" --name my-violet --use

# From a preset name
nvp theme create --from "synthwave" --name my-synth --use
```

---

## File Structure

nvp stores data in `~/.nvp/`:

```
~/.nvp/
├── plugins/           # Installed plugin YAMLs
│   ├── telescope.yaml
│   ├── treesitter.yaml
│   └── ...
└── themes/            # User themes
    └── my-theme.yaml

~/.config/nvim/lua/plugins/nvp/
├── telescope.lua
├── treesitter.lua
└── ...
```

---

## Documentation

Full documentation: [https://rmkohlman.github.io/MaestroNvim/](https://rmkohlman.github.io/MaestroNvim/)

---

## Part of DevOpsMaestro

nvp is part of the DevOpsMaestro toolkit: [https://github.com/rmkohlman/devopsmaestro](https://github.com/rmkohlman/devopsmaestro)

| Tool | Binary | Description |
|------|--------|-------------|
| **DevOpsMaestro** | `dvm` | Workspace and app management with container-native dev environments |
| **NvimOps** | `nvp` | Neovim plugin and theme manager (this tool) |
| **Terminal Operations** | `dvt` | Terminal prompt and configuration management |

---

## License

GPL-3.0 with commercial dual-license. See [LICENSE](LICENSE) for details.
