package repository

import (
	"context"

	"github.com/devaartana/ReviewPiLem/entity"
	"gorm.io/gorm"
)

type (
	GenreRepository interface {
		GetAllGenre(ctx context.Context, tx *gorm.DB) ([]entity.Genre, error)
		Update(ctx context.Context, tx *gorm.DB, genre entity.Genre) (entity.Genre, error)
		Create(ctx context.Context, tx *gorm.DB, name string) (entity.Genre, error)
		FindGenresByFilmID(ctx context.Context, tx *gorm.DB, filmID uint) ([]entity.Genre, error)
	}

	genreRepository struct {
		db *gorm.DB
	}
)

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &genreRepository{
		db: db,
	}
}

func (r *genreRepository) GetAllGenre(ctx context.Context, tx *gorm.DB) ([]entity.Genre, error){
	if tx == nil {
		tx = r.db
	}
	
	var genres []entity.Genre
	if err := tx.WithContext(ctx).Find(&genres).Error; err != nil{
		return []entity.Genre{}, err 
	}

	return genres, nil
}

func (r *genreRepository) Update(ctx context.Context, tx *gorm.DB, genre entity.Genre) (entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Model(&genre).Updates(genre).Error; err != nil {
		return entity.Genre{}, err
	}

	if err := tx.WithContext(ctx).First(&genre, genre.ID).Error; err != nil {
		return entity.Genre{}, err
	}
	
	return genre, nil
}

func (r *genreRepository) Create(ctx context.Context, tx *gorm.DB, name string) (entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	genre := entity.Genre{Name: name}
	if err := tx.WithContext(ctx).Create(&genre).Error; err != nil {
		return entity.Genre{}, err
	}

	return genre, nil
}

func (r *genreRepository) FindGenresByFilmID(ctx context.Context, tx *gorm.DB, filmID uint) ([]entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	var genres []entity.Genre
	err := tx.WithContext(ctx).
		Joins("JOIN film_genres fg ON fg.genre_id = genres.id").
		Where("fg.film_id = ?", filmID).
		Find(&genres).Error
	return genres, err
}