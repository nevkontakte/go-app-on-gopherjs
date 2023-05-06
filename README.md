This is a quick-and-dirty Go workspace that demos using go-app with gopherjs on the client side. All dependency repositories are defined as git submodules.

You need Go 1.18 installed to be able to use GopherJS:

```
$ go install golang.org/dl/go1.18.10@latest
$ go1.18.10 download
$ export GOPHERJS_GOROOT="$(go1.18.10 env GOROOT)"
```

To play with the demos in this workspace:

```
$ git clone git@github.com:nevkontakte/go-app-on-gopherjs.git
$ cd go-app-on-gopherjs
$ git submodule update --init
$ go install github.com/gopherjs/gopherjs@v1.18.0-beta2
$ cd lofimusic
$ make run
# Open http://localhost:4000/
```
