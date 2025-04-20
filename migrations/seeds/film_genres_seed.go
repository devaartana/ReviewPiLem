package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/devaartana/ReviewPiLem/entity"
	"gorm.io/gorm"
)


func ListFilmGenresSeeder (db *gorm.DB) error {
	jsonFile, err := os.Open("migrations/json/film_genres.json")
	if err != nil{
		return err
	}
	defer jsonFile.Close()
	
	jsonData, err := io.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	var listFilmGenre []entity.FilmGenre
	if err := json.Unmarshal(jsonData, &listFilmGenre); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.FilmGenre{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.FilmGenre{}); err != nil {
			return err
		}
	}

	for _, data := range listFilmGenre {
		var filmGenre entity.FilmGenre

		result := db.Where(&entity.FilmGenre{FilmID: data.FilmID, GenreID: data.GenreID}).First(&filmGenre)
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