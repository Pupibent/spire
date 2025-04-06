package main

import (
	"fmt"
	"log"
	"net/http"
	handlers "spire/pkg"

	mux "github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	fmt.Println("Listening on port :4000")

	mux.NotFoundHandler = handlers.NotFound // 404 handler
	mux.HandleFunc("/profile", handlers.UserProfileHandler)
	mux.HandleFunc("/wall", handlers.WallHandler)
	mux.HandleFunc("/friends", handlers.FriendsHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
