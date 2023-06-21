package main

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tombuente/tresor/server/api"
	"github.com/tombuente/tresor/server/service/health"
	"github.com/tombuente/tresor/server/service/snippet"
	"github.com/tombuente/tresor/server/store"
)

//go:embed scheme.sql
var scheme string

//go:embed seed.sql
var seed string

func newApp() chi.Router {
	ctx := context.Background()

	sqliteDb, err := sql.Open("sqlite3", "tresor.db")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := sqliteDb.ExecContext(ctx, scheme); err != nil {
		log.Fatal(err)
	}
	if _, err := sqliteDb.ExecContext(ctx, seed); err != nil {
		log.Fatal(err)
	}

	queries := store.New(sqliteDb)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	healthRouter := api.NewHealthRouter(health.NewService())
	snippetRouter := api.NewSnippetRouter(snippet.NewService(queries))

	apiRouter := chi.NewRouter()
	apiRouter.Mount("/health", healthRouter)
	apiRouter.Mount("/snippets", snippetRouter)

	r.Mount("/api", apiRouter)

	return r
}

func main() {
	server := &http.Server{Addr: "127.0.0.1:8080", Handler: newApp()}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-stop

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out... forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	log.Println("Server is now listening on", server.Addr)
	log.Println("Press Ctrl+C to exit")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}
