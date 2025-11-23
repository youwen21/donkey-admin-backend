package serverx

import (
	"net/http"
)

type FsServerI interface {
	CanServe(urlPath string) bool
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
