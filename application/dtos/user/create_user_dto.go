package user

type CreateUserDTO struct {
	Name string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	AddressID int `json:"addressID"`
}