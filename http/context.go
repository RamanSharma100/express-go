package http

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/ramansharma100/express-go/utils"
)

func (ctx *Context) ParseBody() {
	if ctx.Request.r.Method == "POST" {
		err := ctx.Request.r.ParseForm()
		if err != nil {
			return
		}
		ctx.Request.Body = ctx.Request.r.FormValue("body")
	}
}

func (ctx *Context) GetParams() map[string]string {
	return ctx.Request.AdditionalFields["params"].(map[string]string)
}

func (ctx *Context) GetParam(name string) string {
	return ctx.Request.AdditionalFields["params"].(map[string]string)[name]
}

func (ctx *Context) GetHeader(name string) string {
	return ctx.Request.Headers[name]
}

func (ctx *Context) GetHeaders() map[string]string {
	return ctx.Request.Headers
}

func (ctx *Context) GetBody() any {
	return ctx.Request.GetBody()
}

func (ctx *Context) GetMethod() string {
	return ctx.Request.Method
}

func (ctx *Context) GetUrl() string {
	return ctx.Request.Url
}

func (ctx *Context) GetWriter() *http.ResponseWriter {
	return &ctx.Response.Writer
}

func (ctx *Context) GetStatusCode() int {
	return ctx.Response.StatusCode
}

func (ctx *Context) SetStatusCode(code int) {
	ctx.Response.StatusCode = code
	ctx.Response.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Response.Headers[key] = value
	ctx.Response.Writer.Header().Set(key, value)
}

func (ctx *Context) SetHeaders(headers map[string]string) {
	for key, value := range headers {
		ctx.Response.Headers[key] = value
		ctx.Response.Writer.Header().Set(key, value)
	}
}

func (ctx *Context) Body(body string) {
	ctx.Response.Writer.Write([]byte(body))
}

func (ctx *Context) Json(data any) {
	ctx.Response.Writer.Header().Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(ctx.Response.Writer, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	ctx.Response.Writer.Write(jsonBytes)
}

func (ctx *Context) Send(text string) {
	ctx.Response.Writer.Header().Set("Content-Type", "text/plain")
	ctx.Response.Writer.Write([]byte(text))
}

func (ctx *Context) Status(code int) {
	ctx.Response.StatusCode = code
	ctx.Response.Writer.WriteHeader(code)
}

func (ctx *Context) Render(tmpl string, data any) {
	rootDir := utils.GetRootDirectory()
	fmt.Println("Root Directory:", rootDir)
	filePath := path.Join(rootDir, "templates", tmpl)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		ctx.Response.Writer.WriteHeader(http.StatusInternalServerError)
		ctx.Response.Writer.Write([]byte("Error parsing template: " + err.Error()))
		return
	}

	if data == nil {
		data = map[string]any{}
	}

	ctx.Response.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx.Response.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Response.Writer.Header().Set("Pragma", "no-cache")
	ctx.Response.Writer.Header().Set("Expires", "0")

	err = t.Execute(ctx.Response.Writer, data)
	if err != nil {
		ctx.Response.Writer.WriteHeader(http.StatusInternalServerError)
		ctx.Response.Writer.Write([]byte("Error executing template: " + err.Error()))
		return
	}

}

func (ctx *Context) Redirect(url string) {
	ctx.Response.Writer.Header().Set("Location", url)
	ctx.Response.Writer.WriteHeader(http.StatusFound)
}

func (ctx *Context) GetSearchParams() map[string]string {
	params := make(map[string]string)
	query := ctx.Request.r.URL.Query()
	for key, values := range query {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}
	return params
}

func (ctx *Context) GetSearchParam(name string) string {
	query := ctx.Request.r.URL.Query()
	if values, ok := query[name]; ok && len(values) > 0 {
		return values[0]
	}
	return ""
}
