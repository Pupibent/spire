package models

type User struct {
	Name string
	ID   int
}

var Admin User = User{"Admin", 0001}
