package http

import "net/http"

type Request struct {
	r                *http.Request
	Method           string
	Url              string
	Headers          map[string]string
	Body             string
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

type Server struct {
	Port         int
	Address      string
	Routes       map[string][]Route
	ErrorHandler ErrorHandlerType
}

type HTTPMethod func(path string, handler Handler)

type Route struct {
	Method  []string
	Path    string
	Handler Handler
}
type Router struct {
	Routes map[string][]Route
}
