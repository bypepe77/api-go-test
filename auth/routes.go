package auth

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) SetApiRoutes(r *gin.Engine) {
	r.POST("/auth/register", ctrl.Register)
	r.POST("/auth/login", ctrl.Login)
}
