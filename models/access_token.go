package models

type AccessToken struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}
