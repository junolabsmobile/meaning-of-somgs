package http

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/junolabsmobile/meaning-of-songs/internal/application"
)

func NewRouter(songService *application.SongService) *chi.Mux {
	r := chi.NewRouter()
	r.Use(LoggingMiddleware)
	r.Use(CORSMiddleware)

	h := NewHandler(songService)

	// API routes
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", h.HealthCheck)
		r.Get("/songs", h.ListSongs)
		r.Get("/songs/{id}", h.GetSong)
	})

	// Servir frontend estático desde web/dist/ si existe
	staticDir := "web/dist"
	if _, err := os.Stat(staticDir); err == nil {
		fsys := os.DirFS(staticDir)
		fileServer := http.FileServer(http.FS(fsys))
		// SPA fallback: si no encuentra el archivo, sirve index.html
		r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path[1:]
			if _, err := fs.Stat(fsys, path); err != nil {
				http.ServeFile(w, r, staticDir+"/index.html")
				return
			}
			fileServer.ServeHTTP(w, r)
		})
	}

	return r
}
