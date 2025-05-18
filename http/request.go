package http

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func (req *Request) AddHeader(key string, value interface{}) {
	req.Headers[key] = value.(string)
}

func (req *Request) AddField(key string, value any) {
	req.AdditionalFields[key] = value
}

func (req *Request) GetHeader(key string) string {
	return req.Headers[key]
}

func (req *Request) GetJsonBody() any {
	req.r.ParseForm()
	if req.r.Header.Get("Content-Type") == "application/json" {
		var body any
		err := json.NewDecoder(req.r.Body).Decode(&body)
		if err != nil {
			return nil
		}
		return body
	}
	return nil
}

func (req *Request) GetBody() any {
	req.r.ParseForm()

	if req.r.Header.Get("Content-Type") == "application/json" {
		req.Body = req.GetJsonBody()
		return req.Body
	}

	req.Body = req.r.Form

	return req.Body
}

func (req *Request) GetXMLBody() any {
	if req.r.Header.Get("Content-Type") == "application/xml" {
		var body any
		err := json.NewDecoder(req.r.Body).Decode(&body)
		if err != nil {
			return nil
		}
		return body
	}
	return nil
}

func (req *Request) GetParams() map[string]string {
	params := make(map[string]string)
	re := regexp.MustCompile(`\{([^\s/]+)\}`)
	paramNames := re.FindAllStringSubmatch(req.r.URL.Path, -1)

	pattern := re.ReplaceAllString(req.r.URL.Path, `([^/]+)`)
	pattern = "^" + pattern + "$"
	valueRe := regexp.MustCompile(pattern)

	matches := valueRe.FindStringSubmatch(req.r.URL.Path)
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

func (req *Request) GetParam(name string) string {
	params := req.GetParams()
	if value, ok := params[name]; ok {
		return value
	}
	return ""
}

func (req *Request) GetQueryParams() map[string]string {
	params := make(map[string]string)
	queryParams := req.r.URL.Query()

	for key, values := range queryParams {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}

	return params
}

func (req *Request) GetQueryParam(name string) string {
	queryParams := req.r.URL.Query()
	if values, ok := queryParams[name]; ok && len(values) > 0 {
		return values[0]
	}
	return ""
}

func (req *Request) GetPath() string {
	return req.r.URL.Path
}

func (req *Request) ParseBody() any {
	if req.r.Method == "POST" {
		if req.r.Header.Get("Content-Type") == "application/json" {
			req.Body = req.GetJsonBody()
			return req.Body
		}
		err := req.r.ParseForm()

		if err != nil {
			return nil
		}

		req.Body = req.r.FormValue("body")

		return req.Body
	}

	return ""
}

func (req *Request) GetMethod() string {
	return req.Method
}

func (req *Request) GetUrl() string {
	return req.Url
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
