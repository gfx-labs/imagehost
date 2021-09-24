package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)



var strFlag = flag.String("dir","./data","directory to host")

func main() {
  r := chi.NewRouter()
  r.Use(middleware.Logger)

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("gmorning!"))
  })
  flag.Parse()
  fileDir := http.Dir(filepath.Join(*strFlag))
  FileServer(r, "/map_topdown",fileDir)
  http.ListenAndServe(":10110", r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
