// A go-app hello-world example, designed to work with GopherJS.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	renderStatic = flag.Bool("render_static", false, "Render a static version of the web site instead of starting a web server.")
)

type gopherJSDriver struct{}

func (d *gopherJSDriver) Scripts() []string {
	return []string{"/web/app.js"}
}

func (d *gopherJSDriver) Styles() []string                      { return nil }
func (d *gopherJSDriver) CacheableResources() []string          { return nil }
func (d *gopherJSDriver) PreRenderItems() []app.PreRenderedItem { return nil }

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!")
}

func serve(h *app.Handler) {
	http.Handle("/", h)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func static(h *app.Handler) {
	if err := app.GenerateStaticWebsite("./static", h); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	app.Route("/", &hello{})
	app.RunWhenOnBrowser()
	h := &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Icon: app.Icon{
			Default: "https://raw.githubusercontent.com/maxence-charriere/go-app/master/docs/web/icon.png",
			Large:   "https://raw.githubusercontent.com/maxence-charriere/go-app/master/docs/web/icon.png",
		},
		Driver: &gopherJSDriver{},
	}

	if *renderStatic {
		static(h)
	} else {
		serve(h)
	}
}
