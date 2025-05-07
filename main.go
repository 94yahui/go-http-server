package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string
	Username string
	Email    string
	Status   bool
}

var users = []User{
	{ID: "1", Username: "john_doe", Email: "john@example.com", Status: true},
	{ID: "2", Username: "jane_smith", Email: "jane@example.com", Status: true},
	{ID: "3", Username: "alice_wonderland", Email: "alice@example.com", Status: true},
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func toggleUserStatus(c *gin.Context) {
	id := c.Param("id")
	for i := range users {
		if users[i].ID == id {
			users[i].Status = !users[i].Status
			c.IndentedJSON(http.StatusOK, users[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func main() {
	router := gin.Default()
	router.GET("/", getUser)
	router.GET("/:id", getUserByID)
	router.POST("/", addUser)
	router.PATCH("/:id", toggleUserStatus)
	router.Run("localhost:8080")
}


