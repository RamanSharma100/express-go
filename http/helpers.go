package http

import "strings"

func validateRoute(path string, handler Handler) bool {
	if path == "" {
		panic("Path cannot be empty")
	}
	if handler == nil {
		panic("Handler cannot be empty")
	}

	return true
}

func getParameterizedRoute(path string) (string, []string) {
	Params := []string{}
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			paramName := strings.TrimPrefix(part, ":")
			Params = append(Params, paramName)
			parts[i] = "{" + paramName + "}"
		}
	}

	return strings.Join(parts, "/"), Params
}

func getSearchParams(path string) map[string]string {
	SearchParams := make(map[string]string)
	if strings.Contains(path, "?") {
		parts := strings.Split(path, "?")
		if len(parts) > 1 {
			queryParams := parts[1]
			for _, param := range strings.Split(queryParams, "&") {
				if strings.Contains(param, "=") {
					keyValue := strings.SplitN(param, "=", 2)
					if len(keyValue) == 2 {
						SearchParams[keyValue[0]] = keyValue[1]
					}
				}
			}
		}
	}
	return SearchParams
}

func isParameterizedRoute(path string) bool {
	return strings.ContainsAny(path, ":")
}

func sortRoutesWithParamsLast(routes []Route) []Route {
	for i := 0; i < len(routes); i++ {
		for j := i + 1; j < len(routes); j++ {
			if len(routes[i].Params) > 0 && len(routes[j].Params) == 0 {
				routes[i], routes[j] = routes[j], routes[i]
			}
		}
	}
	return routes
}

func removeQueryParams(path string) string {
	if strings.Contains(path, "?") {
		parts := strings.Split(path, "?")
		return parts[0]
	}
	return path
}
