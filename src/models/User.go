package models

import (
	"api/src/validations"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.check(step); err != nil {
		return err
	}
	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) check(step string) error {
	if user.Name == "" {
		return errors.New("name can not be black")
	}
	if user.Nick == "" {
		return errors.New("nick can not be black")
	}
	if user.Email == "" {
		return errors.New("email can not be black")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return err
	}

	if step == "signup" && user.Password == "" {
		return errors.New("password can not be black")
	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

	if step == "signup" {
		hashPassword, err := validations.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashPassword)
	}
	return nil
}
