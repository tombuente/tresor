package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tombuente/tresor/internal/tresor"
	tresor_sql "github.com/tombuente/tresor/internal/tresor/sql"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "tresor.db")
	defer db.Close()
	if err != nil {
		log.Println("Error opening SQLite3 database: ", err)
		return
	}

	if _, err := db.ExecContext(ctx, tresor_sql.Schema); err != nil {
		log.Println("Error executing database schema: ", err)
		return
	}

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

	tresorService := tresor.NewService(tresor.NewDB(tresor.New(db)))
	tresorRouter := tresor.NewRouter(tresorService)
	r.Mount("/api", tresorRouter)

	server := &http.Server{Addr: "127.0.0.1:8080", Handler: r}

	fmt.Println("Running...")

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
		return
	}
}
