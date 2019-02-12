package main

import (
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
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
func authHandler(w http.ResponseWriter, r *http.Request) {
	opt := option.WithCredentialsFile("secrets/try-firebase-auth_adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	html := HtmlTemplates{
		Dir:     "templates/",
		Layout:  "layout.tmpl",
		Content: "content.tmpl",
	}

	//client, err := app.Auth(context.Background())
	_, err = app.Auth(context.Background())
	if err != nil {
		page := Page{
			Title:  "Error auth",
			Count:  1,
			Header: err.Error(),
		}
		if err := html.Execute(w, page); err != nil {
			panic(err)
		}
		return
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
	http.HandleFunc("/auth", authHandler)
	http.ListenAndServe(":8080", nil)
}
