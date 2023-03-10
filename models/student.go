package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name         string `json:"name" validate:"nonzero"`
	Registration string `json:"registration" validate:"len=6, regexp=^[0-9]*$"`
	Shift        string `json:"shift" validate:"nonzero, len=5"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}

	return nil
}
