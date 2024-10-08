package models;

type Countries struct{
	Id uint `json:"id" gorm:"primarykey;autoIncrement"` //As an unsigned type, uint can only represent non-negative values (0 and positive integers). It does not include negative numbers, which makes it suitable for scenarios where negative values do not make sense
	Country string `json:"country"; gorm:"not null"`;
	CountryCode string `json:"countryCode"; gorm"not null"`
}
