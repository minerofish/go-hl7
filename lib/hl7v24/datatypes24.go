package hl7v24

import "time"

// HL7 v2.4 - CX - Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CX
type CX struct {
	Id                                         string    `hl7:"0"`
	CheckDigit                                 string    `hl7:"1"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD        `hl7:"2"`
	AssigningAuthority                         HD        `hl7:"3"`
	IdentifierTypeCode                         string    `hl7:"4"`
	AssigningFacility                          HD        `hl7:"5"`
	EffectiveDate                              time.Time `hl7:"6"`
	ExpirationDate                             time.Time `hl7:"7"`
}

// HL7 v2.4 - CWE - Coded With Exceptions
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CWE
type CWE struct {
	Identifier                     string `hl7:"0"`
	Text                           string `hl7:"1"`
	NameOfCodingSystem             string `hl7:"2"`
	AlternateIdentifier            string `hl7:"3"`
	AlternateText                  string `hl7:"4"`
	NameOfAlternateCodingSysstem   string `hl7:"5"`
	CodingSystemVersionId          string `hl7:"6"`
	AlternateCodingSystemVersionId string `hl7:"7"`
	OriginalText                   string `hl7:"8"`
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CE
type CE struct {
	Identifier                  string `hl7:"0"`
	Text                        string `hl7:"1"`
	NameOFCodingSystem          string `hl7:"2"`
	AlternateIdentifier         string `hl7:"3"`
	AlternateText               string `hl7:"4"`
	NameOfAlternateCodingSystem string `hl7:"5"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLN
type DLN struct {
	DriverLicenseNumber         string `hl7:"0"`
	IssuingStateProvinceCountry string `hl7:"1"`
	ExpirationDate              string `hl7:"2"`
}

// DR - Date/time Range
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DR
type DR struct {
	RangeStartDateTime time.Time `hl7:"0,longdate"`
	RangeStartEndTime  time.Time `hl7:"1,longdate"`
}

// HL7 v2.4 - ELD - Error
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/ELD
type ELD struct {
	SegmentID            string `hl7:"0"`
	Sequence             int    `hl7:"1"`
	FieldPosition        int    `hl7:"2"`
	CodeIdentifyingError CE     `hl7:"3"`
}

// FN - Family Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/FN
type FN struct {
	Surname              string `hl7:"0"`
	OwnSurnamePrefix     string `hl7:"1"`
	OwnSurname           string `hl7:"2"`
	SurnamePrefixPartner string `hl7:"3"`
	SurnamePartner       string `hl7:"4"`
}

// HL7 v2.4 - HD - Hierarchic Designator
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/HD
type HD struct {
	NamespaceId     string `hl7:"0"`
	UniversalId     string `hl7:"1"`
	UniversalIdType string `hl7:"2"`
}

// HL7 v2.4 - NA - Numeric Array
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NA
type NA struct {
	Value1 float32 `hl7:"0"`
	Value2 float32 `hl7:"1"`
	Value3 float32 `hl7:"2"`
	Value4 float32 `hl7:"3"`
}

// HL7 v2.4 - PI - Person Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PI
type PI struct {
	IDNumber            string `hl7:"0"`
	TypeOfIDNumber      string `hl7:"1"`
	OtherQualifyingInfo string `hl7:"2"`
}

// HL7 v2.4 - SN - Structured Numeric
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SN
type SN struct {
	Comparator        string  `hl7:"0"`
	Num1              float32 `hl7:"0"`
	SeparatorOrSuffix string  `hl7:"0"`
	Num2              float32 `hl7:"0"`
}

// HL7 v2.4 - VID - Version Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/VID
type VID struct {
	VersionID                string `hl7:"0"`
	InternationalizationCode CE     `hl7:"1"`
	InternationalVersionId   CE     `hl7:"2"`
}

// HL7 v2.4 - XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XPN
type XPN struct {
	FamilyName         FN     `hl7:"0"`
	GivenName          string `hl7:"1"`
	SecondGivenName    string `hl7:"2"`
	Suffix             string `hl7:"3"`
	Prefix             string `hl7:"4"`
	Degree             string `hl7:"5"`
	NameTypeCode       string `hl7:"6"`
	NameRepresentation string `hl7:"7"`
	NameContext        CE     `hl7:"8"`
	NameValidityRange  DR     `hl7:"9"`
	NameAssemblyOrder  string `hl7:"10"`
}

// HL7 v2.4 - XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XAD
type XAD struct {
	StreetAddress              string `hl7:"0"`
	OtherDesignation           string `hl7:"1"`
	City                       string `hl7:"2"`
	StateOrProvince            string `hl7:"3"`
	ZipOrPostalCode            string `hl7:"4"`
	Country                    string `hl7:"5"`
	AddressType                string `hl7:"6"`
	OtherGeographicDesignation string `hl7:"7"`
	CountyCode                 string `hl7:"8"`
	CensusTract                string `hl7:"9"`
	AddressRepresentationCode  string `hl7:"10"`
	AddressValidityRange       DR     `hl7:"11"`
}

// HL7 v2.4 - XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XTN
type XTN struct {
	TelephoneNumber                string `hl7:"0"`
	TelecommunicationUseCode       string `hl7:"1"`
	TelecommunicationEquipmentType string `hl7:"2"`
	EmailAddress                   string `hl7:"3"`
	CountryCode                    int    `hl7:"4"`
	AreaCode                       int    `hl7:"5"`
	PhoneNumber                    int    `hl7:"6"`
	Extension                      int    `hl7:"7"`
	AnyText                        string `hl7:"8"`
}

// HL7 v2.4 - XON - Extended Composite Name And Identification Number For Organizations
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XON
type XON struct {
	OrganizationName                           string  `hl7:"0"`
	OrganizationNameTypeCode                   string  `hl7:"1"`
	IdNumber                                   float32 `hl7:"2"`
	CheckDigit                                 string  `hl7:"3"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string  `hl7:"4"`
	AssigningAuthority                         HD      `hl7:"5"`
	IdentifyerTypeCode                         string  `hl7:"6"`
	AssigningFAcilityId                        HD      `hl7:"7"`
	NameRepresentationCode                     string  `hl7:"8"`
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XCN
type XCN struct {
	ID                                         string `hl7:"0"`
	FamilyName                                 FN     `hl7:"1"`
	GivenName                                  string `hl7:"2"`
	MiddleInitialOrName                        string `hl7:"3"`
	Suffix                                     string `hl7:"4"`
	Prefix                                     string `hl7:"5"`
	Degree                                     string `hl7:"6"`
	SourceTable                                string `hl7:"7"`
	AssigningAuthority                         HD     `hl7:"8"`
	NameType                                   string `hl7:"9"`
	IdentifierCheckDigit                       string `hl7:"10"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"11"`
	IdentifierTypeCode                         string `hl7:"12"`
	AssigningFacility                          HD     `hl7:"13"`
	NameRepresentationCode                     string `hl7:"14"`
	NameContext                                CE     `hl7:"15"`
	NameValidityRange                          DR     `hl7:"16"`
	NameAssemblyOrder                          string `hl7:"17"`
}

// CM_AUI - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_AUI
/*type CM_AUI struct {
	AuthorizationNumber string    `hl7:"0"`
	Date                time.Time `hl7:"1,shortdate"`
	Source              string    `hl7:"2"`
}*/

// CM_RMC - Room Coverage
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_RMC
/*type CM_RMC struct {
	RoomType       string  `hl7:"0"`
	AmmountType    string  `hl7:"1"`
	CoverageAmount float32 `hl7:"2"`
}*/

// CM_PTA - Policy Type
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PTA
/*type CM_PTA struct {
	PolicyType  string  `hl7:"0"`
	AmountClass string  `hl7:"1"`
	Amount      float32 `hl7:"2"`
}*/

// CM_DDI - Daily Deductible
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DDI
/*type CM_DDI struct {
	DelayDays    float32 `hl7:"0"`
	Amount       float32 `hl7:"1"`
	NumberOfDays float32 `hl7:"2"`
}*/

// HL7 v2.4 - SPS - Specimen Source
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SPS
type SPS struct {
	SpecimenSourceNameOrCode     CE     `hl7:"0"`
	Additives                    string `hl7:"1"`
	Freetext                     string `hl7:"2"`
	BodySite                     CE     `hl7:"3"`
	SiteModifier                 CE     `hl7:"4"`
	CollectionModifierMethodCode CE     `hl7:"5"`
	SpecimenRole                 CE     `hl7:"5"`
}

// HL7 v2.4 - MOC - Charge To Practise
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MOC
type MOC struct {
	DollarAmount MO `hl7:"0"`
	ChargeCode   CE `hl7:"1"`
}

// HL7 v2.4 - PRL - Parent Result Link
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PRL
type PRL struct {
	ObservationIdentifierOfParentResult CE     `hl7:"0"`
	SubIDOfParentResult                 string `hl7:"1"`
	ObservationResultFromParent         string `hl7:"2"`
}

// HL7 v2.4 - NDL - Observing Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NDL
type NDL struct {
	OPName             CN        `hl7:"0"`
	StartDatetime      time.Time `hl7:"1"`
	EndDatetime        time.Time `hl7:"2"`
	PointOfCare        string    `hl7:"3"`
	Room               string    `hl7:"4"`
	Bed                string    `hl7:"5"`
	Facility           HD        `hl7:"6"`
	LocationStatus     string    `hl7:"7"`
	PersonLocationType string    `hl7:"8"`
	Building           string    `hl7:"9"`
	Floor              string    `hl7:"10"`
}

// HL7 v2.4 - MO - Money
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MO
type MO struct {
	Quantity     int    `hl7:"0"`
	Denomination string `hl7:"1"`
}

// HL7 v2.4 - CN - Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CN
type CN struct {
	IDNumber            string `hl7:"0"`
	FamilyName          string `hl7:"1"`
	GivenName           string `hl7:"2"`
	MiddleInitialOrName string `hl7:"3"`
	Suffix              string `hl7:"4"`
	Prefix              string `hl7:"5"`
	Degree              string `hl7:"6"`
	SourceTable         string `hl7:"7"`
	AssigningAuthority  string `hl7:"8"`
}

// HL7 v2.4 - JCC - Job Code/class
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/JCC
type JCC struct {
	JobCode  string `hl7:"0"`
	JobClass string `hl7:"1"`
}

// HL7 v2.4 - CP - Composite Price
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CP
type CP struct {
	Price      float32 `hl7:"0"`
	PriceType  string  `hl7:"1"`
	FromValue  float32 `hl7:"2"`
	ToValue    float32 `hl7:"3"`
	RangeUnits CE      `hl7:"4"`
	RangeType  string  `hl7:"5"`
}

type Sex string

const (
	Female  Sex = "F"
	Male    Sex = "M"
	Other   Sex = "O"
	Unknown Sex = "U"
)

// HL7 v2.4 - PL - Person Location
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PL
type PL struct {
	PointOfCare         string `hl7:"0"`
	Room                string `hl7:"1"`
	Bed                 string `hl7:"2"`
	Facility            HD     `hl7:"3"`
	LocationStatus      string `hl7:"4"`
	PersonLocationType  string `hl7:"5"`
	Building            string `hl7:"6"`
	Floor               string `hl7:"7"`
	LocationDescription string `hl7:"8"`
}

// HL7 v2.4 - DLD - Discharge Location
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLD
type DLD struct {
	DischargeLocation string    `hl7:"0"`
	EffectiveDate     time.Time `hl7:"1"`
}

// HL7 v2.3 - FC - Financial Class
//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string    `hl7:"0"`
	EffectiveDate  time.Time `hl7:"1"`
}

// HL7 v2.4 - EIP - Parent Order
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EIP
type EIP struct {
	ParentsPlacerOrderNumber string `hl7:"0"`
	ParentsFillerOrderNumber string `hl7:"1"`
}

// HL7 v2.4 - TQ - Timing Quantity
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/TQ
type TQ struct {
	Quantity          CQ        `hl7:"0"`
	Interval          RI        `hl7:"1"`
	Duration          string    `hl7:"2"`
	StartDatetime     time.Time `hl7:"3"`
	EndDatetime       time.Time `hl7:"4"`
	Priority          string    `hl7:"5"`
	Condition         string    `hl7:"6"`
	Text              string    `hl7:"7"`
	Conjunction       string    `hl7:"8"`
	OrderSequencing   OSD       `hl7:"9"`
	OccurenceDuration CE        `hl7:"10"`
	TotalOccurences   float32   `hl7:"11"`
}

// CQ - Composite Quantity With Units
//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CQ
type CQ struct {
	Quantity int    `hl7:"0"`
	Units    string `hl7:"1"`
}

// HL7 v2.4 - RI - Repeat Interval
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/RI
type RI struct {
	RepeatPattern        string `hl7:"0"`
	ExplicitTimeInterval string `hl7:"1"`
}

// HL7 v2.4 - OSD - Order Sequence
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/OSD
type OSD struct {
	SequenceResultsFlag               string `hl7:"0"`
	PlacerOrderNumberEntityIdentifier string `hl7:"1"`
	PlacerOrderNumberNamespaceID      string `hl7:"2"`
	FillerOrderNumberEntityIdentifier string `hl7:"3"`
	FillerOrderNumberNamespaceID      string `hl7:"4"`
	SequenceConditionValue            string `hl7:"5"`
	MaximumNumberOfRepeats            int    `hl7:"6"`
	PlacerOrderNumberUniversalID      string `hl7:"7"`
	PlacerOrderNumberUniversalIDType  string `hl7:"8"`
	FillerOrderNumberUniversalID      string `hl7:"9"`
	FillerOrderNumberUniversalIDType  string `hl7:"10"`
}

// HL7 v2.4 - EI - Entity Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"0"`
	NamespaceId      string `hl7:"1"`
	UniversalId      string `hl7:"2"`
	UniversalIdType  string `hl7:"3"`
}

/*
// CM_PEN - Penalty
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PEN
type CM_PEN struct {
	PenaltyType   string  `hl7:"0"`
	PenaltyAmount float32 `hl7:"1"`
}


// CM_DTN - Day Type And Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DTN
type CM_DTN struct {
	DayType      string  `hl7:"0"`
	NumberOfDays float32 `hl7:"1"`
}

// CM_PCF - Pre-certification Required
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PCF
type CM_PCF struct {
	PreCertificationPatient  string `hl7:"0"`
	PreCertificationRequired string `hl7:"1"`
	PreCertificationWindow   string `hl7:"2"`
}
*/

// HL7 v2.4 - LA1 - Location With Address Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/LA1
type LA1 struct {
	PointOfCare        string `hl7:"0"`
	Room               string `hl7:"1"`
	Bed                string `hl7:"2"`
	Facility           HD     `hl7:"3"`
	LocationStatus     string `hl7:"4"`
	PersonLocationType string `hl7:"5"`
	Building           string `hl7:"6"`
	Floor              string `hl7:"7"`
	Address            AD     `hl7:"8"`
}

// HL7 v2.4 - AD - Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/AD
type AD struct {
	StreetAddress              string `hl7:"0"`
	OtherDesignation           string `hl7:"1"`
	City                       string `hl7:"2"`
	StateOrProvince            string `hl7:"3"`
	PostalCode                 string `hl7:"4"`
	Country                    string `hl7:"5"`
	AddressType                string `hl7:"6"`
	OtherGeographicDesignation string `hl7:"7"`
}

/*
// CM_CCD - Charge Time
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_CCD
type CM_CCD struct {
	WhenToChargeCode string    `hl7:"0"`
	DateTime         time.Time `hl7:"1,longdate"`
}
*/

// HL7 v2.4 - CK - Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CK
type CK struct {
	IDNumber           int    `hl7:"0"`
	CheckDigit         string `hl7:"1"`
	CodeIdentifyingCC  string `hl7:"2"`
	AssigningAuthority HD     `hl7:"3"`
}
