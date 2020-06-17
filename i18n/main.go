package main

import (
	"clevergo.tech/clevergo"
	"clevergo.tech/i18n"
)

var (
	translators *i18n.Translators
)

func index(ctx *clevergo.Context) error {
	translator := i18n.GetTranslator(ctx.Request)
	translator.Fprintf(ctx.Response, "%m", "home")
	return nil
}

func hello(ctx *clevergo.Context) error {
	translator := i18n.GetTranslator(ctx.Request)
	name := ctx.Request.URL.Query().Get("name")
	translator.Fprintf(ctx.Response, "hello %s", name)
	return nil
}

func main() {
	translators = i18n.New(
	// i18n.Fallback("en"), // fallback language, default to English.
	)
	store := i18n.NewFileStore("./translations", i18n.JSONFileDecoder{})
	err := translators.Import(store)
	if err != nil {
		panic(err)
	}

	app := clevergo.New()
	app.Get("/", index)
	app.Get("/hello", hello)
	parsers := []i18n.LanguageParser{
		i18n.NewURLLanguageParser("lang"),    // from URL query
		i18n.NewCookieLanguageParser("lang"), // from cookie
		i18n.HeaderLanguageParser{},          // from Accept-Language header
	}
	app.Use(
		clevergo.WrapHH(i18n.Middleware(translators, parsers...)),
	)
	app.Run(":8080")
}
