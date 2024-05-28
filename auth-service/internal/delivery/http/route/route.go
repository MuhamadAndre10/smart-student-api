package route

import (
	deliver_http "github.com/MuhamadAndre/auth-service/internal/delivery/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"net/http"
)

type Config struct {
	AuthController *deliver_http.AuthController
}

func New(config *Config) http.Handler {

	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.Get("/login", config.AuthController.SignIn)

	return mux
}
