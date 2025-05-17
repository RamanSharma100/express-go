package http

import (
	"net/http"
)

func (s *Server) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	for _, route := range s.Routes[r.Method] {
		if route.Path == r.URL.Path {
			ctx := &Context{
				Request:  NewRequest(r),
				Response: NewResponse(w),
			}
			route.Handler(ctx)
			return
		}
	}

	w.Write([]byte("404 Not Found"))
}
