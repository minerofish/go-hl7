package hl7v23

// BLG - Billing
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/BLG
type BLG struct {
	WhenToCharge CM_CCD `hl7:"1"`
	ChargeType   string `hl7:"2"`
	AccountID    CK     `hl7:"3"`
}
