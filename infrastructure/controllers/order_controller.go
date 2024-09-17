package controllers

import (
	"net/http"
	"ramengo/application/dtos/order"
	"ramengo/application/services"
	"ramengo/infrastructure/security"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(os services.OrderService) OrderController {
	return OrderController{os}
}

func (oc OrderController) Save(c echo.Context) error {
	payload := c.Get(string(security.User)).(*security.JWTClaims)

	dto := new(order.CreateOrderDTO)
	dto.UserID = payload.Id

	saved, responseDTO := oc.orderService.Save(dto)

	if !saved {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusCreated, responseDTO)
}

func (oc OrderController) ChangeStatus(c echo.Context) error {
	dto := new(order.UpdateOrderStatusDTO)

	dto.KitchenID = 1 // todo: implement some authorization token id getter

	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	changed, err := oc.orderService.ChangeStatus(dto)

	if !changed && err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if !changed {
		return c.JSON(http.StatusInternalServerError, false)
	}

	return c.JSON(http.StatusOK, true)
}
