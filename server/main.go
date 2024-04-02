package main

import (
	"awais-go-react-calorie/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()

	// WHAT DOES LOGGER DO HERE!
	router.Use(gin.Logger())
	// WHEN WHICH API WAS CALLED ->
	// IN THE TERMINAL ->
	// FIND OUT ERRORS EXACTLY WHICH API CAUSED THE ERROR!

	// USING CORS!
	router.Use(cors.Default())

	// OUR ROUTES!
	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("/entry/:id", routes.GetEntryById)
	router.GET("/ingredients", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredient/update/:id", routes.UpdateIngredient)
	router.DELETE("/entry/delete/:id", routes.DeleteEntry)

}
