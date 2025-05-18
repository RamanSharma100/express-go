package http

import (
	"fmt"
	"net/http"
	"regexp"
)

func (s *Server) GetParams(routePath, actualPath string) map[string]string {
	params := make(map[string]string)

	re := regexp.MustCompile(`\{([^\s/]+)\}`)
	paramNames := re.FindAllStringSubmatch(routePath, -1)

	pattern := re.ReplaceAllString(routePath, `([^/]+)`)
	pattern = "^" + pattern + "$"
	valueRe := regexp.MustCompile(pattern)

	matches := valueRe.FindStringSubmatch(actualPath)
	if matches == nil {
		return params
	}

	for i, name := range paramNames {
		if len(matches) > i+1 {
			params[name[1]] = matches[i+1]
		}
	}

	return params
}

func (s *Server) GetHeaders(r *http.Request) map[string]string {
	headers := make(map[string]string)
	for key, values := range r.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}
	return headers
}

func (s *Server) GetBasicResponseHeaders(method string) map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Methods"] = "GET"
	headers["Access-Control-Allow-Headers"] = "Content-Type, Authorization"
	return headers
}

func (s *Server) HandleRoutes(w http.ResponseWriter, r *http.Request) {
	for _, route := range s.Routes[r.Method] {

		re := regexp.MustCompile(`\{[^\s/]+\}`)
		path := re.ReplaceAllString(route.Path, "[^/]+")

		if path[0] != '/' {
			path = `/` + path
		}

		re = regexp.MustCompile(fmt.Sprintf(`^%s$`, path))

		if re.MatchString(r.URL.Path) {

			s.Request.Method = r.Method
			s.Request.Url = r.URL.String()
			s.Request.Headers = s.GetHeaders(r)
			s.Request.r = r
			s.Request.AdditionalFields = map[string]any{
				"params": s.GetParams(route.Path, r.URL.Path),
			}

			s.Response.Writer = w
			s.Response.StatusCode = 0
			s.Response.Headers = s.GetBasicResponseHeaders(r.Method)

			ctx := &Context{
				Request:  s.Request,
				Response: s.Response,
			}

			route.Handler(ctx)
			return
		}
	}

	w.Write([]byte("404 Not Found"))
}
