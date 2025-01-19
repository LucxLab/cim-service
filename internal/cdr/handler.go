package cdr

import "net/http"

type Handler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
}
