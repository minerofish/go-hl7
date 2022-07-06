package hl7v23

// Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CX
type CX23 struct {
}

// XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XPN
type XPN23 struct {
}

// XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XAD
type XAD23 struct {
}

// XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XTN
type XTN23 struct {
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CE
type CE23 struct {
	Identifier                  string `hl7:"1.1"`
	Text                        string `hl7:"1.2"`
	NameOFCodingSystem          string `hl7:"1.3"`
	AlternateIdentifier         string `hl7:"1.4"`
	AlternateText               string `hl7:"1.5"`
	NameOfAlternateCodingSystem string `hl7:"1.6"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/DLN
type DLN23 struct {
}

// XON - Extended Composite Name And ID For Organizations
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XON
type XON23 struct {
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XCN
type XCN23 struct {
}

// HD - Hierarchic Designator
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/HD
type HD23 struct {
	NamespaceId     string `hl7:"1.1"`
	UniversalId     string `hl7:"1.2"`
	UniversalIdType string `hl7:"1.3"`
}

type Sex string

const (
	Female  Sex = "F"
	Male    Sex = "M"
	Other   Sex = "O"
	Unknown Sex = "U"
)
