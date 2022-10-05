package movie

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
