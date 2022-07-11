package hl7v23

import "time"

// ORC - Common order segment
//https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/ORC
type ORC struct {
	OrderControl           string    `hl7:"1"`
	PlacerOrderNumber      EI        `hl7:"2"`
	FillerOrderNumber      EI        `hl7:"3"`
	PlacerGroupNumber      EI        `hl7:"4"`
	OrderStatus            string    `hl7:"5"`
	ResponseFlag           string    `hl7:"6"`
	QuantityTiming         TQ        `hl7:"7"`
	ParentOrder            CM_EIP    `hl7:"8"`
	DateTimeOfTransaction  time.Time `hl7:"9"`
	EnteredBy              XCN       `hl7:"10"`
	VerifiedBy             XCN       `hl7:"11"`
	OrderingProvider       XCN       `hl7:"12"`
	EnterersLocation       PL        `hl7:"13"`
	CallBackPhoneNumber    []XTN     `hl7:"14"`
	OrderEffectiveDateTime time.Time `hl7:"15"`
	OrderControlCodeReason CE        `hl7:"16"`
	EnteringOrganization   CE        `hl7:"17"`
	EnteringDevice         CE        `hl7:"18"`
	ActionBy               XCN       `hl7:"19"`
}
