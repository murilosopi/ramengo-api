package repositories

import (
	"ramengo/domain/enums"
	"ramengo/domain/models"
	"time"
)

type KitchenRepository interface {
	FindOrdersByDate(kitchenID int, date time.Time) []*models.OrderModel
	FindOrdersByDiffentStatus(kitchen *models.KitchenModel, status enums.OrderStatus) []*models.OrderModel
	UserNotIncludedForKitchen(userID, kitchenID int) bool
	AddUser(userID, kitchenID int) bool
}

