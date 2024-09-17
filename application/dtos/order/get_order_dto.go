package order

type GetOrderDTO struct {
	Id        int    `json:"id"`
	KitchenId int    `json:"kitchenId"`
	Date      string `json:"date"`
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	Status    string `json:"status"`
	StatusId  int    `json:"statusId"`
}
