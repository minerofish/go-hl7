package hl7v24

import "time"

// HL7 v2.4 - MSH - Message Header
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSH
type MSH struct {
	FieldSeparator string `hl7:"1,fieldseparator" json:"FieldSeparator"`
	// the fieldseparater is regarded as field#1 which doesnt work out with splitting the fields
	// therefore the fields are all offest -1 to the documentation
	EncodingCharacters                  string    `hl7:"1,delimiter" json:"EncodingCharacters"`
	SendingApplication                  HD        `hl7:"2" json:"SendingApplication"`
	SendingFacility                     HD        `hl7:"3" json:"SendingFacility"`
	ReceivingApplication                HD        `hl7:"4" json:"ReceivingApplication"`
	ReceivingFacility                   HD        `hl7:"5" json:"ReceivingFacility"`
	DateTimeOfMessage                   time.Time `hl7:"6,longdate" json:"DateTimeOfMessage"`
	Security                            string    `hl7:"7" json:"Security"`
	MessageType                         string    `hl7:"8.1" json:"MessageType"`
	MessageTriggerEvent                 string    `hl7:"8.2" json:"MessageTriggerEvent"`
	MessageControlID                    string    `hl7:"9" json:"MessageControlID"`
	ProccessingID                       string    `hl7:"10" json:"ProccessingID"`
	VersionID                           string    `hl7:"11" json:"VersionID"`
	SequenceNumber                      int       `hl7:"12" json:"SequenceNumber"`
	ContinuationPointer                 string    `hl7:"13" json:"ContinuationPointer"`
	AcceptAcknowledgementType           string    `hl7:"14" json:"AcceptAcknowledgementType"`
	ApplicationAcknowledgementType      string    `hl7:"15" json:"ApplicationAcknowledgementType"`
	CountryCode                         string    `hl7:"16" json:"CountryCode"`
	CharacterSet                        []string  `hl7:"17" json:"CharacterSet"`
	PrincipalLanguageOfMessage          CE        `hl7:"18" json:"PrincipalLanguageOfMessage"`
	AlternateCharacterSetHandlingScheme string    `hl7:"19" json:"AlternateCharacterSetHandlingScheme"`
	ConformanceStatementID              string    `hl7:"20" json:"ConformanceStatementID"`
}
