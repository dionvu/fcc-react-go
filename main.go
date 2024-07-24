package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// Ptr to MongoDB project collection
var collection *mongo.Collection

func main() {
	// Load environment variables
	config()

	// Connect to MongoDB
	DB_URL := os.Getenv("DB_URL")
	client := connectDB(DB_URL)
	defer client.Disconnect(context.Background())

	// Update collection
	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	setUpRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
