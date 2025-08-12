package main

import (
	"fmt"
	"net/http"

	handlers "github.com/Rahul-Kumar-prog/Autotwt/apps/backend/api/Handlers"
)

type TwitterConfig struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func main() {

	http.HandleFunc("/api/post", handlers.PostRequest)

	fmt.Println("Server is running on post :8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("server is not running")
	}
}
