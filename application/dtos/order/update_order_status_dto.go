package order

type UpdateOrderStatusDTO struct {
	Id        int `json:"id"`
	UserID    int
	KitchenID int
	Status    int
}
