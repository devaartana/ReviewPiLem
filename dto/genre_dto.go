package dto

import (
	"errors"
	"time"
)

const (
	MESSAGE_FAILED_GET_LIST_GENRE       = "failed get list genre"
	MESSAGE_FAILED_GET_GENRE_FROM_PARAM = "failed to get genre id"
	MESSAGE_FAILED_UPDATE_GENRE         = "failed update genre"
	MESSAGE_FAILED_CREATE_GENRE         = "failed create genre"
	MESSAGE_FAILED_

	MESSAGE_SUCCESS_GET_LIST_GENRE = "sucess get list genre"
	MESSAGE_SUCCESS_GET_GENRE      = "success get genre"
	MESSAGE_SUCCESS_CREATE_GENRE   = "success create genre"
	MESSAGE_SUCCESS_
)

var (
	ErrGetGenres = errors.New("failed to get all genre")
	ErrCreateGenre = errors.New("failed to create genre")
	ErrUpdateGenre = errors.New("failed to update genre")
)

type (
	GenreRequest struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	GenreResponse struct {
		ID         uint      `json:"id"`
		Name       string    `json:"name"`
		Created_at time.Time `json:"created_at"`
		Updated_at time.Time `json:"updated_at"`
	}

	GenreCreateRequest struct {
		Name string `json:"name"`
	}
)
