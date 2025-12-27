package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type Course struct {
	ID    string
	Title string `json:"title"`
	Items []Item `json:"items"`
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/select", modeSelectionHandler)
	http.HandleFunc("/study", studyHandler)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func toJSON(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

func render(w http.ResponseWriter, tmplName string, data any) {
	tmpl := template.New("layout").Funcs(template.FuncMap{
		"json": toJSON,
	})
	
	files := []string{"templates/layout.html", "templates/" + tmplName}
	var err error
	tmpl, err = tmpl.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		http.Error(w, "Render error: "+err.Error(), http.StatusInternalServerError)
	}
}

func getCourses() ([]Course, error) {
	var courses []Course
	files, err := os.ReadDir("data")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			content, err := os.ReadFile(filepath.Join("data", file.Name()))
			if err != nil {
				continue
			}

			var course Course
			if err := json.Unmarshal(content, &course); err == nil {
				course.ID = strings.TrimSuffix(file.Name(), ".json")
				courses = append(courses, course)
			}
		}
	}
	return courses, nil
}

func loadCourse(id string) (*Course, error) {
	filename := filepath.Join("data", id+".json")
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var course Course
	if err := json.Unmarshal(content, &course); err != nil {
		return nil, err
	}
	course.ID = id
	return &course, nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	courses, err := getCourses()
	if err != nil {
		http.Error(w, "Unable to load courses", http.StatusInternalServerError)
		return
	}
	render(w, "index.html", courses)
}

func modeSelectionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	course, err := loadCourse(id)
	if err != nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}
	render(w, "mode.html", course)
}

func studyHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	mode := r.URL.Query().Get("mode")

	course, err := loadCourse(id)
	if err != nil {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}

	templateName := ""
	switch mode {
	case "flashcards":
		templateName = "flashcards.html"
	case "quiz":
		templateName = "quiz.html"
	default:
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	render(w, templateName, course)
}