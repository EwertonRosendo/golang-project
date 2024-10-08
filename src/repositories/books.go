package repositories

import (
	"api/src/models"
	"database/sql"
)

type Books struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *Books {
	return &Books{db}
}

func (repository Books) Create(book models.Book) (uint64, error) {
	statement, err := repository.db.Prepare("insert into books (title, author, subtitle, description, published_at, cover, publisher) values (?,?, ?, ?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(book.Title, book.Authors, book.Subtitle, book.Description, book.Published_at, (book.Thumbnail+".jpg"), book.Publisher)

	if err != nil {
		return 0, err
	}

	lastCreatedID, err := result.LastInsertId()
	if err != nil {
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

	row, err := repository.db.Query(
		"select id, title, subtitle, description, author, publisher, published_at, cover from books where id = ?",
		ID,
	)
	if err != nil {
		return models.Book{}, err
	}
	defer row.Close()

	var book models.Book

	for row.Next() {
		if err = row.Scan(
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

func (repository Books) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("delete from books where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repository Books) Update(ID uint64, book models.Book) (error, string) {
	statement, err := repository.db.Prepare("update books set title = ?, subtitle = ?, description = ?, published_at = ?, publisher = ?, author = ? where id = ?")
	if err != nil {
		return err, ""
	}
	defer statement.Close()

	if _, err = statement.Exec(book.Title, book.Subtitle, book.Description, book.Published_at, book.Publisher, book.Authors, ID); err != nil {
		return err, ""
	}
	book, err = repository.FindBookById(ID)
	if err != nil {
		return err, ""
	}
	return nil, book.Thumbnail
}

func (repository Books) UpdateThumbnail(ID uint64, book models.Book) (error) {
	statement, err := repository.db.Prepare("update books set cover = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(book.Title+".jpg", ID); err != nil {
		return err
	}
	book, err = repository.FindBookById(ID)
	if err != nil {
		return err
	}
	return nil
}
