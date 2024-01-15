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
	// maybe useful in dev
	// r.Use(middleware.Logger)

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

	r.Get("/", index(sc.Templates, sc.Config))

	r.Post("/login", rest.Login(sc.Config, rdb))

	http.ListenAndServe(fmt.Sprintf(":%d", sc.Config.Cfg.Port), r)
}
