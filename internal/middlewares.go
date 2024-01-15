package internal

import (
	"net/http"

	"github.com/marcopeocchi/rei/internal/models"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("valeera_session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		sessionID := cookie.Value
		user := models.User{}

		if err := rdb.Get(r.Context(), sessionID).Scan(&user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !user.Authenticated {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
