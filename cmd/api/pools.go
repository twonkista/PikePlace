package main

import "net/http"

func (app *application) listPoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (app *application) openPoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aight"))
}

func (app *application) resolvedPoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bet"))
}

// func (app *application) createNewPoolHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("shi"))
// }
