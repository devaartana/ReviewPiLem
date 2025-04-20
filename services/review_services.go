package service

import (
	"context"

	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/devaartana/ReviewPiLem/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	ReviewServices interface {
		Like(ctx context.Context, reviewID uint, userID uuid.UUID) error
		Dislike(ctx context.Context, reviewID uint, userID uuid.UUID) error
		Delete(ctx context.Context, reviewID uint, userID uuid.UUID) error
		GetAllReviewFilm(ctx context.Context, req dto.PaginationRequest, filmID uint) (dto.ReviewPagitaionResponse, error)
		Create(ctx context.Context, req dto.ReviewRequest) (dto.ReviewResponse, error)
		Update(ctx context.Context, req dto.ReviewRequest) (dto.ReviewResponse, error)
		DeleteReview(ctx context.Context, userID uuid.UUID, filmID uint) error
	}

	reviewServices struct {
		reviewRepo   repository.ReviewRepository
		reactionRepo repository.ReactionReposiotry
		jwtService   JWTService
	}
)

func NewReviewServices(reviewRepo repository.ReviewRepository, reactionRepo repository.ReactionReposiotry, jwtService JWTService) ReviewServices {
	return &reviewServices{
		reviewRepo,
		reactionRepo,
		jwtService,
	}
}

func (s *reviewServices) GetAllReviewFilm(ctx context.Context, req dto.PaginationRequest, filmID uint) (dto.ReviewPagitaionResponse, error) {
	dataWithPaginate, err := s.reviewRepo.GetAllReviewWithPagination(ctx, nil, req, filmID)
	if err != nil {
		return dto.ReviewPagitaionResponse{}, dto.ErrGetAllReviewFilm
	}

	var datas []dto.ReviewResponse
	for _, review := range dataWithPaginate.Reviews {
		data := dto.ReviewResponse {
			ID: review.ID,
			Rating: review.Rating,
			Comment: review.Comment,
		}

		datas = append(datas, data)
	}

	return dto.ReviewPagitaionResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *reviewServices) Create(ctx context.Context, req dto.ReviewRequest) (dto.ReviewResponse, error) {
	review := entity.Review {
		UserID: req.UserID,
		FilmID: req.FilmID,
		Rating: req.Rating,
		Comment: req.Comment,
	}

	created, err := s.reviewRepo.Create(ctx, nil, review)
	if err != nil {
		return dto.ReviewResponse{}, dto.ErrCreateReview
	}

	return dto.ReviewResponse{
		ID: created.ID,
		Rating: created.Rating,
		Comment: created.Comment,
	}, nil
}

func (s *reviewServices) Update(ctx context.Context, req dto.ReviewRequest) (dto.ReviewResponse, error) {
	review, err := s.reviewRepo.GetReview(ctx, nil, req.UserID, req.FilmID)
	if err != nil {
		return dto.ReviewResponse{}, err
	}

	review.Comment = req.Comment
	review.Rating = req.Rating

	updated, err := s.reviewRepo.Update(ctx, nil, review)
	if err != nil {
		return dto.ReviewResponse{}, err
	}

	return dto.ReviewResponse{
		ID: updated.ID,
		Rating: updated.Rating,
		Comment: updated.Comment,
	}, nil
}

func (s *reviewServices) DeleteReview(ctx context.Context, userID uuid.UUID, filmID uint) error {
	review, err := s.reviewRepo.GetReview(ctx, nil, userID, filmID)
	if err != nil {
		return err
	}

	if err := s.reviewRepo.Delete(ctx, nil, review.ID); err != nil {
		return err
	}

	return nil
}

func (s *reviewServices) Like(ctx context.Context, reviewID uint, userID uuid.UUID) error {
	reaction, _, err := s.reactionRepo.CheckReaction(ctx, nil, userID, reviewID)
	reaction.Status = true

	if err == gorm.ErrRecordNotFound {

		if err := s.reactionRepo.Create(ctx, nil, reaction); err != nil {
			return err
		}

		return nil
	}

	if err := s.reactionRepo.Update(ctx, nil, reaction); err != nil {
		return err
	}

	return nil
}

func (s *reviewServices) Dislike(ctx context.Context, reviewID uint, userID uuid.UUID) error {
	reaction, _ , err := s.reactionRepo.CheckReaction(ctx, nil, userID, reviewID)
	reaction.Status = false

	if err == gorm.ErrRecordNotFound {

		if err := s.reactionRepo.Create(ctx, nil, reaction); err != nil {
			return err
		}

		return nil
	}

	if err := s.reactionRepo.Update(ctx, nil, reaction); err != nil {
		return err
	}

	return nil
}

func (s *reviewServices) Delete(ctx context.Context, reviewID uint, userID uuid.UUID) error {
	return s.reactionRepo.Delete(ctx, nil, userID, reviewID)
}
