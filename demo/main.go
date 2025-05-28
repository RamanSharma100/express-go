package main

import (
	"fmt"
	"strconv"

	"github.com/ramansharma100/express-go/demo/routes"
	"github.com/ramansharma100/express-go/http"
)

func main() {
	app := http.New()

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware 1")
		next()
	})

	app.Get("/", func(ctx *http.Context) {
		ctx.Request.AddField("user", "Raman Sharma")
		// ctx.Response.Send("Hello, World!" + ctx.Request.AdditionalFields["user"].(string))
		ctx.Render("index.html", map[string]any{
			"user": ctx.Request.AdditionalFields["user"],
		})
	})

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware 2")
		next()
	})

	app.Get("/:id", func(ctx *http.Context) {
		params := ctx.GetParams()
		ctx.Response.Json(
			map[string]any{
				"params": params,
			},
		)
	})

	app.UseRouter("/company", routes.CompanyRouter())

	app.Get("/user/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(
			map[string]any{
				"params": ctx.GetParams(),
			},
		)
	})

	app.Get("/user/:id/:name", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"context": ctx.GetParams()})
	})

	app.Get("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Send("Hello from user")
	})

	app.Post("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		fmt.Println("Body:", ctx.GetBody())
		ctx.Response.Status(200).Json(map[string]any{"body": ctx.GetBody()})
	})

	app.Put("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		fmt.Println("Body:", ctx.GetBody())
		ctx.Response.Status(200).Json(map[string]any{"body": ctx.Request.GetJsonBody()})
	})

	app.Patch("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		fmt.Println("Body:", ctx.GetBody())
		ctx.Response.Status(200).Json(map[string]any{"body": ctx.Request.GetJsonBody()})
	})

	app.Listen(8000, func(port int, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	})
}
