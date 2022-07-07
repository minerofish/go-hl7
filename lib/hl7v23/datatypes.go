package hl7v23

import "time"

// Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CX
type CX23 struct {
	Id                                         string `hl7:"0.1"`
	CheckDigit                                 string `hl7:"0.2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD23   `hl7:"0.3"`
	AssigningAuthority                         string `hl7:"0.4"`
	AssigningFacility                          HD23   `hl7:"0.5"`
}

// XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XPN
type XPN23 struct {
	FamilyName          string `hl7:"0.1"`
	GivenName           string `hl7:"0.2"`
	MiddleInitialOrName string `hl7:"0.3"`
	Suffix              string `hl7:"0.4"`
	Prefix              string `hl7:"0.5"`
	Degree              string `hl7:"0.6"`
	NameTypeCode        string `hl7:"0.7"`
	NameRepresentation  string `hl7:"0.8"`
}

// XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XAD
type XAD23 struct {
	StreetAddress              string `hl7:"0.1"`
	OtherDesignation           string `hl7:"0.2"`
	City                       string `hl7:"0.3"`
	StateOrProvince            string `hl7:"0.4"`
	ZipOrPostalCode            string `hl7:"0.5"`
	Country                    string `hl7:"0.6"`
	AddressType                string `hl7:"0.7"`
	OtherGeographicDesignation string `hl7:"0.8"`
	CountyCode                 string `hl7:"0.9"`
	CensusTract                string `hl7:"0.10"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XTN
type XTN23 struct {
	TelephoneNumber                string `hl7:"0.1"`
	TelecommunicationUseCode       string `hl7:"0.2"`
	TelecommunicationEquipmentType string `hl7:"0.3"`
	EmailAddress                   string `hl7:"0.4"`
	CountryCode                    int    `hl7:"0.5"`
	AreaCode                       int    `hl7:"0.6"`
	PhoneNumber                    int    `hl7:"0.7"`
	Extension                      int    `hl7:"0.8"`
	AnyText                        string `hl7:"0.9"`
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CE
type CE23 struct {
	Identifier                  string `hl7:"0.1"`
	Text                        string `hl7:"0.2"`
	NameOFCodingSystem          string `hl7:"0.3"`
	AlternateIdentifier         string `hl7:"0.4"`
	AlternateText               string `hl7:"0.5"`
	NameOfAlternateCodingSystem string `hl7:"0.6"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/DLN
type DLN23 struct {
	DriverLicenseNumber         string `hl7:"0.1"`
	IssuingStateProvinceCountry string `hl7:"0.2"`
	ExpirationDate              string `hl7:"0.3"`
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
	NamespaceId     string `hl7:"0.1"`
	UniversalId     string `hl7:"0.2"`
	UniversalIdType string `hl7:"0.3"`
}

type Sex string

const (
	Female  Sex = "F"
	Male    Sex = "M"
	Other   Sex = "O"
	Unknown Sex = "U"
)

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/PL
type PL struct {
	PointOfCare        string `hl7:"0.1"`
	Room               string `hl7:"0.2"`
	Bed                string `hl7:"0.3"`
	Facility           HD23   `hl7:"0.4"`
	LocationStatus     string `hl7:"0.5"`
	PersonLocationType string `hl7:"0.6"`
	Building           string `hl7:"0.7"`
	Floor              string `hl7:"0.8"`
	LocationType       string `hl7:"0.9"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DLD
type CM_DLD struct {
	DischargeLocation string    `hl7:"0.1"`
	EffectiveDate     time.Time `hl7:"0.2"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string    `hl7:"0.1"`
	EffectiveDate  time.Time `hl7:"0.2"`
}
