package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ResponseType string

const (
	NoToken       ResponseType = "No authentication token provided"
	InvalidToken  ResponseType = "Invalid authentication token provided"
	InvalidDomain ResponseType = "Invalid domain requested"
	Internal      ResponseType = "Internal error"
	InvalidPath   ResponseType = "Invalid path specified"
	NoUrl         ResponseType = "No proxy URL provided"
	BadUrl        ResponseType = "URL provided could not be parsed"
)

func Unauthorized(w http.ResponseWriter, r *http.Request, t ResponseType) {
	fmt.Fprintf(os.Stderr, string(t)+": "+getProperties(r))
	http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
}

func Error(w http.ResponseWriter, r *http.Request, response *http.Response) {
	fmt.Fprintf(os.Stderr, response.Status+"("+strconv.Itoa(response.StatusCode)+")"+": "+getProperties(r))
	http.Error(w, http.StatusText(response.StatusCode), response.StatusCode)
}

func Invalid(w http.ResponseWriter, r *http.Request, t ResponseType) {
	fmt.Fprintf(os.Stderr, string(t)+": "+getProperties(r))
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func Log(w http.ResponseWriter, r *http.Request, method string, response string) {
	fmt.Printf("Serving " + method + " response: " + getProperties(r))
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func getBody(r *http.Request) string {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		bodyBytes = []byte{}
	}
	body := string(bodyBytes)
	return body
}

func getProperties(r *http.Request) string {
	return "\r\n" +
		"        Request URL:       " + r.RequestURI + "\r\n" +
		"        Source IP Address: " + getIP(r) + "\r\n" +
		"        Method:            " + r.Method + "\r\n" +
		"        Body:              " + getBody(r) + "\r\n" +
		"        Auth Token:        " + r.Header.Get("X-Token") + "\r\n" +
		"        Proxy URL:         " + r.Header.Get("X-Url") + "\r\n\r\n"
}
