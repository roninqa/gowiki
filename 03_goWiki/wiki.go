package main

// To run this page, you will need to visit localhost:8080/view/test
// and the out put will be 'test Hello World'
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page data structure
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// The title is passed
func loadPage(title string) (*Page, error) {
	// The title must match the .txt file
	filename := title + ".txt"

	// The test.txt file is read
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Oh, snaps! The file does not exist. You will need to create the file.***")
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// You will need to hit the 'view' route and then
// enter the title. The title is the .txt file
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
