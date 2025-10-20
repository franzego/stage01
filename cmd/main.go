package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/franzego/stage01/internal"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()
	r.POST("/strings", internal.PostString)
	r.GET("/strings/:string_value", internal.GetString)

	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		if err := r.Run("0.0.0.0:" + port); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	log.Println("Shutting down gracefully...")
}
