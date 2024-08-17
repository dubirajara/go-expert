package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerRouter struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
type WebServer struct {
	Router        chi.Router
	Handlers      []HandlerRouter
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []HandlerRouter{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, HandlerRouter{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		s.Router.MethodFunc(handler.Method, handler.Path, handler.Handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
