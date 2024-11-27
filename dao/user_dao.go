package dao

import (
	"dto/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type UserDAO struct {
	DB *sqlx.DB
}

func NewUserDAO(DB *sqlx.DB) *UserDAO {
	return &UserDAO{DB: DB}
}

func (dao *UserDAO) Create(user *models.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := dao.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (dao *UserDAO) GetAll() ([]models.User, error) {
	var users []models.User
	query := "SELECT id, username, password FROM users"
	err := dao.DB.Select(&users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *UserDAO) GetByID(id int) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password FROM users WHERE id = $1"
	err := dao.DB.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) Delete(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := dao.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
