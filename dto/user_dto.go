package dto

import (
	"errors"

	"github.com/devaartana/ReviewPiLem/entity"
)

const (
	// Failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY      = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER           = "failed create user"
	MESSAGE_FAILED_GET_LIST_USER           = "failed get list user"
	MESSAGE_FAILED_GET_USER_TOKEN          = "failed get user token"
	MESSAGE_FAILED_TOKEN_NOT_VALID         = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND         = "token not found"
	MESSAGE_FAILED_GET_USER                = "failed get user"
	MESSAGE_FAILED_LOGIN                   = "failed login"
	MESSAGE_FAILED_WRONG_EMAIL_OR_PASSWORD = "wrong email or password"
	MESSAGE_FAILED_UPDATE_USER             = "failed update user"
	MESSAGE_FAILED_DELETE_USER             = "failed delete user"
	MESSAGE_FAILED_PROSES_REQUEST          = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS           = "denied access"
	MESSAGE_FAILED_VERIFY_USERNAME         = "failed verify username"
	MESSAGE_FAILED_GET_USERNAME_FROM_PARAM = "failed get username from param"
	MESSAGE_FAILED_USERNAME_FROMAT         = "username must be 3-20 characters, lowercase, and only contain letters, numbers, _ or -."
	MESSAGE_FAILED_EMAIL_FORMAT            = "invalid email format"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER           = "success create user"
	MESSAGE_SUCCESS_GET_LIST_USER           = "success get list user"
	MESSAGE_SUCCESS_GET_USER                = "success get user"
	MESSAGE_SUCCESS_LOGIN                   = "success login"
	MESSAGE_SUCCESS_UPDATE_USER             = "success update user"
	MESSAGE_SUCCESS_DELETE_USER             = "success delete user"
	MESSAGE_SEND_VERIFICATION_EMAIL_SUCCESS = "success send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL            = "success verify email"
)

var (
	ErrCreateUser             = errors.New("failed to create user")
	ErrGetAllUser             = errors.New("failed to get all user")
	ErrGetUserById            = errors.New("failed to get user by id")
	ErrGetUserByEmail         = errors.New("failed to get user by email")
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrUserAlreadyExists      = errors.New("username already exist")
	ErrUpdateUser             = errors.New("failed to update user")
	ErrUserNotAdmin           = errors.New("user not admin")
	ErrUserNotFound           = errors.New("user not found")
	ErrEmailNotFound          = errors.New("email not found")
	ErrDeleteUser             = errors.New("failed to delete user")
	ErrPasswordNotMatch       = errors.New("password not match")
	ErrEmailOrPassword        = errors.New("wrong email or password")
	ErrAccountNotVerified     = errors.New("account not verified")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrInvalidUsernameFormat  = errors.New("invalid username fromat")
	ErrInvalidEmailFromat     = errors.New("invalid email format")
)

type (
	UserCreateRequest struct {
		Username    string `json:"username" form:"username"`
		DisplayName string `json:"display_name" form:"display_name" `
		Email       string `json:"email" form:"email"`
		Password    string `json:"password" form:"password"`
	}

	UserResponse struct {
		ID          string `json:"id,omitempty"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		Role        string `json:"role"`
		DisplayName string `json:"display_name"`
		Bio         string `json:"bio"`
	}

	UserPaginationResponse struct {
		Data []UserResponse `json:"data"`
		PaginationResponse
	}

	GetAllUserRepositoryResponse struct {
		Users []entity.User `json:"users"`
		PaginationResponse
	}

	UserUpdateRequest struct {
		Name     string `json:"name" form:"name"`
		Username string `json:"username" form:"username"`
	}

	UserUpdateResponse struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Role     string `json:"role"`
		Email    string `json:"email"`
	}

	UserLoginRequest struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}
)
