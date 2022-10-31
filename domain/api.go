package domain

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RouterInitialize() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/generatePassword/size={size}", generatePassword).Methods("GET")
	return r
}

func generatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("accept", "application/json")
	ph := PasswordGenerator{}
	query := mux.Vars(r)["size"]
	size, _ := strconv.Atoi(query)
	if size < 14 || size > 256 {
		err := "password size must between 14 and 256 characters"
		Response(w, http.StatusBadRequest, err)
		return
	}
	Response(w, http.StatusOK, ph.GetPassword(size))
}
