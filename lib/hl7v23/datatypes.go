package hl7v23

// Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CX
type CX23 struct {
	Id                                         string `hl7:"1.1"`
	CheckDigit                                 string `hl7:"1.2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD23   `hl7:"1.3"`
	AssigningAuthority                         string `hl7:"1.4"`
	AssigningFacility                          HD23   `hl7:"1.5"`
}

// XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XPN
type XPN23 struct {
	FamilyName          string `hl7:"1.1"`
	GivenName           string `hl7:"1.2"`
	MiddleInitialOrName string `hl7:"1.3"`
	Suffix              string `hl7:"1.4"`
	Prefix              string `hl7:"1.5"`
	Degree              string `hl7:"1.6"`
	NameTypeCode        string `hl7:"1.7"`
	NameRepresentation  string `hl7:"1.8"`
}

// XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XAD
type XAD23 struct {
	StreetAddress              string `hl7:"1.1"`
	OtherDesignation           string `hl7:"1.2"`
	City                       string `hl7:"1.3"`
	StateOrProvince            string `hl7:"1.4"`
	ZipOrPostalCode            string `hl7:"1.5"`
	Country                    string `hl7:"1.6"`
	AddressType                string `hl7:"1.7"`
	OtherGeographicDesignation string `hl7:"1.8"`
	CountyCode                 string `hl7:"1.9"`
	CensusTract                string `hl7:"1.10"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XTN
type XTN23 struct {
	TelephoneNumber                string `hl7:"1.1"`
	TelecommunicationUseCode       string `hl7:"1.2"`
	TelecommunicationEquipmentType string `hl7:"1.3"`
	EmailAddress                   string `hl7:"1.4"`
	CountryCode                    int    `hl7:"1.5"`
	AreaCode                       int    `hl7:"1.6"`
	PhoneNumber                    int    `hl7:"1.7"`
	Extension                      int    `hl7:"1.8"`
	AnyText                        string `hl7:"1.9"`
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
	DriverLicenseNumber         string `hl7:"1.1"`
	IssuingStateProvinceCountry string `hl7:"1.2"`
	ExpirationDate              string `hl7:"1.3"`
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
