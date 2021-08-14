package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/olingern/golytics/pkg/auth"
)

type LoginJsonBody struct {
	Username string
	Password string
}

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var loginJsonBody LoginJsonBody

	err := decoder.Decode(&loginJsonBody)
	if err != nil {
		json.NewEncoder(w).Encode(&SucessResponse{Success: false})
		return
	}

	result, err := c.db.LogUserIn(loginJsonBody.Username, loginJsonBody.Password)
	if err != nil {
		json.NewEncoder(w).Encode(&SucessResponse{Success: false})
		return
	}

	auth.SaveSession(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&SucessResponse{Success: result})
}
