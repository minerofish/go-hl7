package hl7v24

// HL7 v2.4 - TCD - Test Code Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/TCD
type TCD struct {
	UniversalServiceIdentifier            CE     `hl7:"1" json:"UniversalServiceIdentifier"`
	AntiDilutionFactor                    SN     `hl7:"2" json:"AntiDilutionFactor"`
	RerunDilutionFactor                   SN     `hl7:"3" json:"RerunDilutionFactor"`
	PreDilutionFActor                     SN     `hl7:"4" json:"PreDilutionFActor"`
	EndogenousContentOfPreDilutionDiluent SN     `hl7:"5" json:"EndogenousContentOfPreDilutionDiluent"`
	AutomaticRepeatAlowwed                string `hl7:"6" json:"AutomaticRepeatAlowwed"`
	ReflexAllowed                         string `hl7:"7" json:"ReflexAllowed"`
	AnalyteRepeatStatus                   CE     `hl7:"8" json:"AnalyteRepeatStatus"`
}
