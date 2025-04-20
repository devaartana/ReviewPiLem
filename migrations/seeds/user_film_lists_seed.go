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

func ListUserFilmSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("migrations/json/user_film_lists.json")
	if err != nil {
		return err 
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listUserFilm []entity.UserFilmList
	if err := json.Unmarshal(jsonData, &listUserFilm); err != nil {
		return err 
	}

	hasTable := db.Migrator().HasTable(&entity.UserFilmList{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.UserFilmList{}); err != nil {
			return err
		}
	}
	
	var id string
	result := db.Model(&entity.User{}).Select("id").Where("username = ?", "john pork").Scan(&id)
	if result.Error != nil {
		return result.Error
	}

	for _, data := range listUserFilm{
		var userFilm entity.UserFilmList
	
		data.UserID, err = uuid.Parse(id)
		if err != nil {
			return err
		}

		result := db.Where(&entity.UserFilmList{UserID: data.UserID, FilmID: data.FilmID}).First(&userFilm)
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