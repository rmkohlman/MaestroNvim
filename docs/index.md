# nvp - Neovim Plugin and Theme Manager

**kubectl-style Neovim plugin and theme manager using YAML.**

[![Release](https://img.shields.io/github/v/release/rmkohlman/devopsmaestro)](https://github.com/rmkohlman/devopsmaestro/releases/latest)
[![CI](https://github.com/rmkohlman/devopsmaestro/actions/workflows/ci.yml/badge.svg)](https://github.com/rmkohlman/devopsmaestro/actions/workflows/ci.yml)

---

## What is nvp?

`nvp` is the Neovim plugin and theme management tool from the DevOpsMaestro toolkit. It lets you define plugins and themes as YAML files, apply them from files, URLs, or GitHub repositories, and generate Lua configuration for lazy.nvim.

## Key Features

**Plugins**

- **YAML-based plugins** - Define plugins in YAML, generate Lua for lazy.nvim
- **38+ curated library** - Browse and install pre-configured plugins instantly
- **Plugin packages** - Install grouped collections with a single command
- **External source sync** - Import plugins from LazyVim, AstroNvim, NvChad, and other sources
- **kubectl-style IaC** - `nvp apply -f plugin.yaml`, URL support, GitHub shorthand

**Themes**

- **34+ embedded themes** - All themes available instantly, no installation required
- **21 CoolNight variants** - Blue, purple, green, warm, red/pink, monochrome, and special families
- **13+ popular themes** - Catppuccin, Dracula, Everforest, Gruvbox, Tokyo Night, Nord, and more
- **Parametric generator** - Create custom CoolNight variants from a hue angle, hex color, or preset name

**Infrastructure**

- **Standalone** - Works completely independently, no containers or dvm required
- **Shared database** - Integrates with dvm and dvt via `~/.devopsmaestro/devopsmaestro.db`

---

## Quick Install

=== "Homebrew (Recommended)"

    ```bash
    brew tap rmkohlman/tap

    # Install the full DevOpsMaestro toolkit (includes nvp)
    brew install devopsmaestro

    # Or install nvp standalone
    brew install nvimops

    # Verify
    nvp version
    ```

=== "From Source"

    ```bash
    git clone https://github.com/rmkohlman/devopsmaestro.git
    cd devopsmaestro
    go build -o nvp ./cmd/nvp/
    sudo mv nvp /usr/local/bin/
    ```

---

## Quick Example

```bash
# Initialize nvp
nvp init

# Browse and install plugins
nvp library list
nvp library install telescope treesitter lspconfig

# Install a theme
nvp theme library install coolnight-ocean --use

# Generate Lua files
nvp generate
# Creates ~/.config/nvim/lua/plugins/nvp/*.lua
```

### Infrastructure as Code

```bash
# Apply from file, URL, or GitHub
nvp apply -f plugin.yaml
nvp apply -f https://example.com/plugin.yaml
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml
```

### Plugin YAML Format

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  category: fuzzy-finder
spec:
  repo: nvim-telescope/telescope.nvim
  dependencies:
    - nvim-lua/plenary.nvim
  config: |
    require("telescope").setup({})
```

---

## Next Steps

<div class="grid cards" markdown>

-   **Getting Started**

    ---

    Install nvp and set up your first plugin configuration

    [Installation](getting-started/installation.md)

-   **Plugin Library**

    ---

    Browse and install from 38+ curated plugins

    [Plugin Library](plugins/library.md)

-   **Theme Collection**

    ---

    34+ embedded themes with CoolNight and parametric generator

    [Themes Overview](themes/overview.md)

-   **CoolNight Collection**

    ---

    21 parametrically generated variants for extended coding sessions

    [CoolNight](themes/coolnight.md)

-   **Commands Reference**

    ---

    Complete reference for all nvp commands and flags

    [Commands](commands.md)

-   **YAML Reference**

    ---

    Complete YAML schemas for NvimPlugin, NvimTheme, NvimPackage

    [NvimPlugin](reference/nvim-plugin.md)

</div>
