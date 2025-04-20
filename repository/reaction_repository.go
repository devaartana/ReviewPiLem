package repository

import (
	"context"

	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	ReactionReposiotry interface {
		CheckReaction(ctx context.Context, tx *gorm.DB, userID uuid.UUID, reviewID uint) (entity.Reaction, bool, error)
		Create(ctx context.Context, tx *gorm.DB, react entity.Reaction) error
		Update(ctx context.Context, tx *gorm.DB, react entity.Reaction) error
		Delete(ctx context.Context, tx *gorm.DB, userID uuid.UUID, reviewID uint) error
	}

	reactionRepository struct {
		db *gorm.DB
	}
)

func NewReactionRepository(db *gorm.DB) ReactionReposiotry {
	return &reactionRepository {
		db: db,
	}
}

func (r *reactionRepository) CheckReaction(ctx context.Context, tx *gorm.DB, userID uuid.UUID, reviewID uint) (entity.Reaction, bool, error) {
	if tx == nil {
		tx = r.db
	}
	
	var reaction entity.Reaction
	err := tx.WithContext(ctx).Where("user_id = ? AND review_id = ?", userID, reviewID).First(&reaction).Error
	
	if err != nil {
		return entity.Reaction{UserID: userID, ReviewID: reviewID}, false, err
	}

	return reaction, true, nil
}

func (r *reactionRepository) Create(ctx context.Context, tx *gorm.DB, react entity.Reaction) error {
	if tx == nil {
		tx = r.db
	}

	err := tx.WithContext(ctx).Create(&react).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *reactionRepository) Update(ctx context.Context, tx *gorm.DB, react entity.Reaction) error {
	if tx == nil {
		tx = r.db
	}

	err := tx.WithContext(ctx).
		Model(&entity.Reaction{}).Where("review_id = ? AND user_id = ?", react.ReviewID, react.UserID).Update("status", react.Status).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *reactionRepository) Delete(ctx context.Context, tx *gorm.DB, userID uuid.UUID, reviewID uint) error {
	if tx == nil {
		tx = r.db
	}

	err := tx.WithContext(ctx).Where("user_id = ? AND review_id = ?", userID, reviewID).Delete(&entity.Reaction{}).Error
	if err != nil {
		return err
	}

	return nil
}

