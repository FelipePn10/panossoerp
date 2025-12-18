package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/FelipePn10/panossoerp/internal/infrastructure/config"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database"
	httpmw "github.com/FelipePn10/panossoerp/internal/interfaces/middleware"
	"github.com/go-chi/chi/middleware"
	chimw "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type application struct {
	config *config.Config
	logger *slog.Logger
	db     *database.DB
}

func (app *application) traceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := chimw.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		duration := time.Since(start)

		app.logger.Info("request completed",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int64("duration_ms", duration.Milliseconds()),
			slog.String("client_ip", r.RemoteAddr),
			slog.Int("status", ww.Status()),
		)
	})
}

func (app *application) mount() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(app.traceMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(httpmw.JWT(app.config.JWTSecret, app.logger))
		//r.Get("/", app.)
	})

	// Health check
	r.Get("/health", app.healthHandler)

	return r
}

func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"status":    "ok",
		"timestamp": time.Now().UTC(),
		"service":   "core-api",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (app *application) run(r chi.Router) error {
	addr := app.config.ServerPort
	if addr == "" {
		addr = "5070"
	}
	if !strings.HasPrefix(addr, ":") {
		addr = ":" + addr
	}

	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", addr)
	return srv.ListenAndServe()
}
