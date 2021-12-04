package handlers

import (
	"log"
	"net/http"
)

func handleError(w http.ResponseWriter, r *http.Request, err error, message, path string) {
	if err != nil {
		log.Println(message)
		log.Println(err)
		http.Redirect(w, r, path, http.StatusSeeOther)
		return
	}
}
