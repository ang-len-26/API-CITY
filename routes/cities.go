package routes

import (
	"github.com/ang-len-26/go-open-data-api/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes the routes for the API
// and associates them with their respective handlers.
func RegisterCityRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/cities", controllers.GetCities)
		api.GET("/cities/:id", controllers.GetCityByID)
	}
}
