package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users{
	return &Users{db}
}

func (repository Users) Create (user models.User) (uint64, error){
	statement, err := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")

	if err !=  nil{
		return 0, err 
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err !=  nil{
		return 0, err 
	}

	lastCreatedID, err := result.LastInsertId()
	if err !=  nil{
		return 0, err 
	}

	return uint64(lastCreatedID), nil
}

func (repository Users) SearchUsers(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?", nameOrNick, nameOrNick,
	) // <--- This parenthesis was missing.
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() { // Corrected `rows.next()` to `rows.Next()`
		var user models.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository Users) FindUserById(ID uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"select id, name, nick, email, CreatedAt from users where id = ?",
		ID, 
	)
	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?",)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil{
		return err
	}
	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repository Users) FindByEmail(email string) (models.User, error){
	 row, err := repository.db.Query("select id, password, email, nick, name from users where email = ?", email)

	 if err != nil {
		return models.User{}, err
	 }
	 defer row.Close()

	 var user models.User

	 if row.Next() {
		if err = row.Scan(&user.ID, &user.Password, &user.Email, &user.Nick, &user.Name); err != nil{
			return models.User{}, err
		}
	 }
	 return user, nil

}