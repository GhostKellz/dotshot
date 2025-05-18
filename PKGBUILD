# Maintainer: Your Name <youremail@example.com>

pkgname=dotshot
gitname=dotshot
giturl="https://github.com/ghostkellz/dotshot.git"
pkgver=0.1.0
pkgrel=1
pkgdesc="Lightweight dotfile snapshot & sync tool for Arch-based systems"
arch=('x86_64')
url="https://github.com/ghostkellz/dotshot"
license=('MIT')
depends=('go' 'git')
makedepends=('go')
source=("$pkgname::git+$giturl")
md5sums=('SKIP')

build() {
  cd "$srcdir/$pkgname"
  go build -o dotshot .
}

package() {
  cd "$srcdir/$pkgname"
  install -Dm755 dotshot "$pkgdir/usr/bin/dotshot"
  install -Dm644 dotshot.service "$pkgdir/usr/lib/systemd/user/dotshot.service"
  install -Dm644 config.yaml "$pkgdir/usr/share/dotshot/config.yaml.example"
}
