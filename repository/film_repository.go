package repository

import (
	"context"

	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/entity"
	"gorm.io/gorm"
)


type (
	FilmRepository interface{
		FindById(ctx context.Context, tx *gorm.DB, id uint) (entity.Film, error)
		GetAllFilmWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllFilmRepositoryResponse, error)
		GetImagePath(ctx context.Context, tx *gorm.DB, id uint) (string, error)
	}

	filmRepository struct {
		db *gorm.DB
	}
)

func NewFilmRepository(db *gorm.DB) FilmRepository {
	return &filmRepository {
		db : db,
	}
}

func (r *filmRepository) FindById(ctx context.Context, tx *gorm.DB, id uint) (entity.Film, error) {
	if tx == nil {
		tx = r.db
	}

	var film entity.Film

	err := tx.WithContext(ctx).First(&film, id).Error
	if err != nil {
		return entity.Film{}, err
	}

	return film, err
}

func (r *filmRepository) GetAllFilmWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllFilmRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var films []entity.Film
	var err error
	var count int64

	req.Default()

	query := tx.WithContext(ctx).Model(&entity.Film{})
	if req.Search != "" {
		query = query.Where("LOWER(title) ILIKE ?", "%"+req.Search+"%")
	}
	
	if err := query.Count(&count).Error; err != nil {
		return dto.GetAllFilmRepositoryResponse{}, err
	}

	if err := query.Scopes(Paginate(req)).Find(&films).Error; err != nil {
		return dto.GetAllFilmRepositoryResponse{}, err
	}

	totalPage := TotalPage(count, int64(req.PerPage))
	return dto.GetAllFilmRepositoryResponse{
		Films: films,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *filmRepository) GetImagePath(ctx context.Context, tx *gorm.DB, id uint) (string, error){
	if tx == nil {
		tx = r.db
	}

	var path string
	err := tx.WithContext(ctx).Model(&entity.FilmImage{}).Where("film_id = ?", id).Select("path").Scan(&path).Error
	if err != nil {
		return "", err
	}

	return path, nil
}
