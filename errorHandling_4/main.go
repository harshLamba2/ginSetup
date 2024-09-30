package main;

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"errors"
	"errorHandling/middleware"
)



func main(){
	fmt.Println("hello world");

	router:= gin.Default();

	router.Use(middleware.ErrorHandler())

	router.GET("/", func(c *gin.Context){
		fmt.Println("Hit");

		if 1<10 {
			// c.Error("1 is less Than 10"); // this method expects an argumenet with type error; 
			err:= errors.New("1 is less than 10");
			c.Error(err); // adds/ collects error in the Context which can be processed later
			return;
			// returning the function the ErrorHandler middleware would still work after this function has finidhed executing
		}

		c.JSON(200, gin.H{ //gin.H is just a shorthand for map[string]interface{}
			"success":true, 
			"message":"API Responded",
		});
	});









	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	router.GET("/getUser", func(c *gin.Context){

		var user= User{
			Name:"Harsh",
			Email:"lambaharsh01@gmail.com",
			Age:22,
		};


		c.JSON(200, gin.H{
			"success":true,
			"data": user, 
		});


	})


	router.Run(":8080")
}