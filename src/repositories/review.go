package repositories

import (
	"api/src/models"
	"database/sql"
)

type Reviews struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *Reviews {
	return &Reviews{db}
}

func (repository Reviews) Create(review models.Review) (uint64, error) {
	statement, err := repository.db.Prepare("insert into reviews (book_id, user_id, status, review, rating) values (?, ?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(review.Book.ID, review.User.ID, review.Status, review.Review, review.Rating)

	if err != nil {
		return 0, err
	}

	lastCreatedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastCreatedID), nil
}

func (repository Reviews) SearchReviews() ([]models.Review, error) {

	rows, err := repository.db.Query(
		"SELECT users.id, users.nick, users.name,books.id, books.title, books.cover, books.author,reviews.id, reviews.status, reviews.rating, reviews.review FROM reviews JOIN users ON reviews.user_id = users.id JOIN books ON reviews.book_id = books.id;",
	) // <--- This parenthesis was missing.
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review

	for rows.Next() { // Corrected `rows.next()` to `rows.Next()`
		var review models.Review
		if err = rows.Scan(
			&review.User.ID,
			&review.User.Nick,
			&review.User.Name,
			&review.Book.ID,
			&review.Book.Title,
			&review.Book.Thumbnail,
			&review.Book.Authors,
			&review.ID,
			&review.Status,
			&review.Rating,
			&review.Review,
		); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (repository Reviews) FindReviewsByUser(ID uint64) ([]models.Review, error) {

	rows, err := repository.db.Query(
		"SELECT users.id, users.nick, users.name,books.id, books.title, books.cover, books.author,reviews.id, reviews.status, reviews.rating, reviews.review FROM reviews JOIN users ON reviews.user_id = users.id JOIN books ON reviews.book_id = books.id where users.id = ?;",
		 ID,
	) // <--- This parenthesis was missing.
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review

	for rows.Next() { // Corrected `rows.next()` to `rows.Next()`
		var review models.Review
		if err = rows.Scan(
			&review.User.ID,
			&review.User.Nick,
			&review.User.Name,
			&review.Book.ID,
			&review.Book.Title,
			&review.Book.Thumbnail,
			&review.Book.Authors,
			&review.ID,
			&review.Status,
			&review.Rating,
			&review.Review,
		); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

func (repository Reviews) FindReviewById(ID uint64) (models.Review, error) {

	rows, err := repository.db.Query(
		"SELECT users.id, users.nick, users.name,books.id, books.title, books.cover, books.author,reviews.id, reviews.status, reviews.rating, reviews.review FROM reviews JOIN users ON reviews.user_id = users.id JOIN books ON reviews.book_id = books.id where reviews.id = ?;",
		ID,
	)
	if err != nil {
		return models.Review{}, err
	}
	defer rows.Close()

	var review models.Review

	for rows.Next() {
		if err = rows.Scan(
			&review.User.ID,
			&review.User.Nick,
			&review.User.Name,
			&review.Book.ID,
			&review.Book.Title,
			&review.Book.Thumbnail,
			&review.Book.Authors,
			&review.ID,
			&review.Status,
			&review.Rating,
			&review.Review,
		); err != nil {
			return models.Review{}, err
		}
	}
	return review, nil
}

func (repository Reviews) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM reviews WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repository Reviews) Update(ID uint64, review models.Review) error {
	statement, err := repository.db.Prepare("update reviews set review = ?, status = ?, rating = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(review.Review, review.Status, review.Rating, ID); err != nil {
		return err
	}
	return nil
}
