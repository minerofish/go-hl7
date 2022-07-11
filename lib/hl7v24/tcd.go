package hl7v24

// HL7 v2.4 - TCD - Test Code Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/TCD
type TCD struct {
	UniversalServiceIdentifier            CE     `hl7:"1"`
	AntiDilutionFactor                    SN     `hl7:"2"`
	RerunDilutionFactor                   SN     `hl7:"3"`
	PreDilutionFActor                     SN     `hl7:"4"`
	EndogenousContentOfPreDilutionDiluent SN     `hl7:"5"`
	AutomaticRepeatAlowwed                string `hl7:"6"`
	ReflexAllowed                         string `hl7:"7"`
	AnalyteRepeatStatus                   CE     `hl7:"8"`
}
