package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("ID for new user must be 0")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserByID(ID int) (User, error) {
	for _, u := range users {
		if u.ID == ID {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("no user found with ID %v", ID)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("no user found with ID %v", u.ID)
}

func RemoveUserByID(ID int) error {
	for i, u := range users {
		if u.ID == ID {
			users = append(users[:i], users[1+1:]...)
			return nil
		}
	}
	return fmt.Errorf("no user found with ID %v", ID)
}
