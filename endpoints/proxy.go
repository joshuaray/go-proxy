package endpoints

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"proxy/handlers"
)

func request(w http.ResponseWriter, r *http.Request, method string) *http.Request {
	url := r.Header.Get("X-Url")
	if url == "" {
		handlers.Unauthorized(w, r)
		return nil
	}
	query := r.URL.RawQuery
	if query != "" {
		url = url + "?" + query
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Fprintf(w, "")
		return nil
	}

	req.Header = r.Header.Clone()
	req.Header.Del("X-Url")
	req.Header.Del("X-Token")
	req.Header.Del("Postman-Token")
	return req
}

func execute(w http.ResponseWriter, r *http.Request, req *http.Request) string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		handlers.Error(w, r, res)
		return res.Status
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		handlers.Error(w, r, res)
		return res.Status
	}
	return string(bytes)
}

func Get(w http.ResponseWriter, r *http.Request) {
	req := request(w, r, "GET")
	if req == nil {
		return
	}
	fmt.Fprintf(w, execute(w, r, req))
}

func Post(w http.ResponseWriter, r *http.Request) {
	req := request(w, r, "POST")
	if req == nil {
		return
	}
	req.Body = r.Body
	fmt.Fprintf(w, execute(w, r, req))
}
