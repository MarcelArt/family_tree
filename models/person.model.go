package models

import "gorm.io/gorm"

const personTableName = "people"

type Person struct {
	gorm.Model
	// Insert your fields here
	Name string `gorm:"not null" json:"name"`

	ParentID *uint `json:"parentId"`

	Parent *Parent `json:"-"`
}

type PersonDTO struct {
	DTO
	// Insert your fields here
	Name string `gorm:"not null" json:"name"`

	ParentID *uint `json:"parentId"`

	Parent *Parent `json:"-"`
}

type PersonPage struct {
	// Insert your fields here
	ID       uint   `json:"ID"`
	Name     string `gorm:"not null" json:"name"`
	ParentID *uint  `json:"parentId"`
	FatherID uint   `json:"fatherId"`
	Father   string `json:"father"`
	MotherID uint   `json:"motherId"`
	Mother   string `json:"mother"`
}

func (PersonDTO) TableName() string {
	return personTableName
}
