package myhandlers

import (
	"net/http"
	"os"
)

// Favicon ...
func Favicon(w http.ResponseWriter, r *http.Request) {
	faviconPath := "static" + string(os.PathSeparator) + "pics" + string(os.PathSeparator) + "favicon.ico"
	http.ServeFile(w, r, faviconPath)
}
