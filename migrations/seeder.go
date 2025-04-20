package migrations

import (
	"github.com/devaartana/ReviewPiLem/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListFilmSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListGenresSeeder(db); err != nil {
		return err
	}
	
	if err := seeds.ListFilmImage(db); err != nil {
		return err
	}
	
	if err := seeds.ListFilmGenresSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListUserFilmSeeder(db); err != nil {
		return err
	}	

	if err := seeds.ListReviewFilm(db); err != nil {
		return err
	}
	
	if err := seeds.ListReactionFilm(db); err != nil {
		return err
	}

	return nil
}
