package main

import "net/http"

func (app *application) listPoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
