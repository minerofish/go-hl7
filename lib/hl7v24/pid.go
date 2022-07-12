package hl7v24

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PID
type PID struct {
	ID                       int       `hl7:"1" json:"ID"`
	ExternalID               []CX      `hl7:"2" json:"ExternalID"`
	InternalID               []CX      `hl7:"3" json:"InternalID"`
	AlternateID              []CX      `hl7:"4" json:"AlternateID"`
	Name                     []XPN     `hl7:"5" json:"Name"`
	MothersMaidenName        []XPN     `hl7:"6" json:"MothersMaidenName"`
	DOB                      time.Time `hl7:"7" json:"DOB"`
	Sex                      string    `hl7:"8" json:"Sex"`
	Alias                    []XPN     `hl7:"9" json:"Alias"`
	Race                     []CE      `hl7:"10" json:"Race"`
	Address                  []XAD     `hl7:"11" json:"Address"`
	CountryCode              string    `hl7:"12" json:"CountryCode"`
	PhoneNumber              []XTN     `hl7:"13" json:"PhoneNumber"`
	PhoneNumberBusiness      []XTN     `hl7:"14" json:"PhoneNumberBusiness"`
	PrimaryLanguage          CE        `hl7:"15" json:"PrimaryLanguage"`
	MaritalStatus            CE        `hl7:"16" json:"MaritalStatus"`
	Religion                 CE        `hl7:"17" json:"Religion"`
	PatientAccountNumber     CX        `hl7:"18" json:"PatientAccountNumber"`
	SSNNumber                string    `hl7:"19" json:"SSNNumber"`
	DriversLicenseNumber     DLN       `hl7:"20" json:"DriversLicenseNumber"`
	MothersIdentifier        []CX      `hl7:"21" json:"MothersIdentifier"`
	EthnicGroup              []string  `hl7:"22" json:"EthnicGroup"`
	BirthPlace               string    `hl7:"23" json:"BirthPlace"`
	MultipleBirthIndicator   string    `hl7:"24" json:"MultipleBirthIndicator"`
	BirthOrder               int       `hl7:"25" json:"BirthOrder"`
	Citizenship              []string  `hl7:"26" json:"Citizenship"`
	VeteransMilitaryStatus   CE        `hl7:"27" json:"VeteransMilitaryStatus"`
	NationalityCode          CE        `hl7:"28" json:"NationalityCode"`
	PatientDeathDateAndTime  time.Time `hl7:"29" json:"PatientDeathDateAndTime"`
	PatientDeathIndicator    string    `hl7:"30" json:"PatientDeathIndicator"`
	IdentityUnknownIndicator string    `hl7:"31" json:"IdentityUnknownIndicator"`
	IdentityReliabilityCode  string    `hl7:"32" json:"IdentityReliabilityCode"`
	LastUpdateDateTime       time.Time `hl7:"33" json:"LastUpdateDateTime"`
	LastUpdateFacility       HD        `hl7:"34" json:"LastUpdateFacility"`
	SpeciesCode              CE        `hl7:"35" json:"SpeciesCode"`
	BreedCode                CE        `hl7:"36" json:"BreedCode"`
	Strain                   string    `hl7:"37" json:"Strain"`
	ProductionClassCode      []CE      `hl7:"38" json:"ProductionClassCode"`
}
