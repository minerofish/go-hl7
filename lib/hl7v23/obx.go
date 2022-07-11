package hl7v23

import "time"

// OBX - Observation segment
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/OBX
type OBX struct {
	SetID                        int       `hl7:"1,sequence"`
	ValueType                    string    `hl7:"2"`
	ObservationIdentifier        CE        `hl7:"3"`
	ObservationSubID             string    `hl7:"4"`
	ObervationValue              []string  `hl7:"5"`
	Units                        CE        `hl7:"6"`
	ReferenceRange               string    `hl7:"7"`
	AbnormalFlags                []string  `hl7:"8"`
	Probability                  float32   `hl7:"9"`
	NatureOfAbnormalTest         []string  `hl7:"10"`
	ResultStatus                 string    `hl7:"11"`
	DateLastObservedNormalValues time.Time `hl7:"12,longdate"`
	UserDefinedAccessChecks      string    `hl7:"13"`
	DateTimeOfObervation         time.Time `hl7:"14,longdate"`
	ProducersID                  CE        `hl7:"15"`
	ResponsibleObserver          XCN       `hl7:"16"`
	ObservationMethod            []CE      `hl7:"17"`
}
