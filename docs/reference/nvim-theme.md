# NvimTheme YAML Reference

**Kind:** `NvimTheme`
**APIVersion:** `devopsmaestro.io/v1`

An NvimTheme defines a Neovim colorscheme configuration in YAML. Apply with `nvp apply -f <file>` and export with `nvp theme get <name> -o yaml`.

---

## Full Example

```yaml
apiVersion: devopsmaestro.io/v1
kind: NvimTheme
metadata:
  name: coolnight-synthwave
  description: "CoolNight Synthwave - Retro neon vibes with deep purples and electric blues"
  author: "devopsmaestro"
  category: "dark"
  labels:
    collection: coolnight
    style: synthwave
    brightness: dark
  annotations:
    version: "1.0.0"
    last-updated: "2026-02-19"
spec:
  plugin:
    repo: "rmkohlman/coolnight.nvim"
    branch: "main"
    tag: "v1.0.0"
  style: "synthwave"
  transparent: false
  colors:
    bg: "#0a0a0a"
    fg: "#e1e1e6"
    primary: "#bd93f9"
    secondary: "#ff79c6"
    accent: "#8be9fd"
    error: "#ff5555"
    warning: "#f1fa8c"
    info: "#8be9fd"
    hint: "#50fa7b"
    selection: "#44475a"
    comment: "#6272a4"
    cursor: "#f8f8f2"
    line_number: "#6272a4"
    line_highlight: "#282a36"
    popup_bg: "#282a36"
    popup_border: "#6272a4"
    statusline_bg: "#44475a"
    tabline_bg: "#282a36"
  options:
    italic_comments: true
    bold_keywords: false
    underline_errors: true
    transparent_background: false
    custom_highlights:
      - group: "Keyword"
        style: "bold"
        fg: "#bd93f9"
      - group: "String"
        style: "italic"
        fg: "#f1fa8c"
      - group: "Function"
        style: "bold"
        fg: "#50fa7b"
```

---

## Field Reference

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `apiVersion` | string | Yes | Must be `devopsmaestro.io/v1` |
| `kind` | string | Yes | Must be `NvimTheme` |
| `metadata.name` | string | Yes | Unique identifier for the theme |
| `metadata.description` | string | No | Theme description |
| `metadata.author` | string | No | Theme author |
| `metadata.category` | string | No | Theme category: `dark`, `light`, `monochrome` |
| `metadata.labels` | object | No | Key-value labels for organization |
| `metadata.annotations` | object | No | Key-value annotations |
| `spec.plugin` | object | Yes | Plugin repository information |
| `spec.plugin.repo` | string | Yes | GitHub repository (`owner/repo`) |
| `spec.plugin.branch` | string | No | Git branch |
| `spec.plugin.tag` | string | No | Git tag or version |
| `spec.style` | string | No | Theme variant (plugin-specific) |
| `spec.transparent` | boolean | No | Enable transparent background |
| `spec.colors` | object | No | Color overrides |
| `spec.options` | object | No | Plugin-specific options |

---

## Field Details

### metadata.name (required)

Unique identifier for the theme.

**Conventions:**
- Use kebab-case: `coolnight-synthwave`
- Include the collection name: `coolnight-ocean`, `tokyonight-night`

### metadata.category (optional)

Theme category for organization.

**Valid values:**
- `dark` - Dark background theme
- `light` - Light background theme
- `monochrome` - Black and white theme

### spec.plugin (required)

The Neovim plugin that provides the colorscheme.

```yaml
spec:
  plugin:
    repo: "folke/tokyonight.nvim"      # required
    branch: "main"                     # optional
    tag: "v1.0.0"                     # optional
```

**Popular theme plugins:**
- `folke/tokyonight.nvim`
- `catppuccin/nvim`
- `ellisonleao/gruvbox.nvim`
- `shaunsingh/nord.nvim`
- `Mofiqul/dracula.nvim`

### spec.style (optional)

Theme variant name. Values are plugin-specific.

```yaml
# Tokyo Night
spec:
  style: "night"       # night, storm, day, moon

# Catppuccin
spec:
  style: "mocha"       # mocha, macchiato, frappe, latte

# Gruvbox
spec:
  style: "dark"        # dark, light
```

### spec.transparent (optional)

Enable a transparent background for terminal integration.

```yaml
spec:
  transparent: true
```

### spec.colors (optional)

Override semantic color values. All colors must be valid hex (`#rrggbb`).

```yaml
spec:
  colors:
    bg: "#1a1b26"           # Background
    fg: "#c0caf5"           # Foreground
    primary: "#7aa2f7"      # Primary accent
    secondary: "#bb9af7"    # Secondary accent
    accent: "#7dcfff"       # Tertiary accent
    error: "#f7768e"        # Error messages
    warning: "#e0af68"      # Warning messages
    info: "#7dcfff"         # Info messages
    hint: "#1abc9c"         # Hint messages
    selection: "#33467c"    # Selection highlight
    comment: "#565f89"      # Comments
    cursor: "#c0caf5"       # Cursor color
    line_number: "#3b4261"  # Line numbers
    line_highlight: "#1f2335"
    popup_bg: "#1f2335"
    popup_border: "#27a1b9"
    statusline_bg: "#1f2335"
    tabline_bg: "#1a1b26"
```

### spec.options (optional)

Plugin-specific options and custom highlight groups.

```yaml
spec:
  options:
    italic_comments: true
    bold_keywords: false
    underline_errors: true
    transparent_background: false
    dim_inactive: false
    custom_highlights:
      - group: "Keyword"
        style: "bold"         # bold, italic, underline
        fg: "#bd93f9"
        bg: "#282a36"         # optional
      - group: "String"
        style: "italic"
        fg: "#f1fa8c"
    integrations:
      telescope: true
      gitsigns: true
      lualine: true
```

---

## Theme Collections

### CoolNight Variants (built-in)

All 21 CoolNight themes are embedded and immediately available:

```bash
nvp theme use coolnight-ocean       # 210° blue
nvp theme use coolnight-synthwave   # 280° purple
nvp theme use coolnight-matrix      # 120° green
```

See [CoolNight Collection](../themes/coolnight.md) for all 21 variants.

### Popular Themes (built-in)

```yaml
# Tokyo Night
apiVersion: devopsmaestro.io/v1
kind: NvimTheme
metadata:
  name: tokyonight-night
  category: dark
spec:
  plugin:
    repo: "folke/tokyonight.nvim"
  style: "night"

# Catppuccin
apiVersion: devopsmaestro.io/v1
kind: NvimTheme
metadata:
  name: catppuccin-mocha
  category: dark
spec:
  plugin:
    repo: "catppuccin/nvim"
  style: "mocha"
```

---

## Usage Examples

```bash
# Apply a custom theme
nvp apply -f my-theme.yaml

# Use a theme
nvp theme use my-custom-theme

# Export a theme to YAML
nvp theme get coolnight-ocean -o yaml > ocean-theme.yaml

# Preview a theme in terminal
nvp theme preview coolnight-ocean

# List all themes
nvp theme list

# Generate Lua files with active theme
nvp generate
```

---

## Color Guidelines

Use semantic color names for maintainability:

```yaml
colors:
  primary: "#7aa2f7"      # Good - semantic name
  error: "#f7768e"        # Good - semantic name
```

Ensure sufficient contrast:
- Background and foreground should meet WCAG AA (4.5:1 minimum)
- Keep semantic colors consistent: errors always red tones, warnings always yellow

---

## Validation Rules

- `metadata.name` must be unique across user themes
- `metadata.category` must be `dark`, `light`, or `monochrome`
- `spec.plugin.repo` must be in `owner/repo` format
- All `spec.colors.*` values must be valid hex colors (`#rrggbb` or `#rrggbbaa`)

---

## Related

- [Themes Overview](../themes/overview.md)
- [Theme Library](../themes/library.md)
- [CoolNight Collection](../themes/coolnight.md)
- [Parametric Generator](../themes/parametric.md)
- [NvimPlugin Reference](nvim-plugin.md)
