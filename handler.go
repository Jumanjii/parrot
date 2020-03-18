package main

import "net/http"

// Handler for /storage endpoint.
func storageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/storage" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodGet {
		RenderLink(w, linkStorage)
	}
}

// Handler for /store endpoint.
func storeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/store" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		RenderLinkForm(w)
	case http.MethodPost:
		handleNewLink(w, r)
	}
}

func handleNewLink(w http.ResponseWriter, r *http.Request) {
	link := extractLink(r)
	if link == "" {
		renderError(w)
		return
	}

	StoreNewLink(link)
	http.Redirect(w, r, "/storage", http.StatusFound)
}
