package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

func (req *Request) Validate(data map[string]string, attributes any) map[string]string {
	// validate request like this => $validated = $request->validate([
	//     'title' => 'required|unique:posts|max:255',
	//     'body' => 'required',
	// ]); this is only for demonstration purposes
	if data == nil {
		panic("Validation data cannot be nil")
	}

	errors := make(map[string]string)
	fmt.Println("Request body:", attributes)
	if attributes == nil {
		panic("Request body cannot be nil")
	}
	attrMap := attributes.(map[string]any)

	emailRe := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	for key, rule := range data {
		rule := strings.Split(rule, "|")
		val, exists := attrMap[key]

		for _, r := range rule {
			switch r {
			case "required":
				if !exists || val == nil || val == "" {
					errors[key] = "The" + key + " is required"

				}
			case "string":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a string"
					break
				}
				if _, ok := val.(string); !ok {
					errors[key] = "The " + key + " must be a string"

				}
			case "integer":
			case "number":
			case "int":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a number"
					break
				}
				if _, ok := val.(float64); !ok {
					errors[key] = "The " + key + " must be a number"
				}
			case "boolean":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a boolean"
					break
				}
				if _, ok := val.(bool); !ok {
					errors[key] = "The " + key + " must be a boolean"
				}
			case "array":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be an array"
				}
				if _, ok := val.([]interface{}); !ok {
					errors[key] = "The " + key + " must be an array"
				}
			case "date":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a date"
				}
				if _, ok := val.(string); !ok {
					errors[key] = "The " + key + " must be a date"
				} else {
					// You can add more complex date validation here if needed
					dateRegex := `^\d{4}-\d{2}-\d{2}$`
					dateRe := regexp.MustCompile(dateRegex)
					if !dateRe.MatchString(val.(string)) {
						errors[key] = "The " + key + " must be a valid date in YYYY-MM-DD format"
					}
				}
			case "email":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be an email"
				}
				if _, ok := val.(string); !ok {
					errors[key] = "The " + key + " must be an email"
				} else {
					if !emailRe.MatchString(val.(string)) {
						errors[key] = "The " + key + " must be a valid email address"
					}
				}
			case "datetime":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a datetime"
				}
				if _, ok := val.(string); !ok {
					errors[key] = "The " + key + " must be a datetime"
				} else {
					datetimeRegex := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}$`
					datetimeRe := regexp.MustCompile(datetimeRegex)
					if !datetimeRe.MatchString(val.(string)) {
						errors[key] = "The " + key + " must be a valid datetime in YYYY-MM-DDTHH:MM:SS format"
					}
				}
			case "time":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a time"
				}
				if _, ok := val.(string); !ok {
					errors[key] = "The " + key + " must be a time"
				} else {
					timeRegex := `^\d{2}:\d{2}:\d{2}$`
					timeRe := regexp.MustCompile(timeRegex)
					if !timeRe.MatchString(val.(string)) {
						errors[key] = "The " + key + " must be a valid time in HH:MM:SS format"
					}
				}
			case "url":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a url"
				}
				if _, ok := val.(string); !ok {
					errors[key] = "The " + key + " must be a url"
				} else {
					urlRegex := `^(http|https)://[^\s/$.?#].[^\s]*$`
					urlRe := regexp.MustCompile(urlRegex)
					if !urlRe.MatchString(val.(string)) {
						errors[key] = "The " + key + " must be a valid url"
					}
				}
			case "max":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a max"
					break
				}
				parts := strings.Split(r, ":")
				if len(parts) == 2 {
					maxStr := parts[1]
					maxInt := 0
					if n, err := strconv.Atoi(maxStr); err == nil {
						maxInt = n
					}
					switch v := val.(type) {
					case string:
						if len(v) > maxInt {
							errors[key] = "The " + key + " must be a max of " + maxStr
						}
					case []interface{}:
						if len(v) > maxInt {
							errors[key] = "The " + key + " must be a max of " + maxStr
						}
					case []byte:
						if len(v) > maxInt {
							errors[key] = "The " + key + " must be a max of " + maxStr
						}
					case float64:
						if v > float64(maxInt) {
							errors[key] = "The " + key + " must be a max of " + maxStr
						}
					}
				}
			case "min":
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " must be a min"
					break
				}
				parts := strings.Split(r, ":")
				if len(parts) == 2 {
					minStr := parts[1]
					minInt := 0
					if n, err := strconv.Atoi(minStr); err == nil {
						minInt = n
					}
					switch v := val.(type) {
					case string:
						if len(v) < minInt {
							errors[key] = "The " + key + " must be a min of " + minStr
						}
					case []interface{}:
						if len(v) < minInt {
							errors[key] = "The " + key + " must be a min of " + minStr
						}
					case []byte:
						if len(v) < minInt {
							errors[key] = "The " + key + " must be a min of " + minStr
						}
					case float64:
						if v < float64(minInt) {
							errors[key] = "The " + key + " must be a min of " + minStr
						}
					}
				}
			default:
				if !exists || val == nil || val == "" {
					errors[key] = "The " + key + " is not a valid rule"
				} else {
					errors[key] = "The " + key + " has an unknown validation rule: " + r
				}
			}
		}
	}

	return errors
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
