//http://www.alexedwards.net/blog/serving-static-sites-with-go
//http://stackoverflow.com/questions/12513963/how-to-read-input-from-a-html-form-and-save-it-in-a-file-golang
//http://stackoverflow.com/questions/25685797/html-form-submission-in-golang-template

package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/save", save)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func save(w http.ResponseWriter, r *http.Request) {
	log.Println("INSAVE")
	name := r.FormValue("username")
	log.Println(name)
	log.Println(r.FormValue("password"))
}
func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", r.URL.Path)
	//	var fp string
	//	if r.URL.Path == "/" {
	//		fp = path.Join("templates", "/example.html")
	//	} else {
	//		fp = path.Join("templates", r.URL.Path)
	//	}
	//	log.Printf(" %s | %s", r.URL.Path, fp)

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
