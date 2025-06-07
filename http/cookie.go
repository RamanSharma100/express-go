package http

import (
	"net/http"
	"time"
)

func (ctx *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool, Expires time.Time) {
	http.SetCookie(ctx.Response.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
		Expires:  Expires,
	})
}

func (ctx *Context) GetCookie(name string) (*http.Cookie, error) {
	cookie, err := ctx.Request.r.Cookie(name)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func (ctx *Context) DeleteCookie(name, path, domain string) {
	http.SetCookie(ctx.Response.Writer, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     path,
		Domain:   domain,
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})
}

func (ctx *Context) ClearCookies() {
	for _, cookie := range ctx.Request.r.Cookies() {
		http.SetCookie(ctx.Response.Writer, &http.Cookie{
			Name:     cookie.Name,
			Value:    "",
			Path:     "/",
			Domain:   cookie.Domain,
			MaxAge:   -1,
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		})
	}
}
