package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{
		DB: db,
	}
}

func (taskHandler *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (taskHandler *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := taskHandler.DB.Exec("INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)", task.Title, task.Description, task.Status)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	task.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (taskHandler *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println("id = ", id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var task models.Task
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := taskHandler.DB.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?", task.Title, task.Description, task.Status, id)

	fmt.Println("result = ", result)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.ID = id // Ensure the ID is set in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (taskHandler *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	result, err := taskHandler.DB.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
