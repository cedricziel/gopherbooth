package handlers

import (
	"net/http"
	"appengine/user"
	"appengine"
	"appengine/datastore"
	"text/template"
	"time"

	"model"
	"model/booth"
)

// Handles the /booths index route and switches over the route
func HandleBooths(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetAllBooths(w, r)
		break;
	case "POST":
		handleCreateBooth(w, r)
		break;
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
	}
}

func handleGetAllBooths(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		w.Header().Set("Location", "/")
		return
	}

	q := datastore.NewQuery("Booth").Ancestor(booth.Key(c)).Order("-Date").Limit(10)
	booths := make([]model.Booth, 10)

	if _, err := q.GetAll(c, &booths); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	boothsTemplate, _ := template.ParseFiles("./templates/booths/index.html")

	if err := boothsTemplate.Execute(w, booths); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleCreateBooth(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		w.Header().Set("Location", "/")
		return
	}

	b := model.Booth{
		Author: u.String(),
		Date: time.Now(),
		Name: r.FormValue("name"),
	}

	key := datastore.NewIncompleteKey(c, "Booth", booth.Key(c))
	_, err := datastore.Put(c, key, &b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusCreated)
}
