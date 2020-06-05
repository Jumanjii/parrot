package main

import "net/http"

type handler struct {
	*Renderer
}

func newHandler(r *Renderer) handler {
	return handler{r}
}

// handler for /storage endpoint.
func (h handler) storage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.Renderer.Storage(w, ViewData{Link: linkStorage})
	}
}

// handler for /store endpoint.
func (h handler) store(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Renderer.Form(w)
	case http.MethodPost:
		h.handleNewLink(w, r)
	}
}

// handleNewLink extract link from request and
// save it into the storage.
func (h handler) handleNewLink(w http.ResponseWriter, r *http.Request) {
	link := extractLink(r)
	if link == "" {
		renderError(w)
		return
	}

	StoreNewLink(link)
	http.Redirect(w, r, "/storage", http.StatusFound)
}
