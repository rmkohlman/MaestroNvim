# Theme Library

34+ themes are embedded in the `nvp` binary and available immediately — no installation required. Use them directly with `nvp theme use <name>`.

---

## Using Library Themes

```bash
# List all available themes (library + user)
nvp theme list

# Use any theme directly
nvp theme use coolnight-ocean
nvp theme use tokyonight-night
nvp theme use catppuccin-mocha

# Get theme details
nvp theme get coolnight-ocean
nvp theme get coolnight-ocean -o yaml

# Preview in terminal
nvp theme preview coolnight-ocean
```

---

## CoolNight Collection (21 themes)

The CoolNight Collection is a set of parametrically generated themes optimized for extended coding sessions. See the [CoolNight Collection](coolnight.md) page for the complete guide.

### Blue Family

| Theme | Hue | Best For |
|-------|-----|----------|
| `coolnight-arctic` | 190° | TypeScript, Go, documentation |
| `coolnight-ocean` | 210° | General development (default) |
| `coolnight-midnight` | 240° | Late-night coding, C++ |

### Purple Family

| Theme | Hue | Best For |
|-------|-----|----------|
| `coolnight-violet` | 270° | Web development, CSS |
| `coolnight-synthwave` | 280° | JavaScript, creative coding |
| `coolnight-grape` | 290° | Rust, systems programming |

### Green Family

| Theme | Hue | Best For |
|-------|-----|----------|
| `coolnight-forest` | 110° | Bash, DevOps |
| `coolnight-matrix` | 120° | Terminal work, high contrast |
| `coolnight-mint` | 150° | React, Vue.js, modern JS |

### Warm Family

| Theme | Hue | Best For |
|-------|-----|----------|
| `coolnight-ember` | 20° | Java, Spring Boot |
| `coolnight-sunset` | 30° | HTML, markup languages |
| `coolnight-gold` | 45° | Configuration files, YAML |

### Red and Pink Family

| Theme | Hue | Best For |
|-------|-----|----------|
| `coolnight-crimson` | 0° | Error handling, debugging |
| `coolnight-sakura` | 320° | Design systems |
| `coolnight-rose` | 350° | Personal projects |

### Monochrome Family

| Theme | Description |
|-------|-------------|
| `coolnight-mono-charcoal` | Charcoal gray, minimalist |
| `coolnight-mono-slate` | Slate gray, professional |
| `coolnight-mono-warm` | Warm gray, comfortable |

### Special Variants

| Theme | Inspiration |
|-------|-------------|
| `coolnight-nord` | Arctic blue-gray (Nord-inspired) |
| `coolnight-dracula` | Rich purple (Dracula-inspired) |
| `coolnight-solarized` | Scientific precision (Solarized-inspired) |

---

## Popular Themes (13+)

| Theme | Style | Plugin |
|-------|-------|--------|
| `tokyonight-night` | dark | folke/tokyonight.nvim |
| `tokyonight-storm` | dark | folke/tokyonight.nvim |
| `tokyonight-day` | light | folke/tokyonight.nvim |
| `catppuccin-mocha` | dark | catppuccin/nvim |
| `catppuccin-latte` | light | catppuccin/nvim |
| `catppuccin-frappe` | dark | catppuccin/nvim |
| `catppuccin-macchiato` | dark | catppuccin/nvim |
| `gruvbox-dark` | dark | ellisonleao/gruvbox.nvim |
| `gruvbox-light` | light | ellisonleao/gruvbox.nvim |
| `nord` | dark | shaunsingh/nord.nvim |
| `dracula` | dark | Mofiqul/dracula.nvim |
| `one-dark` | dark | navarasu/onedark.nvim |
| `solarized-dark` | dark | ishan9299/nvim-solarized-lua |

---

## Installing a Library Theme as User Theme

Library themes are read-only. To customize a library theme, export it to YAML, modify it, and apply it under a new name:

```bash
# Export library theme to YAML
nvp theme get tokyonight-night -o yaml > my-tokyonight.yaml

# Edit my-tokyonight.yaml — change name and customize colors

# Apply your custom version
nvp apply -f my-tokyonight.yaml
nvp theme use my-tokyonight
nvp generate
```

---

## Theme Library Commands

```bash
# List available themes in the remote library
nvp theme library list
nvp theme library list --category dark

# Show details of a library theme
nvp theme library show catppuccin-mocha

# Install from the library (saves to ~/.nvp/themes/)
nvp theme library install catppuccin-mocha
nvp theme library install catppuccin-mocha --use

# List theme categories
nvp theme library categories

# List theme tags
nvp theme library tags
```

---

## Next Steps

- [CoolNight Collection](coolnight.md) - All 21 CoolNight variants with usage guide
- [Parametric Generator](parametric.md) - Create custom CoolNight variants
- [NvimTheme YAML Reference](../reference/nvim-theme.md) - Full YAML schema
- [Commands Reference](../commands.md) - Full command reference
