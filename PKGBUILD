pkgname=straw
pkgver=0.1.alpha2
pkgrel=1
pkgdesc="Straw is a CLI tool to easily go-to projects all over your desktop and open a new tmux session in it"
arch=(x86_64)
url="https://github.com/siddhantmadhur/straw"
license=(MIT)
makedepends=(git go)
provides=(straw)
conflicts=(straw)
source=("straw::git+https://github.com/siddhantmadhur/straw.git")
sha256sum=('SKIP')

build() {
    cd straw
    make
}

package() {
    cd straw
    install -Dm755 straw -t "${pkgdir}/usr/bin"
    install -Dm644 LICENCE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}