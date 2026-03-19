# Installation

## Requirements

- **Neovim** - nvp generates configuration for Neovim with lazy.nvim
- **Go 1.25+** (only if building from source)

!!! note "nvp works without containers"
    `nvp` is completely standalone. You do not need Docker, dvm, or any container runtime to use it.

---

## Homebrew (Recommended)

The easiest way to install nvp on macOS or Linux:

```bash
# Add the tap
brew tap rmkohlman/tap

# Install the full DevOpsMaestro toolkit (includes dvm, nvp, and dvt)
brew install devopsmaestro

# Or install nvp standalone (NvimOps only)
brew install nvimops

# Verify installation
nvp version
```

---

## From Source

Build from source if you want the latest development version:

```bash
# Clone the repository
git clone https://github.com/rmkohlman/devopsmaestro.git
cd devopsmaestro

# Build nvp
go build -o nvp ./cmd/nvp/

# Install to your PATH
sudo mv nvp /usr/local/bin/

# Verify installation
nvp version
```

---

## From GitHub Releases

Download pre-built binaries from the [Releases page](https://github.com/rmkohlman/devopsmaestro/releases):

```bash
# Download latest release
VERSION=$(curl -s https://api.github.com/repos/rmkohlman/devopsmaestro/releases/latest | grep '"tag_name"' | sed 's/.*"v\(.*\)".*/\1/')
curl -LO "https://github.com/rmkohlman/devopsmaestro/releases/download/v${VERSION}/devopsmaestro_${VERSION}_$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m | sed 's/x86_64/amd64/').tar.gz"

# Extract and install
tar xzf devopsmaestro_*.tar.gz
sudo mv nvp /usr/local/bin/

# Verify
nvp version
```

Available platforms:
- `devopsmaestro_VERSION_darwin_amd64.tar.gz` (macOS Intel)
- `devopsmaestro_VERSION_darwin_arm64.tar.gz` (macOS Apple Silicon)
- `devopsmaestro_VERSION_linux_amd64.tar.gz` (Linux x64)
- `devopsmaestro_VERSION_linux_arm64.tar.gz` (Linux ARM64)

---

## Shell Completion

Enable tab completion for your shell:

=== "Bash"

    ```bash
    # Add to ~/.bashrc
    eval "$(nvp completion bash)"
    ```

=== "Zsh"

    ```bash
    # Add to ~/.zshrc
    eval "$(nvp completion zsh)"
    ```

=== "Fish"

    ```bash
    # Add to ~/.config/fish/config.fish
    nvp completion fish | source
    ```

After installing completions, restart your shell or source your shell configuration file.

---

## Verify Installation

After installation, verify everything is working:

```bash
# Check version
nvp version

# Initialize nvp (creates ~/.nvp/ directory)
nvp init

# List the plugin library
nvp library list
```

!!! success "Installation Complete"
    If all commands run without errors, nvp is ready to use.

---

## Troubleshooting

### Binary Not Found

If `nvp: command not found`:

```bash
# Check if binary is in PATH
which nvp
echo $PATH

# Add to PATH if needed (replace with your install location)
echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

### Permission Issues

If you get permission errors when installing to `/usr/local/bin`:

```bash
# Option 1: Use sudo
sudo mv nvp /usr/local/bin/

# Option 2: Install to user directory
mkdir -p ~/bin
mv nvp ~/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

---

## Next Steps

- [Quick Start Guide](quickstart.md) - Get up and running in 5 minutes
- [Plugin Library](../plugins/library.md) - Browse 38+ curated plugins
- [Theme Library](../themes/library.md) - Browse 34+ embedded themes
