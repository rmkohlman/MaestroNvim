# Plugins Overview

`nvp` manages Neovim plugins using YAML. Define plugins as Infrastructure as Code, install from the built-in library, and generate Lua configuration files for lazy.nvim.

---

## How It Works

```
YAML Plugin Definitions  →  nvp generate  →  Lua files (lazy.nvim)
  (~/.nvp/plugins/)                          (~/.config/nvim/lua/plugins/nvp/)
```

1. Install plugins from the library or define them in YAML
2. Run `nvp generate`
3. Lua files are created for lazy.nvim to load on Neovim startup

---

## Installing Plugins

### From the Built-in Library

The easiest way — 38+ curated plugins with pre-configured setups:

```bash
# Browse the library
nvp library list
nvp library list --category lsp

# Install plugins
nvp library install telescope
nvp library install treesitter
nvp library install lspconfig
```

### From a YAML File

```bash
nvp apply -f my-plugin.yaml
```

### From a URL

```bash
nvp apply -f https://example.com/plugin.yaml
```

### From GitHub

```bash
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml
```

### From Stdin

```bash
cat plugin.yaml | nvp apply -f -
```

---

## Managing Plugins

```bash
# List installed plugins
nvp list

# Get plugin details
nvp get telescope
nvp get telescope -o yaml

# Enable / disable (without deleting)
nvp enable telescope
nvp disable copilot

# Delete
nvp delete telescope
```

---

## Generating Lua Files

After installing plugins, generate the Lua configuration:

```bash
nvp generate
```

Files go to `~/.config/nvim/lua/plugins/nvp/` by default.

### Custom Output Directory

```bash
nvp generate --output ~/my-nvim-config/lua/plugins
```

---

## Plugin YAML Format

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  category: fuzzy-finder
  description: Highly extendable fuzzy finder
spec:
  repo: nvim-telescope/telescope.nvim
  branch: master                    # Optional
  version: "0.1.5"                  # Optional (tag)
  enabled: true                     # Default: true
  lazy: true                        # Default: true
  event:
    - VimEnter
  cmd:
    - Telescope
  dependencies:
    - nvim-lua/plenary.nvim
    - nvim-tree/nvim-web-devicons
  config: |
    require("telescope").setup({
      defaults = {
        file_ignore_patterns = { "node_modules", ".git" },
      },
    })
  keys:
    - key: "<leader>ff"
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files"
    - key: "<leader>fg"
      action: "<cmd>Telescope live_grep<cr>"
      desc: "Live grep"
```

---

## Lazy Loading

Control when plugins load to minimize Neovim startup time:

```yaml
spec:
  lazy: true          # Don't load on startup
  event:              # Load on these events
    - BufReadPost
    - BufNewFile
  cmd:                # Load when these commands run
    - Telescope
  ft:                 # Load for these filetypes
    - python
    - go
  keys:               # Load when these keys are pressed
    - key: "<leader>ff"
      action: "<cmd>Telescope find_files<cr>"
```

---

## Next Steps

- [Plugin Library](library.md) - Browse all 38+ curated plugins
- [Plugin Packages](packages.md) - Install curated plugin sets
- [Plugin Sources](sources.md) - Configure remote plugin sources
- [Commands Reference](../commands.md) - Full command reference
- [NvimPlugin YAML Reference](../reference/nvim-plugin.md) - All YAML fields explained
