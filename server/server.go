package server

import (
	"github.com/chenlu-chua/penny-wiser/user-service/service"
	"github.com/chenlu-chua/penny-wiser/user-service/svccontainer"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

type Server interface {
	Init(container *svccontainer.DIServiceContainer)
	GetServer() *chi.Mux
	RegisterHandlers() error
}

type server struct {
	mux            *chi.Mux
	userService    service.UserService
	billingService service.BillingService
}

func New() Server {
	mux := chi.NewRouter()

	// init middleware stack here
	mux.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}).Handler)

	return &server{
		mux: mux,
	}
}

func (s *server) Init(container *svccontainer.DIServiceContainer) {
	s.userService = container.UserService
	s.billingService = container.BillingService
}

func (s *server) GetServer() *chi.Mux {
	return s.mux
}

func (s *server) RegisterHandlers() error {

	// Init routes here
	s.mux.Route("/apis/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("Hello World!"))
		})
		r.Post("/users/register", s.handleUserRegister())
	})
	return nil
}
