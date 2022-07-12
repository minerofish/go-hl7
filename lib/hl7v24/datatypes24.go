package hl7v24

import "time"

// HL7 v2.4 - CX - Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CX
type CX struct {
	Id                                         string    `hl7:"0" json:"Id"`
	CheckDigit                                 string    `hl7:"1" json:"CheckDigit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD        `hl7:"2" json:"CodeIdentifyingTheCheckDigitSchemeEmployed"`
	AssigningAuthority                         HD        `hl7:"3" json:"AssigningAuthority"`
	IdentifierTypeCode                         string    `hl7:"4" json:"IdentifierTypeCode"`
	AssigningFacility                          HD        `hl7:"5" json:"AssigningFacility"`
	EffectiveDate                              time.Time `hl7:"6" json:"EffectiveDate"`
	ExpirationDate                             time.Time `hl7:"7" json:"ExpirationDate"`
}

// HL7 v2.4 - CWE - Coded With Exceptions
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CWE
type CWE struct {
	Identifier                     string `hl7:"0" json:"Identifier"`
	Text                           string `hl7:"1" json:"Text"`
	NameOfCodingSystem             string `hl7:"2" json:"NameOfCodingSystem"`
	AlternateIdentifier            string `hl7:"3" json:"AlternateIdentifier"`
	AlternateText                  string `hl7:"4" json:"AlternateText"`
	NameOfAlternateCodingSystem    string `hl7:"5" json:"NameOfAlternateCodingSystem"`
	CodingSystemVersionId          string `hl7:"6" json:"CodingSystemVersionId"`
	AlternateCodingSystemVersionId string `hl7:"7" json:"AlternateCodingSystemVersionId"`
	OriginalText                   string `hl7:"8" json:"OriginalText"`
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CE
type CE struct {
	Identifier                  string `hl7:"0" json:"Identifier"`
	Text                        string `hl7:"1" json:"Text"`
	NameOfCodingSystem          string `hl7:"2" json:"NameOfCodingSystem"`
	AlternateIdentifier         string `hl7:"3" json:"AlternateIdentifier"`
	AlternateText               string `hl7:"4" json:"AlternateText"`
	NameOfAlternateCodingSystem string `hl7:"5" json:"NameOfAlternateCodingSystem"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLN
type DLN struct {
	DriverLicenseNumber         string    `hl7:"0" json:"DriverLicenseNumber"`
	IssuingStateProvinceCountry string    `hl7:"1" json:"IssuingStateProvinceCountry"`
	ExpirationDate              time.Time `hl7:"2,shortdate" json:"ExpirationDate"`
}

// DR - Date/time Range
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DR
type DR struct {
	RangeStartDateTime    time.Time `hl7:"0,longdate" json:"RangeStartDateTime"`
	RangeStartEndDateTime time.Time `hl7:"1,longdate" json:"RangeStartEndDateTime"`
}

// HL7 v2.4 - ELD - Error
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/ELD
type ELD struct {
	SegmentID            string `hl7:"0" json:"SegmentID"`
	Sequence             int    `hl7:"1" json:"Sequence"`
	FieldPosition        int    `hl7:"2" json:"FieldPosition"`
	CodeIdentifyingError CE     `hl7:"3" json:"CodeIdentifyingError"`
}

// FN - Family Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/FN
type FN struct {
	Surname              string `hl7:"0" json:"Surname"`
	OwnSurnamePrefix     string `hl7:"1" json:"OwnSurnamePrefix"`
	OwnSurname           string `hl7:"2" json:"OwnSurname"`
	SurnamePrefixPartner string `hl7:"3" json:"SurnamePrefixPartner"`
	SurnamePartner       string `hl7:"4" json:"SurnamePartner"`
}

// HL7 v2.4 - HD - Hierarchic Designator
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/HD
type HD struct {
	NamespaceId     string `hl7:"0" json:"NamespaceId"`
	UniversalId     string `hl7:"1" json:"UniversalId"`
	UniversalIdType string `hl7:"2" json:"UniversalIdType"`
}

// HL7 v2.4 - NA - Numeric Array
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NA
type NA struct {
	Value1 float32 `hl7:"0" json:"Value1"`
	Value2 float32 `hl7:"1" json:"Value2"`
	Value3 float32 `hl7:"2" json:"Value3"`
	Value4 float32 `hl7:"3" json:"Value4"`
}

// HL7 v2.4 - PI - Person Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PI
type PI struct {
	IDNumber            string `hl7:"0" json:"IDNumber"`
	TypeOfIDNumber      string `hl7:"1" json:"TypeOfIDNumber"`
	OtherQualifyingInfo string `hl7:"2" json:"OtherQualifyingInfo"`
}

// HL7 v2.4 - SN - Structured Numeric
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SN
type SN struct {
	Comparator        string  `hl7:"0" json:"Comparator"`
	Num1              float32 `hl7:"0" json:"Num1"`
	SeparatorOrSuffix string  `hl7:"0" json:"SeparatorOrSuffix"`
	Num2              float32 `hl7:"0" json:"Num2"`
}

// HL7 v2.4 - VID - Version Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/VID
type VID struct {
	VersionID                string `hl7:"0" json:"VersionID"`
	InternationalizationCode CE     `hl7:"1" json:"InternationalizationCode"`
	InternationalVersionId   CE     `hl7:"2" json:"InternationalVersionId"`
}

// HL7 v2.4 - XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XPN
type XPN struct {
	FamilyName         FN     `hl7:"0" json:"FamilyName"`
	GivenName          string `hl7:"1" json:"GivenName"`
	SecondGivenName    string `hl7:"2" json:"SecondGivenName"`
	Suffix             string `hl7:"3" json:"Suffix"`
	Prefix             string `hl7:"4" json:"Prefix"`
	Degree             string `hl7:"5" json:"Degree"`
	NameTypeCode       string `hl7:"6" json:"NameTypeCode"`
	NameRepresentation string `hl7:"7" json:"NameRepresentation"`
	NameContext        CE     `hl7:"8" json:"NameContext"`
	NameValidityRange  DR     `hl7:"9" json:"NameValidityRange"`
	NameAssemblyOrder  string `hl7:"10" json:"NameAssemblyOrder"`
}

// HL7 v2.4 - XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XAD
type XAD struct {
	StreetAddress              string `hl7:"0" json:"StreetAddress"`
	OtherDesignation           string `hl7:"1" json:"OtherDesignation"`
	City                       string `hl7:"2" json:"City"`
	StateOrProvince            string `hl7:"3" json:"StateOrProvince"`
	ZipOrPostalCode            string `hl7:"4" json:"ZipOrPostalCode"`
	Country                    string `hl7:"5" json:"Country"`
	AddressType                string `hl7:"6" json:"AddressType"`
	OtherGeographicDesignation string `hl7:"7" json:"OtherGeographicDesignation"`
	CountyCode                 string `hl7:"8" json:"CountyCode"`
	CensusTract                string `hl7:"9" json:"CensusTract"`
	AddressRepresentationCode  string `hl7:"10" json:"AddressRepresentationCode"`
	AddressValidityRange       DR     `hl7:"11" json:"AddressValidityRange"`
}

// HL7 v2.4 - XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XTN
type XTN struct {
	TelephoneNumber                string `hl7:"0" json:"TelephoneNumber"`
	TelecommunicationUseCode       string `hl7:"1" json:"TelecommunicationUseCode"`
	TelecommunicationEquipmentType string `hl7:"2" json:"TelecommunicationEquipmentType"`
	EmailAddress                   string `hl7:"3" json:"EmailAddress"`
	CountryCode                    int    `hl7:"4" json:"CountryCode"`
	AreaCode                       int    `hl7:"5" json:"AreaCode"`
	PhoneNumber                    int    `hl7:"6" json:"PhoneNumber"`
	Extension                      int    `hl7:"7" json:"Extension"`
	AnyText                        string `hl7:"8" json:"AnyText"`
}

// HL7 v2.4 - XON - Extended Composite Name And Identification Number For Organizations
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XON
type XON struct {
	OrganizationName                           string  `hl7:"0" json:"OrganizationName"`
	OrganizationNameTypeCode                   string  `hl7:"1" json:"OrganizationNameTypeCode"`
	IdNumber                                   float32 `hl7:"2" json:"IdNumber"`
	CheckDigit                                 string  `hl7:"3" json:"CheckDigit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string  `hl7:"4" json:"CodeIdentifyingTheCheckDigitSchemeEmployed"`
	AssigningAuthority                         HD      `hl7:"5" json:"AssigningAuthority"`
	IdentifyerTypeCode                         string  `hl7:"6" json:"IdentifyerTypeCode"`
	AssigningFAcilityId                        HD      `hl7:"7" json:"AssigningFAcilityId"`
	NameRepresentationCode                     string  `hl7:"8" json:"NameRepresentationCode"`
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/XCN
type XCN struct {
	ID                                         string `hl7:"0" json:"ID"`
	FamilyName                                 FN     `hl7:"1" json:"FamilyName"`
	GivenName                                  string `hl7:"2" json:"GivenName"`
	MiddleInitialOrName                        string `hl7:"3" json:"MiddleInitialOrName"`
	Suffix                                     string `hl7:"4" json:"Suffix"`
	Prefix                                     string `hl7:"5" json:"Prefix"`
	Degree                                     string `hl7:"6" json:"Degree"`
	SourceTable                                string `hl7:"7" json:"SourceTable"`
	AssigningAuthority                         HD     `hl7:"8" json:"AssigningAuthority"`
	NameType                                   string `hl7:"9" json:"NameType"`
	IdentifierCheckDigit                       string `hl7:"10" json:"IdentifierCheckDigit"`
	CodeIdentifyingTheCheckDigitSchemeEmployed string `hl7:"11" json:"CodeIdentifyingTheCheckDigitSchemeEmployed"`
	IdentifierTypeCode                         string `hl7:"12" json:"IdentifierTypeCode"`
	AssigningFacility                          HD     `hl7:"13" json:"AssigningFacility"`
	NameRepresentationCode                     string `hl7:"14" json:"NameRepresentationCode"`
	NameContext                                CE     `hl7:"15" json:"NameContext"`
	NameValidityRange                          DR     `hl7:"16" json:"NameValidityRange"`
	NameAssemblyOrder                          string `hl7:"17" json:"NameAssemblyOrder"`
}

// HL7 v2.4 - AUI - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/AUI
type AUI struct {
	AuthorizationNumber string    `hl7:"0" json:"AuthorizationNumber"`
	Date                time.Time `hl7:"1,shortdate" json:"Date"`
	Source              string    `hl7:"2" json:"Source"`
}

// HL7 v2.4 - RMC - Room Coverage
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/RMC
type RMC struct {
	RoomType       string  `hl7:"0" json:"RoomType"`
	AmountType     string  `hl7:"1" json:"AmountType"`
	CoverageAmount float32 `hl7:"2" json:"CoverageAmount"`
}

// HL7 v2.4 - PTA - Policy Type
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PTA
type PTA struct {
	PolicyType  string  `hl7:"0" json:"PolicyType"`
	AmountClass string  `hl7:"1" json:"AmountClass"`
	Amount      float32 `hl7:"2" json:"Amount"`
}

// HL7 v2.4 - DDI - Daily Deductible
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DDI
type DDI struct {
	DelayDays    float32 `hl7:"0" json:"DelayDays"`
	Amount       float32 `hl7:"1" json:"Amount"`
	NumberOfDays float32 `hl7:"2" json:"NumberOfDays"`
}

// HL7 v2.4 - SPS - Specimen Source
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/SPS
type SPS struct {
	SourceNameOrCode             CE     `hl7:"0" json:"SourceNameOrCode"`
	Additives                    string `hl7:"1" json:"Additives"`
	Freetext                     string `hl7:"2" json:"Freetext"`
	BodySite                     CE     `hl7:"3" json:"BodySite"`
	SiteModifier                 CE     `hl7:"4" json:"SiteModifier"`
	CollectionModifierMethodCode CE     `hl7:"5" json:"CollectionModifierMethodCode"`
	Role                         CE     `hl7:"6" json:"Role"`
}

// HL7 v2.4 - MOC - Charge To Practise
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MOC
type MOC struct {
	DollarAmount MO `hl7:"0" json:"DollarAmount"`
	ChargeCode   CE `hl7:"1" json:"ChargeCode"`
}

// HL7 v2.4 - PRL - Parent Result Link
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PRL
type PRL struct {
	ObservationIdentifierOfParentResult CE     `hl7:"0" json:"ObservationIdentifierOfParentResult"`
	SubIDOfParentResult                 string `hl7:"1" json:"SubIDOfParentResult"`
	ObservationResultFromParent         string `hl7:"2" json:"ObservationResultFromParent"`
}

// HL7 v2.4 - NDL - Observing Practitioner
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/NDL
type NDL struct {
	OPName             CN        `hl7:"0" json:"OPName"`
	StartDatetime      time.Time `hl7:"1" json:"StartDatetime"`
	EndDatetime        time.Time `hl7:"2" json:"EndDatetime"`
	PointOfCare        string    `hl7:"3" json:"PointOfCare"`
	Room               string    `hl7:"4" json:"Room"`
	Bed                string    `hl7:"5" json:"Bed"`
	Facility           HD        `hl7:"6" json:"Facility"`
	LocationStatus     string    `hl7:"7" json:"LocationStatus"`
	PersonLocationType string    `hl7:"8" json:"PersonLocationType"`
	Building           string    `hl7:"9" json:"Building"`
	Floor              string    `hl7:"10" json:"Floor"`
}

// HL7 v2.4 - MO - Money
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MO
type MO struct {
	Quantity     int    `hl7:"0" json:"Quantity"`
	Denomination string `hl7:"1" json:"Denomination"`
}

// HL7 v2.4 - CN - Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CN
type CN struct {
	IDNumber            string `hl7:"0" json:"IDNumber"`
	FamilyName          string `hl7:"1" json:"FamilyName"`
	GivenName           string `hl7:"2" json:"GivenName"`
	MiddleInitialOrName string `hl7:"3" json:"MiddleInitialOrName"`
	Suffix              string `hl7:"4" json:"Suffix"`
	Prefix              string `hl7:"5" json:"Prefix"`
	Degree              string `hl7:"6" json:"Degree"`
	SourceTable         string `hl7:"7" json:"SourceTable"`
	AssigningAuthority  string `hl7:"8" json:"AssigningAuthority"`
}

// HL7 v2.4 - JCC - Job Code/class
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/JCC
type JCC struct {
	JobCode  string `hl7:"0" json:"JobCode"`
	JobClass string `hl7:"1" json:"JobClass"`
}

// HL7 v2.4 - CP - Composite Price
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CP
type CP struct {
	Price      float32 `hl7:"0" json:"Price"`
	PriceType  string  `hl7:"1" json:"PriceType"`
	FromValue  float32 `hl7:"2" json:"FromValue"`
	ToValue    float32 `hl7:"3" json:"ToValue"`
	RangeUnits CE      `hl7:"4" json:"RangeUnits"`
	RangeType  string  `hl7:"5" json:"RangeType"`
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
	PointOfCare         string `hl7:"0" json:"PointOfCare"`
	Room                string `hl7:"1" json:"Room"`
	Bed                 string `hl7:"2" json:"Bed"`
	Facility            HD     `hl7:"3" json:"Facility"`
	LocationStatus      string `hl7:"4" json:"LocationStatus"`
	PersonLocationType  string `hl7:"5" json:"PersonLocationType"`
	Building            string `hl7:"6" json:"Building"`
	Floor               string `hl7:"7" json:"Floor"`
	LocationDescription string `hl7:"8" json:"LocationDescription"`
}

// HL7 v2.4 - DLD - Discharge Location
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DLD
type DLD struct {
	DischargeLocation string    `hl7:"0" json:"DischargeLocation"`
	EffectiveDate     time.Time `hl7:"1" json:"EffectiveDate"`
}

// HL7 v2.3 - FC - Financial Class
//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string    `hl7:"0" json:"FinancialClass"`
	EffectiveDate  time.Time `hl7:"1" json:"EffectiveDate"`
}

// HL7 v2.4 - EIP - Parent Order
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EIP
type EIP struct {
	ParentsPlacerOrderNumber string `hl7:"0" json:"ParentsPlacerOrderNumber"`
	ParentsFillerOrderNumber string `hl7:"1" json:"ParentsFillerOrderNumber"`
}

// HL7 v2.4 - TQ - Timing Quantity
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/TQ
type TQ struct {
	Quantity          CQ        `hl7:"0" json:"Quantity"`
	Interval          RI        `hl7:"1" json:"Interval"`
	Duration          string    `hl7:"2" json:"Duration"`
	StartDatetime     time.Time `hl7:"3" json:"StartDatetime"`
	EndDatetime       time.Time `hl7:"4" json:"EndDatetime"`
	Priority          string    `hl7:"5" json:"Priority"`
	Condition         string    `hl7:"6" json:"Condition"`
	Text              string    `hl7:"7" json:"Text"`
	Conjunction       string    `hl7:"8" json:"Conjunction"`
	OrderSequencing   OSD       `hl7:"9" json:"OrderSequencing"`
	OccurenceDuration CE        `hl7:"10" json:"OccurenceDuration"`
	TotalOccurences   float32   `hl7:"11" json:"TotalOccurences"`
}

// CQ - Composite Quantity With Units
//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CQ
type CQ struct {
	Quantity int    `hl7:"0" json:"Quantity"`
	Units    string `hl7:"1" json:"Units"`
}

// HL7 v2.4 - RI - Repeat Interval
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/RI
type RI struct {
	RepeatPattern        string `hl7:"0" json:"RepeatPattern"`
	ExplicitTimeInterval string `hl7:"1" json:"ExplicitTimeInterval"`
}

// HL7 v2.4 - OSD - Order Sequence
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/OSD
type OSD struct {
	SequenceResultsFlag               string `hl7:"0" json:"SequenceResultsFlag"`
	PlacerOrderNumberEntityIdentifier string `hl7:"1" json:"PlacerOrderNumberEntityIdentifier"`
	PlacerOrderNumberNamespaceID      string `hl7:"2" json:"PlacerOrderNumberNamespaceID"`
	FillerOrderNumberEntityIdentifier string `hl7:"3" json:"FillerOrderNumberEntityIdentifier"`
	FillerOrderNumberNamespaceID      string `hl7:"4" json:"FillerOrderNumberNamespaceID"`
	SequenceConditionValue            string `hl7:"5" json:"SequenceConditionValue"`
	MaximumNumberOfRepeats            int    `hl7:"6" json:"MaximumNumberOfRepeats"`
	PlacerOrderNumberUniversalID      string `hl7:"7" json:"PlacerOrderNumberUniversalID"`
	PlacerOrderNumberUniversalIDType  string `hl7:"8" json:"PlacerOrderNumberUniversalIDType"`
	FillerOrderNumberUniversalID      string `hl7:"9" json:"FillerOrderNumberUniversalID"`
	FillerOrderNumberUniversalIDType  string `hl7:"10" json:"FillerOrderNumberUniversalIDType"`
}

// HL7 v2.4 - EI - Entity Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"0" json:"EntityIdentifier"`
	NamespaceId      string `hl7:"1" json:"NamespaceId"`
	UniversalId      string `hl7:"2" json:"UniversalId"`
	UniversalIdType  string `hl7:"3" json:"UniversalIdType"`
}

/*
// CM_PEN - Penalty
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PEN
type CM_PEN struct {
	PenaltyType   string  `hl7:"0" json:""`
	PenaltyAmount float32 `hl7:"1" json:""`
}*/

// HL7 v2.4 - DTN - Day Type And Number
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/DTN
type DTN struct {
	DayType      string  `hl7:"0" json:"DayType"`
	NumberOfDays float32 `hl7:"1" json:"NumberOfDays"`
}

// HL7 v2.4 - PCF - Pre-certification Required
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/PCF
type PCF struct {
	PreCertificationPatient  string    `hl7:"0" json:"PreCertificationPatient"`
	PreCertificationRequired string    `hl7:"1" json:"PreCertificationRequired"`
	PreCertificationWindow   time.Time `hl7:"2,longdate" json:"PreCertificationWindow"`
}

// HL7 v2.4 - LA1 - Location With Address Information
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/LA1
type LA1 struct {
	PointOfCare        string `hl7:"0" json:"PointOfCare"`
	Room               string `hl7:"1" json:"Room"`
	Bed                string `hl7:"2" json:"Bed"`
	Facility           HD     `hl7:"3" json:"Facility"`
	LocationStatus     string `hl7:"4" json:"LocationStatus"`
	PersonLocationType string `hl7:"5" json:"PersonLocationType"`
	Building           string `hl7:"6" json:"Building"`
	Floor              string `hl7:"7" json:"Floor"`
	Address            AD     `hl7:"8" json:"Address"`
}

// HL7 v2.4 - AD - Address
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/AD
type AD struct {
	StreetAddress              string `hl7:"0" json:"StreetAddress"`
	OtherDesignation           string `hl7:"1" json:"OtherDesignation"`
	City                       string `hl7:"2" json:"City"`
	StateOrProvince            string `hl7:"3" json:"StateOrProvince"`
	PostalCode                 string `hl7:"4" json:"PostalCode"`
	Country                    string `hl7:"5" json:"Country"`
	AddressType                string `hl7:"6" json:"AddressType"`
	OtherGeographicDesignation string `hl7:"7" json:"OtherGeographicDesignation"`
}

// HL7 v2.4 - CCD - Charge Time
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CCD
type CCD struct {
	WhenToChargeCode string    `hl7:"1" json:"WhenToChargeCode"`
	DateTime         time.Time `hl7:"2,longdate" json:"DateTime"`
}

// HL7 v2.4 - CK - Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/CK
type CK struct {
	IDNumber           int    `hl7:"0" json:"IDNumber"`
	CheckDigit         string `hl7:"1" json:"CheckDigit"`
	CodeIdentifyingCC  string `hl7:"2" json:"CodeIdentifyingCC"`
	AssigningAuthority HD     `hl7:"3" json:"AssigningAuthority"`
}

// HL7 v2.4 - MOP - Money Or Percentage
// https://hl7-definition.caristix.com/v2/HL7v2.4/DataTypes/MOP
type MOP struct {
	MoneyOrPercentageIndicator string  `hl7:"0" json:"MoneyOrPercentageIndicator"`
	MoneyOrPercentageQuantity  float32 `hl7:"1" json:"MoneyOrPercentageQuantity"`
}
