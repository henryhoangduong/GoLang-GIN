package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestModel struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func main() {
	r := gin.Default()
    
	user := r.Group("user")
	{
		user.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": []TestModel{
					TestModel{
						Id:   1,
						Name: "User1",
					},
					TestModel{
						Id:   2,
						Name: "User1",
					},
				},
			})
		})
		user.GET(":Id", func(c *gin.Context) {
			var id = c.Param("Id")
			fmt.Println("id is recieved is ", id)
			c.JSON(http.StatusOK, gin.H{
				"data": TestModel{
					Id:   1,
					Name: "User1",
				},
			})
		})
	}
    r.Run("localhost:8080")
}
