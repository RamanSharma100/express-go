package http

import (
	"encoding/json"
	"net/http"
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
