package repositories

import "ramengo/domain/models"

type AddressRepository interface {
	Save(*models.AddressModel) bool
}