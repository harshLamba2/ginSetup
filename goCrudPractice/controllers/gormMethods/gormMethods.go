package gormMethods

import(
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FirstMethod(db * gorm.DB) gin.HandlerFunc{
	return func(c *gin.Context){


		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"message":"hello world",
		})


	}

}