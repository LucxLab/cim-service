package health

import "net/http"

type Handler interface {
	GlobalStatus(w http.ResponseWriter, r *http.Request)
}
