package routers

import (
	controllers "restAPI/api/v1/Controllers"

	"github.com/gin-gonic/gin"
)


func RoutesHandler(routes *gin.Engine) {

	/* MOVIES */
	routes.POST("/movies", controllers.PostMovies)
	routes.GET("/movies",controllers.GetMovies)
	routes.PUT("/movies",controllers.PutMovies)
	routes.DELETE("/movies",controllers.DELETEMovies)

	routes.Run(":9090")
}
