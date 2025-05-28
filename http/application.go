package http

type Application struct {
	Listen  func(port int, callback func(int, error))
	Get     HTTPMethod
	Post    HTTPMethod
	Put     HTTPMethod
	Patch   HTTPMethod
	Delete  HTTPMethod
	Options HTTPMethod
	Use     func(middleware Middleware)
}

func New() *Application {
	server := CreateServer()
	return &Application{
		Listen:  server.Listen,
		Get:     server.Get,
		Post:    server.Post,
		Put:     server.Put,
		Patch:   server.Patch,
		Delete:  server.Delete,
		Options: server.Options,
		Use:     server.Use,
	}
}

func (app *Application) Router() *ApplicationRouter {
	return &ApplicationRouter{
		Get:     app.Get,
		Post:    app.Post,
		Put:     app.Put,
		Delete:  app.Delete,
		Patch:   app.Patch,
		Options: app.Options,
		Use:     app.Use,
	}
}

func (app *Application) UseRouter(path string, router *Router) {
	if router == nil {
		return
	}

	if path == "" || path[0] != '/' {
		path = "/" + path
	}

	for _, route := range router.Routes {
		for _, method := range route.Method {
			if route.Path != "/" && route.Path[0] != '/' {
				route.Path = "/" + route.Path
			}
			fullPath := path + route.Path

			switch method {
			case "GET":
				app.Get(fullPath, route.Handler)
			case "POST":
				app.Post(fullPath, route.Handler)
			case "PUT":
				app.Put(fullPath, route.Handler)
			case "DELETE":
				app.Delete(fullPath, route.Handler)
			case "PATCH":
				app.Patch(fullPath, route.Handler)
			case "OPTIONS":
				app.Options(fullPath, route.Handler)
			}
		}
	}
}
