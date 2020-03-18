package main

var linkStorage string

// StoreNewLink updates link stored in memory.
func StoreNewLink(link string) {
	if link == "" {
		return
	}

	linkStorage = link
}
