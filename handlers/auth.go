package handlers

import (
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"golang.org/x/crypto/scrypt"
)

func AuthenticationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-Token")
		if key == "" {
			Unauthorized(w, r, NoToken)
			return
		}
		code := os.Getenv("ProxyAuthKey")
		encryptedToken, err := encrypt([]byte(code[:128]), []byte(time.Now().UTC().Format("20060102")))
		if err != nil {
			Unauthorized(w, r, Internal)
			return
		}
		if key != encryptedToken {
			Unauthorized(w, r, InvalidToken)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func encrypt(token []byte, salt []byte) (string, error) {
	val, err := scrypt.Key(token, salt, 32768, 8, 1, 128)
	if err != nil {
		return "", err
	}
	b64 := base64.URLEncoding.EncodeToString(val)
	return string(b64), nil
}
