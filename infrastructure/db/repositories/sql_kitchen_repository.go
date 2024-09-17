package repositories

import (
	"database/sql"
	"fmt"
	"ramengo/domain/enums"
	"ramengo/domain/models"
	"time"
)

type SQLKitchenRepository struct {
	db *sql.DB
}

func NewSQLKitchenRepository(db *sql.DB) *SQLKitchenRepository {
	return &SQLKitchenRepository{db}
}

func (repo *SQLKitchenRepository) FindOrdersByDate(kitchenID int, date time.Time) []*models.OrderModel {
	slice := make([]*models.OrderModel, 0)

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
				o.kitchen_id = ? and o.date = ?`

	rows, err := repo.db.Query(query, kitchenID, date)

	if err != nil {
		fmt.Println(err)
	} else {

		var (
			order   *models.OrderModel
			user    *models.UserModel
			kitchen *models.KitchenModel
			stringDate string
		)

		for rows.Next() {
			kitchen = &models.KitchenModel{}

			user = &models.UserModel{}
			order = &models.OrderModel{
				Kitchen: kitchen,
				User:    user,
			}

			err := rows.Scan(&order.Id, &kitchen.Id, &stringDate, &user.Id, &user.Name, &order.Status)
			if err != nil {
				fmt.Println(err)
				continue
			}

			order.Date, err = time.Parse(time.DateOnly, stringDate)

			if err != nil {
				fmt.Println(err)
			}

			slice = append(slice, order)
		}
	}

	rows.Next()

	return slice
}

func (repo *SQLKitchenRepository) FindOrdersByDiffentStatus(kitchen *models.KitchenModel, status enums.OrderStatus) []*models.OrderModel {

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
				o.kitchen_id = ? and o.status_id != ?`

	rows, err := repo.db.Query(query, kitchen.Id, status)

	if err != nil {
		fmt.Println(err)
	} else {

		var (
			order   *models.OrderModel
			user    *models.UserModel
			kitchen *models.KitchenModel
			stringDate string
		)

		for rows.Next() {
			kitchen = &models.KitchenModel{}

			user = &models.UserModel{}
			order = &models.OrderModel{
				Kitchen: kitchen,
				User:    user,
			}

			err := rows.Scan(&order.Id, &kitchen.Id, &stringDate, &user.Id, &user.Name, &order.Status)
			if err != nil {
				fmt.Println(err)
				continue
			}

			order.Date, err = time.Parse(time.DateOnly, stringDate)

			if err != nil {
				fmt.Println(err)
			}

			kitchen.Orders = append(kitchen.Orders, order)
		}
	}

	rows.Next()

	return kitchen.Orders
}

func (repo *SQLKitchenRepository) AddUser(userID, kitchenID int) bool {
	query := "INSERT INTO users_kitchens(user_id, kitchen_id) VALUES (?, ?)"

	_, err := repo.db.Exec(query, userID, kitchenID)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (repo *SQLKitchenRepository) UserNotIncludedForKitchen(userID, kitchenID int) bool {
	query := "SELECT COUNT(*) FROM users_kitchens WHERE user_id = ? AND kitchen_id = ?"

	row := repo.db.QueryRow(query, userID, kitchenID)

	var total int

	if err := row.Scan(&total); err != nil {
		fmt.Println(err)
		return false
	}

	return total == 0
}
