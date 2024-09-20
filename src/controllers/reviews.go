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

func AddReview(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)

	fmt.Println("params: ", string(bodyRequest))

	if err != nil {
		responses.ERR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var review models.Review

	if err = json.Unmarshal(bodyRequest, &review); err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewReviewRepository(db)
	review.ID, err = repository.Create(review)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, review)

}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewID, err := strconv.ParseUint(params["review_id"], 10, 64)

	if err != nil {
		responses.ERR(w, http.StatusBadRequest, err)
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewReviewRepository(db)
	if err = repository.Delete(reviewID); err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func SearchReviews(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repositories.NewReviewRepository(db)
	reviews, err := repository.SearchReviews()
	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, reviews)
}
func FindReviewsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["user_id"], 10, 64)

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

	repository := repositories.NewReviewRepository(db)
	reviews, err := repository.FindReviewsByUser(userID)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, reviews)
}

func FindReviewById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	reviewID, err := strconv.ParseUint(params["review_id"], 10, 64)

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

	repository := repositories.NewReviewRepository(db)
	review, err := repository.FindReviewById(reviewID)

	if err != nil {
		responses.ERR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, review)
}
