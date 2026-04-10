# NvimOps Sync

The sync subsystem enables `nvp source sync <source>` to import plugin configurations from popular Neovim distributions including LazyVim, AstroNvim, NvChad, and others.

## Supported Sources

- **lazyvim** - LazyVim configuration
- **astronvim** - AstroNvim configuration
- **nvchad** - NvChad configuration
- **kickstart** - Kickstart.nvim configuration
- **lunarvim** - LunarVim configuration
- **local** - Local filesystem source

## Usage

See [nvimops/sync/sources/README.md](sources/README.md) for source-specific usage examples including filtering by category, dry-run mode, and force-overwrite options.

## How It Works

Each sync source fetches plugin definitions from its upstream (GitHub API, local files, etc.), converts them to DevOpsMaestro YAML plugin format, and writes them to the target directory. A sync result reports how many plugins were created or updated and any errors encountered.

## Future Implementation

Actual source handlers for LazyVim, AstroNvim, and others are being implemented progressively. Each handler fetches plugin configurations from the external source, converts them to DevOpsMaestro's YAML format, writes the YAML files to the target directory, and reports sync results.