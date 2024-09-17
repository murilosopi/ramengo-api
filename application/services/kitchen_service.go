package services

import (
	"errors"
	"ramengo/application/dtos/kitchen"
	"ramengo/application/dtos/order"
	"ramengo/domain/repositories"
	"time"
)

type KitchenService struct {
	kitchenRepo repositories.KitchenRepository
}

func NewKitchenService(repo repositories.KitchenRepository) KitchenService {
	return KitchenService{repo}
}

func (ks *KitchenService) GetCurrentOrders(kitchenID int) (result []order.GetOrderDTO) {
	date := time.Now()
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	orders := ks.kitchenRepo.FindOrdersByDate(kitchenID, date)

	for _, o := range orders {
		result = append(result, order.GetOrderDTO{
			Id:        o.Id,
			KitchenId: o.Kitchen.Id,
			Date:      o.Date.Format(time.DateOnly),
			UserId:    o.User.Id,
			UserName:  o.User.Name,
			Status:    o.Status.String(),
			StatusId:  int(o.Status),
		})
	}

	return
}

func (ks *KitchenService) GetNotReadyOrders(kitchenID int) (result []order.GetOrderDTO) {
	date := time.Now()
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	orders := ks.kitchenRepo.FindOrdersByDate(kitchenID, date)

	for _, o := range orders {
		result = append(result, order.GetOrderDTO{
			Id:        o.Id,
			KitchenId: o.Kitchen.Id,
			Date:      o.Date.Format(time.DateOnly),
			UserId:    o.User.Id,
			UserName:  o.User.Name,
			Status:    o.Status.String(),
			StatusId:  int(o.Status),
		})
	}

	return
}

func (ks *KitchenService) AddUser(dto kitchen.AddUserKitchenDTO) error {
	userNotIncluded := ks.kitchenRepo.UserNotIncludedForKitchen(dto.UserID, dto.KitchenID)

	if !userNotIncluded {
		return errors.New("user already included")
	}

	if !ks.kitchenRepo.AddUser(dto.UserID, dto.KitchenID) {
		return errors.New("failure on including user")
	}

	return nil
}
