package http

type Application struct {
	Listen    func(port int, callback func(int, error))
	Get       HTTPMethod
	Post      HTTPMethod
	Put       HTTPMethod
	Patch     HTTPMethod
	Delete    HTTPMethod
	Options   HTTPMethod
	Use       func(middlewares ...Middleware)
	UseRouter func(prefix string, router *Router)
}

func New() *Application {
	server := CreateServer()
	return &Application{
		Listen:    server.Listen,
		Get:       server.Get,
		Post:      server.Post,
		Put:       server.Put,
		Patch:     server.Patch,
		Delete:    server.Delete,
		Options:   server.Options,
		Use:       server.Use,
		UseRouter: server.UseRouter,
	}
}

func (app *Application) Group(prefix string, middlewares []Middleware, handler func(router *Router)) {
	if prefix == "" || prefix[0] != '/' {
		prefix = "/" + prefix
	}

	router := &Router{
		routes: []Route{},
	}

	router.Use(middlewares...)

	handler(router)

	app.UseRouter(prefix, router)
}
