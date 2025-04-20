package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/devaartana/ReviewPiLem/entity"
	"gorm.io/gorm"
)

func ListGenresSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("migrations/json/genres.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listGenres []entity.Genre
	if err := json.Unmarshal(jsonData, &listGenres); err != nil {
		return err 
	}

	hasTable := db.Migrator().HasTable(&entity.Genre{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Genre{}); err != nil {
			return err 
		}
	}

	for _, data := range listGenres {
		var genre entity.Genre

		result := db.Where(&entity.Genre{Name: data.Name}).First(&genre)
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			if err := db.Create(&data).Error; err!= nil {
				return err
			}
		} else if result.Error != nil {
			return result.Error
		}
	}

	return nil
}