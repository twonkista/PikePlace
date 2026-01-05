package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

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
	// get user by id
	// create user
	// create user (admin only)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {
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
