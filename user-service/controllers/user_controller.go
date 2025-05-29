package controllers

import (
	"net/http"
	"strconv"
	"user-service/storage"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, u := range storage.Users {
		if u.ID == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
}
