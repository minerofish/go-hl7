package hl7v24

// HL7 v2.4 - SID - Substance Identifier
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SID
type SID struct {
	ApplicationOrMethodIdentifier   CE     `hl7:"1" json:"ApplicationOrMethodIdentifier"`
	SubstanceLotNumber              string `hl7:"2" json:"SubstanceLotNumber"`
	SubstanceContainerIdentifier    string `hl7:"3" json:"SubstanceContainerIdentifier"`
	SubstanceManufacturerIdentifier CE     `hl7:"4" json:"SubstanceManufacturerIdentifier"`
}
