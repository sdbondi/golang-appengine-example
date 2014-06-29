package app

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
	http.HandleFunc("/user", userLogin)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	greetings := make([]Greeting, 0, 120)

	if err := LoadAllGreetings(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := NewGreeting(c)
	g.Content = r.FormValue("content")

	if g.Content == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}

	_, err := g.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func userLogin(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "Hello, %v!", u)
}
