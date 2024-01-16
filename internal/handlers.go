package internal

import (
	"html/template"
	"net/http"

	"github.com/marcopeocchi/rei/internal/config"
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
		data["Services"] = sc.Cfg.Services
		data["Wallpaper"] = sc.Cfg.Wallpaper
		data["Opacity"] = sc.Cfg.Opacity

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
