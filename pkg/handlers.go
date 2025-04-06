package handlers

import (
	"fmt"
	"net/http"
	models "spire/pkg/models"
	"spire/pkg/views"
)

// views variables are here.. for now

var ProfileView *views.View
var WallView *views.View
var FriendsView *views.View

var NotFound http.Handler = http.HandlerFunc(PageNotFoundHandler)

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<title>%d</title>", http.StatusNotFound)
	fmt.Fprintf(w, "<b>%d Page Not Found</b>", http.StatusNotFound)
}

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("X-Content-Type-Options", "nosniff")
	ProfileView = views.NewView("E://spire/web/templates/userprofile.html")
	err := ProfileView.Template.ExecuteTemplate(w, "userprofile", models.Admin)
	if err != nil {
		panic(err)
	}

}

func WallHandler(w http.ResponseWriter, r *http.Request) {
	WallView = views.NewView("E://spire/web/templates/wall.html")
	err := WallView.Template.ExecuteTemplate(w, "wall", nil)
	if err != nil {
		panic(err)
	}
}

func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	FriendsView = views.NewView("E://spire/web/templates/friends.html")
	err := FriendsView.Template.ExecuteTemplate(w, "friends", nil)
	if err != nil {
		panic(err)
	}
}
