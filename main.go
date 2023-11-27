package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"

	"github.com/marcopeocchi/valeera/internal"
	"github.com/marcopeocchi/valeera/internal/config"
)

var (
	//go:embed app/dist
	app        embed.FS
	configPath string
)

func init() {
	flag.StringVar(&configPath, "c", "./Valeerafile", "Path of Valeerafile")
	flag.Parse()
}

func main() {
	c := config.New(configPath)

	app, err := fs.Sub(app, "app/dist")
	if err != nil {
		log.Fatalln(err)
	}

	internal.RunBlocking(internal.ServerConfig{
		Frontend: app,
		Config:   c,
	})
}
