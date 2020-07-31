package main

import (
	"io"
	"reflect"
	"strings"

	"clevergo.tech/clevergo"
	"clevergo.tech/jetpackr"
	"clevergo.tech/jetrenderer"
	"github.com/CloudyKit/jet/v5"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	box := packr.New("views", "./views")
	set := jet.NewHTMLSetLoader(jetpackr.New(box))
	set.SetDevelopmentMode(true) // debug
	renderer := jetrenderer.New(set)
	renderer.SetBeforeRender(func(w io.Writer, name string, vars jet.VarMap, data interface{}, ctx *clevergo.Context) error {
		vars.SetFunc("title", jet.Func(func(args jet.Arguments) reflect.Value {
			args.RequireNumOfArguments("title", 1, 1)
			return reflect.ValueOf(strings.Title(args.Get(0).String()))
		}))

		return nil
	})

	app := clevergo.New()
	app.Renderer = renderer
	app.Get("/", func(ctx *clevergo.Context) error {
		return ctx.Render(200, "index.tmpl", map[string]interface{}{
			"message": "hello world",
		})
	})
	app.Run(":8080")
}
