package repository

import (
	"context"

	"github.com/devaartana/ReviewPiLem/entity"
	"gorm.io/gorm"
)

type (
	FilmImagesRepository interface {
		FindImagesByFilmId(ctx context.Context, tx *gorm.DB, filmID uint) ([]entity.FilmImage, error)
		CoverImagesByFilmId(ctx context.Context, tx *gorm.DB, filmID uint) (entity.FilmImage, error)
	}

	filmImagesRepository struct {
		db *gorm.DB
	}
)

func NewFilmImagesRepository(db *gorm.DB) FilmImagesRepository {
	return &filmImagesRepository{
		db: db,
	}
}

func (r *filmImagesRepository) FindImagesByFilmId(ctx context.Context, tx *gorm.DB, filmID uint) ([]entity.FilmImage, error) {
	if tx == nil {
		tx = r.db
	}
	
	var image []entity.FilmImage
	err := tx.WithContext(ctx).Where("film_id = ?", filmID).Find(&image).Error

	return image, err
}

func (r *filmImagesRepository) CoverImagesByFilmId(ctx context.Context, tx *gorm.DB, filmID uint) (entity.FilmImage, error) {
	if tx == nil {
		tx = r.db
	}
	
	var image entity.FilmImage
	err := tx.WithContext(ctx).Where("film_id = ? AND status = ?", filmID, true).First(&image).Error

	return image, err
}