package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var taskDbLock = sync.RWMutex{}

type Task struct {
	Id int `json:"id"`
	TaskRequest
}

type TaskRequest struct {
	Description string `json:"task"`
	Completed   bool   `json:"completed"`
}

var tasksDb = make(map[int]Task)
var taskIdCount = 0

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", ping)
	router.GET("/tasks", getAllTasks)
	router.POST("/tasks", postCreateTask)
	router.PUT("/tasks/:id", putUpdateTask)
	router.DELETE("/tasks/:id", deleteTask)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}

func ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// Get all Tasks
func getAllTasks(ctx *gin.Context) {
	allTasks := make([]Task, 0, len(tasksDb))
	taskDbLock.RLock()
	for _, v := range tasksDb {
		allTasks = append(allTasks, v)
	}
	taskDbLock.RUnlock()
	ctx.IndentedJSON(http.StatusOK, allTasks)
}

// Create a task
func postCreateTask(ctx *gin.Context) {
	taskDbLock.Lock()
	newTaskId := taskIdCount
	taskIdCount += 1
	taskDbLock.Unlock()
	var newTask Task
	if err := ctx.BindJSON(&newTask); err != nil {
		return
	}
	newTask.Id = newTaskId
	taskDbLock.Lock()
	tasksDb[newTaskId] = newTask
	taskDbLock.Unlock()
	ctx.IndentedJSON(http.StatusCreated, newTask)
}

// Update a Task
func putUpdateTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	var updateTask TaskRequest
	if err := ctx.BindJSON(&updateTask); err != nil {
		return
	}
	taskDbLock.RLock()
	val, ok := tasksDb[id]
	taskDbLock.RUnlock()
	if !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"err": "task not found"})
		return
	}
	if val.Description != updateTask.Description {
		val.Description = updateTask.Description
	}
	if val.Completed != updateTask.Completed {
		val.Completed = updateTask.Completed
	}
	taskDbLock.Lock()
	tasksDb[id] = val
	taskDbLock.Unlock()
	ctx.IndentedJSON(http.StatusOK, val)
}

// Delete a Task
func deleteTask(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return
	}
	if _, ok := tasksDb[id]; !ok {
		ctx.JSON(http.StatusNotFound, gin.H{"err": "task does not exist"})
		return
	}
	taskDbLock.Lock()
	delete(tasksDb, id)
	taskDbLock.Unlock()
	ctx.Status(http.StatusNoContent)
}
