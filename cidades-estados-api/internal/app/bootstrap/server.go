package bootstrap

import (
	"fmt"

	"github.com/GeanBarros/internal/app/handlers/locations"
	repositories "github.com/GeanBarros/internal/infraestructure/repositories/location"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	e := gin.Default()
	configureRoutes(e)
	err := e.Run(":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server Started")
}

func configureRoutes(e *gin.Engine) {

	locationRepository := repositories.NewLocationRepository()
	locationHandler := locations.NewLocationHandler(locationRepository)

	g := e.Group("/api/v1")
	{
	    g.GET("/states", locationHandler.GetAllStates)
	}
}