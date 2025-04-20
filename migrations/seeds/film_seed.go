package seeds

import (
	"errors"
	"io"
	"os"

	"github.com/devaartana/ReviewPiLem/entity"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

func ListFilmSeeder(db *gorm.DB) error{
	jsonFile, err := os.Open("migrations/json/films.json")
	if err != nil {
		return err 
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listFilm []entity.Film
	if err := json.Unmarshal(jsonData, &listFilm); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Film{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Film{}); err != nil {
			return err 
		}
	} 
	
	for _, data := range listFilm {
		var film entity.Film

		result := db.Where(&entity.Film{Title: data.Title}).First(&film)
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