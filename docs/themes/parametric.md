# Parametric Theme Generator

The parametric generator creates custom CoolNight theme variants from a hue angle, hex color, or preset name. Use it to create a theme that matches your exact color preference.

---

## How It Works

The generator derives a complete CoolNight color palette from a single input:

- **Hue angle (0–360)** - Position on the color wheel
- **Hex color (`#rrggbb`)** - The hue is extracted from the hex value
- **Preset name** - A named preset maps to a specific hue

All syntax colors, UI colors, and semantic roles are derived from that base hue using mathematically consistent relationships. Every generated theme maintains the same readability and contrast characteristics as the built-in CoolNight variants.

---

## Usage

```bash
nvp theme create --from <value> --name <name> [--use]
```

### By Hue Angle

```bash
# Create by hue (0-360 degrees)
nvp theme create --from "210" --name my-blue-theme
nvp theme create --from "350" --name my-rose-theme
nvp theme create --from "120" --name my-green-theme
nvp theme create --from "280" --name my-purple-theme

# Create and activate immediately
nvp theme create --from "210" --name my-blue-theme --use
```

### By Hex Color

The hue is extracted from the hex color automatically:

```bash
nvp theme create --from "#8B00FF" --name my-violet-theme
nvp theme create --from "#0a7fa8" --name my-teal-theme
nvp theme create --from "#e05c00" --name my-orange-theme
```

### By Preset Name

```bash
nvp theme create --from "synthwave" --name my-synth
nvp theme create --from "ocean" --name ocean-custom
nvp theme create --from "forest" --name my-forest
```

---

## Hue Reference

| Hue Range | Color Family | Example Themes |
|-----------|--------------|----------------|
| 0°–30° | Red to Orange | crimson, ember, sunset |
| 30°–90° | Orange to Yellow | gold, warm yellows |
| 90°–150° | Yellow to Green | forest, matrix, mint |
| 150°–210° | Green to Blue | teal, arctic, ocean |
| 210°–270° | Blue to Purple | midnight, violet |
| 270°–330° | Purple to Pink | synthwave, grape, sakura |
| 330°–360° | Pink to Red | rose, back to crimson |

### Examples by Target Color

```bash
# Teal
nvp theme create --from "165" --name coolnight-teal

# Lime green
nvp theme create --from "75" --name coolnight-lime

# Hot magenta
nvp theme create --from "315" --name coolnight-magenta

# Custom purple from hex
nvp theme create --from "#8B00FF" --name coolnight-violet
```

---

## Flags

| Flag | Required | Description |
|------|----------|-------------|
| `--from <value>` | Yes | Base: hue angle (0–360), hex color (`#rrggbb`), or preset name |
| `--name <name>` | Yes | Name for the generated theme |
| `--use` | No | Set the new theme as active after creation |

---

## After Creating a Theme

```bash
# View the generated theme
nvp theme get my-blue-theme -o yaml

# Use it
nvp theme use my-blue-theme

# Generate Lua files
nvp generate
```

---

## Generated Themes vs Library Themes

| Aspect | Generated Theme | Library Theme |
|--------|----------------|---------------|
| Location | `~/.nvp/themes/` | Embedded in binary |
| Editable | Yes | No (read-only) |
| Access | After creation | Immediate |
| Overrides library | Yes, if same name | — |

---

## Color Palette Structure

Every generated theme uses this semantic mapping:

| Semantic Role | Derived From |
|---------------|-------------|
| `accent` | Primary hue |
| Syntax colors | Hue variations (±30°, ±60°) |
| UI colors | Desaturated primary hue |
| `error` | Red (~0°) regardless of theme |
| `warning` | Yellow (~45°) regardless of theme |

This ensures consistent readability and meaning across all generated variants.

---

## Troubleshooting

### Custom variant not applied

```bash
# Verify the theme was created
nvp theme list | grep my-blue-theme

# Check theme details
nvp theme get my-blue-theme -o yaml

# Regenerate Lua files
nvp generate
```

### Colors don't look right

```bash
# Verify true color support
echo $COLORTERM   # Should print: truecolor

# Preview in terminal before applying
nvp theme preview my-blue-theme
```

---

## Next Steps

- [CoolNight Collection](coolnight.md) - All 21 built-in CoolNight variants
- [Theme Library](library.md) - Complete library of 34+ themes
- [NvimTheme YAML Reference](../reference/nvim-theme.md) - Full YAML schema
- [Commands Reference](../commands.md) - Full command reference
