package hl7v24

import "time"

// HL7 v2.4 - EQU - Equipment Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/EQU
type EQU struct {
	EquipmentIDentifier EI        `hl7:"1" json:"EquipmentIDentifier"`
	EventDateTime       time.Time `hl7:"2,longdate" json:"EventDateTime"`
	EquipmentState      CE        `hl7:"3" json:"EquipmentState"`
	LocalControlState   CE        `hl7:"4" json:"LocalControlState"`
	AlertLevel          CE        `hl7:"5" json:"AlertLevel"`
}
