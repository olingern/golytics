package auth

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func SaveSession(w http.ResponseWriter, r *http.Request) (bool, error) {
	session, err := store.Get(r, "admin-session")

	if err != nil {
		return false, err
	}

	session.Values["username"] = "admin"
	err = session.Save(r, w)

	if err != nil {
		return false, err
	}

	return true, nil

}
