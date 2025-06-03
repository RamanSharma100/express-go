package http

import (
	"net/http"
	"strings"
)

type CorsOptions struct {
	AllowOrigin  string
	AllowMethods string
	AllowHeaders string
	ContentType  string
}

func CORS(options *CorsOptions) Middleware {
	return func(ctx *Context, next func()) {
		origin := "*"
		methods := "GET, POST, PUT, PATCH, DELETE, OPTIONS"
		headers := "Content-Type, Authorization"

		if options != nil {
			if options.AllowOrigin != "" {
				origin = options.AllowOrigin
			}
			if options.AllowMethods != "" {
				methods = options.AllowMethods
			}
			if options.AllowHeaders != "" {
				headers = options.AllowHeaders
			}
			if options.ContentType != "" {
				ctx.Response.Writer.Header().Set("Content-Type", options.ContentType)
			} else {
				ctx.Response.Writer.Header().Set("Content-Type", "application/json")
			}
		}

		if origin != "*" {
			allowedOrigins := strings.Split(origin, ",")
			originAllowed := false
			requestOrigin := ctx.Request.GetHeader("Origin")

			for _, o := range allowedOrigins {
				if strings.TrimSpace(o) == requestOrigin {
					originAllowed = true
					break
				}
			}
			if !originAllowed {
				ctx.Response.Writer.WriteHeader(http.StatusForbidden)
				ctx.Response.Writer.Write([]byte("CORS policy does not allow access from this origin"))
				return
			}
		}

		ctx.Response.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		ctx.Response.Writer.Header().Set("Access-Control-Allow-Methods", methods)
		ctx.Response.Writer.Header().Set("Access-Control-Allow-Headers", headers)

		if ctx.Request.Method == "OPTIONS" {
			ctx.Response.Writer.WriteHeader(http.StatusNoContent)
			return
		}

		next()
	}
}
