package middleware

import (
	"fmt"
	"jwt/models"
	"jwt/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnUnauthorized(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error: []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeUnauthorized,
				ErrorMessage: "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("apikey")
		referer := context.Request.Header.Get("Referer")
		valid, claims := token.VerifyToken(tokenString, referer)
		if !valid {
			ReturnUnauthorized(context)

		}
		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		context.Keys["ComapnyId"] = claims.CompanyId
		context.Keys["Username"] = claims.Username
		context.Keys["Roles"] = claims.Roles
	}
}
func Authorization(validRoles []int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Keys) == 0 {
			ReturnUnauthorized(ctx)
		}
		rolesVal := ctx.Keys["Roles"]
		fmt.Println("roles", rolesVal)
		if rolesVal == nil {
			ReturnUnauthorized(ctx)
		}
		roles := rolesVal.([]int)
		validation := make(map[int]int)
		for _, val := range roles {
			validation[val] = 0
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				ReturnUnauthorized(ctx)
			}
		}

	}
}
