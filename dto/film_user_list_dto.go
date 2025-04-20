package dto

import "github.com/google/uuid"

type ListStatus string

const (
	ListStatusPlanToWatch ListStatus = "plan_to_watch"
	ListStatusWatching    ListStatus = "watching"
	ListStatusCompleted   ListStatus = "completed"
	ListStatusOnHold      ListStatus = "on_hold"
	ListStatusDropped     ListStatus = "dropped"
)

const (
	MESSAGE_FAILED_ADD_FILM_TO_LIST     = "failed to add film to list"
	MESSAGE_FAILED_UPDATE_FILM_TO_LIST  = "failed to update film in list"
	MESSAGE_FAILED_DELETE_FILM_TO_LIST  = "failed to delete film in list"
	MESSAGE_FAILED_GET_ALL_FILM_TO_LIST = "failed to get all film in list"
	MESSAGE_FAILED_GET_FILM_ID          = "failed to get film id"

	MESSAGE_SUCCESS_ADD_FILM_TO_LIST     = "success to add film to list"
	MESSAGE_SUCCESS_UPDATE_FILM_TO_LIST  = "success to update film in list"
	MESSAGE_SUCCESS_DELETE_FILM_TO_LIST  = "success to delete film in list"
	MESSAGE_SUCCESS_GET_ALL_FILM_TO_LIST = "success to get all film in list"
)

type (
	UserFilmListRequest struct {
		UserID     uuid.UUID  `json:"user_id"`
		FilmID     uint       `json:"film_id"`
		Status     ListStatus `json:"status"`
		Visibility bool       `json:"visibility"`
	}

	UserFilmResponse struct {
		UserID uuid.UUID  `json:"user_id"`
		FilmID uint       `json:"film_id"`
		Status ListStatus `json:"status"`
	}
)
