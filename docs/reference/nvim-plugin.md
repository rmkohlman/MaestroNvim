# NvimPlugin YAML Reference

**Kind:** `NvimPlugin`
**APIVersion:** `devopsmaestro.io/v1`

An NvimPlugin defines a Neovim plugin configuration in YAML. Apply with `nvp apply -f <file>` and export with `nvp get <name> -o yaml`.

---

## Full Example

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  description: "Fuzzy finder over lists"
  category: "navigation"
  tags: ["fuzzy-finder", "telescope", "navigation", "files"]
  labels:
    maintainer: "nvim-telescope"
    language: "lua"
  annotations:
    version: "0.1.4"
    last-updated: "2026-02-19"
    documentation: "https://github.com/nvim-telescope/telescope.nvim"
spec:
  repo: "nvim-telescope/telescope.nvim"
  branch: "master"
  version: "0.1.4"
  lazy: true
  priority: 1000
  event: ["VeryLazy"]
  ft: ["lua", "vim"]
  cmd: ["Telescope"]
  keys:
    - key: "<leader>ff"
      mode: "n"
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files"
    - key: "<leader>fg"
      mode: "n"
      action: "<cmd>Telescope live_grep<cr>"
      desc: "Live grep"
    - key: "<leader>fb"
      mode: ["n", "v"]
      action: "<cmd>Telescope buffers<cr>"
      desc: "Find buffers"
  dependencies:
    - "nvim-lua/plenary.nvim"
    - repo: "nvim-tree/nvim-web-devicons"
      build: ""
      config: false
    - repo: "nvim-telescope/telescope-fzf-native.nvim"
      build: "make"
  build: "make"
  config: |
    require('telescope').setup({
      defaults = {
        file_ignore_patterns = {"node_modules", ".git/"},
        layout_strategy = 'horizontal',
        layout_config = {
          width = 0.95,
          height = 0.85,
          preview_cutoff = 120,
        },
      },
    })
    require('telescope').load_extension('fzf')
  init: |
    vim.g.telescope_theme = 'dropdown'
  opts:
    defaults:
      prompt_prefix: "> "
      selection_caret: "> "
      path_display: ["truncate"]
```

---

## Field Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `apiVersion` | string | Yes | Must be `devopsmaestro.io/v1` |
| `kind` | string | Yes | Must be `NvimPlugin` |
| `metadata.name` | string | Yes | Unique identifier for the plugin |
| `metadata.description` | string | No | Plugin description |
| `metadata.category` | string | No | Plugin category |
| `metadata.tags` | array | No | Tags for filtering |
| `metadata.labels` | object | No | Key-value labels |
| `metadata.annotations` | object | No | Key-value annotations |
| `spec.repo` | string | Yes | GitHub repository (`owner/repo`) |
| `spec.branch` | string | No | Git branch |
| `spec.version` | string | No | Git tag or version |
| `spec.lazy` | boolean | No | Enable lazy loading |
| `spec.priority` | integer | No | Load priority (higher = earlier) |
| `spec.event` | array | No | Load on these Neovim events |
| `spec.ft` | array | No | Load on these filetypes |
| `spec.cmd` | array | No | Load when these commands are used |
| `spec.keys` | array | No | Key mappings (also triggers load) |
| `spec.dependencies` | array | No | Plugin dependencies |
| `spec.build` | string | No | Build command after install |
| `spec.config` | string | No | Lua configuration code (runs after load) |
| `spec.init` | string | No | Lua initialization code (runs before load) |
| `spec.opts` | object | No | Options passed to plugin's setup function |

---

## Field Details

### metadata.name (required)

Unique identifier for the plugin. Used as the filename under `~/.nvp/plugins/`.

**Conventions:**
- Use the plugin's common name: `telescope`, `lspconfig`, `treesitter`
- For customized configs: `telescope-custom`, `lsp-golang`

### metadata.category (optional)

Category for organization and library filtering.

**Common categories:**
- `navigation` - File and buffer navigation
- `lsp` - Language Server Protocol
- `completion` - Code completion
- `syntax` - Syntax highlighting
- `git` - Git integration
- `ui` - User interface enhancements
- `editing` - Text editing
- `debugging` - Debug support
- `testing` - Test integration

### spec.repo (required)

GitHub repository in `owner/repo` format.

```yaml
spec:
  repo: "nvim-telescope/telescope.nvim"
```

### spec.branch and spec.version (optional)

Pin to a specific branch or version tag.

```yaml
spec:
  branch: "master"      # Use a specific branch
  # OR
  version: "0.1.4"      # Use a specific tag
```

### spec.priority (optional)

Load priority. Higher numbers load earlier. Useful for colorschemes that must load before other plugins.

```yaml
spec:
  priority: 1000
```

### Lazy Loading

#### spec.lazy (optional)

```yaml
spec:
  lazy: true
```

#### spec.event (optional)

Load on Neovim events:

```yaml
spec:
  event: ["BufReadPre", "BufNewFile"]
  # OR
  event: ["VeryLazy"]
```

**Common events:**
- `VeryLazy` - After startup
- `BufReadPre` - Before reading a buffer
- `BufNewFile` - On new file
- `InsertEnter` - Entering insert mode

#### spec.ft (optional)

Load for specific filetypes:

```yaml
spec:
  ft: ["go", "lua", "python"]
```

#### spec.cmd (optional)

Load when specific commands are used:

```yaml
spec:
  cmd: ["Telescope", "Tele"]
```

#### spec.keys (optional)

Load on key mapping press. Also registers the key mappings.

```yaml
spec:
  keys:
    - key: "<leader>ff"
      mode: "n"                        # n, i, v, x, o, c
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files"
    - key: "<C-p>"
      mode: ["n", "i"]                 # Multiple modes
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files (Ctrl+P)"
```

### spec.dependencies (optional)

Plugin dependencies loaded before this plugin.

```yaml
spec:
  dependencies:
    # Simple string format
    - "nvim-lua/plenary.nvim"

    # Detailed format
    - repo: "nvim-tree/nvim-web-devicons"
      build: ""
      config: false
    - repo: "nvim-telescope/telescope-fzf-native.nvim"
      build: "make"
```

### spec.build (optional)

Build command run after plugin installation or update.

```yaml
spec:
  build: "make"             # Shell command
  # OR
  build: ":TSUpdate"        # Neovim command
  # OR
  build: "npm install"      # Node command
```

### spec.config (optional)

Lua code executed after the plugin loads.

```yaml
spec:
  config: |
    require('telescope').setup({
      defaults = {
        file_ignore_patterns = {"node_modules"},
      }
    })
```

### spec.init (optional)

Lua code executed before the plugin loads.

```yaml
spec:
  init: |
    vim.g.telescope_theme = 'dropdown'
```

### spec.opts (optional)

Options passed directly to the plugin's setup function.

```yaml
spec:
  opts:
    defaults:
      prompt_prefix: "> "
    pickers:
      find_files:
        theme: "dropdown"
```

---

## Examples by Category

### Navigation Plugin

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
  category: navigation
spec:
  repo: "nvim-telescope/telescope.nvim"
  keys:
    - key: "<leader>ff"
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files"
```

### LSP Plugin

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: lspconfig
  category: lsp
spec:
  repo: "neovim/nvim-lspconfig"
  event: ["BufReadPre", "BufNewFile"]
```

### Completion Plugin

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: nvim-cmp
  category: completion
spec:
  repo: "hrsh7th/nvim-cmp"
  event: "InsertEnter"
  dependencies:
    - "hrsh7th/cmp-nvim-lsp"
    - "hrsh7th/cmp-buffer"
    - "hrsh7th/cmp-path"
```

### Language-Specific Plugin

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: vim-go
  category: language
  tags: ["go", "golang"]
spec:
  repo: "fatih/vim-go"
  ft: ["go"]
```

---

## Usage Examples

```bash
# Apply from a YAML file
nvp apply -f my-plugin.yaml

# Apply from GitHub
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml

# List installed plugins
nvp list

# Get plugin details
nvp get telescope -o yaml

# Export a plugin
nvp get telescope -o yaml > telescope-config.yaml
```

---

## Validation Rules

- `metadata.name` must be unique across all plugins
- `spec.repo` must be in `owner/repo` format
- `spec.keys[].mode` must be valid Neovim mode(s): `n`, `i`, `v`, `x`, `o`, `c`
- `spec.config` and `spec.init` must be valid Lua code
- `spec.priority` must be a positive integer

---

## Related

- [Plugins Overview](../plugins/overview.md)
- [Plugin Library](../plugins/library.md)
- [Plugin Packages](../plugins/packages.md) and [NvimPackage Reference](nvim-package.md)
- [NvimTheme Reference](nvim-theme.md)
