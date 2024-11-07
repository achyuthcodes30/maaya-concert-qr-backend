package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	PRN       string
	SRN       string
	Email     string
	Name      string
	Semester  string
	Branch    string
	IsPresent bool
}
