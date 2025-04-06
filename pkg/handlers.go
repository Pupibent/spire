package handlers

import (
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
	// fmt.Fprintf(w, "<title>%d</title>", http.StatusNotFound)
	// fmt.Fprintf(w, "<b>%d Page Not Found</b>", http.StatusNotFound)
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("X-Content-Type-Options", "nosniff")
	ProfileView = views.NewView("bootstrap", "E://spire/web/templates/userprofile.html")
	err := ProfileView.Template.ExecuteTemplate(w, ProfileView.Layout, models.Admin)
	if err != nil {
		panic(err)
	}

}

func WallHandler(w http.ResponseWriter, r *http.Request) {
	WallView = views.NewView("bootstrap", "E://spire/web/templates/wall.html")
	err := WallView.Template.ExecuteTemplate(w, WallView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	FriendsView = views.NewView("bootstrap", "E://spire/web/templates/friends.html")
	err := FriendsView.Template.ExecuteTemplate(w, FriendsView.Layout, nil)
	if err != nil {
		panic(err)
	}
}
