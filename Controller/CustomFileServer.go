package Controller

import (
	logger "GoLibs/logs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type customFileServer struct {
	root            http.Dir
	NotFoundHandler func(http.ResponseWriter, *http.Request)
}

func (fs *customFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Branco("customFileServer.ServeHTTP", "Inicializando")

	dir := string(fs.root)
	if dir == "" {
		dir = "."
	}

	logger.Branco("customFileServer.ServeHTTP", "Montando diretório", dir)

	// add prefix and clean
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	upath = path.Clean(upath)

	// path to file
	name := path.Join(dir, filepath.FromSlash(upath))

	//check if file exists
	f, err := os.Open(name)
	if err != nil {
		if os.IsNotExist(err) {
			fs.NotFoundHandler(w, r)
			return
		}
	}

	defer f.Close()
	http.ServeFile(w, r, name)
}

func CustomFileServer(root http.Dir, NotFoundHandler http.HandlerFunc) http.Handler {
	return &customFileServer{root: root, NotFoundHandler: NotFoundHandler}
}
