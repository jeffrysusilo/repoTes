package models

import "gorm.io/gorm"

type Terminal struct {
    gorm.Model
    Name string `gorm:"unique"`
}
