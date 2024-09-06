package repositories

import (
	"api/src/models"
	"database/sql"
)

type Books struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *Books{
	return &Books{db}
}

func (repository Books) Create (book models.Book) (uint64, error){
	statement, err := repository.db.Prepare("insert into books (title, author, subtitle, description, published_at, cover, publisher) values (?,?, ?, ?, ?, ?, ?)")

	if err !=  nil{
		return 0, err 
	}
	defer statement.Close()

	result, err := statement.Exec(book.Title, book.Authors, book.Subtitle, book.Description, book.Published_at, book.Thumbnail, book.Publisher)

	if err !=  nil{
		return 0, err 
	}

	lastCreatedID, err := result.LastInsertId()
	if err !=  nil{
		return 0, err 
	}

	return uint64(lastCreatedID), nil
}

func (repository Books) SearchBooks() ([]models.Book, error) {

	rows, err := repository.db.Query(
		"select id, title, subtitle, description, author, publisher, published_at, cover from books",
	) // <--- This parenthesis was missing.
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() { // Corrected `rows.next()` to `rows.Next()`
		var book models.Book
		if err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Subtitle,
			&book.Description,
			&book.Authors,
			&book.Publisher,
			&book.Published_at,
			&book.Thumbnail,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (repository Books) FindBookById(ID uint64) (models.Book, error) {

	rows, err := repository.db.Query(
		"select id, title, subtitle, description, author, publisher, published_at, cover from books where id = ?",
		ID,
	) 
	if err != nil {
		return models.Book{}, err
	}
	defer rows.Close()

	var book models.Book

	for rows.Next() { 
		if err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Subtitle,
			&book.Description,
			&book.Authors,
			&book.Publisher,
			&book.Published_at,
			&book.Thumbnail,
		); err != nil {
			return models.Book{}, err
		}
	}
	return book, nil
}