package http

import "net/http"

func (res *Response) Send(text string) {
	w := res.Writer
	w.Header().Set("Content-Type", "text/plain")
	for key, value := range res.Headers {
		w.Header().Set(key, value)
	}
	if res.StatusCode == 0 {
		res.StatusCode = http.StatusOK
	}

	w.WriteHeader(res.StatusCode)

	w.Write([]byte(text))
}

func (res *Response) Json(data any) {
	w := res.Writer
	w.Header().Set("Content-Type", "application/json")
	for key, value := range res.Headers {
		w.Header().Set(key, value)
	}
	if res.StatusCode == 0 {
		res.StatusCode = http.StatusOK
	}

	w.WriteHeader(res.StatusCode)
	w.Write([]byte(data.(string)))

}

func (res *Response) AddHeader(key string, value interface{}) {
	w := res.Writer
	w.Header().Set(key, value.(string))
}

func (res *Response) Status(code int) *Response {
	res.StatusCode = code
	return res
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		Writer:  w,
		Headers: make(map[string]string),
	}
}
