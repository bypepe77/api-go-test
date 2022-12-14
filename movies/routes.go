package movie

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) SetApiRoutes(r *gin.Engine) {
	r.GET("/movies", ctrl.GetMovies)
	r.GET("/movies/:id", ctrl.GetMovie)
}
