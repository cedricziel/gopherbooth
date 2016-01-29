package gopherbooth

import (
	"net/http"

	"appengine"
	"appengine/user"
	"handlers"
	"text/template"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/booths", handlers.HandleBooths)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	template, _ := template.ParseFiles("templates/index.html")
	template.Execute(w, u)
}
