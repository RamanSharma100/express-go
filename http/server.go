package http

import (
	"net/http"
	"strconv"
)

func (s *Server) Listen(port int, callback func(int, error)) {

	mux := http.NewServeMux()

	s.Port = port

	addr := s.Address

	if addr == "" {
		addr = "localhost"
	}

	if port == 0 {
		port = 8080
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.HandleRoutes(w, r)
	})

	go func() {
		err := http.ListenAndServe(addr+":"+strconv.Itoa(port), mux)
		if err != nil && callback != nil {
			callback(port, err)
		}
	}()

	if callback != nil {
		callback(port, nil)
	}

	select {} // this will prevent the main goroutine from exiting
}

func basicErrorHandler(ctx *Context, err error) {
	ctx.Response.StatusCode = http.StatusOK
	ctx.Response.Json(map[string]any{
		"error": err.Error(),
	})
}

func CreateServer() *Server {
	return &Server{
		Routes:       make(map[string][]Route),
		ErrorHandler: basicErrorHandler,
		Request: &Request{
			Headers:          make(map[string]string),
			AdditionalFields: make(map[string]any),
			Body:             nil,
		},
		Response: &Response{
			Headers: make(map[string]string),
		},
		Middlewares: []Middleware{
			Logs(&LogOptions{
				Enable: true,
				Format: "{{.Method}} {{.Path}} - {{.StatusCode}}",
			}),
		},
	}
}

func (s *Server) SetErrorHandler(handler ErrorHandlerType) {
	s.ErrorHandler = handler
}
