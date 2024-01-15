package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"

	"github.com/marcopeocchi/valeera/internal"
	"github.com/marcopeocchi/valeera/internal/config"
)

//go:generate npm run build

var (
	//go:embed tmpl/* tmpl/layouts/*
	files embed.FS
	tmpls map[string]*template.Template

	//go:embed static
	static embed.FS

	configPath    string
	wallpaperPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "./config.yml", "path of configuration file")
	flag.StringVar(&wallpaperPath, "bg", "./static/wallpaper.avif", "path of background image")
	flag.Parse()
}

func parseTemplates() error {
	tmpls = make(map[string]*template.Template)

	tmplFiles, err := fs.ReadDir(files, "tmpl")
	if err != nil {
		return errors.New("cannot open templates directory")
	}

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		parsed, err := template.ParseFS(files, "tmpl/"+tmpl.Name(), "tmpl/layouts/*.html")
		if err != nil {
			return fmt.Errorf("cannot parse template %s, err: %w", tmpl.Name(), err)
		}

		tmpls[tmpl.Name()] = parsed
	}

	return nil
}

func main() {
	cfg := config.New(configPath)

	if err := parseTemplates(); err != nil {
		log.Fatalln(err)
	}

	internal.RunBlocking(internal.ServerConfig{
		TmplFS:    files,
		Templates: &tmpls,
		StaticFS:  static,
		Config:    cfg,
	})
}
