package server

import (
	"github.com/chenlu-chua/penny-wiser/user-service/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Server interface {
	GetRouter() *chi.Mux
	RegisterHandlers() error
}

type server struct {
	mux            *chi.Mux
	userService    service.UserService
	billingService service.BillingService
}

func NewRouter() Server {
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

func (s *server) GetRouter() *chi.Mux {
	return s.mux
}

func (s *server) RegisterHandlers() error {

	// Init routes here
	s.mux.Route("/apis/v1", func(r chi.Router) {
		r.Post("/users/register", s.handleUserRegister())
	})
	return nil
}
