//go:generate mockgen -package=mocks -destination=../mocks/handlers.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/handlers Handlers

package handlers

import (
	"net/http"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/wire"
)

//Module to denote wire binding function
var Module = wire.NewSet(
	New,
)

//Handlers interface to describe BookClubHandlers struct receiver functions
type Handlers interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

//BookClubHandler struct to hold relevant inner data members and hold functions for pure handler logic
type BookClubHandler struct {
	service service.Service
	router  *chi.Mux
}

//New ... constructor
func New(service service.Service) (*BookClubHandler, error) {
	handlers := &BookClubHandler{service: service}
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	registerEndpoint("/health", router.Get, handlers.HealthCheck)
	handlers.router = router

	return handlers, nil
}

// ServeHTTP serves a http request given a response builder and request
func (bh *BookClubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bh.router.ServeHTTP(w, r)
}

//registerEndpoint registers an endpoint to the router for a specified method type and handlerFunction
func registerEndpoint(endpoint string, routeMethod func(pattern string, handlerFn http.HandlerFunc), handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	routeMethod(endpoint, http.HandlerFunc(handlerFunc).ServeHTTP)
}
