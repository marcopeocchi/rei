package main

import (
	"embed"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io/fs"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/marcopeocchi/rei/internal"
	"github.com/marcopeocchi/rei/internal/config"
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

		parsed, err := template.ParseFS(
			files,
			"tmpl/"+tmpl.Name(),
			"tmpl/layouts/*.html",
			"tmpl/fragments/*.html",
		)
		if err != nil {
			return fmt.Errorf("cannot parse template %s, err: %w", tmpl.Name(), err)
		}

		log.Println("parsed", tmpl.Name())

		tmpls[tmpl.Name()] = parsed
	}

	return nil
}

func main() {
	cfg := config.New(configPath)

	if err := parseTemplates(); err != nil {
		log.Fatalln(err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Println("modified cofig file")
					cfg.Load(configPath)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	if err := watcher.Add(configPath); err != nil {
		log.Fatalln(err)
	}

	internal.RunBlocking(internal.ServerConfig{
		TmplFS:    files,
		Templates: &tmpls,
		StaticFS:  static,
		Config:    cfg,
	})
}
