package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type userRequest struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type Store struct {
	mu     sync.Mutex
	users  []User
	nextID int
}

func NewStore() *Store {
	return &Store{nextID: 1}
}

func (s *Store) allUsers() []User {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]User, len(s.users))
	copy(out, s.users)
	return out
}

func (s *Store) usersByMinAge(minAge int) []User {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []User
	for _, u := range s.users {
		if u.Age >= minAge {
			out = append(out, u)
		}
	}
	return out
}

func (s *Store) getUser(id int) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, u := range s.users {
		if u.ID == id {
			return u, true
		}
	}
	return User{}, false
}

func (s *Store) createUser(req userRequest) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	u := User{
		ID:    s.nextID,
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}
	s.nextID++
	s.users = append(s.users, u)
	return u
}

func (s *Store) updateUser(id int, req userRequest) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.users {
		if s.users[i].ID == id {
			s.users[i].Name = req.Name
			s.users[i].Age = req.Age
			s.users[i].Email = req.Email
			return s.users[i], true
		}
	}
	return User{}, false
}

func (s *Store) deleteUser(id int) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.users {
		if s.users[i].ID == id {
			deleted := s.users[i]
			s.users = append(s.users[:i], s.users[i+1:]...)
			return deleted, true
		}
	}
	return User{}, false
}

func main() {
	store := NewStore()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		usersHandler(w, r, store)
	})
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		userByIDHandler(w, r, store)
	})

	fmt.Println("server started: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request, store *Store) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Get("minAge") != "" {
			getUsersByAge(w, r, store)
		} else {
			writeJSON(w, http.StatusOK, store.allUsers())
		}
	case http.MethodPost:
		createUser(w, r, store)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func userByIDHandler(w http.ResponseWriter, r *http.Request, store *Store) {
	id, ok := idFromPath(r.URL.Path)
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUser(w, id, store)
	case http.MethodPut:
		updateUser(w, r, id, store)
	case http.MethodDelete:
		deleteUser(w, id, store)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func getUsersByAge(w http.ResponseWriter, r *http.Request, store *Store) {
	minAge, err := strconv.Atoi(r.URL.Query().Get("minAge"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid minAge")
		return
	}
	writeJSON(w, http.StatusOK, store.usersByMinAge(minAge))
}

func getUser(w http.ResponseWriter, id int, store *Store) {
	user, ok := store.getUser(id)
	if !ok {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}
	writeJSON(w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request, store *Store) {
	req, ok := decodeUserRequest(w, r)
	if !ok {
		return
	}
	user := store.createUser(req)
	writeJSON(w, http.StatusCreated, user)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int, store *Store) {
	req, ok := decodeUserRequest(w, r)
	if !ok {
		return
	}
	user, ok := store.updateUser(id, req)
	if !ok {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}
	writeJSON(w, http.StatusOK, user)
}

func deleteUser(w http.ResponseWriter, id int, store *Store) {
	user, ok := store.deleteUser(id)
	if !ok {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}
	writeJSON(w, http.StatusOK, user)
}

func decodeUserRequest(w http.ResponseWriter, r *http.Request) (userRequest, bool) {
	defer r.Body.Close()

	var req userRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return userRequest{}, false
	}

	if strings.TrimSpace(req.Name) == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return userRequest{}, false
	}

	if req.Age < 0 {
		writeError(w, http.StatusBadRequest, "age must be greater than or equal to 0")
		return userRequest{}, false
	}

	return req, true
}

func idFromPath(path string) (int, bool) {
	idText := strings.TrimPrefix(path, "/users/")
	if idText == "" || strings.Contains(idText, "/") {
		return 0, false
	}

	id, err := strconv.Atoi(idText)
	if err != nil || id <= 0 {
		return 0, false
	}

	return id, true
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
