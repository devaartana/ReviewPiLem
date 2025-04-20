package service

import (
	"context"
	"math"

	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/repository"
)

type (
	FilmServices interface {
		GetDetailFilmById(ctx context.Context, id uint) (dto.DetailFilmResponse, error)
		GetAllFilmWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.FilmPaginationRespons, error)
		GetImagePath(ctx context.Context, id uint) (string, error)
	}

	filmServices struct {
		filmRepo      repository.FilmRepository
		genreRepo     repository.GenreRepository
		filmImageRepo repository.FilmImagesRepository
		reviewRepo    repository.ReviewRepository
	}
)

func NewFilmServices(filmRepo repository.FilmRepository, genreRepo repository.GenreRepository, filmImageRepo repository.FilmImagesRepository, reviewRepo repository.ReviewRepository) FilmServices {
	return &filmServices{
		filmRepo:      filmRepo,
		genreRepo:     genreRepo,
		filmImageRepo: filmImageRepo,
		reviewRepo:    reviewRepo,
	}
}

func (s *filmServices) GetDetailFilmById(ctx context.Context, id uint) (dto.DetailFilmResponse, error) {
	film, err := s.filmRepo.FindById(ctx, nil, id)
	if err != nil {
		return dto.DetailFilmResponse{}, dto.ErrGetIdParam
	}

	genres, err := s.genreRepo.FindGenresByFilmID(ctx, nil, id)
	if err != nil {
		return dto.DetailFilmResponse{}, err
	}

	filmImages, err := s.filmImageRepo.FindImagesByFilmId(ctx, nil, id)
	if err != nil {
		return dto.DetailFilmResponse{}, err
	}

	genreResponses := []dto.GenreResponse{}
	for _, g := range genres {
		genreResponses = append(genreResponses, dto.GenreResponse{
			ID:         g.ID,
			Name:       g.Name,
			Created_at: g.CreatedAt,
			Updated_at: g.UpdatedAt,
		})
	}

	filmImagesResponse := []dto.FilmImageResponse{}
	for _, fg := range filmImages {
		filmImagesResponse = append(filmImagesResponse, dto.FilmImageResponse{
			ID:     fg.ID,
			Path:   fg.Path,
			Status: fg.Status,
		})
	}

	return dto.DetailFilmResponse{
		ID:            film.ID,
		Title:         film.Title,
		Synopsis:      film.Synopsis,
		Status:        dto.FilmStatus(film.Status),
		TotalEpisodes: film.TotalEpisodes,
		ReleaseDate:   film.ReleaseDate,
		Genres:        genreResponses,
		CreatedAt:     film.CreatedAt,
		Images:        filmImagesResponse,
	}, nil
}

func (s *filmServices) GetAllFilmWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.FilmPaginationRespons, error) {
	dataWithPaginate, err := s.filmRepo.GetAllFilmWithPagination(ctx, nil, req)
	if err != nil {
		return dto.FilmPaginationRespons{}, dto.ErrGetAllFilm
	}

	var datas []dto.FilmResponse
	for _, film := range dataWithPaginate.Films {
		genres, err := s.genreRepo.FindGenresByFilmID(ctx, nil, film.ID)
		if err != nil {
			return dto.FilmPaginationRespons{}, err
		}

		genreResponses := []dto.GenreResponse{}
		for _, g := range genres {
			genreResponses = append(genreResponses, dto.GenreResponse{
				ID:         g.ID,
				Name:       g.Name,
				Created_at: g.CreatedAt,
				Updated_at: g.UpdatedAt,
			})
		}

		cover, err := s.filmImageRepo.CoverImagesByFilmId(ctx, nil, film.ID)
		if err != nil {
			return dto.FilmPaginationRespons{}, err
		}

		rating, err := s.reviewRepo.CalculateRatingByFilmID(ctx, nil, film.ID)
		if err != nil {
			return dto.FilmPaginationRespons{}, err
		}

		data := dto.FilmResponse{
			ID:            film.ID,
			Title:         film.Title,
			Status:        dto.FilmStatus(film.Status),
			TotalEpisodes: film.TotalEpisodes,
			ReleaseDate:   film.ReleaseDate,
			Genres:        genreResponses,
			Rating:        math.Round(rating * 100) / 100,
			Images:        dto.FilmImageResponse(cover),
		}

		datas = append(datas, data)
	}

	return dto.FilmPaginationRespons{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func(s *filmServices) GetImagePath(ctx context.Context, id uint) (string, error) {
	return s.filmRepo.GetImagePath(ctx, nil, id)
}
