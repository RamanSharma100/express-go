package http

import (
	"fmt"
)

func (s *Server) AddRoute(path string, handler Handler, method []string) {
	if validateRoute(path, handler) {
		for _, m := range method {
			if _, ok := s.Routes[m]; !ok {
				s.Routes[m] = []Route{}
			}

			Params := []string{}

			if isParameterizedRoute(path) {
				path, Params = getParameterizedRoute(path)
			}

			fmt.Printf("🚀 [%s] %s Route loaded\n", m, path)

			searchParams := getSearchParams(path)
			path = removeQueryParams(path)

			s.Routes[m] = append(s.Routes[m], Route{
				Method:       method,
				Path:         path,
				Handler:      handler,
				Params:       Params,
				SearchParams: searchParams,
				Middlewares:  append([]Middleware{}, s.Middlewares...),
			})
			s.Routes[m] = sortRoutesWithParamsLast(s.Routes[m])
		}
	}
}

func (s *Server) addRouteWithMiddleware(path string, handler Handler, method []string, middlewares ...Middleware) {
	if validateRoute(path, handler) {
		for _, m := range method {
			if _, ok := s.Routes[m]; !ok {
				s.Routes[m] = []Route{}
			}

			Params := []string{}

			fmt.Printf("🚀 [%s] %s Route loaded\n", m, path)

			if isParameterizedRoute(path) {
				path, Params = getParameterizedRoute(path)
			}

			searchParams := getSearchParams(path)
			path = removeQueryParams(path)

			allMiddlewares := s.Middlewares

			s.Routes[m] = append(s.Routes[m], Route{
				Method:       method,
				Path:         path,
				Handler:      handler,
				Params:       Params,
				SearchParams: searchParams,
				Middlewares:  append(allMiddlewares, middlewares...),
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
		if validateRoute(path, route.Handler) {
			for _, m := range route.Method {
				if _, ok := s.Routes[m]; !ok {
					s.Routes[m] = []Route{}
				}

				Params := []string{}

				fmt.Printf("🚀 [%s] %s Route loaded\n", m, path)

				if isParameterizedRoute(path) {
					path, Params = getParameterizedRoute(path)
				}

				searchParams := getSearchParams(path)
				path = removeQueryParams(path)

				middlewares := s.Middlewares

				s.Routes[m] = append(s.Routes[m], Route{
					Method:       route.Method,
					Path:         path,
					Handler:      route.Handler,
					Params:       Params,
					SearchParams: searchParams,
					Middlewares:  append(middlewares, router.middlewares...),
				})
				s.Routes[m] = sortRoutesWithParamsLast(s.Routes[m])
			}
		}
	}
}

func (s *Server) Get(path string, handler Handler) *RouteChain {
	s.AddRoute(path, handler, []string{"GET"})
	return &RouteChain{
		server: s,
		path:   path,
		method: []string{"GET"},
	}

}

func (s *Server) Post(path string, handler Handler) *RouteChain {
	s.AddRoute(path, handler, []string{"POST"})
	return &RouteChain{
		server: s,
		path:   path,
		method: []string{"POST"},
	}

}

func (s *Server) Put(path string, handler Handler) *RouteChain {
	s.AddRoute(path, handler, []string{"PUT"})
	return &RouteChain{
		server: s,
		path:   path,
		method: []string{"PUT"},
	}

}

func (s *Server) Delete(path string, handler Handler) *RouteChain {
	s.AddRoute(path, handler, []string{"DELETE"})
	return &RouteChain{
		server: s,
		path:   path,
		method: []string{"DELETE"},
	}

}

func (s *Server) Patch(path string, handler Handler) *RouteChain {
	s.AddRoute(path, handler, []string{"PATCH"})
	return &RouteChain{
		server: s,
		path:   path,
		method: []string{"PATCH"},
	}

}

func (s *Server) Options(path string, handler Handler) *RouteChain {
	s.AddRoute(path, handler, []string{"OPTIONS"})
	return &RouteChain{
		server: s,
		path:   path,
		method: []string{"OPTIONS"},
	}

}

func (s *Server) Head(path string, handler Handler) *RouteChain {
	s.AddRoute(path, handler, []string{"HEAD"})
	return &RouteChain{
		server: s,
		path:   path,
		method: []string{"GET"},
	}
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
		s.addRouteWithMiddleware(fullPath, route.Handler, route.Method, route.Middlewares...)
	}
}
