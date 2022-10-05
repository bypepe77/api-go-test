package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

	moviesCtrl := movie.NewController(log)

	//Create API routes and middlewaress
	moviesCtrl.SetApiRoutes(router)

	portString := os.Getenv("PORT")

	port, err := strconv.Atoi(portString)
	if err != nil || port == 0 {
		port = 8080
	}

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

}
