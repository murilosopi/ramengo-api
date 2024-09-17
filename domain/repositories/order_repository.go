package repositories

import "ramengo/domain/models"

type OrderRepository interface {
	Save(*models.OrderModel) bool
	ChangeStatus(*models.OrderModel) bool
	FindById(id int) *models.OrderModel
}