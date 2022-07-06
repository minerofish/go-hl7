package hl7v23

// PD1 - Patient Demographic
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PD1
type PD1 struct {
	LivingDependency    []string `hl7:"1"`
	LivingArrangement   string   `hl7:"2"`
	PrimaryFacility     []XON23  `hl7:"3"`
	PrimaryCarePrvoider []XCN23  `hl7:"4"`
	StudentIndicator    string   `hl7:"5"`
	Handicap            string   `hl7:"6"`
	LivingWill          string   `hl7:"7"`
	OrganDonor          string   `hl7:"8"`
	SeparateBill        string   `hl7:"9"`
	DuplicatePatient    []CX     `hl7:"10"`
	PublicityIndicator  CE       `hl7:"11"`
	ProtectionIndicator string   `hl7:"12"`
}
