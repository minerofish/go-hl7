package hl7v24

import "time"

// HL7 v2.4 - EQU - Equipment Detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/EQU
type EQU struct {
	EquipmentIDentifier EI        `hl7:"1"`
	EventDateTime       time.Time `hl7:"2,longdate"`
	EquipmentState      CE        `hl7:"3"`
	LocalControlState   CE        `hl7:"4"`
	AlertLevel          CE        `hl7:"5"`
}
