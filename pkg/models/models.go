package models

type Profile struct {
	Name   string
	ID     int
	Avatar string
}

const adminAvatar = "https://assets.jobzmall.com/stores/3577522/images/a0a1bdfc-f844-438a-a89d-4eb218af79c8.jpg"

var Admin Profile = Profile{"Admin", 0001, adminAvatar}
