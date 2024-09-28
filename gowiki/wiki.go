package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Covered in this: https://go.dev/doc/articles/wiki/

// Creating a data structure with load and save methods
// Using the net/http package to build web applications
// Using the html/template package to process HTML templates
// Using the regexp package to validate user input
// Using closures

// Let's start by defining the data structures.
// A wiki consists of a series of interconnected pages,
// each of which has a title and a body (the page content).

// Here, we define Page as a struct with two fields representing the title and body.
type Page struct {
	Title string
	// The type []byte means "a byte slice"
	Body []byte
}

// The Page struct describes how page data will be stored in memory.
// But what about persistent storage? We can address that by creating a save method on Page:

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// In addition to saving pages,
// we also want to load pages

// func loadPage(title string) *Page {
// 	filename := title + ".txt"
// 	body, _ := os.ReadFile(filename)
// 	return &Page{Title: title, Body: body}
// }

// But what happens if ReadFile encounters an error?
// For example, the file might not exist.
// We should not ignore such errors.
// Let's modify the function to return *Page and error.

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// Render a template for page title and body
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, p)
}

// func main() {
// 	p1 := &Page{Title: "Test Page", Body: []byte("Hello, World !.\n - From page")}

// 	p1.save()
// 	p2, _ := loadPage("Test Page")
// 	fmt.Println(string(p2.Body))
// }

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, 404)
	}
	// t, _ := template.ParseFiles("templtes/view.html")
	// t.Execute(w, p)
	tmpl := "view.html"
	renderTemplate(w, tmpl, p)
	// fmt.Fprintf(w, "<h1>%s</h1><br><div>%s</div>", p.Title, string(p.Body))

}

// func editHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/edit/"):]

// 	p, err := loadPage(title)

// 	if err != nil {
// 		p = &Page{Title: title}
// 	}

// 	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
// 		"<form action=\"/save/\" method=\"POST\""+
// 		"<textarea name=\"body\"> %s </textarea><br>"+
// 		"<input type=\"submit\" value=\"save\">"+
// 		"</form>",
// 		p.Title, p.Body)
// }

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	// t, _ := template.ParseFiles("templates/edit.html")
	// t.Execute(w, p)
	tmpl := "edit.html"
	renderTemplate(w, tmpl, p)

}

// func saveHandler(w http.ResponseWriter, r *http.Request) {

// }

func main() {

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
