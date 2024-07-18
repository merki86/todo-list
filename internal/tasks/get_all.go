package tasks

import (
	"fmt"
	"net/http"
	"todo_list/internal/storage"
)

func GetAll(s *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := s.GetAllTasks()
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get the tasks: %v", err), http.StatusInternalServerError)
			return
		}

		for _, task := range tasks {
			fmt.Fprintf(w, "<p>ID: %d, Title: %s, Done: %v</p>", task.ID, task.Title, task.Done)
		}
	}
}
