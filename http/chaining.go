package http

import (
	"fmt"
	"regexp"
	"strings"
)

func (rc *RouteChain) Name(name string) {
	if name == "" {
		panic("Route name cannot be empty")
	}

	if rc.router != nil {
		for i, route := range rc.router.routes {
			if route.Path == rc.path && fmt.Sprintf("%v", route.Method) == fmt.Sprintf("%v", rc.method) {
				rc.router.routes[i].Name = name
			}
		}
		return
	}

	for _, m := range rc.method {
		routes := rc.server.Routes[m]
		for i := range routes {
			routePattern := "^" + strings.ReplaceAll(routes[i].Path, "{", "(?P<") // convert {param} to regex group
			routePattern = strings.ReplaceAll(routePattern, "}", ">[^/]+)") + "$"
			matched, _ := regexp.MatchString(routePattern, rc.path)
			if (routes[i].Path == rc.path || matched) && fmt.Sprintf("%v", routes[i].Method) == fmt.Sprintf("%v", rc.method) {
				routes[i].Name = name
				rc.server.Routes[m][i].Name = name
			}
		}
	}

}
