package main;
import (
	"middlewarePractice/middleware"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main(){

	fmt.Println("Initialized");

	
	router:= gin.Default(); // gin.Default() includes the Logger and Recovery middleware by default


	// Apply the AuthMiddleware from the middleware package to the route group
	private:= router.Group("/private");

	private.Use(middleware.AuthMiddleware());

	private.GET("/data", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"Authenticated",
			"success":true,
		});
	});




	router.GET("/", func(c *gin.Context){

		c.JSON(200, gin.H{
			"message":"Successfull Configued",
		})
	});

	router.Run(":8080"); //r.Run() // By default, it serves on :8080

}