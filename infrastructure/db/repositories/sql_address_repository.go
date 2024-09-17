package repositories

import (
	"database/sql"
	"ramengo/domain/models"
)

type SQLAddressRepository struct {
	db *sql.DB
}

func NewSQLAddressRepository(db *sql.DB) *SQLAddressRepository {
	return &SQLAddressRepository{ db }
}

func (repo *SQLAddressRepository) Save(address *models.AddressModel) bool {

	query := "INSERT INTO address (street, `number`, zipcode) VALUES(?, ?, ?)"

	result, err := repo.db.Exec(query, address.Street, address.Number, address.ZipCode)

	if err != nil {
		return false
	}

	id, err := result.LastInsertId()

	if err != nil {
		return false
	}
	address.Id = int(id)

	return true
}