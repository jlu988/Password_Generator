package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.WriteHeader(statusCode)

	if data == nil || statusCode == http.StatusNoContent {
		return nil
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("encoding: %+v", err)
	}

	return nil
}
