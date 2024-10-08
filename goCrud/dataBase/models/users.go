package models;

type Users struct{
	Id uint `json:"id" gorm:"primarykey;autoIncrement"`
	Name string `json:"name" gorm:"not null"`
	Age int `json:"age" gorm:"not null"`
	Email string `json:"email" gorm:"not null; unique"`
	Country Countries `gorm:"foreignKey:CountryCode"`
}

