package main

import (
	"fmt"
	"strconv"

	"github.com/ramansharma100/express-go/http"
)

func main() {
	app := http.New()

	app.Get("/", func(ctx *http.Context) {
		ctx.Request.AddField("user", "Raman Sharma")
		ctx.Response.Send("Hello, World!" + ctx.Request.AdditionalFields["user"].(string))
	})

	app.Listen(8000, func(port int, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	})
}
