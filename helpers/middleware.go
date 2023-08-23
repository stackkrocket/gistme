package helpers

import (
	"mime"
	"net/http"
)

func CheckHeaderGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if ct != "" {
			mt, _, err := mime.ParseMediaType(ct)
			if err != nil {
				http.Error(w, "Malformed content type", http.StatusBadRequest)
				return
			}

			if mt != "text/html" {
				http.Error(w, "Content type must be text/html", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func CheckHeaderPost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ct := r.Header.Get("Content-Type")
		if ct != "" {
			mt, _, err := mime.ParseMediaType(ct)
			if err != nil {
				http.Error(w, "Malformed content type", http.StatusBadRequest)
				return
			}

			if mt != "application/x-www-form-urlencoded" {
				http.Error(w, "Content type must be application/x-www-form-urlencoded", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)

	})
}

//Authorize user before granting access to Requested API

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("access_token")

		if clientToken == "" {
			panic("Could not retrieve credentials. Make sure you are logged in")
		}

		claims, err := ValidateToken(clientToken)
		if err != "" {
			panic("credentials not validated")
		}

		r.Header.Set("email", claims.Email)
		r.Header.Set("name", claims.Name)
		r.Header.Set("uid", claims.Uid)

		next.ServeHTTP(w, r)
	})
}
