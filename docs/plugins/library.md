# Plugin Library

The nvp built-in library contains 38+ curated Neovim plugins with pre-configured setups ready to install.

---

## Browsing the Library

```bash
# List all available plugins
nvp library list

# Filter by category
nvp library list --category lsp
nvp library list --category fuzzy-finder
nvp library list --category git

# List all categories
nvp library categories

# List all tags
nvp library tags

# Show details about a plugin
nvp library show telescope
```

---

## Installing from the Library

```bash
nvp library install telescope
nvp library install treesitter
nvp library install lspconfig
```

After installing, generate Lua files:

```bash
nvp generate
```

---

## Available Plugins

### Core Dependencies

| Plugin | Description |
|--------|-------------|
| `plenary` | Lua utility functions (dependency for many plugins) |
| `nvim-web-devicons` | File icons |

### Fuzzy Finding and Navigation

| Plugin | Description |
|--------|-------------|
| `telescope` | Fuzzy finder for files, grep, buffers, and more |
| `harpoon` | Quick file navigation with bookmarks |

### Syntax and Parsing

| Plugin | Description |
|--------|-------------|
| `treesitter` | Advanced syntax highlighting and code understanding |
| `treesitter-textobjects` | Text objects based on tree-sitter syntax |

### LSP and Completion

| Plugin | Description |
|--------|-------------|
| `lspconfig` | Language Server Protocol configuration |
| `mason` | LSP/DAP/Linter installer |
| `nvim-cmp` | Intelligent autocompletion |
| `cmp-nvim-lsp` | LSP completion source |
| `cmp-buffer` | Buffer completion source |
| `cmp-path` | Filesystem path completion source |
| `luasnip` | Snippet engine |

### Git Integration

| Plugin | Description |
|--------|-------------|
| `gitsigns` | Git decorations and hunk navigation |
| `fugitive` | Git commands inside Neovim |
| `diffview` | Side-by-side git diff viewer |

### UI and Interface

| Plugin | Description |
|--------|-------------|
| `lualine` | Status line |
| `bufferline` | Buffer and tab line |
| `which-key` | Keybinding discovery and help |
| `alpha-nvim` | Dashboard / start screen |
| `neo-tree` | File tree explorer |
| `dressing` | Improved UI for inputs and selects |
| `notify` | Better notifications |

### Editing and Text Manipulation

| Plugin | Description |
|--------|-------------|
| `comment` | Easy line and block commenting |
| `surround` | Surround text with pairs |
| `autopairs` | Auto-close brackets and quotes |
| `conform` | Code formatting |
| `nvim-lint` | Linting integration |

### Terminal and AI

| Plugin | Description |
|--------|-------------|
| `toggleterm` | Floating and split terminal management |
| `copilot` | GitHub Copilot integration |
| `copilot-cmp` | Copilot completion source for nvim-cmp |

### Language-Specific

| Plugin | Description |
|--------|-------------|
| `rust-tools` | Enhanced Rust development tools |
| `go` | Go development with gopls and tools |
| `typescript-tools` | TypeScript and JavaScript support |

---

## Overriding a Library Plugin

Library plugins come with default configurations. To customize a plugin, apply a YAML with your own `config` block — it will override the library version:

```yaml
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
        layout_config = {
          width = 0.9,
          height = 0.8,
        },
      },
    })
```

```bash
nvp apply -f my-telescope.yaml
nvp generate
```

---

## Next Steps

- [Plugin Packages](packages.md) - Install curated plugin sets
- [Plugin Sources](sources.md) - Configure remote plugin sources
- [NvimPlugin YAML Reference](../reference/nvim-plugin.md) - All YAML fields explained
- [Commands Reference](../commands.md) - Full command reference
