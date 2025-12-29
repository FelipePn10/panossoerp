package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/FelipePn10/panossoerp/internal/application/usecase"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/config"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/repository/product"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/repository/user"
	"github.com/FelipePn10/panossoerp/internal/interfaces/http/handler"
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

	queries := app.db.Queries()

	userRepo := user.NewRepositoryUserSQLC(queries)
	registerUserUC := usecase.NewRegisterUserUseCase(userRepo)
	loginUserUC := usecase.NewLoginUserUseCase(userRepo)

	userHandler := handler.NewUserHandler(
		registerUserUC,
		loginUserUC,
		app.config.JWTSecret,
	)

	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.RegisterUserHandler)
		r.Post("/login", userHandler.LoginHandler)
	})

	productRepo := product.NewRepositorySQLC(queries)
	createProductUC := usecase.NewCreateProductUseCase(productRepo)
	productHandler := handler.NewCreateProductHandler(createProductUC)

	r.Route("/products", func(r chi.Router) {
		r.Post("/create", productHandler.CreateProduct)

	})

	r.Group(func(r chi.Router) {
		r.Use(httpmw.JWT(app.config.JWTSecret, app.logger))

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
