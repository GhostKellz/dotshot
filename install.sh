#!/bin/bash
set -e

BIN=dotshot
PREFIX=${PREFIX:-/usr/local}
CONFIG_DIR="$HOME/.config/dotshot"

# Build binary
if [ ! -f $BIN ]; then
  go build -o $BIN .
fi

# Install binary
install -Dm755 $BIN "$PREFIX/bin/$BIN"

# Install example config
mkdir -p "$CONFIG_DIR"
cp -n config.yaml "$CONFIG_DIR/config.yaml"

# Install systemd user service
mkdir -p "$HOME/.config/systemd/user"
cp dotshot.service "$HOME/.config/systemd/user/dotshot.service"

cat <<EOF

Installed $BIN to $PREFIX/bin/$BIN
Example config at $CONFIG_DIR/config.yaml
Systemd user service at ~/.config/systemd/user/dotshot.service

To enable background sync:
  systemctl --user daemon-reload
  systemctl --user enable --now dotshot.service

EOF
