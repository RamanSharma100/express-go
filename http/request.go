package http

import (
	"net/http"
)

func (req *Request) AddHeader(key string, value interface{}) {
	req.Headers[key] = value.(string)
}

func (req *Request) AddField(key string, value any) {
	req.AdditionalFields[key] = value
}

func (req *Request) ParseBody() {
	if req.r.Method == "POST" {
		err := req.r.ParseForm()
		if err != nil {
			return
		}
		req.Body = req.r.FormValue("body")
	}
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		r:                r,
		Method:           r.Method,
		Url:              r.URL.String(),
		Headers:          make(map[string]string),
		Body:             r.FormValue("body"),
		AdditionalFields: make(map[string]any),
	}
}
