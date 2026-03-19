# Themes Overview

`nvp` includes 34+ embedded themes that are ready to use immediately — no installation required. Use any library theme directly, or create custom CoolNight variants with the parametric generator.

---

## Quick Start

```bash
# List all available themes
nvp theme list

# Use any theme directly
nvp theme use coolnight-ocean        # Deep blue (default)
nvp theme use tokyonight-night
nvp theme use catppuccin-mocha

# Create a custom CoolNight variant
nvp theme create --from "210" --name my-blue-theme --use

# Generate Lua files
nvp generate
```

---

## Library Themes vs User Themes

| Type | Location | Access | Customizable |
|------|----------|--------|--------------|
| **Library** | Embedded in binary | Automatic, no install needed | No (read-only) |
| **User** | `~/.nvp/themes/` | Manual via `nvp apply` | Yes (full control) |

**Override behavior:** A user theme with the same name takes precedence over the library version.

---

## 34+ Built-in Themes

### CoolNight Collection (21 themes)

The CoolNight Collection is a set of parametrically generated themes designed for extended coding sessions. All 21 variants are available immediately.

Popular variants:

| Theme | Hue | Description |
|-------|-----|-------------|
| `coolnight-ocean` | 210° | Deep blue — professional default |
| `coolnight-arctic` | 190° | Ice blue, crisp and clean |
| `coolnight-midnight` | 240° | Dark blue, intense focus |
| `coolnight-synthwave` | 280° | Retro neon purple |
| `coolnight-matrix` | 120° | High-contrast green |
| `coolnight-sunset` | 30° | Warm orange |
| `coolnight-rose` | 350° | Rose pink |
| `coolnight-mono-slate` | — | Minimalist grayscale |

See the [CoolNight Collection](coolnight.md) for all 21 variants.

### Popular Themes (13+)

| Theme | Style | Description |
|-------|-------|-------------|
| `tokyonight-night` | dark | Standard Tokyo Night |
| `tokyonight-storm` | dark | Stormy blue variant |
| `tokyonight-day` | light | Light Tokyo Night |
| `catppuccin-mocha` | dark | Soothing dark pastel |
| `catppuccin-latte` | light | Warm light pastel |
| `catppuccin-frappe` | dark | Medium pastel |
| `catppuccin-macchiato` | dark | Dark warm pastel |
| `gruvbox-dark` | dark | Retro warm |
| `gruvbox-light` | light | Retro light |
| `nord` | dark | Arctic bluish |
| `dracula` | dark | Dark purple |
| `one-dark` | dark | Dark blue |
| `solarized-dark` | dark | Blue-green |

---

## Parametric Generator

Create custom CoolNight variants using a hue angle (0–360), a hex color, or a preset name:

```bash
# Create by hue angle
nvp theme create --from "210" --name my-blue-theme
nvp theme create --from "350" --name my-rose-theme

# Create from a hex color
nvp theme create --from "#8B00FF" --name my-violet-theme

# Create from a preset name
nvp theme create --from "synthwave" --name my-synth

# Create and set as active immediately
nvp theme create --from "280" --name my-purple --use
```

See [Parametric Generator](parametric.md) for detailed usage.

---

## Using Themes

```bash
# Set the active theme
nvp theme use coolnight-ocean

# View the active theme
nvp theme get

# View a specific theme
nvp theme get coolnight-ocean
nvp theme get coolnight-ocean -o yaml

# Preview in terminal
nvp theme preview coolnight-ocean

# Delete a user theme
nvp theme delete my-custom-theme
```

---

## Theme YAML Format

Apply a custom theme from YAML:

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimTheme
metadata:
  name: my-custom-theme
  description: My personal colorscheme
  category: dark
spec:
  plugin:
    repo: folke/tokyonight.nvim
  style: night
  transparent: false
  colors:
    bg: "#1a1b26"
    fg: "#c0caf5"
    accent: "#7aa2f7"
```

```bash
nvp apply -f my-custom-theme.yaml
nvp theme use my-custom-theme
nvp generate
```

---

## Generated Files

When you run `nvp generate`, theme files are created in:

```
~/.config/nvim/lua/
├── theme/
│   ├── init.lua        # Theme setup and helpers
│   └── palette.lua     # Color palette module
└── plugins/nvp/
    └── colorscheme.lua # Lazy.nvim plugin spec
```

---

## Switching Themes

```bash
nvp theme use gruvbox-dark
nvp generate
# Restart Neovim to see changes
```

---

## Next Steps

- [Theme Library](library.md) - Complete library theme listing
- [CoolNight Collection](coolnight.md) - All 21 CoolNight variants
- [Parametric Generator](parametric.md) - Create custom variants
- [NvimTheme YAML Reference](../reference/nvim-theme.md) - Full YAML schema
- [Commands Reference](../commands.md) - Full command reference
