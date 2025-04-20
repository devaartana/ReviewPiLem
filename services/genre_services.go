package service

import (
	"context"

	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/devaartana/ReviewPiLem/repository"
)

type (
	GenreService interface {
		GetAllGenre(ctx context.Context) ([]dto.GenreResponse, error)
		Update(ctx context.Context, req dto.GenreRequest) (dto.GenreResponse, error) 
		Create(ctx context.Context, req dto.GenreCreateRequest) (dto.GenreResponse, error)
	}

	genreService struct {
		genreRepo  repository.GenreRepository
		jwtService JWTService
	}
)

func NewGenreService(genreRepo repository.GenreRepository, jwtService JWTService) GenreService {
	return &genreService {
		genreRepo: genreRepo,
		jwtService: jwtService,
	}
}


func (s *genreService) GetAllGenre(ctx context.Context) ([]dto.GenreResponse, error){

	genres, err := s.genreRepo.GetAllGenre(ctx, nil)
	if err != nil {
		return []dto.GenreResponse{}, dto.ErrGetGenres
	}

	var result []dto.GenreResponse
	for _, genre := range genres {
		result = append(result, dto.GenreResponse{
			ID:         genre.ID,
			Name:       genre.Name,
			Created_at: genre.CreatedAt,
			Updated_at: genre.UpdatedAt,
		})
	}
	
	return result, nil
}

func (s *genreService) Create(ctx context.Context, req dto.GenreCreateRequest) (dto.GenreResponse, error) {
	genre, err := s.genreRepo.Create(ctx, nil, req.Name)
	if err != nil {
		return dto.GenreResponse{}, dto.ErrCreateGenre
	}

	return dto.GenreResponse{
		ID: genre.ID,
		Name: genre.Name,
		Created_at: genre.CreatedAt,
		Updated_at: genre.UpdatedAt,
	}, nil 
}

func (s *genreService) Update(ctx context.Context, req dto.GenreRequest) (dto.GenreResponse, error) {
	
	data := entity.Genre {
		ID: req.ID,
		Name: req.Name,
	}

	genre, err := s.genreRepo.Update(ctx, nil, data)
	if err != nil {
		return dto.GenreResponse{}, dto.ErrUpdateGenre
	}

	return dto.GenreResponse{
		ID: genre.ID,
		Name: genre.Name,
		Created_at: genre.CreatedAt,
		Updated_at: genre.UpdatedAt,
	}, nil
}