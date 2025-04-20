package controller

import (
	"net/http"
	"strconv"

	"github.com/devaartana/ReviewPiLem/dto"
	service "github.com/devaartana/ReviewPiLem/services"
	"github.com/devaartana/ReviewPiLem/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	ReviewController interface {
		Like(ctx *gin.Context)
		Dislike(ctx *gin.Context)
		Delete(ctx *gin.Context)
		GetAllReviewFilm(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		DeleteReview(ctx *gin.Context)
	}

	reviewController struct {
		reviewServices service.ReviewServices
	}
)

func NewReviewController(s service.ReviewServices) ReviewController {
	return &reviewController{
		reviewServices: s,
	}
}

func (c *reviewController) GetAllReviewFilm(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_FILM_ID, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var req dto.PaginationRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.reviewServices.GetAllReviewFilm(ctx.Request.Context(), req, uint(id))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_LIST_REVIEW, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.Response{
		Status:  true,
		Message: dto.MESSAGE_SUCCESS_GET_LIST_USER,
		Data:    result.Data,
		Meta:    result.PaginationResponse,
	}

	ctx.JSON(http.StatusOK, res)
}

func (r *reviewController) Create(ctx *gin.Context) {
	userIdRaw := ctx.MustGet("user_id")
	userId, err := uuid.Parse(userIdRaw.(string))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var review dto.ReviewRequest
	if err := ctx.ShouldBind(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	review.UserID = userId

	result, err := r.reviewServices.Create(ctx, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_REVIEW, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_REVIEW, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) Update(ctx *gin.Context) {
	userIdRaw := ctx.MustGet("user_id")
	userId, err := uuid.Parse(userIdRaw.(string))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	var review dto.ReviewRequest
	if err := ctx.ShouldBind(&review); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_REVIEW, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	review.UserID = userId

	result, err := c.reviewServices.Update(ctx, review)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_UPDATE_REVIEW, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_UPDATE_REVIEW, result)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) DeleteReview(ctx *gin.Context) {
	userIdRaw := ctx.MustGet("user_id")
	userId, err := uuid.Parse(userIdRaw.(string))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_FILM_ID, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.reviewServices.DeleteReview(ctx, userId, uint(id)); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE_REVIEW, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE_REVIEW, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) Like(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userIdRaw, ok := ctx.Get("user_id")
	if !ok {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	userId, err := uuid.Parse(userIdRaw.(string))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if err := c.reviewServices.Like(ctx, uint(id), userId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REACTION, dto.ErrReaction.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REACTION, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) Dislike(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userIdRaw, ok := ctx.Get("user_id")
	if !ok {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	userId, err := uuid.Parse(userIdRaw.(string))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if err := c.reviewServices.Dislike(ctx, uint(id), userId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_REACTION, dto.ErrReaction.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_REACTION, nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *reviewController) Delete(ctx *gin.Context) {
	paramId := ctx.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrParam.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userIdRaw, ok := ctx.Get("user_id")
	if !ok {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	userId, err := uuid.Parse(userIdRaw.(string))
	if err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_GET_USER_ID, dto.ErrGetUserID.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if err := c.reviewServices.Delete(ctx, uint(id), userId); err != nil {
		res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_DELETE, dto.ErrReaction.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(dto.MESSAGE_SUCCESS_DELETE, nil)
	ctx.JSON(http.StatusOK, res)
}
