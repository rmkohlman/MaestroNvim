# NvimPackage YAML Reference

**Kind:** `NvimPackage`
**APIVersion:** `devopsmaestro.io/v1`

An NvimPackage defines a named collection of Neovim plugins. Apply with `nvp apply -f <file>` and export with `nvp package get <name> -o yaml`.

---

## Full Example

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: golang-dev
  description: "Complete Go development environment with LSP, debugging, and testing"
  category: "language"
  tags: ["go", "golang", "lsp", "debugging", "testing"]
  labels:
    language: "go"
    maintainer: "devopsmaestro"
  annotations:
    version: "1.0.0"
    last-updated: "2026-03-18"
spec:
  extends: "core"
  plugins:
    - neovim/nvim-lspconfig
    - williamboman/mason.nvim
    - williamboman/mason-lspconfig.nvim
    - hrsh7th/nvim-cmp
    - hrsh7th/cmp-nvim-lsp
    - fatih/vim-go
    - ray-x/go.nvim
    - ray-x/guihua.lua
    - mfussenegger/nvim-dap
    - leoluz/nvim-dap-go
    - rcarriga/nvim-dap-ui
    - nvim-neotest/neotest
    - nvim-neotest/neotest-go
    - stevearc/conform.nvim
  enabled: true
```

---

## Field Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `apiVersion` | string | Yes | Must be `devopsmaestro.io/v1` |
| `kind` | string | Yes | Must be `NvimPackage` |
| `metadata.name` | string | Yes | Unique identifier for the package |
| `metadata.description` | string | No | Package description |
| `metadata.category` | string | No | Package category |
| `metadata.tags` | array | No | Tags for filtering and searching |
| `metadata.labels` | object | No | Key-value labels |
| `metadata.annotations` | object | No | Key-value annotations |
| `spec.extends` | string | No | Name of parent package to inherit plugins from |
| `spec.plugins` | array | Yes | List of plugin references to include |
| `spec.enabled` | boolean | No | Enable or disable the package (default: `true`) |

---

## Field Details

### metadata.name (required)

Unique identifier for the package. Used as the key in the package store.

**Conventions:**
- Use kebab-case: `golang-dev`, `typescript-full`, `python-data`
- Include purpose or language: `rust-dev`, `web-frontend`, `data-science`
- Be specific: `react-typescript` rather than `js-stuff`

### metadata.category (optional)

Category for organization and browsing.

**Common categories:**
- `core` - Base or foundation packages
- `language` - Language-specific packages
- `framework` - Framework-specific packages (React, Vue, etc.)
- `purpose` - Purpose-specific packages (data-science, devops, etc.)
- `specialty` - Specialized or niche packages

### metadata.tags (optional)

Tags for filtering and discovery.

```yaml
metadata:
  tags: ["go", "golang", "lsp", "debugging", "testing"]
```

### spec.extends (optional)

Name of an existing package to inherit plugins from. The parent package's plugins are loaded first, then this package's plugins are appended.

```yaml
spec:
  extends: "core"
```

**Inheritance chain:**

```
core -> golang-dev -> golang-web
```

- `core` — base plugins (telescope, treesitter, lspconfig, etc.)
- `golang-dev` extends `core` — adds Go-specific plugins
- `golang-web` extends `golang-dev` — adds HTTP/REST tooling

Package inheritance must not create circular dependencies.

### spec.plugins (required)

List of plugin references to include in the package.

```yaml
spec:
  plugins:
    - neovim/nvim-lspconfig        # GitHub owner/repo format
    - williamboman/mason.nvim
    - fatih/vim-go
```

**Accepted formats:**
- Full GitHub repository: `neovim/nvim-lspconfig`
- Short name referencing an installed plugin: `telescope`

Plugins referenced by short name must already exist in the local plugin store.

### spec.enabled (optional)

Enable or disable the package. When `false`, the package is stored but not included when generating Lua. Defaults to `true`.

```yaml
spec:
  enabled: false
```

---

## Usage Examples

### Apply a Package

```bash
# Apply from a local file
nvp apply -f golang-dev.yaml

# Apply from a URL
nvp apply -f https://packages.example.com/golang-dev.yaml

# Apply from the plugin source library (if configured)
nvp source sync
nvp apply -f golang-dev.yaml
```

### List and Inspect Packages

```bash
# List all packages
nvp package list

# Alias
nvp pkg list

# Get package details
nvp package get golang-dev

# Export to YAML
nvp package get golang-dev -o yaml

# Save to file
nvp package get golang-dev -o yaml > golang-dev.yaml
```

### Enable and Disable

```bash
# Disable a package without deleting it
nvp package disable golang-dev

# Re-enable it
nvp package enable golang-dev

# Delete a package
nvp package delete golang-dev
```

### Generate Lua After Changes

Always regenerate Lua after applying or modifying packages:

```bash
nvp generate
```

---

## Built-in Package Examples

### Core Package

Base package providing essential IDE-like functionality:

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: core
  description: "Essential plugins for any Neovim setup"
  category: core
spec:
  plugins:
    - nvim-telescope/telescope.nvim
    - nvim-treesitter/nvim-treesitter
    - neovim/nvim-lspconfig
    - hrsh7th/nvim-cmp
    - lewis6991/gitsigns.nvim
    - folke/which-key.nvim
    - nvim-lualine/lualine.nvim
```

### Language Packages

#### Go Development

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: golang-dev
  category: language
  tags: ["go", "golang"]
spec:
  extends: core
  plugins:
    - fatih/vim-go
    - ray-x/go.nvim
    - leoluz/nvim-dap-go
```

#### Python Development

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: python-dev
  category: language
  tags: ["python"]
spec:
  extends: core
  plugins:
    - nvim-neotest/neotest-python
    - mfussenegger/nvim-dap-python
```

#### TypeScript Development

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: typescript-dev
  category: language
  tags: ["typescript", "javascript"]
spec:
  extends: core
  plugins:
    - nvim-neotest/neotest-jest
    - pmizio/typescript-tools.nvim
```

### Framework Packages

#### React Development

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPackage
metadata:
  name: react-dev
  category: framework
  tags: ["react", "jsx", "typescript"]
spec:
  extends: typescript-dev
  plugins:
    - windwp/nvim-ts-autotag
    - JoosepAlviste/nvim-ts-context-commentstring
```

---

## Package Inheritance Examples

### Linear Chain

```yaml
# Base
core:
  plugins: [telescope, treesitter, lspconfig]

# Language layer extends base
golang-dev:
  extends: core
  plugins: [vim-go, go.nvim, nvim-dap-go]
  # Effective: telescope + treesitter + lspconfig + vim-go + go.nvim + nvim-dap-go

# Framework layer extends language layer
golang-web:
  extends: golang-dev
  plugins: [rest.nvim]
  # Effective: all golang-dev plugins + rest.nvim
```

---

## Validation Rules

- `metadata.name` must be unique across all packages
- `spec.extends` must reference an existing package by name
- `spec.plugins` must contain valid plugin references
- Package inheritance must not create circular dependencies

---

## Related

- [Plugin Packages](../plugins/packages.md)
- [Plugin Library](../plugins/library.md)
- [NvimPlugin Reference](nvim-plugin.md)
- [NvimTheme Reference](nvim-theme.md)
- [Commands Reference](../commands.md)
