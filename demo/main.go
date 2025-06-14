package main

import (
	"fmt"
	"strconv"

	"github.com/ramansharma100/express-go/demo/routes"
	"github.com/ramansharma100/express-go/http"
)

func middlewareTest1(ctx *http.Context, next func()) {
	fmt.Println("Middleware Group Test 1")
	next()
}

func middlewareTest2(ctx *http.Context, next func()) {
	fmt.Println("Middleware Group Test 2")
	next()
}

func main() {
	app := http.New()

	app.Static("/static", "/demo/static")

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware Global 1")
		next()
	})

	// rate limiter
	app.Use(http.RateLimit(&http.RateLimitOptions{
		Limit:     10,
		Window:    60, // in seconds
		Remaining: 10,
	}))

	// cors middleware
	app.Use(http.CORS(&http.CorsOptions{
		AllowOrigin: "*",
	}))

	// custom error handler
	app.SetErrorHandler(func(ctx *http.Context, err error) {
		fmt.Println("Error Handler:", err)
		ctx.Response.Status(500).Json(map[string]any{
			"error":   err.Error(),
			"message": "An error occurred from custom error handler",
		})
	})

	app.Get("/", func(ctx *http.Context) {
		ctx.Request.AddField("user", "Raman Sharma")
		http.Logger().Info("Request received for / endpoint")
		// ctx.Response.Send("Hello, World!" + ctx.Request.AdditionalFields["user"].(string))
		ctx.Render("index.html", map[string]any{
			"user": ctx.Request.AdditionalFields["user"],
		})
	}).Name("home")

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware Global 2")
		next()
	})

	app.Get("/error", func(ctx *http.Context) {
		panic("This is a test error")
	})

	// Now define your routes
	app.Post("/upload", func(ctx *http.Context) {
		files, err := ctx.GetUploadedFiles()
		if err != nil {
			ctx.Response.Status(400).Json(map[string]any{"error": err.Error()})
			return
		}
		ctx.Response.Status(200).Json(map[string]any{"files": files})
	})

	// session management routes
	app.Get("/session-form", func(ctx *http.Context) {
		ctx.Render("session-form.html", nil)
	})

	app.Post("/set-session", func(ctx *http.Context) {
		body := ctx.Request.GetBody()

		m, ok := body.(map[string]any)
		if !ok {
			ctx.Response.Status(400).Send("Invalid JSON body")
			return
		}
		username, ok := m["username"].(string)
		if !ok || username == "" {
			ctx.Response.Status(400).Send("Username must be a non-empty string")
			return
		}
		ctx.SetSessionData("username", username)
		ctx.Response.Headers["Content-Type"] = "text/html"
		ctx.Response.Send("Session set! <a href=\"/get-session\">Check session value</a>")
	})

	app.Get("/get-session", func(ctx *http.Context) {
		val, ok := ctx.GetSessionData("username")
		if ok {
			ctx.Response.Send("Session username: " + val.(string) + `<br><a href="/session-form">Back</a>`)
		} else {
			ctx.Response.Send("No session value set for 'username'. <a href=\"/session-form\">Back</a>")
		}
	})

	// use groups to add
	app.Group("/test", []http.Middleware{middlewareTest1, middlewareTest2}, func(router *http.Router) {
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

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware Global 3")
		next()
	})

	app.Get("/:id", func(ctx *http.Context) {
		params := ctx.GetParams()
		ctx.Response.Json(
			map[string]any{
				"params":     params,
				"queryPrams": ctx.GetSearchParams(),
			},
		)
	}).Name("getById")

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

	app.Patch("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		body := ctx.Request.GetJsonBody()
		// validate the request body
		errors := ctx.Request.Validate(map[string]string{
			"name":  "required|string",
			"email": "required|email",
		}, body)

		if len(errors) > 0 {
			ctx.Response.Status(400).Json(map[string]any{"errors": errors})
			return
		}

		ctx.Response.Status(200).Json(map[string]any{"body": body})
	})

	app.Listen(8000, func(port int, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	})
}
