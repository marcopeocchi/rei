package pkg

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
)

var (
	config *SafeConfig
)

func RunBlocking(ctx context.Context) {
	frontendFS := http.FileServer(http.FS(
		ctx.Value("frontend").(fs.FS),
	))

	config = ctx.Value("config").(*SafeConfig)

	mux := http.NewServeMux()

	mux.Handle("/", frontendFS)

	mux.HandleFunc("/config", handleConfig)
	mux.HandleFunc("/temp", handleTemp)
	mux.HandleFunc("/top", handleTop)

	http.ListenAndServe(fmt.Sprintf(":%d", config.cfg.Port), CORS(mux))
}
