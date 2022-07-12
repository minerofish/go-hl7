package hl7v24

import "time"

// HL7 v2.4 - ORC - Common Order
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/ORC
type ORC struct {
	OrderControl                  string    `hl7:"1" json:"OrderControl"`
	PlacerOrderNumber             EI        `hl7:"2" json:"PlacerOrderNumber"`
	FillerOrderNumber             EI        `hl7:"3" json:"FillerOrderNumber"`
	PlacerGroupNumber             EI        `hl7:"4" json:"PlacerGroupNumber"`
	OrderStatus                   string    `hl7:"5" json:"OrderStatus"`
	ResponseFlag                  string    `hl7:"6" json:"ResponseFlag"`
	QuantityTiming                []TQ      `hl7:"7" json:"QuantityTiming"`
	ParentOrder                   EIP       `hl7:"8" json:"ParentOrder"`
	DateTimeOfTransaction         time.Time `hl7:"9" json:"DateTimeOfTransaction"`
	EnteredBy                     []XCN     `hl7:"10" json:"EnteredBy"`
	VerifiedBy                    []XCN     `hl7:"11" json:"VerifiedBy"`
	OrderingProvider              []XCN     `hl7:"12" json:"OrderingProvider"`
	EnterersLocation              PL        `hl7:"13" json:"EnterersLocation"`
	CallBackPhoneNumber           []XTN     `hl7:"14" json:"CallBackPhoneNumber"`
	OrderEffectiveDateTime        time.Time `hl7:"15" json:"OrderEffectiveDateTime"`
	OrderControlCodeReason        CE        `hl7:"16" json:"OrderControlCodeReason"`
	EnteringOrganization          CE        `hl7:"17" json:"EnteringOrganization"`
	EnteringDevice                CE        `hl7:"18" json:"EnteringDevice"`
	ActionBy                      []XCN     `hl7:"19" json:"ActionBy"`
	AdvancedBeneficiaryNoticeCode CE        `hl7:"20" json:"AdvancedBeneficiaryNoticeCode"`
	OrderingFacilityName          []XON     `hl7:"21" json:"OrderingFacilityName"`
	OrderingFacilityPhoneNumber   []XTN     `hl7:"22" json:"OrderingFacilityPhoneNumber"`
	OrderingProviderAddress       []XAD     `hl7:"23" json:"OrderingProviderAddress"`
	OrderStatusModifier           CWE       `hl7:"24" json:"OrderStatusModifier"`
}
