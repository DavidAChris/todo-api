package query

import (
	"davidc/todo-api/database"
	"davidc/todo-api/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func SelectAllTodos() ([]models.Task, error) {
	allTasks := []models.Task{}
	var readErr error
	database.Read(func(db *sqlx.DB) {
		readErr = db.Select(&allTasks, "SELECT * FROM todos;")
		log.Printf("Read Tasks: %v\n", allTasks)
	})
	if readErr != nil {
		return nil, readErr
	}
	return allTasks, nil
}

func CreateTask(task models.TaskRequest) error {
	var writeErr error
	database.Write(func(db *sqlx.DB) error {
		tx := db.MustBegin()
		_, writeErr = tx.NamedExec("INSERT INTO todos (task) VALUES (:task)", &task)
		tx.Commit()
		return writeErr
	})
	return writeErr
}
