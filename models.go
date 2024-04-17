package main

import "gorm.io/gorm"

type Composer struct {
	gorm.Model
	Name         string `json:"name"`
	Home_Country string `json:"home_country"`
	Pieces       []Piece
}

type Piece struct {
	gorm.Model
	Name        string `json:"name"`
	Alt_Name    string `json:"alt_name"`
	Difficulty  int    `json:"difficulty"`
	Composer_ID int    `json:"composer_id" gorm:"foreignkey:ComposerRefer"`
}
