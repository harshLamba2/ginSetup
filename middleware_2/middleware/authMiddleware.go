package middleware; // make it into a package if you want to interchangably use it

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{ // You need to capitalize the middleware package func names for Gin to access it in other modules
	return func(c *gin.Context){
		token:= c.GetHeader("Authorization");

		if token!="Bearer"{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":"Invalid Token",
			});
			return;
		}

		c.Next();
	}
	
}