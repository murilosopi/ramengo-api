package repositories

import (
	"database/sql"
	"fmt"
	"ramengo/domain/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewSQLAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (repo *AuthRepository) FindUserByEmail(email string) *models.UserModel {
	query := `SELECT
				u.id as user_id,
				u.name as user_name,
				u.email as user_email,
				u.password as user_password,
				a.id as address_id,
				a.street as address_street,
				a.number as address_number,
				a.zipcode as address_zipcode
			FROM
				users u
			JOIN address a ON
				u.address_id = a.id
			WHERE
				u.email = ?`

	row := repo.db.QueryRow(query, email)

	if row == nil {
		return nil
	}

	user := new(models.UserModel)
	user.Address = new(models.AddressModel)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Address.Id, &user.Address.Street, &user.Address.Number, &user.Address.ZipCode)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return user
}

func (repo *AuthRepository) GetUserKitchenByID(userID int, kitchenID int) *models.KitchenModel {
	query := `SELECT
				k.id as kitchen_id,
				a.id as address_id,
				a.street as address_street,
				a.number as address_number,
				a.zipcode as address_zipcode
			FROM
				kitchens k
			JOIN users_kitchens uk ON
				k.id = uk.kitchen_id
			JOIN address a on
				a.id = k.address_id 
			WHERE
				uk.user_id = ?
				and uk.kitchen_id = ?`

	row := repo.db.QueryRow(query, userID, kitchenID)

	if row == nil {
		return nil
	}

	kitchen := new(models.KitchenModel)
	kitchen.Address = new(models.AddressModel)

	err := row.Scan(&kitchen.Id, &kitchen.Address.Id, &kitchen.Address.Street, &kitchen.Address.Number, &kitchen.Address.ZipCode)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return kitchen
}
