package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type CorsOptions struct {
	AllowOrigin  string
	AllowMethods string
	AllowHeaders string
	ContentType  string
}

type LogOptions struct {
	Enable bool
	Format string
}

type RateLimitOptions struct {
	Limit     int
	Window    int
	Remaining int
}

var rateLimitStore = make(map[string]int)
var rateLimitMutex sync.Mutex

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
				ctx.Response.Status(http.StatusForbidden)
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

func Logs(options *LogOptions) Middleware {
	return func(ctx *Context, next func()) {
		next()
		if options != nil && options.Enable {
			format := options.Format
			if format == "" {
				format = "{{.Method}} {{.Path}} - {{.StatusCode}}"
			}

			logData := map[string]any{
				"Method":     ctx.Request.Method,
				"Path":       ctx.Request.r.URL.Path,
				"StatusCode": strconv.Itoa(ctx.Response.StatusCode),
			}

			logMessage := format
			for key, value := range logData {
				logMessage = strings.ReplaceAll(logMessage, "{{."+key+"}}", value.(string))
			}

			fmt.Println(logMessage)
		}

	}
}

func RateLimit(options *RateLimitOptions) Middleware {
	return func(ctx *Context, next func()) {
		if options == nil || options.Limit <= 0 || options.Window <= 0 {
			next()
			return
		}

		//TODO: Implement redis for production
		clientIP := ctx.Request.r.RemoteAddr

		rateLimitMutex.Lock()
		count := rateLimitStore[clientIP]
		if count >= options.Limit {
			options.Remaining = 0
		} else {
			options.Remaining = options.Limit - count
			rateLimitStore[clientIP] = count + 1
		}
		rateLimitMutex.Unlock()

		if options.Remaining <= 0 {
			ctx.Response.Writer.WriteHeader(http.StatusTooManyRequests)
			ctx.Response.Status(http.StatusTooManyRequests)
			ctx.Response.Writer.Write([]byte("Rate limit exceeded"))
			return
		}

		options.Remaining--
		next()
	}
}

func uploadFiles(uploadDir string) Middleware {
	return func(ctx *Context, next func()) {
		if ctx.Request.r.Method == "POST" || ctx.Request.r.Method == "PUT" || ctx.Request.r.Method == "PATCH" {
			contentType := ctx.Request.r.Header.Get("Content-Type")
			if len(contentType) > 0 && contentType[:19] == "multipart/form-data" {
				_, _ = ctx.saveUploadedFiles(uploadDir)
			}
		}
		next()
	}
}
