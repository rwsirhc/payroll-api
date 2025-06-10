package provider

import (
	config "github.com/rwsirhc/payroll-api/config"
	models "github.com/rwsirhc/payroll-api/models/user"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func GetUserByID(id string) (*models.User, error) {
	var u models.User
	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func InsertUser(user *models.User) error {
	return config.DB.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		user.Name, user.Email,
	).Scan(&user.ID)
}

func UpdateUserByID(id string, user *models.User) (int64, error) {
	res, err := config.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func DeleteUserByID(id string) (int64, error) {
	res, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
