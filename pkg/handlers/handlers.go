package handlers

import (
	"net/http"
	models "spire/pkg/models"
	"spire/pkg/views"
)

const (
	BS          string = "bootstrap"
	ProfilePath string = "E://spire/web/templates/profile.html"
	WallPath    string = "E://spire/web/templates/wall.html"
	FriendsPath string = "E:/spire/web/templates/friends.html"
	FAQPath     string = "E:/spire/web/templates/faq.html"
	NewUserPath string = "E:/spire/pkg/views/users/new.html"
)

var NotFound http.Handler = http.HandlerFunc(PageNotFoundHandler)

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	ProfileView := views.NewView(BS, ProfilePath)
	err := ProfileView.Render(w, models.Admin)
	if err != nil {
		panic(err)
	}

}

func WallHandler(w http.ResponseWriter, r *http.Request) {
	WallView := views.NewView(BS, WallPath)
	err := WallView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	FriendsView := views.NewView(BS, FriendsPath)
	err := FriendsView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func FAQHandler(w http.ResponseWriter, r *http.Request) {
	FAQView := views.NewView(BS, FAQPath)
	err := FAQView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}
