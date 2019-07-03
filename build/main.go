package main

import (
  "html/template"
  "net/http"
  "fmt"
  "log"
  "os"
)

var (
  tpl *template.Template
)

func init() {
  tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
  http.HandleFunc("/", indexHandler)

  // Serve static files out of the public directory.
	// By configuring a static handler in app.yaml, App Engine serves all the
	// static content itself. As a result, the following two lines are in
	// effect for development only.
	public := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
	http.Handle("/public/", public)

  // [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	// [END setting_port]
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
    return
  }
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
