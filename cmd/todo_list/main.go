package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"todo_list/internal/config"
	"todo_list/internal/storage"
	"todo_list/internal/tasks"
)

func main() {
	cfg := config.MustRead()

	stg := storage.New(cfg.StoragePath)
	_ = stg // will be used later

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", tasks.GetAll(stg))
	r.Post("/add", tasks.Add(stg))

	srv := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      r,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	log.Printf("Starting the server on http://%s [%s]", cfg.Server.Address, cfg.Env)
	log.Fatal(srv.ListenAndServe())
}
