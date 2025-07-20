package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"chatapp/handlers"
	"chatapp/middleware"
	"chatapp/utils"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	utils.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/api/users", middleware.AuthMiddleware(handlers.GetAllUsers)).Methods("GET")
	r.HandleFunc("/api/profile", middleware.AuthMiddleware(handlers.GetProfile)).Methods("GET")
	r.HandleFunc("/api/profile", middleware.AuthMiddleware(handlers.UpdateProfile)).Methods("PUT")
	r.HandleFunc("/api/messages/{userId}", middleware.AuthMiddleware(handlers.GetMessages)).Methods("GET")
	r.HandleFunc("/ws", middleware.SocketMiddleware(handlers.WebSocketHandler))

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	}).Handler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
