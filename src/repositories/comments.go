package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Comments struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) *Comments {
	return &Comments{db}
}

func (repository Comments) Create(comment models.Comment) (uint64, error) {
	statement, err := repository.db.Prepare("insert into comments (review_id, user_id, comment) values (?, ?, ?)")
	fmt.Println(comment)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(comment.Review.ID, comment.User.ID, comment.Comment)

	if err != nil {
		return 0, err
	}

	lastCreatedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastCreatedID), nil
}

func (repository Comments) SearchComments(ID uint64) ([]models.Comment, error) {
	fmt.Println("review id: ", ID)
	rows, err := repository.db.Query(
		"SELECT comments.id, comments.comment, comments.CreatedAt, reviews.id, users.nick, users.id FROM comments JOIN users ON comments.user_id = users.id JOIN reviews ON comments.review_id = reviews.id where reviews.id =  ? ;",
		ID,
	) // <--- This parenthesis was missing.
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment

	for rows.Next() { // Corrected `rows.next()` to `rows.Next()`
		var comment models.Comment
		if err = rows.Scan(
			&comment.ID,
			&comment.Comment,
			&comment.CreatedAt,
			&comment.Review.ID,
			&comment.User.Nick,
			&comment.User.ID,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository Comments) Delete(ID uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM comments WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}
	return nil
}