package internal

import (
	"html/template"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/google/uuid"
	"github.com/marcopeocchi/rei/internal/config"
	"github.com/marcopeocchi/rei/internal/models"
)

func index(tmpls *map[string]*template.Template, sc *config.SafeConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, ok := (*tmpls)["index.html"]
		if !ok {
			http.Error(w, "cannot read template", http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["ServerName"] = sc.Cfg.ServerName
		data["Services"] = sc.Cfg.Services
		data["Wallpaper"] = sc.Cfg.Wallpaper
		data["Opacity"] = sc.Cfg.Opacity
		data["Scheme"] = sc.Cfg.Scheme

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getLogin(tmpls *map[string]*template.Template, sc *config.SafeConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, ok := (*tmpls)["login.html"]
		if !ok {
			http.Error(w, "cannot read template", http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["Opacity"] = sc.Cfg.Opacity
		data["Wallpaper"] = sc.Cfg.Wallpaper

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func postLogin(tmpls *map[string]*template.Template, sc *config.SafeConfig, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		r.ParseForm()

		user := models.User{}

		user.Name = r.Form.Get("username")
		user.Password = r.Form.Get("password")

		if user.Name != sc.Cfg.Username && user.Password != sc.Cfg.Password {
			tmpl, ok := (*tmpls)["login.html"]
			if !ok {
				http.Error(w, "cannot read template", http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["Opacity"] = sc.Cfg.Opacity
			data["Wallpaper"] = sc.Cfg.Wallpaper
			data["Error"] = "Wrong username or password!"

			if err := tmpl.Execute(w, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			return
		}

		user.Authenticated = true

		var (
			sessionID = uuid.NewString()
			ttl       = time.Hour
		)

		err := rdb.Set(r.Context(), sessionID, user, ttl).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "valeera_session",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(ttl),
		})

		http.Redirect(w, r, "/web", http.StatusTemporaryRedirect)
	}
}
