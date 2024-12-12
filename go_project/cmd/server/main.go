package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"

	"go_project/internal/config"
	"go_project/internal/database"
	"go_project/internal/graphql"
	"go_project/internal/routes"
	"go_project/internal/websocket"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the router
	r := gin.Default()

	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Connect to the database and create tables
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Create tables if they don't exist
	if err := database.CreateTables(db); err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	// Initialize routes
	api := r.Group("/api")
	routes.RegisterRoutes(api, db)

	// GraphQL setup with playground
	h := handler.New(&handler.Config{
		Schema:     &graphql.Schema,
		Pretty:     true,
		GraphiQL:   true, // Enable the GraphiQL interface
		Playground: true, // Enable the Playground interface
	})

	// Handle both GET and POST for GraphQL
	r.Any("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	// WebSocket setup
	r.GET("/ws", gin.WrapF(websocket.HandleConnections))
	go websocket.HandleMessages()

	// Start the server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// Add middleware
//r.Use(middleware.Logger())
//r.Use(middleware.ErrorHandler())
//r.Use(middleware.Auth())
