package domain

import "github.com/JscorpTech/jst-go/models"

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserRepository interface {
	IsAlready(phone string) bool
	Create(firstName, LastName, phone, password string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	FindById(id int) (*models.User, error)
}
