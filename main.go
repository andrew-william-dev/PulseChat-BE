package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"chatapp/handlers"
	"chatapp/middleware"
	"chatapp/utils"
)

func main() {
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
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	}).Handler(r) 

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
