package util

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

//check for valid admin token
func JWTAuth() gin.HandlerFunc{
	return func(context *gin.Context){
		err :=ValidateJWT(context)
		if err!=nil{
			context.JSON(http.StatusUnauthorized, gin.H{"error":"Authentication required"})
			context.Abort()
			return
		}
		error := ValidateAdminRoleJWT(context)
		if error !=nil{
			context.JSON(http.StatusUnauthorized, gin.H{"error":"Only Administrator is allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}

//Check for valid cutsotmer token
func JWTAuthCustomer() gin.HandlerFunc{
	return func(context *gin.Context){
		err := ValidateJWT(context)
		if err!=nil{
			context.JSON(http.StatusUnauthorized, gin.H{"error":"Authentication required"})
			context.Abort()
			return
		}
		error :=ValidateCustomerRoleJWT(context)
		if error !=nil{
			context.JSON(http.StatusUnauthorized, gin.H{"error":"Only registered customers are allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}