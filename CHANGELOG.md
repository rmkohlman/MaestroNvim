# Changelog

All notable changes to MaestroNvim will be documented in this file.

## [v0.2.9] - 2026-04-23

### Documentation

- Documented `spec.inherits` field on NvimTheme reference page (resolves rmkohlman/devopsmaestro#429)

## [v0.2.8] - 2026-04-16

### Bug Fixes

- Fix Neovim clipboard error inside container workspaces (#381) — clipboard setting is now wrapped in a conditional provider check so Neovim starts cleanly when no clipboard provider (xclip, xsel, pbcopy, wl-copy) is available
