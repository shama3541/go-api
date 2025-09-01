package main

import "net/http"

func (app *application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Everything looks goood"))
}
