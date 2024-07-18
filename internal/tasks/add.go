package tasks

import (
	"fmt"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

type Request struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type TaskAdder interface {
	AddTask(title string, done bool) error
}

func Add(taskAdder TaskAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			fmt.Errorf("failed to decode the request: %w", err)
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "Failed to decode the request",
			})
			return
		}

		err = taskAdder.AddTask(req.Title, req.Done)
		if err != nil {
			render.JSON(w, r, Response{
				Status: "Error",
				Error:  err.Error(),
			})
			return
		}

		log.Printf("Task added: %s [%v]", req.Title, req.Done)
		render.JSON(w, r, Response{
			Status: "OK",
		})
	}
}
