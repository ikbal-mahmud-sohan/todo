package main

import (
	"fmt"
	"log"
	"myGo/todo/handler"
	"net/http"
)


func main() {
	h := handler.New()
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/todos/create",h.Todos )
	fmt.Println("Server starting")
	if err := http.ListenAndServe("127.0.0.1:5731", nil); err != nil {
		log.Fatal(err)
	}
}





