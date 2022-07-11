package hl7v24

import "time"

// HL7 v2.4 - ROL - Role
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ROL
type ROL struct {
	RoleInstanceID       EI        `hl7:"1"`
	ActionCode           string    `hl7:"2"`
	RoleROL              CE        `hl7:"3"`
	RolePerson           []XCN     `hl7:"4"`
	RoleBeginDateTime    time.Time `hl7:"5"`
	RoleEndDateTime      time.Time `hl7:"6"`
	RoleDuration         CE        `hl7:"7"`
	RoleActionReason     CE        `hl7:"8"`
	ProviderType         []CE      `hl7:"9"`
	OrganizationUnitType CE        `hl7:"10"`
	OfficeHomeAddress    []XAD     `hl7:"11"`
	Phone                []XTN     `hl7:"12"`
}
