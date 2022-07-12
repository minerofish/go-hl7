package hl7v24

import "time"

// HL7 v2.4 - ROL - Role
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ROL
type ROL struct {
	RoleInstanceID       EI        `hl7:"1" json:"RoleInstanceID"`
	ActionCode           string    `hl7:"2" json:"ActionCode"`
	RoleROL              CE        `hl7:"3" json:"RoleROL"`
	RolePerson           []XCN     `hl7:"4" json:"RolePerson"`
	RoleBeginDateTime    time.Time `hl7:"5" json:"RoleBeginDateTime"`
	RoleEndDateTime      time.Time `hl7:"6" json:"RoleEndDateTime"`
	RoleDuration         CE        `hl7:"7" json:"RoleDuration"`
	RoleActionReason     CE        `hl7:"8" json:"RoleActionReason"`
	ProviderType         []CE      `hl7:"9" json:"ProviderType"`
	OrganizationUnitType CE        `hl7:"10" json:"OrganizationUnitType"`
	OfficeHomeAddress    []XAD     `hl7:"11" json:"OfficeHomeAddress"`
	Phone                []XTN     `hl7:"12" json:"Phone"`
}
