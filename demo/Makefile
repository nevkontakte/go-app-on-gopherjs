tools:
	go install -v github.com/gopherjs/gopherjs@v1.18.0-beta2

build: tools
	gopherjs build -v -o web/app.js .
	go build -v -o demo .

run: build
	./demo

static: build
	./demo -render_static
	cp web/app.js static/web/app.js

.PHONY: tools, build, run, static