package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hit the get user handler"))
}

func (a *api) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created user"))
}

func main() {
	api := &api{
		addr: ":8080",
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.GetUserHandler)
	mux.HandleFunc("POST /users", api.CreateUserHandler)
	log.Println("Staring server on PORT 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error starting the server")
	}

}
