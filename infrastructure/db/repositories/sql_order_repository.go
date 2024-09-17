package repositories

import (
	"database/sql"
	"fmt"
	"ramengo/domain/models"
)

type SQLOrderRepository struct {
	db *sql.DB
}

func NewSQLOrderRepository(db *sql.DB) *SQLOrderRepository {
	return &SQLOrderRepository{db}
}

func (repo *SQLOrderRepository) Save(order *models.OrderModel) bool {
	query := `INSERT INTO orders (kitchen_id, status_id, user_id) VALUES(?, ?, ?)`
	res, err := repo.db.Exec(query, order.Kitchen.Id, order.Status, order.User.Id)

	if err != nil {
		fmt.Println(err)
		return false
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
		return false
	}

	order.Id = int(id)

	return true
}

func (repo *SQLOrderRepository) ChangeStatus(order *models.OrderModel) bool {
	query := "UPDATE orders SET status_id = ? WHERE id = ?"

	res, err := repo.db.Exec(query, order.Status, order.Id)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if rows, err := res.RowsAffected(); err != nil {
		fmt.Println(err)
		return false
	} else {
		return int(rows) == 1
	}
}

func (repo *SQLOrderRepository) FindById(id int) *models.OrderModel {
	query := "SELECT id, kitchen_id, status_id, user_id FROM orders WHERE id = ?"
	row := repo.db.QueryRow(query, id)

	order := &models.OrderModel{
		Kitchen: new(models.KitchenModel),
		User:    new(models.UserModel),
	}

	err := row.Scan(&order.Id, &order.Kitchen.Id, &order.Status, &order.User.Id)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return order
}