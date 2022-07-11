package hl7v23

import "time"

type AL1 struct {
	SetId              string    `hl7:"1"`
	AllergyType        string    `hl7:"2"`
	AllergyCode        CE        `hl7:"3"`
	AllergySeverity    string    `hl7:"4"`
	AllergyReaction    string    `hl7:"5"`
	IdentificationDate time.Time `hl7:"6,shortdate"`
}
