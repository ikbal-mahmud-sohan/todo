package handler

import (
	"fmt"
	"log"
	"net/http"
)
func (h *Handler) Todos(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil{
		log.Fatal(err)
	}
	taskName := r.FormValue("task")
	fmt.Fprintf(rw,"Task Name: %s", taskName)
}