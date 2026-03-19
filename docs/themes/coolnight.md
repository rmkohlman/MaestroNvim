# CoolNight Theme Collection

The CoolNight Collection is a set of 21 parametrically generated themes designed for consistent, professional color schemes optimized for extended development sessions.

---

## Overview

CoolNight themes are designed with:

- **Reduced eye strain** - Carefully calibrated contrast ratios
- **Consistent syntax highlighting** - Uniform color semantics across all variants
- **Professional appearance** - Suitable for presentations and screen sharing
- **Parametric generation** - Mathematically derived color relationships
- **Wide variety** - 21 variants covering all color preferences

All 21 themes are embedded in the nvp binary and available immediately:

```bash
nvp theme list | grep coolnight   # See all 21 variants
nvp theme use coolnight-ocean     # Use immediately
```

---

## Color Philosophy

### Color Science

CoolNight themes use **HSL color space** for predictable hue relationships:
- Consistent lightness across variants for uniform readability
- Optimal contrast ratios meeting WCAG AA standards (4.5:1 minimum)
- Semantic color mapping — similar code elements use related hues

### Design Principles

1. **Hierarchy** - Different code elements have clear visual weight
2. **Harmony** - All colors work together aesthetically
3. **Function** - Colors convey meaning (errors=red, strings=green, etc.)
4. **Consistency** - Same color rules across all 21 variants

---

## Complete Collection

### Blue Family (Ocean Tones)

| Theme | Hue | Character | Best For |
|-------|-----|-----------|----------|
| `coolnight-arctic` | 190° | Ice blue, crisp and clean | TypeScript, Go, documentation |
| `coolnight-ocean` | 210° | Deep blue, default variant | General development, Python |
| `coolnight-midnight` | 240° | Dark blue, intense focus | Late-night coding, C++ |

```bash
nvp theme use coolnight-ocean
nvp theme use coolnight-arctic
nvp theme use coolnight-midnight
```

### Purple Family (Creative Tones)

| Theme | Hue | Character | Best For |
|-------|-----|-----------|----------|
| `coolnight-violet` | 270° | Soft violet, gentle on eyes | Web development, CSS |
| `coolnight-synthwave` | 280° | Neon purple, retro vibes | JavaScript, creative coding |
| `coolnight-grape` | 290° | Rich grape, sophisticated | Rust, systems programming |

```bash
nvp theme use coolnight-synthwave
nvp theme use coolnight-violet
nvp theme use coolnight-grape
```

### Green Family (Natural Tones)

| Theme | Hue | Character | Best For |
|-------|-----|-----------|----------|
| `coolnight-forest` | 110° | Forest green, earthy | Bash scripts, DevOps |
| `coolnight-matrix` | 120° | Matrix green, high contrast | Terminal work, cybersec |
| `coolnight-mint` | 150° | Fresh mint, modern | React, Vue.js, modern JS |

```bash
nvp theme use coolnight-matrix
nvp theme use coolnight-forest
nvp theme use coolnight-mint
```

### Warm Family (Energy Tones)

| Theme | Hue | Character | Best For |
|-------|-----|-----------|----------|
| `coolnight-ember` | 20° | Glowing ember, energetic | Java, Spring Boot |
| `coolnight-sunset` | 30° | Warm orange, inviting | HTML, markup languages |
| `coolnight-gold` | 45° | Golden yellow, premium | Configuration files, YAML |

```bash
nvp theme use coolnight-sunset
nvp theme use coolnight-ember
nvp theme use coolnight-gold
```

### Red and Pink Family (Passionate Tones)

| Theme | Hue | Character | Best For |
|-------|-----|-----------|----------|
| `coolnight-crimson` | 0° | Deep crimson, bold | Error handling, debugging |
| `coolnight-sakura` | 320° | Cherry blossom, elegant | Design systems |
| `coolnight-rose` | 350° | Rose pink, romantic | Personal projects, blogs |

```bash
nvp theme use coolnight-rose
nvp theme use coolnight-crimson
nvp theme use coolnight-sakura
```

### Monochrome Family (Focus Tones)

| Theme | Character | Best For |
|-------|-----------|----------|
| `coolnight-mono-charcoal` | Charcoal gray, minimalist | Distraction-free coding |
| `coolnight-mono-slate` | Slate gray, professional | Enterprise development |
| `coolnight-mono-warm` | Warm gray, comfortable | Long coding sessions |

```bash
nvp theme use coolnight-mono-slate
nvp theme use coolnight-mono-charcoal
nvp theme use coolnight-mono-warm
```

### Special Variants (Inspired Themes)

| Theme | Inspiration | Character |
|-------|-------------|-----------|
| `coolnight-nord` | Nord theme | Arctic blue-gray |
| `coolnight-dracula` | Dracula theme | Rich purple |
| `coolnight-solarized` | Solarized theme | Scientific precision |

```bash
nvp theme use coolnight-nord
nvp theme use coolnight-dracula
nvp theme use coolnight-solarized
```

---

## Usage Recommendations

### By Development Context

**Terminal-heavy workflows:**
- `coolnight-matrix` - High contrast green
- `coolnight-mono-charcoal` - Minimal distractions

**Web development:**
- `coolnight-mint` - Modern, fresh feel
- `coolnight-synthwave` - Creative, vibrant

**Systems programming:**
- `coolnight-midnight` - Deep focus
- `coolnight-grape` - Sophisticated

**Documentation writing:**
- `coolnight-arctic` - Clean, readable
- `coolnight-mono-warm` - Easy on eyes

**Presentations and screen sharing:**
- `coolnight-ocean` - Professional default
- `coolnight-sunset` - Warm, welcoming

### By Time of Day

**Morning:** `coolnight-arctic`, `coolnight-mint`

**Daytime:** `coolnight-ocean`, `coolnight-forest`

**Evening:** `coolnight-sunset`, `coolnight-ember`

**Late night:** `coolnight-midnight`, `coolnight-mono-slate`

---

## Semantic Color Palette

Every CoolNight theme uses a consistent semantic color mapping:

| Semantic Role | Purpose | Example Elements |
|---------------|---------|------------------|
| `bg` | Background | Editor background, panels |
| `fg` | Foreground | Default text, variables |
| `accent` | Primary accent | Cursor, selection, highlights |
| `comment` | Comments | `// comments`, `# comments` |
| `keyword` | Language keywords | `function`, `class`, `if` |
| `string` | String literals | `"hello"`, `'world'` |
| `function` | Function names | `myFunction()`, method calls |
| `type` | Type annotations | `String`, `int`, class names |
| `constant` | Constants | `true`, `false`, `null`, numbers |
| `error` | Error indicators | Diagnostics, squiggles |
| `warning` | Warning indicators | Warning messages |
| `selection` | Text selection | Selected text background |
| `border` | UI borders | Window borders, splits |

---

## Troubleshooting

### Theme not applying

```bash
# Check if theme exists
nvp theme list | grep coolnight-ocean

# Verify theme content
nvp theme get coolnight-ocean -o yaml

# Regenerate
nvp generate
```

### Colors look wrong

```bash
# Check terminal true color support
echo $COLORTERM   # Should print: truecolor

# Preview the theme in terminal
nvp theme preview coolnight-ocean
```

---

## Next Steps

- [Parametric Generator](parametric.md) - Create custom CoolNight variants
- [Theme Library](library.md) - All 34+ available themes
- [NvimTheme YAML Reference](../reference/nvim-theme.md) - Full YAML schema
