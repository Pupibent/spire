package handlers

import (
	"net/http"
	models "spire/pkg/models"
	"spire/pkg/views"
)

const (
	bs          string = "bootstrap"
	profilePath string = "E://spire/web/templates/userprofile.html"
	wallPath    string = "E://spire/web/templates/wall.html"
	friendsPath string = "E:/spire/web/templates/friends.html"
	FAQPath     string = "E:/spire/web/templates/faq.html"
)

var NotFound http.Handler = http.HandlerFunc(PageNotFoundHandler)

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("X-Content-Type-Options", "nosniff")
	ProfileView := views.NewView(bs, profilePath)
	err := ProfileView.Render(w, models.Admin)
	if err != nil {
		panic(err)
	}

}

func WallHandler(w http.ResponseWriter, r *http.Request) {
	WallView := views.NewView(bs, wallPath)
	err := WallView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	FriendsView := views.NewView(bs, friendsPath)
	err := FriendsView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

func FAQHandler(w http.ResponseWriter, r *http.Request) {
	FAQView := views.NewView(bs, FAQPath)
	err := FAQView.Render(w, nil)
	if err != nil {
		panic(err)
	}
}

// TODO: FAQ handler
