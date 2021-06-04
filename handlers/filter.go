package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var whiteList = [...]string{"betmgm.com", "fanduel.com", "draftkings.com"}

func FilterHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestUrl := r.Header.Get("X-Url")
		if requestUrl == "" {
			Unauthorized(w, r)
			return
		}
		url, err := url.Parse(requestUrl)
		if err != nil {
			fmt.Println("Unable to parse request url: " + requestUrl)
			Unauthorized(w, r)
			return
		}
		hostParts := strings.Split(url.Host, ".")
		hostLen := len(hostParts)
		host := hostParts[hostLen-2] + "." + hostParts[hostLen-1]
		if err != nil {
			fmt.Println("Unable to parse url host: " + requestUrl)
			Unauthorized(w, r)
			return
		}

		host = strings.ToLower(host)
		isValid := false
		for i := 0; i < len(whiteList); i++ {
			if strings.ToLower(whiteList[i]) == host {
				isValid = true
				break
			}
		}

		if !isValid {
			Unauthorized(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
