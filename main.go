package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	newRouter := mux.NewRouter()
	newRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello goapi")
	})

	http.ListenAndServe(":5000", handlers.CORS()(newRouter))
}
