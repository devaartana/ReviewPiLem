package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/devaartana/ReviewPiLem/entity"
	"gorm.io/gorm"
)

func ListUserSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("migrations/json/users.json")
	if err != nil {
		return err 
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var listUser []entity.User
	if err := json.Unmarshal(jsonData, &listUser); err != nil {
		return err 
	}

	hasTable := db.Migrator().HasTable(&entity.User{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err 
		}
	} 

	for _, data := range listUser {
		var user entity.User

		result := db.Where(&entity.User{Username: data.Username}).First(&user)
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