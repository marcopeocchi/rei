package internal

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"

	"github.com/marcopeocchi/rei/internal/config"
	"github.com/marcopeocchi/rei/internal/rest"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	DB:       0,
	Addr:     os.Getenv("REDIS_ADDR"),
	Username: os.Getenv("REDIS_USER"),
	Password: os.Getenv("REDIS_PASS"),
})

type ServerConfig struct {
	TmplFS   fs.FS
	StaticFS fs.FS

	Templates *map[string]*template.Template

	Config *config.SafeConfig
}

func RunBlocking(sc ServerConfig) {
	r := chi.NewRouter()

	r.Use(cors)

	r.Mount("/static", http.FileServer(http.FS(sc.StaticFS)))

	r.Route("/api", func(r chi.Router) {
		if sc.Config.Cfg.Authentication {
			r.Use(authenticated)
		}
		r.Get("/temp", rest.Temp)
		r.Get("/top", rest.Top)
		r.Get("/topFmt", rest.TopFmt)
		r.Get("/config", rest.Config(sc.Config))
	})

	r.Route("/web", func(r chi.Router) {
		if sc.Config.Cfg.Authentication {
			r.Use(authenticated)
		}
		r.Handle("/", index(sc.Templates, sc.Config))
		r.Get("/login", getLogin(sc.Templates, sc.Config))
		r.Post("/login", postLogin(sc.Templates, sc.Config, rdb))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web", http.StatusTemporaryRedirect)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", sc.Config.Cfg.Port), r)
}
