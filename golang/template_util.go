package main

import (
	"html/template"
	"net/http"
)

type HtmlTemplates struct {
	Dir     string
	Layout  string
	Content string
}

func (t *HtmlTemplates) Files() []string {
	files := []string{}
	if t.Layout != "" {
		files = append(files, t.Dir+t.Layout)
	}
	if t.Content != "" {
		files = append(files, t.Dir+t.Content)
	}
	return files
}

func (t *HtmlTemplates) Execute(w http.ResponseWriter, bind interface{}) error {
	tmpl := template.Must(template.ParseFiles(t.Files()...))
	if err := tmpl.Execute(w, bind); err != nil {
		return err
	}
	return nil
}
