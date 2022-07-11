package hl7v24

import "time"

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PID
type PID struct {
	ID                      int       `hl7:"1"`
	ExternalID              []CX      `hl7:"2"`
	InternalID              []CX      `hl7:"3"`
	AlternateID             []CX      `hl7:"4"`
	Name                    []XPN     `hl7:"5"`
	MothersMaidenName       []XPN     `hl7:"6"`
	DOB                     time.Time `hl7:"7"`
	Sex                     string    `hl7:"8"`
	Alias                   []XPN     `hl7:"9"`
	Race                    []CE      `hl7:"10"`
	Address                 []XAD     `hl7:"11"`
	CountryCode             string    `hl7:"12"`
	PhoneNumber             []XTN     `hl7:"13"`
	PhoneNumberBusiness     []XTN     `hl7:"14"`
	PrimaryLanguage         CE        `hl7:"15"`
	MaritalStatus           CE        `hl7:"16"`
	Religion                CE        `hl7:"17"`
	PatientAccountNumber    CX        `hl7:"18"`
	SSNNumber               string    `hl7:"19"`
	DriversLicenseNumber    DLN       `hl7:"20"`
	MothersIdentifier       []CX      `hl7:"21"`
	EthnicGroup             []string  `hl7:"22"`
	BirthPlace              string    `hl7:"23"`
	MultipleBirthIndicator  string    `hl7:"24"`
	BirthOrder              int       `hl7:"25"`
	Citizenship             []string  `hl7:"26"`
	VeteransMilitaryStatus  CE        `hl7:"27"`
	NationalityCode         CE        `hl7:"28"`
	PatientDeathDateAndTime time.Time `hl7:"29"`
	PatientDeathIndicator   string    `hl7:"30"`
	IdentityUnkownIndicator string    `hl7:"31"`
	IdentityReliabilityCode string    `hl7:"32"`
	LastUpdateDateTime      time.Time `hl7:"33"`
	LastUpdateFacility      HD        `hl7:"34"`
	SpeciesCode             CE        `hl7:"35"`
	BreedCode               CE        `hl7:"36"`
	Strain                  string    `hl7:"37"`
	ProductionClassCode     []CE      `hl7:"38"`
}
