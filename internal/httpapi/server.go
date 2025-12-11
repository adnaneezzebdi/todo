package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/internal/tasks"
)

type Server struct {
	Repo *tasks.SQLRepository
}

func Startserver(repo *tasks.SQLRepository) {
	s := &Server{Repo: repo}

	http.HandleFunc("/tasks", s.handleTasks)
	http.HandleFunc("/tasks/complete", s.handleComplete)
	http.HandleFunc("/tasks/delete", s.handleDelete)

	println("Server HTTP attivo su http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func (s *Server) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		ts, err := s.Repo.List()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(ts)

	case http.MethodPost:
		var body struct {
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "json non valido", 400)
			return
		}

		_, err := s.Repo.Add(body.Title)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "task aggiunta!",
		})
	}
}

func (s *Server) handleComplete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if err := s.Repo.Complete(id); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "task completata!",
	})
}

func (s *Server) handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	if err := s.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "task eliminata!",
	})
}
