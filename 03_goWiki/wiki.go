package main

//Using net/http to serve wiki pages
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Page struct
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[1:]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func vhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Chue. I love %s!", r.URL.Path[1:])
}

func main() {
	// adding some troubleshooting
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", vhandler)
	// this route is not working...in ubuntu
	// http.HandleFunc("/view/", viewHandler)
	// http.ListenAndServe(":8080", nil)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
