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

// HelperJSON returns data in JSON format
func HelperJSON(w http.ResponseWriter, r *http.Request, data any) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
