pkgname=straw
pkgver=0.1.alpha2
pkgrel=1
pkgdesc="CLI tool to easily go-to projects all over your desktop and open a new tmux session in it"
arch=(x86_64)
url="https://github.com/siddhantmadhur/straw"
license=(MIT)
makedepends=(git go)
provides=(straw)
conflicts=(straw)
source=("straw::git+https://github.com/siddhantmadhur/straw.git")
sha256sums=(f4236ac8dd7009d02d9541828be5c748cc976f81a7c280b0efdc32f57eb629a2)

build() {
    cd straw
    make
}

package() {
    cd straw
    install -Dm755 straw -t "${pkgdir}/usr/bin"
    install -Dm644 LICENCE -t "${pkgdir}/usr/share/licenses/${pkgname}"
}
