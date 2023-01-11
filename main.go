package main

import (
	"context"
	"embed"
	"flag"
	"io/fs"
	"log"
	"valeera/m/pkg"
)

type ContextKey interface{}

var (
	//go:embed frontend/dist
	vueFS      embed.FS
	configPath string
	config     pkg.SafeConfig
)

func init() {
	config = pkg.SafeConfig{}
	flag.StringVar(&configPath, "c", "./Valeerafile", "Path of Valeerafile")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	config.Load(configPath)

	frontend, err := fs.Sub(vueFS, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	ctx = context.WithValue(ctx, ContextKey("config"), &config)
	ctx = context.WithValue(ctx, ContextKey("frontend"), frontend)

	pkg.RunBlocking(ctx)
}
