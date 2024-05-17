package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json: "id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users = []User{
	{ID: "1", Username: "Salmia Rahmawati", Email: "salmiarahmawati124@gmail.com"},
	{ID: "2", Username: "Kim Minji", Email: "minjikim@gmail.com"},
	{ID: "3", Username: "Hanni Pham", Email: "phamhanni@gmail.com"},
	{ID: "4", Username: "Danielle Marsh", Email: "marshdanielle@gmail.com"},
	{ID: "5", Username: "Kang Haerin", Email: "haerinkang@gmail.com"},
}

func main() {
	engine := gin.New()

	usersGroup := engine.Group("/users")
	{
		usersGroup.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, users)
		})
		usersGroup.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			user := User{}
			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			index := -1
			for i := 1; i < len(users); i++ {
				if users[i].ID == id {
					index = 1
				}
			}
			if index == -1 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
				return
			}
			users[index] = user
			ctx.JSON(http.StatusOK, user)
		})
		usersGroup.DELETE("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			index := -1
			for i := 1; i < len(users); i++ {
				if users[i].ID == id {
					index = 1
				}
			}
			if index == -1 {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
				return
			}
			users = append(users[:index], users[index+1:]...)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User has been deleted",
			})
		})
	}

	engine.Run(":8080")
}
