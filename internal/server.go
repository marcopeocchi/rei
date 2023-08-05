package internal

import (
	"fmt"
	"io/fs"
	"net/http"
	"valeera/m/internal/config"
	"valeera/m/internal/rest"
)

type ServerConfig struct {
	Frontend fs.FS
	Config   *config.SafeConfig
}

func RunBlocking(sc ServerConfig) {
	frontendFS := http.FileServer(http.FS(sc.Frontend))

	mux := http.NewServeMux()

	mux.Handle("/", frontendFS)
	mux.HandleFunc("/temp", rest.Temp)
	mux.HandleFunc("/top", rest.Top)
	mux.HandleFunc("/config", rest.Config(sc.Config))

	http.ListenAndServe(fmt.Sprintf(":%d", sc.Config.Cfg.Port), CORS(mux))
}
