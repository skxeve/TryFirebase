package main

import (
	"net/http"
)

type Page struct {
	Title  string
	Count  int
	Header string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	html := HtmlTemplates{
		Dir:     "templates/",
		Layout:  "layout.tmpl",
		Content: "empty_content.tmpl",
	}
	page := Page{
		Title: "Hello Golang!",
		Count: 1,
	}

	if err := html.Execute(w, page); err != nil {
		panic(err)
	}
}

func viewHandler2(w http.ResponseWriter, r *http.Request) {
	html := HtmlTemplates{
		Dir:     "templates/",
		Layout:  "layout.tmpl",
		Content: "content.tmpl",
	}
	page := Page{
		Title:  "Hello Golang!",
		Count:  2,
		Header: "Hoge",
	}

	if err := html.Execute(w, page); err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/next", viewHandler2)
	http.ListenAndServe(":8080", nil)
}
