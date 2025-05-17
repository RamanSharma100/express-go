package http

func (s *Server) ValidateRoute(path string, handler Handler) bool {
	if path == "" {
		panic("Path cannot be empty")
	}
	if handler == nil {
		panic("Handler cannot be empty")
	}

	return true

}

func (s *Server) AddRoute(path string, handler Handler, method []string) {
	if s.ValidateRoute(path, handler) {
		for _, m := range method {
			if _, ok := s.Routes[m]; !ok {
				s.Routes[m] = []Route{}
			}
			s.Routes[m] = append(s.Routes[m], Route{
				Method:  method,
				Path:    path,
				Handler: handler,
			})
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
