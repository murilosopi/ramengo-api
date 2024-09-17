package models

type KitchenModel struct {
	Id      int
	Address *AddressModel
	Orders  []*OrderModel
	Users   []*UserModel
}
