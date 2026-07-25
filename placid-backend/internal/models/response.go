package models

type UserResponse struct {
	Uuid      string `json:"uuid"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"isAdmin"`
	IsPremium bool   `json:"isPremium"`
	CreatedAt string `json:"createdAt"`
}
