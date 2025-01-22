package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id    int    `uri:"id"`
	Name  string `uri:"name"`
	email string `uri:"email"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var userObj user
		if err := c.ShouldBindQuery(&userObj); err == nil {
			fmt.Printf("usr obj - %+v \n", userObj)
		} else {
			fmt.Printf("error - %+v \n", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	r.Run("localhost:8080")
}
