package repository

import (
	"context"

	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	ReviewRepository interface {
		CalculateRatingByFilmID(ctx context.Context, tx *gorm.DB, film_id uint) (float64, error)
		GetAllReviewWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, film_id uint) (dto.GetAllReviewWithPaginationResponse, error)
		GetReview(ctx context.Context, tx *gorm.DB, userID uuid.UUID, filmID uint) (entity.Review, error)
		Create(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error)
		Update(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error)
		Delete(ctx context.Context, tx *gorm.DB, id uint) error
	}

	reviewRepository struct {
		db *gorm.DB
	}
)

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{
		db: db,
	}
}

func (r *reviewRepository) CalculateRatingByFilmID(ctx context.Context, tx *gorm.DB, film_id uint) (float64, error) {
	if tx == nil {
		tx = r.db
	}

	var result struct {
		AverageRating float64
	}

	err := tx.WithContext(ctx).Table("reviews").Select("AVG(rating) as average_rating").Where("film_id = ?", film_id).Scan(&result).Error
	if err != nil {
		return 0, err
	}

	return result.AverageRating, nil
}

func (r *reviewRepository) GetAllReviewWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, film_id uint) (dto.GetAllReviewWithPaginationResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var review []entity.Review
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Review{})

	query = query.Where("film_id = ?", film_id)

	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllReviewWithPaginationResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Find(&review).Error; err != nil {
		return dto.GetAllReviewWithPaginationResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllReviewWithPaginationResponse{
		Reviews: review,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *reviewRepository) GetReview(ctx context.Context, tx *gorm.DB, userID uuid.UUID, filmID uint) (entity.Review, error) {
	if tx == nil {
		tx = r.db
	}

	var review entity.Review
	err := tx.WithContext(ctx).Where("user_id = ? AND film_id = ?", userID, filmID).First(&review).Error
	if err != nil {
		return entity.Review{}, err
	}

	return review, nil
}

func (r *reviewRepository) Create(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error) {
	if tx == nil {
		tx = r.db
	}

	var film entity.Film
	if err := tx.WithContext(ctx).First(&film, review.FilmID).Error; err != nil {
		return entity.Review{}, err
	}

	if film.Status == entity.FilmStatusNotYetAired {
		return entity.Review{}, dto.ErrInvalidStatus
	}

	var userListFilm entity.UserFilmList
	if err := tx.WithContext(ctx).Where("user_id = ? AND film_id = ?", review.UserID, review.FilmID).First(&userListFilm).Error; err != nil {
		return entity.Review{}, err
	}

	if userListFilm.Status == entity.ListStatusPlanToWatch {
		return entity.Review{}, dto.ErrInvalidStatus
	}

	if err := tx.WithContext(ctx).Create(&review).Error; err != nil {
		return entity.Review{}, err
	}

	return review, nil
}

func (r *reviewRepository) Update(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&review).Error; err != nil {
		return entity.Review{}, err
	}

	return review, nil
}

func (r *reviewRepository) Delete(ctx context.Context, tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entity.Review{}).Error; err != nil {
		return err
	}

	return nil
}
