package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name     string `json:"name"`
	FileURL  string `json:"file_url"`
}
