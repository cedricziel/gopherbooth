package util

import (
	"net/http"
	"appengine"
	"appengine/user"
)

// Redirects the user if not logged in
func RedirectIfNotLoggedIn(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
