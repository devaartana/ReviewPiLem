package service

import (
	"context"

	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/devaartana/ReviewPiLem/repository"
	"github.com/devaartana/ReviewPiLem/utils"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error)
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		GetUserById(ctx context.Context, userId string) (dto.UserResponse, error)
		GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error)
		GetUserByUsername(ctx context.Context, username string) (dto.UserResponse, error)
	}

	userService struct {
		userRepo   repository.UserRepository
		jwtService JWTService
	}
)

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *userService) Register(ctx context.Context, req dto.UserCreateRequest) (dto.UserResponse, error) {
	if err := utils.ValidateUsername(req.Username); err != nil {
		return dto.UserResponse{}, dto.ErrInvalidUsernameFormat
	}
	_, flag, _ := s.userRepo.CheckUsername(ctx, nil, req.Username)
	if flag {
		return dto.UserResponse{}, dto.ErrUserAlreadyExists
	}

	if err := utils.ValidateEmail(req.Email); err != nil {
		return dto.UserResponse{}, dto.ErrInvalidEmailFromat
	}
	_, flag, _ = s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.User{
		Username:    req.Username,
		DisplayName: req.DisplayName,
		Role:        constants.ENUM_ROLE_USER,
		Email:       req.Email,
		Password:    req.Password,
		Bio:         "",
	}

	userReg, err := s.userRepo.Register(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:          userReg.ID.String(),
		Username:    userReg.Username,
		DisplayName: userReg.DisplayName,
		Role:        userReg.Role,
		Email:       userReg.Email,
		Bio:         userReg.Bio,
	}, nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, flag, err := s.userRepo.CheckUsername(ctx, nil, req.Username)
	if err != nil || !flag {
		return dto.UserLoginResponse{}, dto.ErrUserNotFound
	}

	checkPassword, err := utils.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	token := s.jwtService.GenerateToken(user.ID.String(), user.Role)

	return dto.UserLoginResponse{
		Token: token,
		Role:  user.Role,
	}, nil
}

func (s *userService) GetUserById(ctx context.Context, userId string) (dto.UserResponse, error) {

	user, err := s.userRepo.GetUserById(ctx, nil, userId)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		DisplayName: user.DisplayName,
		Bio:         user.Bio,
		Email:       user.Email,
		Role:        user.Role,
	}, nil
}

func (s *userService) GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error) {
	dataWithPaginate, err := s.userRepo.GetAllUserWithPagination(ctx, nil, req)
	if err != nil {
		return dto.UserPaginationResponse{}, dto.ErrGetAllUser
	}

	var datas []dto.UserResponse
	for _, user := range dataWithPaginate.Users {
		data := dto.UserResponse{
			ID:          user.ID.String(),
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Email:       user.Email,
			Role:        user.Role,
			Bio:         user.Bio,
		}

		datas = append(datas, data)
	}

	return dto.UserPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (dto.UserResponse, error) {
	user, flag, err := s.userRepo.CheckUsername(ctx, nil, username)
	if !flag || err != nil || user.Role == "admin"{
		return dto.UserResponse{}, dto.ErrUserNotFound
	}

	return dto.UserResponse{
		Username: user.Username,
		DisplayName: user.DisplayName,
		Role: user.Role,
		Bio: user.Bio,
		Email: user.Email,
	}, nil
} 