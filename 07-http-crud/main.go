package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type userRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{}
var nextID = 1

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userByIDHandler)

	fmt.Println("server started: http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("server error:", err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w)
	case http.MethodPost:
		createUser(w, r)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := idFromPath(r.URL.Path)
	if !ok {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	switch r.Method {
	case http.MethodGet:
		getUser(w, id)
	case http.MethodPut:
		updateUser(w, r, id)
	case http.MethodDelete:
		deleteUser(w, id)
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func getUsers(w http.ResponseWriter) {
	writeJSON(w, http.StatusOK, users)
}

func getUser(w http.ResponseWriter, id int) {
	for _, user := range users {
		if user.ID == id {
			writeJSON(w, http.StatusOK, user)
			return
		}
	}

	writeError(w, http.StatusNotFound, "user not found")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	req, ok := decodeUserRequest(w, r)
	if !ok {
		return
	}

	user := User{
		ID:   nextID,
		Name: req.Name,
		Age:  req.Age,
	}
	nextID++
	users = append(users, user)

	writeJSON(w, http.StatusCreated, user)
}

func updateUser(w http.ResponseWriter, r *http.Request, id int) {
	req, ok := decodeUserRequest(w, r)
	if !ok {
		return
	}

	for i := range users {
		if users[i].ID == id {
			users[i].Name = req.Name
			users[i].Age = req.Age
			writeJSON(w, http.StatusOK, users[i])
			return
		}
	}

	writeError(w, http.StatusNotFound, "user not found")
}

func deleteUser(w http.ResponseWriter, id int) {
	for i := range users {
		if users[i].ID == id {
			deleted := users[i]
			users = append(users[:i], users[i+1:]...)
			writeJSON(w, http.StatusOK, deleted)
			return
		}
	}

	writeError(w, http.StatusNotFound, "user not found")
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
