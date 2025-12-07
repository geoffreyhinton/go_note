package main

import (
	"net/http"

	"github.com/geoffreyhinton/go_note/api"
	"github.com/geoffreyhinton/go_note/store"

	"github.com/gorilla/mux"
)

func main() {
	store.InitDBConn()

	r := mux.NewRouter().StrictSlash(true)

	api.RegisterUserRoutes(r)
	api.RegisterAuthRoutes(r)

	http.ListenAndServe("localhost:8080", r)
}
