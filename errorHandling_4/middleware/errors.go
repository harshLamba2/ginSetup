package middleware;

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func ErrorHandler() gin.HandlerFunc { 
	return func(c *gin.Context){

		c.Next(); // processing

		fmt.Println("error middleware has been hit.");

		if( len(c.Errors) > 0 ){
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":c.Errors.String(),
			})
		}

	}

}