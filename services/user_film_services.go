package service

import (
	"context"

	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/devaartana/ReviewPiLem/repository"
	"github.com/google/uuid"
)

type (
	UserFilmServices interface {
		GetUserList(ctx context.Context, user_id uuid.UUID) ([]dto.UserFilmResponse, error)
		Create(ctx context.Context, req dto.UserFilmListRequest) (dto.UserFilmResponse, error)
		Update(ctx context.Context, req dto.UserFilmListRequest) (dto.UserFilmResponse, error)
		Delete(ctx context.Context, userID uuid.UUID, filmID uint) error
	}

	userFilmServices struct {
		userFilmRepo repository.UserFilmRepository
		jwtService   JWTService
	}
)

func NewUserFilmServices(userFilmRepo repository.UserFilmRepository, jwtService JWTService) UserFilmServices {
	return &userFilmServices{
		userFilmRepo,
		jwtService,
	}
}

func (s *userFilmServices) GetUserList(ctx context.Context, user_id uuid.UUID) ([]dto.UserFilmResponse, error) {
	userFilms, err := s.userFilmRepo.GetUserList(ctx, nil, user_id)
	if err != nil {
		return []dto.UserFilmResponse{}, err
	}

	var res []dto.UserFilmResponse
	for _, userFilm := range userFilms {
		if userFilm.Visibility {
			res = append(res, dto.UserFilmResponse{
				UserID: userFilm.UserID,
				FilmID: userFilm.FilmID,
				Status: dto.ListStatus(userFilm.Status),
			})
		}
	}

	return res, nil
}

func (s *userFilmServices) Create(ctx context.Context, req dto.UserFilmListRequest) (dto.UserFilmResponse, error) {
	userFilm := entity.UserFilmList{
		UserID:     req.UserID,
		FilmID:     req.FilmID,
		Status:     entity.ListStatus(req.Status),
		Visibility: req.Visibility,
	}

	createdUserFilm, err := s.userFilmRepo.Create(ctx, nil, userFilm)
	if err != nil {
		return dto.UserFilmResponse{}, err
	}

	return dto.UserFilmResponse{
		UserID: createdUserFilm.UserID,
		FilmID: createdUserFilm.FilmID,
		Status: dto.ListStatus(createdUserFilm.Status),
	}, nil
}

func (s *userFilmServices) Update(ctx context.Context, req dto.UserFilmListRequest) (dto.UserFilmResponse, error) {
	userFilm, err := s.userFilmRepo.GetUserFilm(ctx, nil, req.UserID, req.FilmID)
	if err != nil {
		return dto.UserFilmResponse{}, err
	}

	userFilm.Status = entity.ListStatus(req.Status)
	userFilm.Visibility = req.Visibility

	updated, err := s.userFilmRepo.Update(ctx, nil, userFilm)
	if err != nil {
		return dto.UserFilmResponse{}, err
	}

	return dto.UserFilmResponse{
		UserID: updated.UserID,
		FilmID: updated.FilmID,
		Status: dto.ListStatus(updated.Status),
	}, nil
}

func (s *userFilmServices) Delete(ctx context.Context, userID uuid.UUID, filmID uint) error {
	_, err := s.userFilmRepo.GetUserFilm(ctx, nil, userID, filmID)
	if err != nil {
		return err
	}

	if err := s.userFilmRepo.Delete(ctx, nil, userID, filmID); err != nil {
		return err
	}

	return nil
}
