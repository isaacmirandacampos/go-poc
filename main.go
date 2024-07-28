package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	EMAIL string `json:"email"`
}

var users = []User{
	{ID: 1, NAME: "John Doe", EMAIL: "john.doe@example.com"},
	{ID: 2, NAME: "Jane Smith", EMAIL: "jane.smith@example.com"},
}

func main() {
	gin_server := gin.Default()
	gin_server.POST("users", func(ctx *gin.Context) {
		var user User

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, userExists := findUserByID(user.ID)
		if userExists {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "User already exists",
			})
			return
		}
		users = append(users, user)

		ctx.JSON(http.StatusOK, gin.H{
			"id":    user.ID,
			"name":  user.NAME,
			"email": user.EMAIL,
		})
	})

	gin_server.GET("users/:user_id", func(ctx *gin.Context) {
		user_id_str := ctx.Param("user_id")
		user_id, err := strconv.Atoi(user_id_str)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "user_id must be a number",
			})
			return
		}
		user, found := findUserByID(user_id)

		if !found {
			ctx.JSON(404, gin.H{
				"message": "Not found",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"id":    user.ID,
			"name":  user.NAME,
			"email": user.EMAIL,
		})
	})
	gin_server.Run(":8080")
}

func findUserByID(userID int) (User, bool) {
	for _, user := range users {
		if user.ID == userID {
			return user, true
		}
	}
	return User{}, false
}
