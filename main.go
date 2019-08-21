package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb
const uploadPath = "./tmp"
const headerBulma = `
<!DOCTYPE html>
<html>
	<head>
	    <meta charset="utf-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1">
	    <title>Momo</title>
	    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.4/css/bulma.min.css">
	    <script defer src="https://use.fontawesome.com/releases/v5.4.0/js/all.js"></script>
	</head>
	<body>
		<section class="section">
`
const footerBulma = `
		</section>
	</body>
</html>
`

func renderView(content string) string {
	return headerBulma + content + footerBulma
}

const uploadView = `
	<div class="container">
	<form method="post">
	<h1 class="title">
		Enter your link
      	</h1>

	<div class="field">
	  <div class="control has-icons-left">
		<span class="icon is-small is-left">
			<i class="fas fa-at"></i>
	    	</span>
	    <input class="input is-primary" type="text" name="link">
	  </div>
	</div>

	 <div class="control">
	   <button class="button is-primary">Save</button>
	 </div>
	</form>
	</div>
`

const downloadView = `
	<div class="container">
		%s
	</div>
`

func getDownloadLinkView(filename string) []byte {
	view := renderView(fmt.Sprintf(downloadView, filename))
	return []byte(view)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, renderView(uploadView))
		case http.MethodPost:
			uploadFileHandler(w, r)
		}

	})

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {

	// parse and validate file and post parameters
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		renderError(w, "Something bad happens", http.StatusBadRequest)
		return
	}

	//TODO Sanitize r.PostForm
	spew.Println(r.PostForm)
	w.Write(getDownloadLinkView(r.PostForm["link"][0]))
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
