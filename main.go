package main

import (
	"flag"
	"fmt"
	"net/http"
	"proxy/endpoints"
	"proxy/handlers"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods("OPTIONS").HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	r.Use(handlers.AuthenticationHandler)
	r.Use(handlers.WhiteListHandler(parseConfig()))
	r.HandleFunc("/proxy", endpoints.Get).Methods("get")
	r.HandleFunc("/proxy", endpoints.Post).Methods("post")
	r.PathPrefix("/").HandlerFunc(endpoints.NotFound).Methods("get", "post", "put", "patch", "delete")

	go func() {
		err := http.ListenAndServe(":33333", r)
		if err != nil {
			fmt.Println(err)
		}
	}()
	select {}
}

func parseConfig() []string {
	whitelist := flag.String("whitelist", "", "comma delimited list of white-listed domains")
	flag.Parse()
	if whitelist == nil {
		return nil
	}
	return strings.Split(*whitelist, ",")
}
