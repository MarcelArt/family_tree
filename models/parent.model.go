package models

import "gorm.io/gorm"

const parentTableName = "parents"

type Parent struct {
	gorm.Model
	// Insert your fields here
	HusbandID uint `gorm:"not null" json:"husbandId"`
	WifeID    uint `gorm:"not null" json:"wifeId"`

	Husband *Person `json:"-"`
	Wife    *Person `json:"-"`
}

type ParentDTO struct {
	DTO
	// Insert your fields here
	HusbandID uint `gorm:"not null" json:"husbandId"`
	WifeID    uint `gorm:"not null" json:"wifeId"`

	Husband *Person `json:"-"`
	Wife    *Person `json:"-"`
}

type ParentPage struct {
	// Insert your fields here
	HusbandID uint `gorm:"not null" json:"husbandId"`
	WifeID    uint `gorm:"not null" json:"wifeId"`

	Husband string `json:"husband"`
	Wife    string `json:"wife"`
}

func (ParentDTO) TableName() string {
	return parentTableName
}
