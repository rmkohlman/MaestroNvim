# Plugin Packages

The nvp library organizes plugins into curated sets. Instead of installing plugins one by one, you can browse the library by category and install groups of related plugins that are known to work well together.

---

## What Are Plugin Packages?

An `NvimPackage` is a named collection of plugins. Packages support inheritance — a language-specific package can extend a base `core` package and add its own plugins on top.

Key benefits:

- **Complete setups** - Tested plugin combinations that work together
- **Consistent key bindings** - Uniform behavior across plugins
- **Inheritance** - Build on core packages without repeating plugin lists

---

## The Core Package

The `core` package is the foundation — 6 essential plugins for IDE-like functionality:

| Plugin | Category | Description |
|--------|----------|-------------|
| `treesitter` | syntax | Modern syntax highlighting and code understanding |
| `telescope` | fuzzy-finder | Fuzzy finder for files, grep, and buffers |
| `which-key` | ui | Keybinding discovery and help |
| `lspconfig` | lsp | Language Server Protocol configuration |
| `nvim-cmp` | completion | Intelligent autocompletion |
| `gitsigns` | git | Git integration with inline status |

---

## Installing from the Library

```bash
# Browse all available plugins
nvp library list

# Filter by category
nvp library list --category lsp
nvp library list --category fuzzy-finder

# See available categories
nvp library categories

# Install individual plugins
nvp library install telescope
nvp library install treesitter
nvp library install lspconfig
nvp library install nvim-cmp
nvp library install gitsigns
nvp library install which-key
```

---

## Defining a Custom Package

You can define your own `NvimPackage` in YAML and apply it:

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: golang-dev
  description: Go development environment
  category: language
  tags: ["go", "golang", "lsp"]
spec:
  extends: core
  plugins:
    - neovim/nvim-lspconfig
    - williamboman/mason.nvim
    - fatih/vim-go
    - ray-x/go.nvim
    - mfussenegger/nvim-dap
    - leoluz/nvim-dap-go
```

```bash
nvp apply -f golang-dev.yaml
nvp generate
```

### Package Inheritance

Packages can extend other packages:

```
core  →  golang-dev  →  golang-web
(base)   (adds Go LSP)  (adds HTTP tools)
```

```yaml
spec:
  extends: golang-dev   # Inherits all of golang-dev's plugins
  plugins:
    - NTBBloodbath/rest.nvim
```

---

## Managing Installed Plugins

```bash
# List installed plugins
nvp list

# Get details on a plugin
nvp get telescope
nvp get telescope -o yaml

# Enable or disable without deleting
nvp enable telescope
nvp disable copilot

# Delete a plugin
nvp delete telescope
```

---

## Override a Plugin Configuration

Apply a customized YAML to change a plugin's settings:

```bash
cat > my-telescope.yaml << 'EOF'
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
spec:
  repo: nvim-telescope/telescope.nvim
  config: |
    require("telescope").setup({
      defaults = {
        layout_strategy = "horizontal",
        layout_config = { width = 0.9, height = 0.8 },
      },
    })
EOF

nvp apply -f my-telescope.yaml
nvp generate
```

---

## Lazy Loading

All library plugins come with lazy loading pre-configured:

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

## Troubleshooting

### Plugin errors after generating

```bash
# Check Neovim health
nvim --headless -c 'checkhealth' -c 'quit'

# Regenerate clean configuration
rm -rf ~/.config/nvim/lua/plugins/nvp
nvp generate
```

### See what is installed

```bash
nvp list
nvp get telescope -o yaml
```

---

## Next Steps

- [Plugin Library](library.md) - Browse all 38+ curated plugins
- [Plugin Sources](sources.md) - Configure remote plugin sources
- [NvimPackage YAML Reference](../reference/nvim-package.md) - Full NvimPackage schema
- [Commands Reference](../commands.md) - Full command reference
