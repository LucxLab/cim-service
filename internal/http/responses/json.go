package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Json(w http.ResponseWriter, httpStatus int, data interface{}) error {
	return writeJson(w, httpStatus, data)
}

func BadRequest(w http.ResponseWriter) {
	data := map[string]string{"message": "The request contains invalid data"}
	_ = writeJson(w, http.StatusBadRequest, data)
}

func UnexpectedError(w http.ResponseWriter, error error) {
	data := map[string]string{"message": "An unexpected error occurred"}

	fmt.Println(error)
	_ = writeJson(w, http.StatusInternalServerError, data)
}

func writeJson(w http.ResponseWriter, httpStatus int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	return nil
}
