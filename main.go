package main

import (
	"fmt"
	"net/http"
	"proxy/endpoints"
	"proxy/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods("OPTIONS").HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	r.Use(handlers.AuthenticationHandler)
	r.Use(handlers.FilterHandler)
	r.HandleFunc("/proxy", endpoints.Get).Methods("get")
	r.HandleFunc("/proxy", endpoints.Post).Methods("post")

	go func() {
		err := http.ListenAndServe(":33333", r)
		if err != nil {
			fmt.Println(err)
		}
	}()
	select {}
}
