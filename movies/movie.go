package movie

import (
	"fmt"
	"net/http"
	"test-api/models"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetMovies(c *gin.Context) {

	author := models.Author{Name: "Pepe"}
	//no-lint
	createAuthor := ctrl.db.Create(&author)

	fmt.Println(createAuthor)

	movie := models.Movie{Title: "PepeLand", Year: "2022", Rating: 8, AuthorID: author.ID, Author: &author}

	//no-lint

	createMovie := ctrl.db.Create(&movie)

	if createMovie.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": createMovie.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})

}

func (ctrl *Controller) GetMovie(c *gin.Context) {
	var movie models.Movie

	err := ctrl.db.Preload("Author").First(&movie, c.Param("id"))

	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})

}
