package main

import (
	"net/http"
	//"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
  )

func main() {
	//setup router
	router := gin.Default()

	//serve frontend static files
	//router.Use(static.Serve("/",static.LocalFile("./gin-client/public",true)))

	//setup API route group
	api := router.Group("/api") 
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {[
				{
					"Name": "Negroni",
					"Ingredients": [
						{"Name": "London dry gin", "Quantity": "1 oz"},
						{"Name": "Sweet Vermouth", "Quantity": "1 oz"},
						{"Name": "Campari", "Quantity": "1 oz"}
					],
					"Instructions": [
						"Stir with ice for 20-30 seconds",
						"Strain into coupe glass",
						"Garnish with orange peel"
					]
				},
				{
					"Name": "Gin & Tonic",
					"Ingredients": [
						{"Name": "London dry gin", "Quantity": "1 oz"},
						{"Name": "Tonic Water", "Quantity": "1 oz"},
						{"Name": "Lime wedge", "Quantity": "1"}
					],
					"Instructions": [
						"Pour over ice",
						"Garnish with lime wedge"
					]
				}
			]})
		})
	}

	//start & run server
	router.Run(":5000")
}
