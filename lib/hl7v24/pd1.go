package hl7v24

import "time"

// PD1 - Patient Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PD1
type PD1 struct {
	LivingDependency                  []string  `hl7:"1" json:"LivingDependency"`
	LivingArrangement                 string    `hl7:"2" json:"LivingArrangement"`
	PrimaryFacility                   []XON     `hl7:"3" json:"PrimaryFacility"`
	PrimaryCarePrvoider               []XCN     `hl7:"4" json:"PrimaryCarePrvoider"`
	StudentIndicator                  string    `hl7:"5" json:"StudentIndicator"`
	Handicap                          string    `hl7:"6" json:"Handicap"`
	LivingWill                        string    `hl7:"7" json:"LivingWill"`
	OrganDonor                        string    `hl7:"8" json:"OrganDonor"`
	SeparateBill                      string    `hl7:"9" json:"SeparateBill"`
	DuplicatePatient                  []CX      `hl7:"10" json:"DuplicatePatient"`
	PublicityIndicator                CE        `hl7:"11" json:"PublicityIndicator"`
	ProtectionIndicator               string    `hl7:"12" json:"ProtectionIndicator"`
	ProtectionIndicatorEffectiveDate  time.Time `hl7:"13" json:"ProtectionIndicatorEffectiveDate"`
	PlaceOfWorship                    []XON     `hl7:"14" json:"PlaceOfWorship"`
	AdvanceDirectiveCode              []CE      `hl7:"15" json:"AdvanceDirectiveCode"`
	ImmunizationRegistryStatus        string    `hl7:"16" json:"ImmunizationRegistryStatus"`
	ImmunizationRegistryEffectiveDate time.Time `hl7:"17" json:"ImmunizationRegistryEffectiveDate"`
	PublicityCodeEffectiveDate        time.Time `hl7:"18" json:"PublicityCodeEffectiveDate"`
	MilitaryBranch                    string    `hl7:"19" json:"MilitaryBranch"`
	MilitaryRank                      string    `hl7:"10" json:"MilitaryRank"`
	MilitaryStatus                    string    `hl7:"21" json:"MilitaryStatus"`
}
