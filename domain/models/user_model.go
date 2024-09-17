package models

type UserModel struct {
	Id       int
	Name     string
	Email    string
	Password []byte
	Address  *AddressModel
	Kitchens []*KitchenModel
	Orders   []*OrderModel
}
