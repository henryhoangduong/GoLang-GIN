package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestModel struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"id" binding:"required"`
}

func main() {
	r := gin.Default()
	v1 := r.Group("v1")
	v1.Use(func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authentication")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "no token",
			})
			return
		}
		if authHeader != "abc" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
		} else {
			if len(ctx.Keys) == 0 {
				ctx.Keys = make(map[string]interface{})
			}
			ctx.Keys["user"] = authHeader
			ctx.Next()
		}
	})
}
