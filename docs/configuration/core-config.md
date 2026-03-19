# Neovim Core Configuration

`nvp` manages your Neovim plugin and theme configuration. The generated files integrate with lazy.nvim and follow a consistent directory structure.

---

## Configuration Storage

nvp stores all configuration in two locations:

| Location | Contents |
|----------|----------|
| `~/.nvp/` | Plugin and theme YAML definitions |
| `~/.config/nvim/lua/plugins/nvp/` | Generated Lua files for lazy.nvim |

### The `~/.nvp/` Directory

```
~/.nvp/
├── plugins/           # Installed plugin YAML files
│   ├── telescope.yaml
│   ├── treesitter.yaml
│   ├── lspconfig.yaml
│   └── ...
└── themes/            # User theme YAML files
    ├── active.yaml    # Currently active theme
    └── my-custom.yaml # User-defined themes
```

### The Generated Lua Directory

```
~/.config/nvim/lua/plugins/nvp/
├── telescope.lua
├── treesitter.lua
├── lspconfig.lua
├── nvim-cmp.lua
└── ...

~/.config/nvim/lua/theme/
├── init.lua        # Theme setup and helpers
└── palette.lua     # Color palette module
```

---

## Initialization

Run once to create the `~/.nvp/` directory structure:

```bash
nvp init
```

---

## Configuration File

### Initialize Config

```bash
nvp config init
```

Creates a default configuration file.

### Show Current Configuration

```bash
nvp config show
```

### Generate a Default Configuration File

```bash
nvp config generate
```

### Edit Configuration in Editor

```bash
nvp config edit
```

---

## Generating Lua Files

After installing or modifying plugins, regenerate the Lua files:

```bash
# Generate all plugins and active theme
nvp generate

# Generate to a custom directory
nvp generate --output ~/my-nvim-config/lua/plugins

# Generate only the theme
nvp theme generate
```

nvp always overwrites the entire `~/.config/nvim/lua/plugins/nvp/` directory on each run, so the generated files always reflect your current installed plugins.

---

## Sharing Configuration

### Export Your Configuration

Export any plugin or theme to YAML for sharing:

```bash
# Export a plugin
nvp get telescope -o yaml > telescope.yaml

# Export the active theme
nvp theme get -o yaml > my-theme.yaml
```

### Import Configuration

Apply shared configurations:

```bash
# From a local file
nvp apply -f shared-plugin.yaml

# From a URL
nvp apply -f https://example.com/plugin.yaml

# From GitHub
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml
```

---

## Shared Database

nvp stores metadata in `~/.devopsmaestro/devopsmaestro.db`, which is shared with `dvm` and `dvt` when those tools are installed. This allows consistent theme management across the DevOpsMaestro ecosystem.

nvp works standalone without `dvm` — the shared database is created automatically on first use.

---

## Next Steps

- [YAML Schema](yaml-schema.md) - Complete YAML format reference for all resource types
- [Commands Reference](../commands.md) - Full command reference
- [Plugin Overview](../plugins/overview.md) - Managing plugins
- [Themes Overview](../themes/overview.md) - Managing themes
