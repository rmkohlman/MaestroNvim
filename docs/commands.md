# Commands Reference

Complete reference for all `nvp` commands.

---

## Global Flags

These flags are available on every command:

| Flag | Description |
|------|-------------|
| `--config <path>` | Path to config file |
| `-v, --verbose` | Enable debug logging |
| `--log-file <path>` | Write logs to file (JSON format) |
| `--no-color` | Disable color output |
| `-h, --help` | Show help for any command |

---

## Initialization

### `nvp init`

Initialize the nvp store.

```bash
nvp init
```

Creates the `~/.nvp/` directory structure.

---

## Plugin Management

### `nvp list`

List all installed plugins.

```bash
nvp list
```

### `nvp get`

Get details of an installed plugin.

```bash
nvp get <name> [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `-o, --output <format>` | Output format: `json`, `yaml`, `table` |

**Examples:**

```bash
nvp get telescope
nvp get telescope -o yaml
nvp get telescope -o json
```

### `nvp enable`

Enable an installed plugin.

```bash
nvp enable <name>
```

**Example:**

```bash
nvp enable telescope
```

### `nvp disable`

Disable an installed plugin without deleting it.

```bash
nvp disable <name>
```

**Example:**

```bash
nvp disable copilot
```

### `nvp delete`

Delete an installed plugin.

```bash
nvp delete <name>
```

**Example:**

```bash
nvp delete telescope
```

---

## Applying Resources

### `nvp apply`

Apply a plugin or theme from a source.

```bash
nvp apply -f <source>
```

**Source types:**

| Type | Example |
|------|---------|
| Local file | `plugin.yaml` |
| URL | `https://example.com/plugin.yaml` |
| GitHub shorthand | `github:user/repo/path/plugin.yaml` |
| Stdin | `-` |

**Examples:**

```bash
# From a local file
nvp apply -f my-plugin.yaml

# From a URL
nvp apply -f https://example.com/plugin.yaml

# From GitHub
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml

# From stdin
cat plugin.yaml | nvp apply -f -
```

---

## Plugin Library

### `nvp library list`

List plugins available in the built-in library.

```bash
nvp library list [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--category <cat>` | Filter by category |
| `--tags <tags>` | Filter by tags |

**Examples:**

```bash
nvp library list
nvp library list --category lsp
nvp library list --category fuzzy-finder
```

### `nvp library show`

Show details of a library plugin.

```bash
nvp library show <name>
```

**Example:**

```bash
nvp library show telescope
```

### `nvp library install`

Install a plugin from the library.

```bash
nvp library install <name>
```

**Examples:**

```bash
nvp library install telescope
nvp library install treesitter
nvp library install lspconfig
```

### `nvp library categories`

List available plugin categories in the library.

```bash
nvp library categories
```

### `nvp library tags`

List available plugin tags in the library.

```bash
nvp library tags
```

---

## Source Management

### `nvp source list`

List configured plugin sources.

```bash
nvp source list
```

### `nvp source show`

Show details of a source.

```bash
nvp source show <name>
```

### `nvp source sync`

Sync plugins from a configured source.

```bash
nvp source sync <name> [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--dry-run` | Preview what would be synced without applying |
| `-l, --selector <key=value>` | Filter by label selector |
| `--tag <tag>` | Filter by tag |
| `-f, --force` | Overwrite existing plugins |
| `-o, --output <format>` | Output format: `json`, `yaml`, `table` |

**Examples:**

```bash
nvp source sync my-team-plugins
nvp source sync my-team-plugins --dry-run
nvp source sync my-team-plugins --tag lsp
nvp source sync my-team-plugins --force
```

---

## Theme Management

### `nvp theme list`

List all available themes (library + user).

```bash
nvp theme list
```

### `nvp theme get`

Get theme details. If no name is provided, shows the active theme.

```bash
nvp theme get [name] [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `-o, --output <format>` | Output format: `json`, `yaml`, `table` |

**Examples:**

```bash
nvp theme get
nvp theme get coolnight-ocean
nvp theme get coolnight-ocean -o yaml
```

### `nvp theme use`

Set the active theme.

```bash
nvp theme use <name>
```

**Examples:**

```bash
nvp theme use coolnight-ocean
nvp theme use catppuccin-mocha
nvp theme use gruvbox-dark
```

### `nvp theme create`

Create a custom CoolNight variant using the parametric generator.

```bash
nvp theme create --from <value> --name <name> [flags]
```

The `--from` value can be:
- A hue angle in degrees (`"210"`)
- A hex color (`"#8B00FF"`)
- A named preset (`"synthwave"`, `"ocean"`, `"forest"`)

**Flags:**

| Flag | Required | Description |
|------|----------|-------------|
| `--from <value>` | Yes | Base: hue (0–360), hex (`#rrggbb`), or preset name |
| `--name <name>` | Yes | Name for the new theme |
| `--use` | No | Set the new theme as active after creation |

**Examples:**

```bash
nvp theme create --from "210" --name my-blue-theme
nvp theme create --from "#8B00FF" --name my-violet-theme
nvp theme create --from "synthwave" --name my-synth --use
```

### `nvp theme preview`

Preview a theme's colors in the terminal.

```bash
nvp theme preview <name>
```

**Example:**

```bash
nvp theme preview coolnight-ocean
nvp theme preview catppuccin-mocha
```

### `nvp theme delete`

Delete a user theme.

```bash
nvp theme delete <name>
```

**Example:**

```bash
nvp theme delete my-custom-theme
```

### `nvp theme generate`

Generate Lua files for the active theme only.

```bash
nvp theme generate
```

---

## Theme Library

### `nvp theme library list`

List themes available in the library.

```bash
nvp theme library list [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--category <cat>` | Filter by category |

**Examples:**

```bash
nvp theme library list
nvp theme library list --category dark
```

### `nvp theme library show`

Show details of a library theme.

```bash
nvp theme library show <name>
```

**Example:**

```bash
nvp theme library show catppuccin-mocha
```

### `nvp theme library install`

Install a theme from the library to `~/.nvp/themes/`.

```bash
nvp theme library install <name> [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `--use` | Set the theme as active after installing |

**Examples:**

```bash
nvp theme library install catppuccin-mocha
nvp theme library install catppuccin-mocha --use
```

### `nvp theme library categories`

List available theme categories.

```bash
nvp theme library categories
```

### `nvp theme library tags`

List available theme tags.

```bash
nvp theme library tags
```

---

## Configuration

### `nvp config init`

Initialize the nvp configuration file.

```bash
nvp config init
```

### `nvp config show`

Show current configuration.

```bash
nvp config show
```

### `nvp config generate`

Generate a default configuration file.

```bash
nvp config generate
```

### `nvp config edit`

Open the configuration file in your editor.

```bash
nvp config edit
```

---

## Generation

### `nvp generate`

Generate Lua files from installed plugins and the active theme.

```bash
nvp generate [flags]
```

**Flags:**

| Flag | Description |
|------|-------------|
| `-o, --output <dir>` | Output directory (default: `~/.config/nvim/lua/plugins/nvp`) |

**Examples:**

```bash
nvp generate
nvp generate --output ~/my-nvim-config/lua/plugins
```

### `nvp generate-lua`

Alias for `nvp generate`.

```bash
nvp generate-lua [flags]
```

---

## Shell Completion

### `nvp completion`

Generate shell completion scripts.

```bash
nvp completion <shell>
```

**Supported shells:** `bash`, `zsh`, `fish`, `powershell`

**Examples:**

=== "Bash"

    ```bash
    nvp completion bash > /etc/bash_completion.d/nvp
    ```

=== "Zsh"

    ```bash
    nvp completion zsh > "${fpath[1]}/_nvp"
    ```

=== "Fish"

    ```bash
    nvp completion fish > ~/.config/fish/completions/nvp.fish
    ```

---

## Version

### `nvp version`

Show version information.

```bash
nvp version
```

---

## Common Workflows

### Install Plugins from Library

```bash
nvp library list
nvp library list --category lsp
nvp library install telescope
nvp library install treesitter
nvp library install lspconfig
nvp generate
```

### Apply from a YAML File

```bash
nvp apply -f my-plugin.yaml
nvp apply -f https://example.com/my-plugin.yaml
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml
nvp generate
```

### Create and Use a Custom Theme

```bash
# By hue angle
nvp theme create --from "280" --name my-synthwave --use

# By hex color
nvp theme create --from "#8B00FF" --name my-violet --use

# Generate Lua files
nvp generate
```

### Switch Themes

```bash
nvp theme use catppuccin-mocha
nvp generate
```

### Install Theme from Library

```bash
nvp theme library list
nvp theme library install catppuccin-mocha --use
nvp generate
```
