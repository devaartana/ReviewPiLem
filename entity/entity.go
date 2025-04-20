package entity

import (
	"time"

	"github.com/devaartana/ReviewPiLem/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilmStatus string

const (
	FilmStatusNotYetAired    FilmStatus = "not_yet_aired"
	FilmStatusAiring         FilmStatus = "airing"
	FilmStatusFinishedAiring FilmStatus = "finished_airing"
)

type ListStatus string

const (
	ListStatusPlanToWatch ListStatus = "plan_to_watch"
	ListStatusWatching    ListStatus = "watching"
	ListStatusCompleted   ListStatus = "completed"
	ListStatusOnHold      ListStatus = "on_hold"
	ListStatusDropped     ListStatus = "dropped"
)

type Genre struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(30);unique;not null" json:"name"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

type Film struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	Title         string      `gorm:"type:varchar(100);not null" json:"title"`
	Synopsis      string      `gorm:"type:text;not null" json:"synopsis"`
	Status        FilmStatus  `gorm:"film_status;not null" json:"status"`
	TotalEpisodes int         `gorm:"not null" json:"total_episodes"`
	ReleaseDate   time.Time   `gorm:"type:date;not null" json:"release_date"`
	CreatedAt     time.Time   `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	Genres        []Genre     `gorm:"many2many:film_genres;" json:"genres"`
	Images        []FilmImage `gorm:"foreignKey:FilmID" json:"images"`
}

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Username    string    `gorm:"type:varchar(30);unique;not null" json:"username"`
	Email       string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	DisplayName string    `gorm:"type:varchar(30);not null" json:"display_name"`
	Bio         string    `gorm:"type:text" json:"bio"`
	Password    string    `gorm:"type:varchar(100);not null" json:"password"`
	Role        string    `gorm:"type:varchar(100);not null" json:"role"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

type FilmGenre struct {
	FilmID  uint `gorm:"primaryKey" json:"film_id"`
	GenreID uint `gorm:"primaryKey" json:"genre_id"`
}

type FilmImage struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FilmID uint   `gorm:"not null" json:"film_id"`
	Path   string `gorm:"type:varchar(255);not null" json:"path"`
	Status bool   `gorm:"default:false" json:"status"`
}

type UserFilmList struct {
	UserID     uuid.UUID  `gorm:"type:uuid;primaryKey;not null" json:"user_id"`
	FilmID     uint       `gorm:"primaryKey;not null" json:"film_id"`
	Status     ListStatus `gorm:"type:list_status;not null" json:"status"`
	Visibility bool       `gorm:"default:true" json:"visibility"`
	CreatedAt  time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Review struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"user_id"`
	FilmID    uint      `gorm:"primaryKey;not null" json:"film_id"`
	Rating    int       `gorm:"not null" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

type Reaction struct {
	ReviewID uint      `gorm:"primaryKey;not null" json:"review_id"`
	UserID   uuid.UUID `gorm:"primaryKey;not null" json:"user_id"`
	Status   bool      `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
