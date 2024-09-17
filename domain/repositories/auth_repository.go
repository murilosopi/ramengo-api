package repositories

import "ramengo/domain/models"

type AuthRepository interface {
	FindUserByEmail(email string) *models.UserModel
	GetUserKitchenByID(userID int, kitchenID int) *models.KitchenModel
}
