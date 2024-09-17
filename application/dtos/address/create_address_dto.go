package address

type CreateAddressDTO struct {
	Street  string `json:"street" validate:"required"`
	ZipCode string `json:"zipCode" validate:"required"`
	Number  int    `json:"number" validate:"required"`
}
