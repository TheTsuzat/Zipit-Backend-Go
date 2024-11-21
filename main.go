package main

import (
	"log"
	"os"

	"github.com/Tsuzat/zipit/src/config"
	"github.com/Tsuzat/zipit/src/db"
	"github.com/Tsuzat/zipit/src/routes"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	// Initialize variables
	config.Init()
	fiberlog.Info("Initializing Config and Global Variables")

	// Initialize database
	if err := db.InitDatabase(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fiberlog.Info("Database initialized successfully")
	defer config.DB.Close()

	// Initialize redis
	if err := db.InitRedis(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fiberlog.Info("Redis initialized successfully")
	defer config.RDB.Close()

	// Initialize server
	config.APP = fiber.New()

	// Initialize routes
	routes.InitRoutes()
	log.Fatal(config.APP.Listen(":8080"))
}
