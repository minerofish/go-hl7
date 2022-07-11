package hl7v24

// HL7 v2.4 - ERR - Error
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ERR
type ERR struct {
	ERrorCodeAndLocation ELD `hl7:"1"`
}
