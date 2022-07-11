package hl7v24

import "time"

// PD1 - Patient Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PD1
type PD1 struct {
	LivingDependency                  []string  `hl7:"1"`
	LivingArrangement                 string    `hl7:"2"`
	PrimaryFacility                   []XON     `hl7:"3"`
	PrimaryCarePrvoider               []XCN     `hl7:"4"`
	StudentIndicator                  string    `hl7:"5"`
	Handicap                          string    `hl7:"6"`
	LivingWill                        string    `hl7:"7"`
	OrganDonor                        string    `hl7:"8"`
	SeparateBill                      string    `hl7:"9"`
	DuplicatePatient                  []CX      `hl7:"10"`
	PublicityIndicator                CE        `hl7:"11"`
	ProtectionIndicator               string    `hl7:"12"`
	ProtectionIndicatorEffectiveDate  time.Time `hl7:"13"`
	PlaceOfWorship                    []XON     `hl7:"14"`
	AdvanceDirectiveCode              []CE      `hl7:"15"`
	ImmunizationRegistryStatus        string    `hl7:"16"`
	ImmunizationRegistryEffectiveDate time.Time `hl7:"17"`
	PublicityCodeEffectiveDate        time.Time `hl7:"18"`
	MilitaryBranch                    string    `hl7:"19"`
	MilitaryRank                      string    `hl7:"10"`
	MilitaryStatus                    string    `hl7:"21"`
}
