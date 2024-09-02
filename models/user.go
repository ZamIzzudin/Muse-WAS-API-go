// models/user.go
package models

import (
	"database/sql"
	"muse-dashboard-api/config"
	"muse-dashboard-api/utilities"
)

type User struct {
	ID  string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

type CredentialsAuth struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT id, username, email FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserByID retrieves a user by ID
func GetUserByID(id string) (User, error) {
	var user User
	err := config.DB.QueryRow("SELECT id, username, email FROM user WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	return user, nil
}

// CreateUser creates a new user in the database
func CreateUser(user User) error {
	_, err := config.DB.Exec("INSERT INTO user (username, email) VALUES (?, ?)", user.Username, user.Email)
	return err
}

// UpdateUser updates an existing user in the database
func UpdateUser(id string, user User) error {
	_, err := config.DB.Exec("UPDATE user SET username = ?, email = ? WHERE id = ?", user.Username, user.Email, id)
	return err
}

// DeleteUser deletes a user from the database
func DeleteUser(id string) error {
	_, err := config.DB.Exec("DELETE FROM user WHERE id = ?", id)
	return err
}

func Login(credential CredentialsAuth) (User, error){
	var credentialStored CredentialsAuth
	var userData User

	err := config.DB.QueryRow("SELECT id, username, password, email FROM user WHERE username = ?", credential.Username).Scan(&credentialStored.Username, &credentialStored.Password, &userData.ID, &userData.Username, &userData.Email)
	
	if err != nil{
		if err := utilities.ComparePassword(credentialStored.Password, credential.Password); err != nil {
			return userData, err
		}
	}
	return userData, err
}