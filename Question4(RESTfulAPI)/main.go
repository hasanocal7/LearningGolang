package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: "1", Name: "Hasan", Age: 22},
	{ID: "2", Name: "Mert", Age: 22},
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

// Create User
func createUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	users = append(users, user)
	c.IndentedJSON(http.StatusCreated, user)
}

// Get All Users
func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// Get User by ID
func getUserByID(id string) (*User, error) {
	for i, u := range users {
		if u.ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("User not found")
}

// Get User
func getUser(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

// Update User Age
func updateUserAge(c *gin.Context) {
	id := c.Param("id")
	user, err := getUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	user.Age++
	c.IndentedJSON(http.StatusOK, user)
}

// Delete User
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func main() {
	r := gin.Default()
	r.GET("/", Hello)
	r.POST("/users", createUser)
	r.GET("/users", getAllUsers)
	r.GET("/users/:id", getUser)
	r.PATCH("/users/:id", updateUserAge)
	r.DELETE("/users/:id", deleteUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
