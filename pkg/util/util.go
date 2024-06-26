package util

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func ParseTime(o *time.Time, t string) error {
	v, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return err
	}

	*o = v

	return nil
}

func ReadBytes(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// HelperJSOM return data in json format
func HelperJSOM(w http.ResponseWriter, r *http.Request, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// HandleError return error message in json format
func HandleError(w http.ResponseWriter, r *http.Request, error error) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ErrorResponse{Error: error.Error()}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
