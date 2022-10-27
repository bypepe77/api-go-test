package main

import (
	"os"
	"strconv"
	"test-api/auth"
	database "test-api/databaseSettings"
	movie "test-api/movies"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	log.Info("Starting server...")

	router := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//database connection
	db, _ := database.DatabaseConection()

	moviesCtrl := movie.NewController(log, db)
	authCtrl := auth.NewController(log, db)
	
	//Create API routes and middlewaress
	moviesCtrl.SetApiRoutes(router)
	authCtrl.SetApiRoutes(router)

	portString := os.Getenv("PORT")

	port, err := strconv.Atoi(portString)
	if err != nil || port == 0 {
		port = 8080
	}

	router.Run(":" + portString)
	/*
		server := &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      router,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		}

		go func() {
			log.Info("server started on port ", port)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()

	*/

}
