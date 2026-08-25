package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"github.com/dada/recipe-api/config"
	"github.com/dada/recipe-api/handlers"
)

func main() {
	// Load .env
	godotenv.Load()

	databaseURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Init Database
	if err := config.InitDB(databaseURL); err != nil {
		log.Fatal(err)
	}
	defer config.CloseDB()

	// Create Fiber app
	app := fiber.New()

	// CORS middleware
	app.Use(cors.New())

	// Routes
	api := app.Group("/api")
	recipes := api.Group("/recipes")

	recipes.Get("", handlers.GetRecipes)          // GET all
	recipes.Get("/:id", handlers.GetRecipe)       // GET one
	recipes.Post("", handlers.CreateRecipe)       // CREATE
	recipes.Put("/:id", handlers.UpdateRecipe)    // UPDATE
	recipes.Delete("/:id", handlers.DeleteRecipe) // DELETE

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	fmt.Printf("🚀 Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
