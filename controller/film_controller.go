package controller

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/devaartana/ReviewPiLem/dto"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/devaartana/ReviewPiLem/utils"
	"github.com/gin-gonic/gin"
)

type (
	FilmController interface {
		GetFilmDetail(ctx *gin.Context)
		GetFilm(ctx *gin.Context)
		GetImage(ctx *gin.Context)
	}

	filmController struct {
		filmServices service.FilmServices
	}
)

func NewFilmController(f service.FilmServices) FilmController {
	return &filmController{
		filmServices: f,
	}
}

func (c *filmController) GetFilmDetail(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_GENRE_FROM_PARAM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.filmServices.GetDetailFilmById(ctx, uint(id))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_FILM_DETAIL, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_FILM_DETAIL, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *filmController) GetFilm(ctx *gin.Context) {
	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.filmServices.GetAllFilmWithPagination(ctx.Request.Context(), req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_FILM, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	resp := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *filmController) GetImage(ctx *gin.Context) {
	imageId := ctx.Param("id")

	id, err := strconv.Atoi(imageId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_IMAGE_ID, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	path, err := c.filmServices.GetImagePath(ctx, uint(id))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_IMAGE_PATH, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	fullImagePath := "./" + path

	file, err := os.Open(fullImagePath)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_OPEN_IMAGE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}
	defer file.Close()

	ctx.Header("Content-Type", "image/jpeg")

	_, err = io.Copy(ctx.Writer, file)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_SEND_IMAGE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
}
