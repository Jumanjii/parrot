package main

import "net/http"

// extract link from posted form.
func extractLink(r *http.Request) string {
	err := r.ParseForm()
	if err != nil {
		return ""
	}

	formLink := r.PostForm["link"]
	if len(formLink) != 1 {
		return ""
	}

	//TODO Sanitize user input
	// https://github.com/microcosm-cc/bluemonday

	return formLink[0]
}
