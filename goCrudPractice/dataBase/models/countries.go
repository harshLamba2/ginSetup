package models;

type Countries struct{
	Id uint `json:"id" gorm:"primarykey;autoIncrement"` //As an unsigned type, uint can only represent non-negative values (0 and positive integers). It does not include negative numbers, which makes it suitable for scenarios where negative values do not make sense
	Country string `json:"country"; gorm:"not null"`;
	CountryCode string `json:"countryCode"; gorm"not null"`
	Colonizer string `json:"colonizer"; gorm"type:varchar(100);"` // type usually adapted to be the one assigned to the struct attribute. however to assign a certain length to the datatype explicitly define a datatype like varchar(100)
	Government string `json:"government"; gorm"type:varchar(100);"` // type usually adapted to be the one assigned to the struct attribute. however to assign a certain length to the datatype explicitly define a datatype like varchar(100)
}
