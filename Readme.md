# express-go [Work in Progress]

[![GitHub](https://img.shields.io/github/license/ramansharma100/express-go)](LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/ramansharma100/express-go)](https://github.com/ramansharma100/express-go/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/ramansharma100/express-go)](https://goreportcard.com/report/github.com/ramansharma100/express-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/ramansharma100/express-go)](https://pkg.go.dev/github.com/ramansharma100/express-go)

A simple and lightweight HTTP server framework for Go, inspired by Node.js Express. It provides a minimalistic API for building web applications and APIs with ease.

> - This is for educational purposes and is not intended for production use. It is a work in progress and may not be fully functional or secure.
> - The API is inspired by Express.js, but it is not a direct port. Some features may be different or missing.

## Features

- Simple routing (GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD)
- Context-based request and response handling [More work needed]
- Easy to use API inspired by Node.js Express
- Add custom headers to requests and responses
- Add additional fields to requests
- Easily extendable and lightweight
- JSON response support
- Plain text response support
- JSON parsing support
- Parameterized routes

## Upcoming Features

This is lot of work in progress and will be updated frequently. Some of the upcoming features include:

- Nested routing
- Rendering HTML templates
- Support for query parameters
- Support for URL parameters
- Support for cookies
- Middleware support
- Custom Error handling
- Static file serving
- Session management
- Global error handling
- CORS support
- Rate limiting
- Request validation
- File uploads [multer like]
- URL encoding
- WebSocket support
- Support for mixins
- Support for plugins
- Layout support

## Upcoming Extended Features

- CLI Support
- Migrations Support
- AI Studio Support
- Will Add more...

## Installation

```bash
go get github.com/ramansharma100/express-go # may not work yet
```

or

```bash
# or
# clone the repo and run
go build
go install
```

## Usage

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/ramansharma100/express-go/http"
)

func main() {
	app := http.New()

	app.Get("/", func(ctx *http.Context) {
		ctx.Response.Send("Hello, World!")
	})

	app.Listen(8000, func(port int, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	})
}
```

## API

### Application

- `app.Get(path string, handler Handler)`
- `app.Post(path string, handler Handler)`
- `app.Put(path string, handler Handler)`
- `app.Patch(path string, handler Handler)`
- `app.Delete(path string, handler Handler)`
- `app.Options(path string, handler Handler)`
- `app.Listen(port int, callback func(int, error))`

### Handler

A handler receives a `*http.Context`:

```go
func(ctx *http.Context) {
    // ctx.Request, ctx.Response
}
```

### Request

- `ctx.Request.Method` - HTTP method
- `ctx.Request.Url` - URL string
- `ctx.Request.Body` - Request body (for POST)
- `ctx.Request.Headers` - Map of headers
- `ctx.Request.AddHeader(key, value string)` - Add custom header
- `ctx.Request.AddField(key, value string)` - Add custom field to `request.AdditionalFields`
- `ctx.Request.ParseBody()` - Parse request body (for POST)
- `ctx.Request.r` - Get the underlying `http.Request`

### Response

- `ctx.Response.Send(text string)` - Send plain text response
- `ctx.Response.Json(data any)` - Send JSON response
- `ctx.Response.Status(code int)` - Set status code
- `ctx.Response.AddHeader(key, value string)` - Add custom header
- `ctx.Response.Writer` - Get the underlying `http.ResponseWriter`

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

Feel free to open issues for suggestions, bugs, or questions.

## License

This project is licensed under the terms of the [MIT License](LICENSE).
