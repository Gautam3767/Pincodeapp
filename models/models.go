package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Pincode model
type Pincode struct {
	ID                uint   `gorm:"primaryKey"`
	Officename        string `json:"officename"`
	Pincode           string `json:"pincode"`
	OfficeType        string `json:"officeType"`
	Deliverystatus    string `json:"deliverystatus"`
	Divisionname      string `json:"divisionname"`
	Regionname        string `json:"regionname"`
	Circlename        string `json:"circlename"`
	Taluk             string `json:"taluk"`
	Districtname      string `json:"districtname"`
	Statename         string `json:"statename"`
	Telephone         string `json:"telephone"`
	RelatedSuboffice  string `json:"relatedSuboffice"`
	RelatedHeadoffice string `json:"relatedHeadoffice"`
}

// Initialize the database
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("pincodes.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Pincode{})
}
