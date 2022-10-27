package models

import "time"

type Movie struct {
	ID        int64     `gorm:"primary_key"`
	Title     string    `gorm:"type:varchar(100);not null"`
	Year      string    `gorm:"type:varchar(4);not null"`
	Rating    float32   `gorm:"type:decimal(2,1);not null;default:5"`
	AuthorID  int64     `gorm:"not null"`
	Author    *Author   `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type Author struct {
	ID        int64     `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
