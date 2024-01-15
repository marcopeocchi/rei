package internal

import (
	"html/template"
	"net/http"

	"github.com/marcopeocchi/valeera/internal/config"
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
		data["Arr"] = sc.Cfg.Arr
		data["Downloaders"] = sc.Cfg.Downloaders
		data["Media"] = sc.Cfg.Media
		data["System"] = sc.Cfg.System

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
