package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

type application struct {
	config config
	db     *gorm.DB
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		r.Route("/pools", func(r chi.Router) {
			r.Get("/list", app.listPoolsHandler)
			r.Get("/open", app.openPoolsHandler)
			r.Get("/resolved", app.resolvedPoolsHandler)
		})
	})
	// list all pools
	// list open pools (or list pools with status=open)
	// list resolved pools (or status=resolved)
	// create a new pool
	// place wager
	// cancel wager
	// list wagers for a user
	// list results for a pool
	// close pool (usually due to time expiration)
	// resolve pools
	// get pool by id
	// authenticate user
	// create user
	// create user (admin only)

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server started at %s", app.config.addr)

	return srv.ListenAndServe()
}
