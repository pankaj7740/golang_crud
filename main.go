package main

import (
	"crud_basic/db"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	db.ConnectDatabase()

	db.AutoMigrate()
	
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })
	fmt.Println("Server is running on http://localhost:3000")
    http.ListenAndServe(":3000", r)
}
