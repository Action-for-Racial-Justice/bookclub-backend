// Package classification Bookclub API.
//
// Documentation for Bookclub API
//
//	Schemes: http
//	BasePath: /v1
//	Version: 0.0.1
//
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//	swagger:meta

package handlers

//go:generate mockgen -package=mocks -destination=../mocks/handlers.go github.com/Action-for-Racial-Justice/bookclub-backend/internal/handlers Handlers

import (
	"net/http"
	"time"

	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/models"
	"github.com/Action-for-Racial-Justice/bookclub-backend/internal/service"
	"github.com/felixge/httpsnoop"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/google/wire"
)

//Module to denote wire binding function
var Module = wire.NewSet(
	New,
)

//Handlers interface to describe BookClubHandlers struct receiver functions
type Handlers interface {
	EndUserSession(w http.ResponseWriter, r *http.Request)
	JoinClub(w http.ResponseWriter, r *http.Request)
	GetClubs(w http.ResponseWriter, r *http.Request)
	GetClubData(w http.ResponseWriter, r *http.Request)
	GetUserData(w http.ResponseWriter, r *http.Request)
	GetUserClubs(w http.ResponseWriter, r *http.Request)
	GetSSOToken(w http.ResponseWriter, r *http.Request)
	GetBookData(w http.ResponseWriter, r *http.Request)
	HealthCheck(w http.ResponseWriter, r *http.Request)
	LeaveClub(w http.ResponseWriter, r *http.Request)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	CreateClub(w http.ResponseWriter, r *http.Request)
}

//BookClubHandler struct to hold relevant inner data members and hold functions for pure handler logic
type BookClubHandler struct {
	service service.Service
	router  *chi.Mux
}

//New ... constructor
func New(service service.Service, logger *zap.Logger) (*BookClubHandler, error) {
	handlers := &BookClubHandler{service: service}
	router := chi.NewRouter()
	router.Use(accessLogger(logger))

	router.Use(cors.Handler(setCorsOptions()))

	// health endpoint
	registerEndpoint("/health", router.Get, handlers.HealthCheck) // integration tested

	// book endpoints
	registerEndpoint("/v1/book", router.Get, handlers.GetBookData)

	// user endpoints
	registerEndpoint("/v1/user", router.Get, handlers.GetArjBackendUserData) //integration tested
	registerEndpoint("/v1/user/clubs", router.Post, handlers.GetUserClubs)
	registerEndpoint("/v1/user/session", router.Post, handlers.GetSSOToken)      // integration tested
	registerEndpoint("/v1/user/session", router.Delete, handlers.EndUserSession) // integration tested

	// club endpoints
	registerEndpoint("/v1/club", router.Get, handlers.GetClubs)
	registerEndpoint("/v1/club/create", router.Post, handlers.CreateClub)
	registerEndpoint("/v1/club/id", router.Post, handlers.GetClubData)
	registerEndpoint("/v1/club/join", router.Post, handlers.JoinClub)
	registerEndpoint("/v1/club/leave", router.Post, handlers.LeaveClub)

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

//setCorsOptions acts as a setter function for the cors.Options struct
func setCorsOptions() cors.Options {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
	return corsOptions
}

func curateJSONError(errs ...error) models.ErrorResponse {

	errList := make([]string, 0)

	for _, err := range errs {
		errList = append(errList, err.Error())
	}
	return models.ErrorResponse{ErrList: errList}
}

func accessLogger(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			response := httpsnoop.CaptureMetrics(next, w, r)
			host := r.Header.Get("x-forwarded-for")
			if len(host) == 0 {
				host = r.RemoteAddr
			}

			contentLength := r.Header.Get("Content-Length")
			if len(contentLength) == 0 {
				contentLength = "-"
			}

			referer := r.Header.Get("referer")
			if len(referer) == 0 {
				referer = "-"
			}

			userAgent := r.Header.Get("user-agent")
			if len(userAgent) == 0 {
				userAgent = "-"
			}

			xCorrelationID := r.Header.Get("X-Correlation-Id")
			if len(xCorrelationID) == 0 {
				xCorrelationID = "-"
			}

			var logFunction func(msg string, fields ...zapcore.Field)
			logFunction = logger.Info
			logFunction("access", zap.String("host", host),
				zap.String("time", time.Now().Format("2006-01-02 15:04:05.999Z")),
				zap.String("user-name", "-"), zap.String("content-length", contentLength),
				zap.String("user-agent", userAgent), zap.String("referer", referer),
				zap.Int("status-code", response.Code), zap.String("request-method", r.Method),
				zap.String("correlation-id", xCorrelationID),
				zap.String("request-uri", r.URL.RequestURI()), zap.Duration("response-time", response.Duration))
		})
	}
}
