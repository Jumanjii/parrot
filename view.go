package main

import (
	"fmt"
	"io"
)

// RenderLinkForm render the form allowing to store a link.
func RenderLinkForm(w io.Writer) {
	fmt.Fprintf(w, renderView(linkFormView))
}

// RenderLink generates and return storage webpage.
func RenderLink(w io.Writer, link string) {
	v := fmt.Sprintf(linkStoredView, link, link)
	fmt.Fprintf(w, renderView(v))
}

// docHeader defines the head of the HTML document.
// Inject:
//     - Bulma CSS framework (from CDN)
//     - Font Awesome
const docHeader = `
<!DOCTYPE html>
<html>
	<head>
	    <meta charset="utf-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1">
	    <title>Parrot</title>
	    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.8.0/css/bulma.min.css">
	    <script defer src="https://use.fontawesome.com/releases/v5.12.1/js/all.js"></script>
	</head>
	<body>
		<section class="section">
`

// docFooter defines the footer of the HTML document.
const docFooter = `
		</section>
	</body>
</html>
`

//linkFormView print the form allowing to save a link.
const linkFormView = `
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

// linkStoredView print the link stored.
const linkStoredView = `
	<div class="container">
		<a href="%s">%s</a>
	</div>
`

// renderView generates the HTML document containing the content.
func renderView(content string) string {
	return docHeader + content + docFooter
}
