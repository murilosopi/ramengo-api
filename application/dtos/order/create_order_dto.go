package order

type CreateOrderDTO struct {
	UserID int
}

type CreateOrderResponseDTO struct {
	Id int
	UserID int
	KitchenID int
}