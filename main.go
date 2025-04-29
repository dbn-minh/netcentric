package main

import (
	"fmt"
	"net/http"

	"netcentric-nosql/config"
	"netcentric-nosql/controllers"
)

func main() {
	config.ConnectDB()

	controllers.InitUserCollection()
	controllers.InitMessageCollection()
	controllers.InitConversationCollection()

	// Register routes
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/user", controllers.CreateUser)

	http.HandleFunc("/messages", controllers.GetMessages)
	http.HandleFunc("/message", controllers.CreateMessage)

	http.HandleFunc("/conversations", controllers.GetConversations)
	http.HandleFunc("/conversation", controllers.CreateConversation)

	fmt.Println("🚀 Server is running at http://localhost:8080 ...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("❌ Server failed to start: %v\n", err)
	}
}
