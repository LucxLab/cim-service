package health

import "net/http"

/*
// Interfaces
*/

// Handler defines methods that a Health handler must implement to handle HTTP requests.
type Handler interface {
	GlobalStatus(w http.ResponseWriter, r *http.Request)
}
