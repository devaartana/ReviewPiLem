package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ListReactionFilm(db *gorm.DB) error {
	jsonFile, err := os.Open("migrations/json/reaction.json")
	if err != nil {
		return err 
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listReviewFilm []entity.Reaction
	if err := json.Unmarshal(jsonData, &listReviewFilm); err != nil {
		return err 
	}

	hasTable := db.Migrator().HasTable(&entity.Reaction{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Reaction{}); err != nil {
			return err
		}
	}
	
	var id1 string
	result1 := db.Model(&entity.User{}).Select("id").Where("username = ?", "john pork").Scan(&id1)
	if result1.Error != nil {
		return result1.Error
	}

	var id2 string
	result2 := db.Model(&entity.User{}).Select("id").Where("username = ?", "johndoe").Scan(&id2)
	if result2.Error != nil {
		return result2.Error
	}

	data1, err := uuid.Parse(id1)
	if err != nil {
		return err
	}
	listReviewFilm[0].UserID = data1
	
	data2, err := uuid.Parse(id2)
	if err != nil {
		return err
	}
	listReviewFilm[1].UserID = data2


	for _, data := range listReviewFilm {
		var existingReview entity.Reaction
		err := db.Where(&entity.Reaction{ReviewID: data.ReviewID, UserID: data.UserID}).First(&existingReview).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}