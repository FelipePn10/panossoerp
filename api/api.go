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
	generatemask "github.com/FelipePn10/panossoerp/internal/infrastructure/repository/generate_mask"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/repository/product"
	productquestion "github.com/FelipePn10/panossoerp/internal/infrastructure/repository/product_question"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/repository/questions"
	questionsoptions "github.com/FelipePn10/panossoerp/internal/infrastructure/repository/questions_options"
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
	r.Use(middleware.StripSlashes)
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

	// product
	productRepo := product.NewRepositoryProductSQLC(queries)

	createProductUC := usecase.NewCreateProductUseCase(productRepo)
	deleteProductUC := usecase.NewDeleteProductUseCase(productRepo)
	findProductByNameAndCodeUC := usecase.NewFindProductByNameAndCode(productRepo)

	productCreateHandler := handler.NewCreateProductHandler(createProductUC)
	productDeleteHandler := handler.NewDeleteProductHandler(deleteProductUC)
	findProductByNameAndCodeHandler := handler.NewFindProductByNameAndCodeHandler(findProductByNameAndCodeUC)

	// question
	questionRepo := questions.NewRepositoryQuestionSQLC(queries)

	createQuestionUC := usecase.NewQuestionUserUseCase(questionRepo)
	deleteQuestionUC := usecase.NewDeleteQuestionUseCase(questionRepo)
	findQuestionByNameUC := usecase.NewFindQuestionByName(questionRepo)

	questionCreateHandler := handler.NewQuestionHandler(createQuestionUC)
	questionDeleteHandler := handler.NewDeleteQuestionHandler(deleteQuestionUC)
	findQuestionByNameHandler := handler.NewFindQuestionByName(findQuestionByNameUC)

	// question option
	questionOptionRepo := questionsoptions.NewRepositoryQuestionOptionSQLC(queries)

	createQuestionOptionUC := usecase.NewCreateQuestionOptionUseCase(questionOptionRepo)
	deleteQuestionOptionUC := usecase.NewDeleteQuestionOptionUseCase(questionOptionRepo)

	questionOptionCreateHandler := handler.NewCreateQuestionOptionHandler(createQuestionOptionUC)
	questionOptionDeleteHandler := handler.NewDeleteQuestionOptionHandler(deleteQuestionOptionUC)

	// associate question in product
	productByQuestionProductRepo := productquestion.NewAssociateQuestionProductRepositorySQLC(queries)

	associateByQuestionProductUC := usecase.NewAssociateByQuestionProductUseCase(productByQuestionProductRepo)
	associateByQuestionProductHandler := handler.NewAssociateByQuestionProductHandler(associateByQuestionProductUC)

	// generate mask product
	generateMaskProduct := generatemask.NewRepositoryGenerateMaskSQLC(queries)

	generateMaskProductUC := usecase.NewGenerateMaskProductUseCase(generateMaskProduct)
	generateMaskProductHandler := handler.NewGeneratMaskProductHandler(generateMaskProductUC)

	// routes
	r.Group(func(r chi.Router) {
		r.Use(httpmw.JWT(app.config.JWTSecret, app.logger))

		r.Route("/api/products", func(r chi.Router) {
			r.With(httpmw.RequireRole("ADMIN", "USER")).Post("/create", productCreateHandler.CreateProduct)
			r.With(httpmw.RequireRole("ADMIN", "USER")).Delete("/{id}", productDeleteHandler.DeleteProduct)
			r.With(httpmw.RequireRole("ADMIN", "USER")).Get("/", findProductByNameAndCodeHandler.FindByNameAndCodeHandler)
		})
		r.Route("/api/questions", func(r chi.Router) {
			r.With(httpmw.RequireRole("ADMIN", "USER")).Post("/questions/create", questionCreateHandler.CreateQuestion)
			r.With(httpmw.RequireRole("ADMIN", "USER")).Delete("/{id}", questionDeleteHandler.DeleteQuestion)
			r.With(httpmw.RequireRole("ADMIN", "USER")).Get("/", findQuestionByNameHandler.FindQuestionByName)
			r.Route("/options", func(r chi.Router) {
				r.With(httpmw.RequireRole("ADMIN", "USER")).Post("/create-option", questionOptionCreateHandler.CreateQuestionOptionHandler)
				r.With(httpmw.RequireRole("ADMIN", "USER")).Delete("/{id}", questionOptionDeleteHandler.DeleteQuestionOption)
			})
			r.With(httpmw.RequireRole("ADMIN", "USER")).Post("/associate", associateByQuestionProductHandler.AssociateQuestions)
		})
		r.Route("/api/mask", func(r chi.Router) {
			r.With(httpmw.RequireRole("ADMIN", "USER")).Post("/generate", generateMaskProductHandler.GenerateMask)
		})
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
