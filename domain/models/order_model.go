package models

import (
	"ramengo/domain/enums"
	"time"
)

type OrderModel struct {
	Id      int
	Kitchen *KitchenModel
	Status  enums.OrderStatus
	User    *UserModel
	Date    time.Time
}
