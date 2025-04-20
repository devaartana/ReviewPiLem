package repository

import (
	"context"

	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type (
	UserFilmRepository interface {
		Create(ctx context.Context, tx *gorm.DB, userFilm entity.UserFilmList) (entity.UserFilmList, error)
		Update(ctx context.Context, tx *gorm.DB, userFilm entity.UserFilmList) (entity.UserFilmList, error)
		Delete(ctx context.Context, tx *gorm.DB, user_id uuid.UUID, film_id uint) error
		GetUserList(ctx context.Context, tx *gorm.DB, user_id uuid.UUID) ([]entity.UserFilmList, error)
		GetUserFilm(ctx context.Context, tx *gorm.DB, user_id uuid.UUID, film_id uint) (entity.UserFilmList, error)
	}

	userFilmRepository struct {
		db *gorm.DB
	}
)

func NewUserFilmRepository(db *gorm.DB) UserFilmRepository {
	return &userFilmRepository {
		db: db,
	}
}


func (r *userFilmRepository) Create(ctx context.Context, tx *gorm.DB, userFilm entity.UserFilmList) (entity.UserFilmList, error) {
	if tx == nil {
		tx = r.db
	}
	err := tx.WithContext(ctx).Create(&userFilm).Error
	if err != nil {
		return entity.UserFilmList{}, err
	}
	return userFilm, nil
}

func (r *userFilmRepository) Update(ctx context.Context, tx *gorm.DB, userFilm entity.UserFilmList) (entity.UserFilmList, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Updates(&userFilm).Error; err != nil {
		return entity.UserFilmList{}, err
	}

	return userFilm, nil
}

func (r *userFilmRepository) Delete(ctx context.Context, tx *gorm.DB, user_id uuid.UUID, film_id uint) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("user_id = ? AND film_id = ?", user_id, film_id).Delete(&entity.UserFilmList{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *userFilmRepository) GetUserList(ctx context.Context, tx *gorm.DB, user_id uuid.UUID) ([]entity.UserFilmList, error) {
	if tx == nil {
		tx = r.db
	}
	var userFilms []entity.UserFilmList
	
	if err := tx.WithContext(ctx).Where("user_id = ?", user_id).Find(&userFilms).Error; err != nil {
		return nil, err
	}

	return userFilms, nil
}

func (r *userFilmRepository) GetUserFilm(ctx context.Context, tx *gorm.DB, user_id uuid.UUID, film_id uint) (entity.UserFilmList, error) {
	if tx == nil {
		tx = r.db
	}
	var userFilm entity.UserFilmList
	
	if err := tx.WithContext(ctx).Where("user_id = ? AND film_id = ?", user_id, film_id).First(&userFilm).Error; err != nil {
		return entity.UserFilmList{}, err
	}

	return userFilm, nil
}