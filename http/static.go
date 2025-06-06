package http

import (
	"fmt"
	"net/http"
	"path"
	"runtime"
	"strings"
)

func (s *Server) Static(prefix string, dir string) {
	if prefix == "" || prefix[0] != '/' {
		prefix = "/" + prefix
	}
	routePrefix := strings.TrimRight(prefix, "/")

	if dir == "" {
		panic("Directory cannot be empty for static file server")
	}

	_, b, _, _ := runtime.Caller(0)
	baseDir := path.Join(path.Dir(b), "../")
	staticDir := path.Clean(path.Join(baseDir, dir))

	s.Get(routePrefix, func(ctx *Context) {
		http.Redirect(ctx.Response.Writer, ctx.Request.r, routePrefix+"/", http.StatusMovedPermanently)
	})

	s.Get(routePrefix+"/", func(ctx *Context) {
		http.StripPrefix(routePrefix+"/", http.FileServer(http.Dir(staticDir))).ServeHTTP(ctx.Response.Writer, ctx.Request.r)
	})

	s.Get(routePrefix+"/:filepath", func(ctx *Context) {
		filePath := ctx.GetParam("filepath")
		fmt.Println("Requested file path:", filePath)
		if filePath == "" {
			http.NotFound(ctx.Response.Writer, ctx.Request.r)
			return
		}

		cleanPath := path.Clean(filePath)
		if strings.HasPrefix(cleanPath, "..") || path.IsAbs(cleanPath) {
			http.Error(ctx.Response.Writer, "Invalid file path", http.StatusBadRequest)
			return
		}

		fullPath := path.Join(staticDir, cleanPath)
		fmt.Println("Serving static file:", fullPath)
		http.ServeFile(ctx.Response.Writer, ctx.Request.r, fullPath)
	})

}
