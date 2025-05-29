package routes

import (
	"fmt"

	"github.com/ramansharma100/express-go/http"
)

func middlewareTest1(ctx *http.Context, next func()) {
	fmt.Println("Middleware Router Group Test 1")
	next()
}

func middlewareTest2(ctx *http.Context, next func()) {
	fmt.Println("Middleware Router Group Test 2")
	next()
}

func CompanyRouter() *http.Router {
	router := http.NewRouter()

	router.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware  Router Global 2")
		next()
	})

	// // use groups to add
	router.Group("/test", []http.Middleware{middlewareTest1, middlewareTest2}, func(router *http.Router) {
		router.Get("/", func(ctx *http.Context) {
			ctx.Response.AddHeader("Content-Type", "application/json")
			ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
			ctx.Response.Status(200).Json(map[string]any{"message": "Hello from test group!"})
		})

		router.Get("/info", func(ctx *http.Context) {
			ctx.Response.AddHeader("Content-Type", "application/json")
			ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
			ctx.Response.Status(200).Json(map[string]any{"info": "Test group information."})
		})
	})

	router.Get("/", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"message": "Welcome to the company!"})
	})

	router.Get("/info", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"info": "Company information goes here."})
	})

	router.Post("/create", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		body := ctx.GetBody()
		ctx.Response.Status(201).Json(map[string]any{"message": "Company created successfully", "data": body})
	})

	router.Put("/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		params := ctx.GetParams()
		id := params["id"]
		body := ctx.GetBody()
		ctx.Response.Status(200).Json(map[string]any{"message": "Company updated successfully", "id": id, "data": body})
	})

	router.Patch("/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		params := ctx.GetParams()
		id := params["id"]
		body := ctx.GetBody()
		ctx.Response.Status(200).Json(map[string]any{"message": "Company patched successfully", "id": id, "data": body})
	})

	router.Delete("/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		params := ctx.GetParams()
		id := params["id"]
		ctx.Response.Status(200).Json(map[string]any{"message": "Company deleted successfully", "id": id})
	})

	return router
}
