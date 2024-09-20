package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"fmt"
	//"os"
	//"api/src/services"
	"strconv"
	"encoding/json"
	"io"
	"net/http"
	"github.com/gorilla/mux"
)

func SearchComments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	reviewID, err := strconv.ParseUint(params["review_id"], 10, 64)

	fmt.Println(reviewID)
	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCommentRepository(db)
	comments, err := repository.SearchComments(reviewID)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, comments)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	fmt.Println("params: ", string(bodyRequest))

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var comment models.Comment

	if err = json.Unmarshal(bodyRequest, &comment); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println("Comment: ", comment)

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCommentRepository(db)
	comment.ID, err = repository.Create(comment)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, comment)
}

func DeleteComment (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	commentID, err := strconv.ParseUint(params["comment_id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCommentRepository(db)
	if err = repository.Delete(commentID); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
