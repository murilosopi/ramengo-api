package repositories

import (
	"ramengo/domain/models"
)

type UserRepository interface {
	Save(*models.UserModel) bool
	OrderHistory(userId int) []*models.OrderModel
	VerifyEmailAvailable(email string) bool
}
