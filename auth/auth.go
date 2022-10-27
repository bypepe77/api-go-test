package auth

import (
	"net/http"
	"test-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (ctrl *Controller) Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	userCreated := ctrl.db.Create(&user)

	if userCreated.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating user"})
		return

	}
	c.JSON(200, gin.H{"data": user})

}

func (ctrl *Controller) Login(c *gin.Context) {

}
