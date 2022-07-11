package hl7v24

// HL7 v2.4 - SID - Substance Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SID
type SID struct {
	ApplicationOrMethodIdentifier   CE     `hl7:"1"`
	SubstanceLotNumber              string `hl7:"2"`
	SubstanceContainerIdentifier    string `hl7:"3"`
	SubstanceManufacturerIdentifier CE     `hl7:"4"`
}
