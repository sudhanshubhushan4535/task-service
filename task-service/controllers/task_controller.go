package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"task-service/models"
	"task-service/storage"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	task.ID = storage.NextId
	storage.NextId++
	storage.Tasks = append(storage.Tasks, task)
	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	status := strings.ToLower(c.Query("status"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	var filtered []models.Task
	for _, t := range storage.Tasks {
		if strings.ToLower(t.Status) == status || t.Status == "" {
			filtered = append(filtered, t)
		}
	}
	start := (page - 1) * limit
	end := start + limit
	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	c.JSON(http.StatusOK, filtered[start:end])
}

func GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, t := range storage.Tasks {
		if t.ID == id {
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updated models.Task
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	for i, t := range storage.Tasks {
		if t.ID == id {
			storage.Tasks[i].Title = updated.Title
			storage.Tasks[i].Status = updated.Status
			c.JSON(http.StatusOK, t)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range storage.Tasks {
		if t.ID == id {
			storage.Tasks = append(storage.Tasks[:i], storage.Tasks[i+1:]...)
			c.JSON(http.StatusOK, http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
