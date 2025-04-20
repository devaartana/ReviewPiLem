package controller

import (
	"net/http"
	"strconv"

	"github.com/devaartana/ReviewPiLem/dto"
	"github.com/devaartana/ReviewPiLem/services"
	"github.com/devaartana/ReviewPiLem/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	UserFilmController interface {
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Delete(ctx *gin.Context)
		GetUserList(ctx *gin.Context)
	}

	userFilmController struct {
		userFilmService service.UserFilmServices
	}
)

func NewUserFilmController(s service.UserFilmServices) UserFilmController {
	return &userFilmController{
		userFilmService: s,
	}
}

func (c *userFilmController) Create(ctx *gin.Context) {
	user_id := ctx.MustGet("user_id").(string)

	userID, err := uuid.Parse(user_id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	
	var userFilm dto.UserFilmListRequest
	if err := ctx.ShouldBind(&userFilm); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userFilm.UserID = userID

	result, err := c.userFilmService.Create(ctx.Request.Context(), userFilm)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_ADD_FILM_TO_LIST, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_ADD_FILM_TO_LIST, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userFilmController) Update(ctx *gin.Context) {
	user_id := ctx.MustGet("user_id").(string)

	userID, err := uuid.Parse(user_id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var userFilm dto.UserFilmListRequest
	if err := ctx.ShouldBind(&userFilm); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userFilm.UserID = userID

	result, err := c.userFilmService.Update(ctx.Request.Context(), userFilm)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_FILM_TO_LIST, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_FILM_TO_LIST, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *userFilmController) Delete(ctx *gin.Context) {
	user_id := ctx.MustGet("user_id").(string)

	userID, err := uuid.Parse(user_id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	film_id := ctx.Param("id")
	filmID, err := strconv.Atoi(film_id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_FILM_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if err := c.userFilmService.Delete(ctx.Request.Context(), userID, uint(filmID)); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_FILM_TO_LIST, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_FILM_TO_LIST, nil)
	ctx.JSON(http.StatusOK, res)

}

func (c *userFilmController) GetUserList(ctx *gin.Context) {
	user_id := ctx.Param("id")
	userID, err := uuid.Parse(user_id)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	result, err := c.userFilmService.GetUserList(ctx, userID)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_ALL_FILM_TO_LIST, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_ALL_FILM_TO_LIST, result)
	ctx.JSON(http.StatusOK, res)
}
