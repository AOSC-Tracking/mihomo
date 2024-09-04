package route

import (
	"fmt"
	"net/http"

	"github.com/metacubex/mihomo/component/updater"
	"github.com/metacubex/mihomo/log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func upgradeRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/ui", updateUI)
	r.Post("/geo", updateGeoDatabases)
	return r
}

func updateUI(w http.ResponseWriter, r *http.Request) {
	err := updater.UpdateUI()
	if err != nil {
		log.Warnln("%s", err)
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, newError(fmt.Sprintf("%s", err)))
		return
	}

	render.JSON(w, r, render.M{"status": "ok"})
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
