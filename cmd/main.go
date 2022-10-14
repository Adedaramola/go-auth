package main

import (
	"fmt"
	"log"
	"os"

	"github.com/adedaramola/golang-auth/datastore"
	"github.com/adedaramola/golang-auth/internal/pkg/server"
	"github.com/adedaramola/golang-auth/internal/transport"
	"github.com/adedaramola/golang-auth/services"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(".env file is missing")
	}
}

func main() {
	db, err := datastore.NewConnection(env("DB_URL", ""), true)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// register application services
	userService := services.NewUserService(db)

	// register route handlers
	h := transport.NewHandler(userService)

	// start http server
	server := server.NewServer(9000)
	server.SetupRoutes(h.RegisterRoutes())

	fmt.Println("Server started and running on http://localhost:9000")

	server.Listen()
}

func env(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
