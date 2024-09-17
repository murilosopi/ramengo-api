package services

import (
	"ramengo/application/dtos/address"
	"ramengo/domain/models"
	"ramengo/domain/repositories"
)

type AddressService struct {
	adddressRepo repositories.AddressRepository
}

func NewAddressService(repo repositories.AddressRepository) AddressService {
	return AddressService{repo}
}

func (as *AddressService) Save(dto *address.CreateAddressDTO) (success bool, id int)  {
	model := models.AddressModel{
		Street: dto.Street,
		ZipCode: dto.ZipCode,
		Number: dto.Number,
	}

	success = as.adddressRepo.Save(&model)

	if success {
		id = model.Id
	}

	return
}
