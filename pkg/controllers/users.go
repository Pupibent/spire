package controllers

import (
	"fmt"
	"net/http"
	"spire/pkg/handlers"
	"spire/pkg/views"
)

type Credentials struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{NewView: views.NewView(handlers.BS, handlers.NewUserPath)}
}

// Method New calls to method Render of View to create and render a new View
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Method Create is used to process an Authorization form from a user with
// METHOD POST
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	AdmCreds := &Credentials{}
	if err := DecodeF(r, AdmCreds); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Email is %s\n", AdmCreds.Email)
	fmt.Fprintf(w, "Password is %s\n", AdmCreds.Password)
}
