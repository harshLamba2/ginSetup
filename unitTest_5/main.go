package main;

import (
	"net/http"
	"tests/controller"
	"github.com/gin-gonic/gin"
)

func main(){

	router:=gin.Default();

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"success":true,
		});
	});

	router.GET("/getUser", controller.GetUser)

	router.Run(":8080")

}