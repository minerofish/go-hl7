package hl7v23

// RXO - Pharmacy prescription order segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/RXO
type RXO struct {
	RequestedGiveGode             CE      `hl7:"1"`
	RequestedGiveAmountMinimung   float32 `hl7:"2"`
	RequestedGiveAmountMaximum    float32 `hl7:"3"`
	RequestedGiveUnits            CE      `hl7:"4"`
	RequestedDosageForm           CE      `hl7:"5"`
	ProvidersPharmacyInstructions []CE    `hl7:"6"`
	DeliverToLocation             []LA1   `hl7:"7"`
	AllowSubstitutions            string  `hl7:"8"`
	RequestedDispenseCode         CE      `hl7:"9"`
	RequestedDispenseAmount       float32 `hl7:"10"`
	RequestedDispenseUnits        CE      `hl7:"11"`
	NumberOfRefills               float32 `hl7:"12"`
	OrderingProvidersDEA          XCN     `hl7:"13"`
	PharmacistSupplierVerifierID  XCN     `hl7:"14"`
	NeedsHumanReview              string  `hl7:"15"`
	RequestedGivePer              string  `hl7:"16"`
	RequestedGiveStrength         float32 `hl7:"17"`
	RequestedGiveStrengthUnits    CE      `hl7:"18"`
	Indication                    CE      `hl7:"19"`
	RequestedGiveRateAmount       string  `hl7:"20"`
	RequestedGiveRateUnits        CE      `hl7:"21"`
}
