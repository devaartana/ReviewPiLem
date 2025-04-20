package dto

import (
	"errors"

	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/google/uuid"
)

const (
	MESSAGE_FAILED_GET_USER_ID         = "failed get user id"
	MESSAGE_FAILED_GET_REVIEW_ID_PARAM = "failed get review id in param"
	MESSAGE_FAILED_REACTION            = "failed to react"
	MESSAGE_FAILED_DELETE              = "failed to delete"
	MESSAGE_FAILED_GET_LIST_REVIEW     = "failed to get list review"
	MESSAGE_FAILED_GET_REVIEW          = "failed to get review"
	MESSAGE_FAILED_CREATE_REVIEW       = "failed create review"
	MESSAGE_FAILED_UPDATE_REVIEW		= "failed update review"
	MESSAGE_FAILED_DELETE_REVIEW = "failed delete review"

	MESSAGE_SUCCESS_REACTION      = "success to react"
	MESSAGE_SUCCESS_DELETE        = "success to delete"
	MESSAGE_SUCCESS_CREATE_REVIEW = "success to create review"
	MESSAGE_SUCCESS_UPDATE_REVIEW = "success to update review"
	MESSAGE_SUCCESS_DELETE_REVIEW = "success to delete review"
)

var (
	ErrCreateReview     = errors.New("failed create review")
	ErrGetAllReviewFilm = errors.New("error get all review")
	ErrGetUserID        = errors.New("failed get user id from ctx")
	ErrParam            = errors.New("failed read param")
	ErrReaction         = errors.New("failed to react")
	ErrInvalidStatus    = errors.New("cannot create review for a film that has not yet aired")
)

type (
	ReviewRequest struct {
		ID      uint      `json:"id"`
		UserID  uuid.UUID `json:"user_id"`
		FilmID  uint      `json:"film_id"`
		Rating  int       `json:"rating"`
		Comment string    `json:"comment"`
	}

	ReviewResponse struct {
		ID      uint   `json:"id"`
		Rating  int    `json:"rating"`
		Comment string `json:"comment"`
	}

	GetAllReviewWithPaginationResponse struct {
		Reviews []entity.Review `json:"review"`
		PaginationResponse
	}

	ReviewPagitaionResponse struct {
		Data []ReviewResponse `json:"data"`
		PaginationResponse
	}
)
