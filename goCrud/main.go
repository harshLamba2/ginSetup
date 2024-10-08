package main;

import (
	"github.com/gin-gonic/gin"
	//"crud/database"
	"crud/routes"
)

func main(){
	router:=gin.Default();

	routes.CountriesRoute(router);

	router.Run(":8080");
}