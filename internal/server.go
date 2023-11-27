package internal

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/marcopeocchi/valeera/internal/config"
	"github.com/marcopeocchi/valeera/internal/rest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	DB:       0,
	Addr:     os.Getenv("REDIS_ADDR"),
	Username: os.Getenv("REDIS_USER"),
	Password: os.Getenv("REDIS_PASS"),
})

type ServerConfig struct {
	Frontend fs.FS
	Config   *config.SafeConfig
}

func RunBlocking(sc ServerConfig) {
	fe := http.FileServer(http.FS(sc.Frontend))

	r := chi.NewRouter()

	r.Use(cors)
	r.Use(middleware.Logger)
	r.Mount("/", fe)

	r.Route("/api", func(r chi.Router) {
		if sc.Config.Cfg.Authentication {
			r.Use(authenticated)
		}
		r.Get("/temp", rest.Temp)
		r.Get("/top", rest.Top)
		r.Get("/config", rest.Config(sc.Config))
	})

	r.Post("/login", rest.Login(sc.Config, rdb))

	http.ListenAndServe(fmt.Sprintf(":%d", sc.Config.Cfg.Port), r)
}
