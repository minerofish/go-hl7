package hl7v24

import "time"

// MSH - Message header segment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSH
type MSH struct {
	FieldSeparator string `hl7:"1,fieldseparator"`
	// the fieldseparater is regarded as field#1 which doesnt work out with splitting the fields
	// therefore the fields are all offest -1 to the documentation
	EncodingCharacters                  string    `hl7:"1,delimiter"`
	SendingApplication                  HD        `hl7:"2"`
	SendingFacility                     HD        `hl7:"3"`
	ReceivingApplication                HD        `hl7:"4"`
	ReceivingFacility                   HD        `hl7:"5"`
	DateTimeOfMessage                   time.Time `hl7:"6,longdate"`
	Security                            string    `hl7:"7"`
	MessageType                         string    `hl7:"8.1"`
	MessageTriggerEvent                 string    `hl7:"8.2"`
	MessageControlID                    string    `hl7:"9"`
	ProccessingID                       string    `hl7:"10"`
	VersionID                           string    `hl7:"11"`
	SequenceNumber                      int       `hl7:"12"`
	ContinuationPointer                 string    `hl7:"13"`
	AcceptAcknowledgementType           string    `hl7:"14"`
	ApplicationAcknowledgementType      string    `hl7:"15"`
	CountryCode                         string    `hl7:"16"`
	CharacterSet                        []string  `hl7:"17"`
	PrincipalLanguageOfMessage          CE        `hl7:"18"`
	AlternateCharacterSetHandlingScheme string    `hl7:"19"`
	ConformanceStatementID              string    `hl7:"20"`
}
