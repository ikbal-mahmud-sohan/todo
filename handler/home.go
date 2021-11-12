package handler

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Task        string `json:"task"`
	IsCompleted bool   `json:"is_completed"`
}
func (h *Handler) Home(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	todos := []Todo{
		{
			Task:        "This is task 1",
			IsCompleted: false,
		},
		{
			Task:        "This is task 2",
			IsCompleted: true,
		},
	}
	json.NewEncoder(rw).Encode(todos)
}