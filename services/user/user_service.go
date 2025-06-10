package service

import (
	"database/sql"

	models "github.com/rwsirhc/payroll-api/models/user"
	provider "github.com/rwsirhc/payroll-api/providers/user"
)

func FetchAllUsers() ([]models.User, error) {
	return provider.GetAllUsers()
}

func FetchUserByID(id string) (*models.User, error) {
	user, err := provider.GetUserByID(id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func CreateUser(user *models.User) error {
	return provider.InsertUser(user)
}

func ModifyUser(id string, user *models.User) (bool, error) {
	count, err := provider.UpdateUserByID(id, user)
	return count > 0, err
}

func RemoveUser(id string) (bool, error) {
	count, err := provider.DeleteUserByID(id)
	return count > 0, err
}
