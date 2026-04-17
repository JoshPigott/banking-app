package main

import (
	"banking-app/internal/database"
	"banking-app/internal/server"
	"banking-app/internal/services"
	"fmt"
	"net/http"
)

func main() {
	err := database.InitDB()
	if err != nil {
		fmt.Printf("Error connection with database %v\n", err)
		return
	}
	defer database.DB.Close()
	services.CleanUpSessions()

	// I am also going to have to server static files here look fabiens code
	server := &http.Server{
		Addr:    ":8080",
		Handler: server.Router(),
	}

	fmt.Println("server starting on port on http://localhost:8080")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
