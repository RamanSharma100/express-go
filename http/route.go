package http

import (
	"fmt"
	"strings"
)

func (s *Server) ValidateRoute(path string, handler Handler) bool {
	if path == "" {
		panic("Path cannot be empty")
	}
	if handler == nil {
		panic("Handler cannot be empty")
	}

	return true

}

func (s *Server) getParameterizedRoute(path string) (string, []string) {
	Params := []string{}
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			paramName := strings.TrimPrefix(part, ":")
			Params = append(Params, paramName)
			parts[i] = "{" + paramName + "}"
		}
	}

	return strings.Join(parts, "/"), Params
}

func isParameterizedRoute(path string) bool {
	return strings.ContainsAny(path, ":")
}

func sortRoutesWithParamsLast(routes []Route) []Route {
	for i := 0; i < len(routes); i++ {
		for j := i + 1; j < len(routes); j++ {
			if len(routes[i].Params) > 0 && len(routes[j].Params) == 0 {
				routes[i], routes[j] = routes[j], routes[i]
			}
		}
	}
	return routes
}

func (s *Server) AddRoute(path string, handler Handler, method []string) {
	if s.ValidateRoute(path, handler) {
		for _, m := range method {
			if _, ok := s.Routes[m]; !ok {
				s.Routes[m] = []Route{}
			}

			Params := []string{}

			fmt.Printf("🚀 %s Route loaded\n", path)

			if isParameterizedRoute(path) {
				path, Params = s.getParameterizedRoute(path)
			}

			s.Routes[m] = append(s.Routes[m], Route{
				Method:      method,
				Path:        path,
				Handler:     handler,
				Params:      Params,
				Middlewares: append([]Middleware{}, s.Middlewares...),
			})
			s.Routes[m] = sortRoutesWithParamsLast(s.Routes[m])
		}
	}
}

func (s *Server) AddRouteWithRouter(path string, router *Router) {
	if router == nil {
		return
	}

	if path == "" || path[0] != '/' {
		path = "/" + path
	}

	for _, route := range router.routes {
		if s.ValidateRoute(path, route.Handler) {
			for _, m := range route.Method {
				if _, ok := s.Routes[m]; !ok {
					s.Routes[m] = []Route{}
				}

				Params := []string{}

				fmt.Printf("🚀 %s Route loaded\n", path)

				if isParameterizedRoute(path) {
					path, Params = s.getParameterizedRoute(path)
				}

				middlewares := s.Middlewares

				s.Routes[m] = append(s.Routes[m], Route{
					Method:      route.Method,
					Path:        path,
					Handler:     route.Handler,
					Params:      Params,
					Middlewares: append(middlewares, route.Middlewares...),
				})
				s.Routes[m] = sortRoutesWithParamsLast(s.Routes[m])
			}
		}
	}
}

func (s *Server) Get(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"GET"})
}

func (s *Server) Post(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"POST"})
}

func (s *Server) Put(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"PUT"})
}

func (s *Server) Delete(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"DELETE"})
}

func (s *Server) Patch(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"PATCH"})
}

func (s *Server) Options(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"OPTIONS"})
}

func (s *Server) Head(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"HEAD"})
}

func (s *Server) Add(path string, handler Handler) {
	s.AddRoute(path, handler, []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"})
}

func (s *Server) Use(middlewares ...Middleware) {
	if middlewares == nil {
		panic("Middleware cannot be nil")
	}
	s.Middlewares = append(s.Middlewares, middlewares...)
}

func (s *Server) UseRouter(path string, router *Router) {
	if router == nil {
		return
	}

	if path == "" || path[0] != '/' {
		path = "/" + path
	}

	for _, route := range router.routes {
		if route.Path != "/" && route.Path[0] != '/' {
			route.Path = "/" + route.Path
		}
		fullPath := path + route.Path

		s.AddRoute(fullPath, route.Handler, route.Method)
	}
}
