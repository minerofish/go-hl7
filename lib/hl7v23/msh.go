package hl7v23

import "time"

// MSH - Message header segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/MSH
type MSH struct {
	FieldSeparator                 string    `hl7:"1"`
	EncodingCharacters             string    `hl7:"2,delimiter"`
	SendingApplication             HD23      `hl7:"3"`
	SendingFacility                HD23      `hl7:"4"`
	ReceivingApplication           HD23      `hl7:"5"`
	ReceivingFacility              HD23      `hl7:"6"`
	DateTimeOfMessage              time.Time `hl7:"7"`
	Security                       string    `hl7:"8"`
	MessageType                    string    `hl7:"9.1"`
	MessageTriggerEvent            string    `hl7:"9.2"`
	MessageControlID               string    `hl7:"10"`
	ProccessingID                  string    `hl7:"11"`
	VersionID                      string    `hl7:"12"`
	SequenceNumber                 int       `hl7:"13"`
	ContinuationPointer            string    `hl7:"14"`
	AcceptAcknowledgementType      string    `hl7:"15"`
	ApplicationAcknowledgementType string    `hl7:"16"`
	CountryCode                    string    `hl7:"17"`
	CharacterSet                   []string  `hl7:"18"`
	PrincipalLanguageOfMessage     CE23      `hl7:"19"`
}
