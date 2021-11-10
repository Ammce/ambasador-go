package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id          uint   `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    []byte `json:"-"`
	IsAmbasador bool   `json:"-"`
}

func (u *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = hashedPassword
}

func (u *User) CompareHashAndPassword(password []byte) error {
	if err := bcrypt.CompareHashAndPassword(u.Password, password); err != nil {
		return err
	}
	return nil
}
