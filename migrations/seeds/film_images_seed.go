package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/devaartana/ReviewPiLem/entity"
	"gorm.io/gorm"
)

func ListFilmImage (db *gorm.DB) error {
	jsonFile, err := os.Open("migrations/json/film_images.json")
	if err != nil {
		return err 
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listFilmImage []entity.FilmImage
	if err := json.Unmarshal(jsonData, &listFilmImage); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.FilmImage{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.FilmImage{}); err != nil {
			return err
		}
	} 

	for _, data := range listFilmImage {
		var filmImage entity.FilmImage 
		result := db.Where(&entity.FilmImage{Path: data.Path}).First(&filmImage)
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