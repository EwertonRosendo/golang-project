package models

import (
	"time"
	"strings"
	"errors"	
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

func (user *User) Prepare() error{
	if err := user.check(); err != nil{
		return err
	}
	user.format()
	return nil
}

func (user *User) check() error{
	if user.Name == ""{
		return errors.New("name can not be black")
	}
	if user.Nick == ""{
		return errors.New("nick can not be black")
	}
	if user.Email == ""{
		return errors.New("email can not be black")
	}
	if user.Password == ""{
		return errors.New("password can not be black")
	}
	return nil
}

func (user *User) format(){
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick) 
}