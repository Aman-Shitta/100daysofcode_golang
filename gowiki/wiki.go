package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
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

// Validation

// The function regexp.MustCompile will parse and compile the regular expression,
// and return a regexp.Regexp. MustCompile is distinct from Compile in that it
// will panic if the expression compilation fails, while Compile returns an
// error as a second parameter.
var validPath = regexp.MustCompile("^/(edit|view|save)/([A-Za-z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)

	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil // Title is the second sub-expression
}

var templates = template.Must(
	template.ParseFiles("tmpl/edit.html", "tmpl/view.html"),
)

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

	filename := "data/" + title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

// Render a template for page title and body
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles(tmpl)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func main() {
// 	p1 := &Page{Title: "Test Page", Body: []byte("Hello, World !.\n - From page")}

// 	p1.save()
// 	p2, _ := loadPage("Test Page")
// 	fmt.Println(string(p2.Body))
// }

// Function definition

// func viewHandler(w http.ResponseWriter, r *http.Request, title string)
// func editHandler(w http.ResponseWriter, r *http.Request, title string)
// func saveHandler(w http.ResponseWriter, r *http.Request, title string)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// here we will extract the page title from the Request,
		// and call the provided handler fn

		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }
	p, err := loadPage(title)

	// handle non-existent pages
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, 404)
	}
	// t, _ := template.ParseFiles("templtes/view.html")
	// t.Execute(w, p)
	tmpl := "view"
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

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/edit/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }

	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	// t, _ := template.ParseFiles("templates/edit.html")
	// t.Execute(w, p)
	tmpl := "edit"
	renderTemplate(w, tmpl, p)

}

// Saving pages

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	// title, err := getTitle(w, r)
	// if err != nil {
	// 	return
	// }
	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func rootRedirectHandler(w http.ResponseWriter, r *http.Request, title string) {
}

func main() {

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
