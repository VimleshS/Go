//http://www.alexedwards.net/blog/serving-static-sites-with-go
//http://stackoverflow.com/questions/12513963/how-to-read-input-from-a-html-form-and-save-it-in-a-file-golang
//http://stackoverflow.com/questions/25685797/html-form-submission-in-golang-template

package main

import (
	"html/template"
	//	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type SurveyCombo struct {
	Id         string
	SurveyName string
}

type Page struct {
	Title      string
	SurveyList []SurveyCombo
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/save", save)
	http.HandleFunc("/input", inputHandler)
	http.HandleFunc("/down", downloadHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Disposition", "attachment; filename=Export.csv")
	w.Header().Set("Content-Type", "text/csv")
	//	file * os.File
	//	file, _ := os.Create("E:\\Go\\text.csv")
	//	defer file.Close()

	buffer, err := ioutil.ReadFile("E:\\Go\\test.csv")
	if err != nil {
		log.Println(err.Error())
	}
	w.Write(buffer)
}

func save(w http.ResponseWriter, r *http.Request) {
	log.Println("INSAVE")
	name := r.FormValue("username")
	log.Println(name)
	log.Println(r.FormValue("password"))
	http.Redirect(w, r, "/input", http.StatusFound)
}

func inputHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{}
	p.Title = "New"
	t := SurveyCombo{Id: "1", SurveyName: "11"}
	p.SurveyList = append(p.SurveyList, t)
	t = SurveyCombo{Id: "2", SurveyName: "22"}
	p.SurveyList = append(p.SurveyList, t)

	log.Println("I'm here")
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "input.html")
	log.Println(lp)
	log.Println(fp)
	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
	log.Println(p)
	if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

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
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}
