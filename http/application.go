package http

type Application struct {
	Listen  func(port int, callback func(int, error))
	Get     HTTPMethod
	Post    HTTPMethod
	Put     HTTPMethod
	Patch   HTTPMethod
	Delete  HTTPMethod
	Options HTTPMethod
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
	}
}
