# Quick Start

Get `nvp` configured and your Neovim plugins generating Lua in under 5 minutes.

---

## Prerequisites

- `nvp` installed — see [Installation](installation.md)
- Neovim 0.9+ with [lazy.nvim](https://github.com/folke/lazy.nvim) set up

---

## Step 1: Initialize

Run once to create the `~/.nvp/` directory structure:

```bash
nvp init
```

---

## Step 2: Install Plugins from the Library

Browse the built-in library of 38+ curated plugins and install what you need:

```bash
# See everything available
nvp library list

# Filter by category
nvp library list --category lsp
nvp library list --category fuzzy-finder

# Install individual plugins
nvp library install telescope
nvp library install treesitter
nvp library install lspconfig
nvp library install nvim-cmp
nvp library install gitsigns
```

---

## Step 3: Pick a Theme

34+ themes are built in and ready to use immediately — no installation required:

```bash
# See all available themes
nvp theme list

# Use any library theme directly
nvp theme use coolnight-ocean        # Deep blue (default)
nvp theme use tokyonight-night       # Tokyo Night
nvp theme use catppuccin-mocha       # Catppuccin dark

# Or create a custom CoolNight variant
nvp theme create --from "280" --name my-synthwave --use
```

---

## Step 4: Generate Lua Files

Generate the Lua configuration files for lazy.nvim:

```bash
nvp generate
```

Files are created in `~/.config/nvim/lua/plugins/nvp/`.

---

## Step 5: Restart Neovim

Launch Neovim. lazy.nvim will pick up the generated files and install your plugins on startup.

```bash
nvim
```

---

## Complete Example

```bash
# 1. Initialize
nvp init

# 2. Install a set of plugins
nvp library install telescope
nvp library install treesitter
nvp library install lspconfig
nvp library install nvim-cmp
nvp library install gitsigns
nvp library install which-key

# 3. Set a theme
nvp theme use coolnight-ocean

# 4. Generate Lua files
nvp generate

# 5. Launch Neovim
nvim
```

---

## Apply from YAML

You can also define plugins in YAML and apply them directly:

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

Minimal plugin YAML:

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimPlugin
metadata:
  name: telescope
spec:
  repo: nvim-telescope/telescope.nvim
  lazy: true
  cmd:
    - Telescope
  keys:
    - key: "<leader>ff"
      action: "<cmd>Telescope find_files<cr>"
      desc: "Find files"
```

---

## Managing Installed Plugins

```bash
# List installed plugins
nvp list

# Get details on a plugin
nvp get telescope
nvp get telescope -o yaml

# Disable without deleting
nvp disable copilot

# Re-enable
nvp enable copilot

# Delete
nvp delete telescope
```

---

## File Structure

After running `nvp generate`:

```
~/.nvp/
├── plugins/           # Installed plugin YAMLs
│   ├── telescope.yaml
│   ├── treesitter.yaml
│   └── ...
└── themes/
    └── active.yaml    # Active theme

~/.config/nvim/lua/plugins/nvp/
├── telescope.lua
├── treesitter.lua
├── lspconfig.lua
└── ...
```

---

## Next Steps

- [Plugin Library](../plugins/library.md) - Browse all 38+ curated plugins
- [Plugin Packages](../plugins/packages.md) - Install curated plugin sets
- [Themes Overview](../themes/overview.md) - Explore the full theme library
- [Commands Reference](../commands.md) - Complete command reference
