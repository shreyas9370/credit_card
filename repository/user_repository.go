package repository

import (
	"credit-card-notification/models"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (name, email, password, card_number) VALUES (?, ?, ?, ?)"
	_, err := repo.DB.Exec(query, user.Name, user.Email, user.Password, user.CardNumber)
	if err != nil {
		return err
	}
	return nil
}
