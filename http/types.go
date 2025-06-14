package http

import "net/http"

type Request struct {
	r                *http.Request
	Method           string
	Url              string
	Headers          map[string]string
	Body             any
	AdditionalFields map[string]any
}

type Response struct {
	Headers    map[string]string
	Writer     http.ResponseWriter
	StatusCode int
}

type Context struct {
	Request  *Request
	Response *Response
}

type Handler func(*Context)

type ErrorHandlerType func(*Context, error)

type Middleware func(ctx *Context, next func())

type Server struct {
	Port         int
	Address      string
	Request      *Request
	Response     *Response
	ErrorHandler ErrorHandlerType
	Routes       map[string][]Route
	Middlewares  []Middleware
}

type HTTPMethod func(path string, handler Handler) *RouteChain

type Route struct {
	Method       []string
	Path         string
	Handler      Handler
	Params       []string
	SearchParams map[string]string
	Middlewares  []Middleware
	Name         string
}

type RouteChain struct {
	server *Server
	router *Router
	path   string
	method []string
}
