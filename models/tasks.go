package models

type Task struct {
	Id uint64 `json:"id" db:"id"`
	TaskRequest
}

type TaskRequest struct {
	Description string `json:"task" db:"task"`
	Completed   uint8  `json:"completed" db:"completed"`
}
