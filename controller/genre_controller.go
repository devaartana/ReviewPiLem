package controller

import (
	"net/http"
	"strconv"

	"github.com/devaartana/ReviewPiLem/dto"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/devaartana/ReviewPiLem/utils"
	"github.com/gin-gonic/gin"
)

type (
	GenreController interface {
		GetAllGenre(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	genreController struct {
		genreService service.GenreService
	}
)

func NewGenreController(s service.GenreService) GenreController{
	return &genreController{
		genreService: s,
	}
}

func (c *genreController) GetAllGenre(ctx *gin.Context) {
	result, err := c.genreService.GetAllGenre(ctx)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_GENRE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_LIST_GENRE, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *genreController) Update(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_GENRE_FROM_PARAM, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var req dto.GenreRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	req.ID = uint(id)

	result, err := c.genreService.Update(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_GENRE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_GET_GENRE, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *genreController) Create(ctx *gin.Context) {
	var req dto.GenreCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	
	result, err := c.genreService.Create(ctx, req)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_GENRE, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)		
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_GENRE, result)
	ctx.JSON(http.StatusCreated, res)
}