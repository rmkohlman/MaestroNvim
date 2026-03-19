# Plugin Sources

Plugin sources allow you to configure remote repositories of plugin definitions that `nvp` can sync from. This is useful for team sharing or maintaining a centralized plugin configuration repository.

---

## What Are Sources?

A source is a configured remote location (such as a GitHub repository) that contains plugin YAML files. Once a source is configured, you can sync its plugins into your local nvp store.

---

## Listing Sources

```bash
nvp source list
```

---

## Showing a Source

```bash
nvp source show <name>
```

---

## Syncing from a Source

Pull and apply plugin definitions from a configured source:

```bash
nvp source sync <name>
```

### Sync Flags

| Flag | Description |
|------|-------------|
| `--dry-run` | Preview what would be synced without applying |
| `-l, --selector <key=value>` | Filter resources by label selector |
| `--tag <tag>` | Filter resources by tag |
| `-f, --force` | Overwrite existing plugins |
| `-o, --output <format>` | Output format: `json`, `yaml`, `table` |

**Examples:**

```bash
# Sync all plugins from a source
nvp source sync my-team-plugins

# Preview without applying
nvp source sync my-team-plugins --dry-run

# Sync only plugins with a specific tag
nvp source sync my-team-plugins --tag lsp

# Sync with label selector
nvp source sync my-team-plugins --selector team=backend

# Force overwrite existing plugins
nvp source sync my-team-plugins --force
```

---

## Using GitHub as a Source

The most common pattern is using a GitHub repository as a plugin source. You can apply individual plugin YAMLs from GitHub directly using `nvp apply`:

```bash
# Apply a plugin from a GitHub repository
nvp apply -f github:rmkohlman/nvim-yaml-plugins/plugins/telescope.yaml

# Apply from a specific branch
nvp apply -f https://raw.githubusercontent.com/rmkohlman/nvim-yaml-plugins/main/plugins/telescope.yaml
```

---

## Team Plugin Sharing

A common workflow for teams sharing plugin configurations:

1. Create a repository of plugin YAMLs (e.g., `my-team/nvim-plugins`)
2. Configure it as a source
3. Team members sync from it

```bash
# Team member syncs the shared plugins
nvp source sync team-plugins

# Preview first
nvp source sync team-plugins --dry-run

# Generate Lua files
nvp generate
```

---

## Next Steps

- [Plugin Overview](overview.md) - Plugin management basics
- [Plugin Library](library.md) - Built-in library of 38+ plugins
- [Commands Reference](../commands.md) - Full command reference
