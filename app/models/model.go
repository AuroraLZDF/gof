package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

type Model struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
