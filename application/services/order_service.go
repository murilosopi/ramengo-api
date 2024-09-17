package services

import (
	"errors"
	"fmt"
	"ramengo/application/dtos/order"
	"ramengo/domain/enums"
	"ramengo/domain/models"
	"ramengo/domain/repositories"
)

type OrderService struct {
	orderRepo           repositories.OrderRepository
	notificationService NotificationService
}

func NewOrderService(repo repositories.OrderRepository, notificationService NotificationService) OrderService {
	return OrderService{repo, notificationService}
}

func (os *OrderService) Save(dto *order.CreateOrderDTO) (success bool, responseDTO order.CreateOrderResponseDTO) {
	orderModel := models.OrderModel{
		User:    &models.UserModel{Id: dto.UserID},
		Kitchen: &models.KitchenModel{Id: 1}, // todo: calc geolocalization distance
		Status:  enums.ConfirmedStatus,
	}

	success = os.orderRepo.Save(&orderModel)

	if success {
		message := fmt.Sprintf("Your order #%v is confirmed", orderModel.Id)
		os.notificationService.Send(orderModel.User, message)

		responseDTO = order.CreateOrderResponseDTO{
			Id:        orderModel.Id,
			KitchenID: orderModel.Kitchen.Id,
			UserID:    orderModel.User.Id,
		}
	}

	return
}

func (os *OrderService) ChangeStatus(dto *order.UpdateOrderStatusDTO) (bool, error) {

	// validate DTO status
	if !enums.IsValidOrderStatus(dto.Status) {
		return false, errors.New("invalid order status")
	}

	orderModel := os.orderRepo.FindById(dto.Id) // get order data

	// start validation block
	if orderModel == nil {
		return false, errors.New("order not found")
	}

	if orderModel.Kitchen.Id != dto.KitchenID {
		return false, errors.New("order not found for this kitchen")
	}

	newStatus := enums.OrderStatus(dto.Status)

	if newStatus < orderModel.Status {
		return false, errors.New("can not roll back the order status")
	}

	if newStatus == orderModel.Status {
		return false, errors.New("current status")
	}
	// end validation block

	orderModel.Status = newStatus

	saved := os.orderRepo.ChangeStatus(orderModel) // save new order status

	if saved {
		message := fmt.Sprintf("Your order #%v has changed status: %s", orderModel.Id, orderModel.Status)

		os.notificationService.Send(orderModel.User, message)
	}

	return saved, nil
}
