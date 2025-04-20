package dto

import (
	"errors"
	"time"

	"github.com/devaartana/ReviewPiLem/entity"
)

type FilmStatus string

const (
	FilmStatusNotYetAired    FilmStatus = "not_yet_aired"
	FilmStatusAiring         FilmStatus = "airing"
	FilmStatusFinishedAiring FilmStatus = "finished_airing"
)

const (
	MESSAGE_FAILED_GET_FILM_DETAIL = "failed get detail film"
	MESSAGE_FAILED_GET_LIST_FILM   = "failed get list film"

	MESSAGE_SUCCESS_GET_FILM_DETAIL = "success get detail film"
	MESSAGE_SUCCESS_GET_LIST_FILM   = "success get list film"
)

var (
	ErrGetIdParam = errors.New("failed to get id from param")
	ErrGetAllFilm = errors.New("failed to get film list")
)

type (
	FilmResponse struct {
		ID            uint              `json:"id"`
		Title         string            `json:"title"`
		Status        FilmStatus        `json:"status"`
		TotalEpisodes int               `json:"total_episodes"`
		ReleaseDate   time.Time         `json:"release_date"`
		Rating        float64            `json:"rating"`
		Genres        []GenreResponse   `json:"genres"`
		Images        FilmImageResponse `json:"images"`
	}

	FilmPaginationRespons struct {
		Data []FilmResponse
		PaginationResponse
	}

	GetAllFilmRepositoryResponse struct {
		Films []entity.Film `json:"films"`
		PaginationResponse
	}

	DetailFilmResponse struct {
		ID            uint                `json:"id"`
		Title         string              `json:"title"`
		Synopsis      string              `json:"synopsis"`
		Status        FilmStatus          `json:"status"`
		TotalEpisodes int                 `json:"total_episodes"`
		ReleaseDate   time.Time           `json:"release_date"`
		CreatedAt     time.Time           `json:"created_at"`
		Genres        []GenreResponse     `json:"genres"`
		Images        []FilmImageResponse `json:"images"`
	}
)
