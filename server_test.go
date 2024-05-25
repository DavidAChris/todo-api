package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func routerAndHttpTest() (*gin.Engine, *httptest.ResponseRecorder) {
	return setupRouter(), httptest.NewRecorder()
}

func TestPingRoute(t *testing.T) {
	router, w := routerAndHttpTest()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestCreateTask(t *testing.T) {
	newTask := TaskRequest{
		Description: "Hello",
		Completed:   false,
	}
	newTaskJson, _ := json.Marshal(newTask)
	router, w := routerAndHttpTest()
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(newTaskJson))
	req.Header.Add("Content-Type", "application/json")
	defer req.Body.Close()
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}
