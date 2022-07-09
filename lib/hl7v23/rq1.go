package hl7v23

// RQ1 - Requisition detail-1 segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RQ1
type RQ1 struct {
	AnticipatedPrice     string `hl7:"1"`
	ManufacturedId       CE     `hl7:"2"`
	ManufacturersCatalog string `hl7:"3"`
	VendorID             CE     `hl7:"4"`
	VendorCatalog        string `hl7:"5"`
	Taxable              string `hl7:"6"`
	SubstituteAllowed    string `hl7:"7"`
}
