package hl7v24

// HL7 v2.4 - ACK - General acknowledgment message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ACK
type ACK struct {
	MSH                    MSH `hl7:"MSH"`
	MessageAcknowledgement MSA `hl7:"MSA"`
	Error                  ERR `hl7:"ERR"`
}
