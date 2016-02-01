package main

import (
	"net/http"

	"appengine"
	"appengine/user"
	"handlers"
	"text/template"
	"util"
)

type Page struct {
	User *user.User
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/booths", handlers.HandleBooths)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/login", handleLogin)
}

// Handles the index route
func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	page := Page{
		User: u,
	}

	template, _ := template.ParseFiles("templates/index.html")
	template.Execute(w, page)
}

// Logs the user out of the application
func handleLogout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	util.RedirectIfNotLoggedIn(w, r)

	if u == nil {
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	logoutUrl, _ := user.LogoutURL(c, "/")
	w.Header().Set("Location", logoutUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

// Handles the AppEngine
func handleLogin(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)

	if u == nil {
		url, err := user.LoginURL(c, "/")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusContinue)
}
