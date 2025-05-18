# dotshot

[![Go Build](https://github.com/ghostkellz/dotshot/actions/workflows/go.yml/badge.svg)](https://github.com/ghostkellz/dotshot/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghostkellz/dotshot)](https://goreportcard.com/report/github.com/ghostkellz/dotshot)
[![Neovim](https://img.shields.io/badge/Editor-Neovim-brightgreen?logo=neovim)](https://neovim.io)
[![Zsh](https://img.shields.io/badge/Shell-Zsh-blue?logo=gnu-bash)](https://www.zsh.org)
[![Starship](https://img.shields.io/badge/Prompt-Starship-yellow?logo=starship)](https://starship.rs)

> Lightweight dotfile snapshot & sync tool for Arch-based systems (or anyone who needs Git-tracked config backups)

---

### 🧠 What is `dotshot`?

**dotshot** is a small, focused CLI app for syncing and snapshotting dotfiles, like:

- Neovim configs
- Zsh (`~/.zshrc`, `~/.config/zsh/`)
- WezTerm themes and settings
- Any other custom config files you want version-controlled

It helps keep your system's critical user-level config files in sync with a tracked Git repo — ideal for power users managing dotfiles manually or through a central `~/arch` or `~/dotfiles` repo.

---

### ⚡ Features

- 🗂️ Track and sync dotfiles to a defined repo structure
- 🔍 Dry-run mode for safe previews
- 🧱 Optional Git commit automation
- 🧭 Designed for integration into larger tools like `ghostctl`
- 🛠 Configurable via TOML/YAML (via [Viper](https://github.com/spf13/viper))

---

### 📦 Installation

```bash
git clone https://github.com/ghostkellz/dotshot.git
cd dotshot
go build -o dotshot .
```

---

### 🚀 Usage

```bash
# One-time setup
./dotshot init

# Add paths to sync
./dotshot add nvim ~/.config/nvim
./dotshot add zsh ~/.config/zsh ~/.zshrc

# Sync current config state to repo
./dotshot sync

# Show what would sync
./dotshot sync --dry-run

# Commit and push changes
./dotshot commit -m "Update nvim + wezterm"
```

---

### 🔧 Planned Features

- `watch` mode for auto-tracking changes
- Automatic tagging/versioning
- Integration with `ghostctl` as a subcommand
- System snapshot metadata (hostname, kernel, date)

