package hl7v24

// HL7 v2.4 - MSA - Message Acknowledgment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/MSA
type MSA struct {
	AcknowledgementCode       string `hl7:"1"`
	MessageControlID          string `hl7:"2"`
	TextMessage               string `hl7:"3"`
	ExpectedSequenceNumber    int    `hl7:"4"`
	DelayedAcknowledgmentType string `hl7:"5"`
	ErrorCondition            CE     `hl7:"6"`
}
