package repositories

import (
	"database/sql"
	"fmt"
	"ramengo/domain/models"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Save(user *models.UserModel) bool {
	query := "INSERT INTO users (name, email, password, address_id) VALUES(?, ?, ?, ?)"

	result, err := repo.db.Exec(query, user.Name, user.Email, user.Password, user.Address.Id)

	if err != nil {
		fmt.Println(err)
		return false
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return false
	}

	user.Id = int(id)

	return true
}

func (repo *UserRepository) OrderHistory(userId int) []*models.OrderModel {
	query := `SELECT 
				o.id AS order_id,
				o.kitchen_id,
				o.date,
				u.id AS user_id,
				u.name AS user_name,
				o.status_id
			FROM 
				orders o
			JOIN 
				users u ON o.user_id = u.id
			WHERE 
				u.id = ?`

	rows, err := repo.db.Query(query, userId)

	var result []*models.OrderModel
	if err != nil {
		fmt.Println(err)
		return result
	}

	for rows.Next() {
		var (
			stringDate string
			kitchen    = &models.KitchenModel{}
			user       = &models.UserModel{}
			order      = &models.OrderModel{
				User:    user,
				Kitchen: kitchen,
			}
		)

		err := rows.Scan(&order.Id, &kitchen.Id, &stringDate, &user.Id, &user.Name, &order.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}

		order.Date, err = time.Parse(time.DateOnly, stringDate)

		if err != nil {
			fmt.Println(err)
		}

		result = append(result, order)
	}

	return result
}

func (repo *UserRepository) VerifyEmailAvailable(email string) bool {
	query := `SELECT 
				COUNT(*)
			FROM 
				users
			WHERE 
				email = ?`

	row := repo.db.QueryRow(query, email)

	var count int

	if err := row.Scan(&count); err != nil {
		fmt.Println(err)
		return false
	}

	return count == 0
}
