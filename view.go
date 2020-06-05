package main

import (
	"fmt"
	"html/template"
	"io"
)

// ViewData hosts data used for rendering.
type ViewData struct {
	Link string
}

// Renderer is used to render view of the application.
type Renderer struct {
	form    *template.Template
	storage *template.Template
}

// NewRenderer returns a new Renderer with all templates loaded.
func NewRenderer() (*Renderer, error) {
	form, err := template.ParseFiles("html/page.html", "html/form.html")
	if err != nil {
		return nil, fmt.Errorf("could not parse form html files: %w", err)
	}

	storage, err := template.ParseFiles("html/page.html", "html/storage.html")
	if err != nil {
		return nil, fmt.Errorf("could not parse storage html files: %w", err)
	}

	return &Renderer{
		form:    form,
		storage: storage,
	}, nil
}

// Form render the form view into given writer.
func (r *Renderer) Form(w io.Writer) {
	err := r.form.ExecuteTemplate(w, "page", nil)
	if err != nil {
		panic(err)
	}
}

// Storage render the storage view into given writer.
func (r *Renderer) Storage(w io.Writer, v ViewData) {
	err := r.storage.ExecuteTemplate(w, "page", v)
	if err != nil {
		panic(err)
	}
}
