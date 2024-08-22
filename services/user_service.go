package services

import (
	"credit-card-notification/models"
	"credit-card-notification/repository"
	"credit-card-notification/utils"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(name, email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	cardNumber := generateCardNumber()

	user := models.User{
		Name:       name,
		Email:      email,
		Password:   hashedPassword,
		CardNumber: cardNumber,
	}
	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}
	err = utils.SendEmail(user.Email, "Your Credit Card Details", "Your Card number is:"+cardNumber)
	if err != nil {
		log.Println("Failed to send email:", err)
	}
	return nil

}

func generateCardNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%012d", rand.Intn(1000000000000))
}
