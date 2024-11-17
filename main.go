package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ChanchalS7/practice_api/database"
	"github.com/ChanchalS7/practice_api/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	loadEnv()

	// Load database configuration, perform migrations, and seed data
	loadDatabase()

	// Start the server
	serveApplication()
}

// Load environment variables from .env file
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

// Initialize database, run migrations, and seed data
func loadDatabase() {
	// Initialize the database connection
	database.InitDb()

	// Run migrations
	err := database.Db.AutoMigrate(&model.Role{}, &model.User{})
	if err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database migrated successfully")

	// Seed data
	seedData()
}

// Seed initial data into the database
func seedData() {
	// Seed roles
	roles := []model.Role{
		{Name: "admin", Description: "Administrator role"},
		{Name: "customer", Description: "Authenticated customer role"},
		{Name: "anonymous", Description: "Unauthenticated customer role"},
	}

	// Seed a default admin user
	users := []model.User{
		{
			Username: os.Getenv("ADMIN_USERNAME"),
			Email:    os.Getenv("ADMIN_EMAIL"),
			Password: os.Getenv("ADMIN_PASSWORD"),
			RoleID:   1, // Assuming RoleID 1 corresponds to "admin"
		},
	}

	// Save roles to the database
	for _, role := range roles {
		database.Db.FirstOrCreate(&role, model.Role{Name: role.Name})
	}
	log.Println("Roles seeded successfully")

	// Save users to the database
	for _, user := range users {
		database.Db.FirstOrCreate(&user, model.User{Email: user.Email})
	}
	log.Println("Users seeded successfully")
}

// Start the application server
func serveApplication() {
	router := gin.Default()
	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
