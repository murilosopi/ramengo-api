package controllers

import (
	"net/http"
	"ramengo/application/dtos/kitchen"
	"ramengo/application/dtos/order"
	"ramengo/application/services"
	"ramengo/domain/enums"
	"ramengo/infrastructure/security"

	"github.com/labstack/echo/v4"
)

type KitchenController struct {
	KitchenService services.KitchenService
	OrderService   services.OrderService
}

func NewKitchenController(s services.KitchenService, o services.OrderService) KitchenController {
	return KitchenController{s, o}
}

func (kc KitchenController) GetCurrentOrders(c echo.Context) error {
	payload := c.Get(string(security.Kitchen)).(*security.JWTClaims)

	idKitchen := payload.Id

	slice := kc.KitchenService.GetCurrentOrders(idKitchen)

	return c.JSON(http.StatusAccepted, slice)
}

func (kc KitchenController) CancelNotReadyOrders(c echo.Context) error {
	payload := c.Get(string(security.Kitchen)).(*security.JWTClaims)

	kitchenID := payload.Id

	orders := kc.KitchenService.GetNotReadyOrders(kitchenID)

	var mapUpdatedOrders = make(map[int]bool, len(orders))

	for _, o := range orders {
		order := order.UpdateOrderStatusDTO{
			Id:        o.Id,
			UserID:    o.UserId,
			KitchenID: o.KitchenId,
			Status:    int(enums.CancelledStatus),
		}

		mapUpdatedOrders[order.Id], _ = kc.OrderService.ChangeStatus(&order)
	}

	return c.JSON(http.StatusOK, mapUpdatedOrders)
}

func (kc KitchenController) AddUser(c echo.Context) error {
	payload := c.Get(string(security.Kitchen)).(*security.JWTClaims)

	userKitchenDTO := &kitchen.AddUserKitchenDTO{
		KitchenID: payload.Id,
	}

	if err := c.Bind(userKitchenDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := kc.KitchenService.AddUser(*userKitchenDTO); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, true)
}
