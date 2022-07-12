package hl7v24

import (
	"time"
)

// HL7 v2.4 - AL1 - Patient allergy information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/AL1
type AL1 struct {
	SetId               CE        `hl7:"1" json:"SetId"`
	AllergenTypeCode    CE        `hl7:"2" json:"AllergenTypeCode"`
	AllergenCode        CE        `hl7:"3" json:"AllergenCode"`
	AllergySeverityCode CE        `hl7:"4" json:"AllergySeverityCode"`
	AllergyReactionCode string    `hl7:"5" json:"AllergyReactionCode"`
	IdentificationDate  time.Time `hl7:"6,shortdate" json:"IdentificationDate"`
}
