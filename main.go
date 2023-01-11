package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed frontend/dist
	vueFS      embed.FS
	configPath string
	config     SafeConfig
)

func init() {
	config = SafeConfig{cfg: Config{Port: 8686}}
	flag.StringVar(&configPath, "c", "./Valeerafile", "Path of Valeerafile")
	flag.Parse()
}

func main() {
	frontend, err := fs.Sub(vueFS, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	config.Load(configPath)

	frontendFS := http.FileServer(http.FS(frontend))

	mux := http.NewServeMux()
	mux.Handle("/", frontendFS)
	mux.HandleFunc("/config", handleConfig)
	mux.HandleFunc("/temp", handleTemp)
	mux.HandleFunc("/top", handleTop)

	http.ListenAndServe(fmt.Sprintf(":%d", config.cfg.Port), CORS(mux))
}
