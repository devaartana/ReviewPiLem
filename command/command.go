package command

import (
	"log"
	"os"

	"github.com/devaartana/ReviewPiLem/constants"
	"github.com/devaartana/ReviewPiLem/migrations"
	"github.com/samber/do"
	"gorm.io/gorm"
)


func Commands(injector *do.Injector) bool {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	
	migrate := false
	down := false
	seed := false
	run := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--seed" {
			seed = true
		}
		if arg == "--run" {
			run = true
		}
		if arg == "--down" {
			down = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration completed successfully")
	}

	if down {
		if err := migrations.Down(db); err != nil {
			log.Fatalf("error dropping table: %v", err)
		}
		log.Println("drop completed successfully")
	}

	if seed {
		if err := migrations.Seeder(db); err != nil {
			log.Fatalf("error migration seeder: %v", err)
		}
		log.Println("seeder completed successfully")
	}

	if run {
		return true
	}

	return false
}
