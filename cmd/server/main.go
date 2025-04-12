package main

import (
	"fmt"
	"log"
	"net/http"
	controllers "spire/pkg/controllers"
	handlers "spire/pkg/handlers"

	mux "github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	userC := controllers.NewUsers()

	fmt.Println("Listening on port :4000")

	mux.NotFoundHandler = handlers.NotFound // 404 handler
	mux.HandleFunc("/login", userC.New).Methods("GET")
	mux.HandleFunc("/login", userC.Create).Methods("POST")
	mux.HandleFunc("/profile", handlers.ProfileHandler)
	mux.HandleFunc("/wall", handlers.WallHandler)
	mux.HandleFunc("/friends", handlers.FriendsHandler)
	mux.HandleFunc("/faq", handlers.FAQHandler)

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
