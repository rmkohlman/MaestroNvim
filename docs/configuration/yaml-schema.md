# YAML Schema Reference

nvp uses Kubernetes-style YAML for all resources: `apiVersion`, `kind`, and `metadata`/`spec`. Resources can be applied with `nvp apply -f <file>` and exported with `nvp get <name> -o yaml`.

---

## Base Structure

All resources follow this pattern:

```yaml
apiVersion: devopsmaestro.io/v1
kind: <ResourceType>
metadata:
  name: <resource-name>
  description: Optional description
  labels: {}
  annotations: {}
spec:
  # Resource-specific fields
```

---

## Resource Types

nvp supports three resource kinds:

| Kind | Purpose |
|------|---------|
| `NvimPlugin` | A Neovim plugin configuration |
| `NvimTheme` | A Neovim colorscheme theme |
| `NvimPackage` | A named collection of plugins |

---

## NvimPlugin

Represents a Neovim plugin configuration.

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  description: "Fuzzy finder over lists"
  category: "navigation"
  tags: ["fuzzy-finder", "telescope", "navigation"]
spec:
  repo: "nvim-telescope/telescope.nvim"
  branch: "master"                          # Optional
  version: "0.1.4"                         # Optional git tag
  lazy: true
  priority: 1000                            # Higher = loads earlier
  event: ["VeryLazy"]
  ft: ["lua", "vim"]
  cmd: ["Telescope"]
  keys:
    - key: "<leader>ff"
      mode: "n"
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files"
  dependencies:
    - "nvim-lua/plenary.nvim"
    - repo: "nvim-tree/nvim-web-devicons"
      build: ""
      config: false
  build: "make"
  config: |
    require('telescope').setup({
      defaults = {
        file_ignore_patterns = {"node_modules"},
        layout_strategy = 'horizontal',
      }
    })
  init: |
    vim.g.telescope_theme = 'dropdown'
  opts:
    defaults:
      prompt_prefix: "> "
      selection_caret: "> "
```

See the [NvimPlugin YAML Reference](../reference/nvim-plugin.md) for all fields.

---

## NvimTheme

Represents a Neovim colorscheme theme.

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimTheme
metadata:
  name: my-custom-theme
  description: "My beautiful custom theme"
  author: "your-name"
  category: "dark"                          # dark, light, monochrome
spec:
  plugin:
    repo: "folke/tokyonight.nvim"
    branch: "main"
    tag: "v1.0.0"
  style: "night"                            # Theme variant (plugin-specific)
  transparent: false
  colors:
    bg: "#1a1b26"
    fg: "#c0caf5"
    accent: "#7aa2f7"
    comment: "#565f89"
    error: "#f7768e"
    warning: "#e0af68"
    info: "#7dcfff"
    hint: "#1abc9c"
    selection: "#33467c"
    border: "#29a4bd"
  options:
    italic_comments: true
    bold_keywords: false
    underline_errors: true
    custom_highlights:
      - group: "Comment"
        style: "italic"
        fg: "#565f89"
```

See the [NvimTheme YAML Reference](../reference/nvim-theme.md) for all fields.

---

## NvimPackage

Represents a named collection of plugins.

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: golang-dev
  description: "Complete Go development environment"
  category: "language"
  tags: ["go", "golang", "lsp", "debugging"]
spec:
  extends: "core"                           # Inherit from another package
  plugins:
    - neovim/nvim-lspconfig
    - williamboman/mason.nvim
    - fatih/vim-go
    - ray-x/go.nvim
    - mfussenegger/nvim-dap
    - leoluz/nvim-dap-go
  enabled: true                             # Default: true
```

See the [NvimPackage YAML Reference](../reference/nvim-package.md) for all fields.

---

## Multi-Document YAML

You can combine multiple resources in a single file using `---` separators:

```yaml
---
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
spec:
  repo: nvim-telescope/telescope.nvim
---
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: treesitter
spec:
  repo: nvim-treesitter/nvim-treesitter
---
apiVersion: devopsmaestro.io/v1
kind: NvimTheme
metadata:
  name: my-theme
spec:
  plugin:
    repo: folke/tokyonight.nvim
  style: night
```

```bash
nvp apply -f full-config.yaml
```

---

## Applying Resources

```bash
# From a local file
nvp apply -f my-plugin.yaml

# From a URL
nvp apply -f https://example.com/plugin.yaml

# From GitHub shorthand
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml

# From stdin
cat plugin.yaml | nvp apply -f -
```

---

## Exporting Resources

```bash
# Export a plugin
nvp get telescope -o yaml

# Export the active theme
nvp theme get -o yaml

# Export a specific theme
nvp theme get coolnight-ocean -o yaml
```

---

## Validation

nvp validates YAML on apply:

- Required fields: `apiVersion`, `kind`, `metadata.name`
- `apiVersion` must be `devopsmaestro.io/v1`
- `kind` must be `NvimPlugin`, `NvimTheme`, or `NvimPackage`
- `spec.repo` must be in `owner/repo` format
- Color values must be valid hex (`#rrggbb`)
- `metadata.category` for themes must be `dark`, `light`, or `monochrome`

---

## Next Steps

- [NvimPlugin YAML Reference](../reference/nvim-plugin.md) - Complete plugin field reference
- [NvimTheme YAML Reference](../reference/nvim-theme.md) - Complete theme field reference
- [NvimPackage YAML Reference](../reference/nvim-package.md) - Complete package field reference
- [Commands Reference](../commands.md) - Full command reference
