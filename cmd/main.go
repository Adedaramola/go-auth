package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adedaramola/golang-auth/internal/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(".env file is missing")
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Ping).Methods("GET")
	router.HandleFunc("/register", RegisterUser).Methods("POST")
	router.HandleFunc("/login", AttemptToAuthenticate).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("GET")

	_, err := database.NewConnection(env("DB_URL", ""), true)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	log.Println("Database connection established")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", env("APP_PORT", "5000")),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Handler:      router,
	}

	log.Printf("Server started and running at %s\n", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func env(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
