package handlers

import "net/http"

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func Error(w http.ResponseWriter, r *http.Request, response *http.Response) {
	http.Error(w, http.StatusText(response.StatusCode), response.StatusCode)
}
