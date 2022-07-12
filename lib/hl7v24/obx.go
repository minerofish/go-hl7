package hl7v24

import "time"

// HL7 v2.4 - OBX - Observation/Result
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OBX
type OBX struct {
	SetID                        int       `hl7:"1,sequence" json:"SetID"`
	ValueType                    string    `hl7:"2" json:"ValueType"`
	ObservationIdentifier        CE        `hl7:"3" json:"ObservationIdentifier"`
	ObservationSubID             string    `hl7:"4" json:"ObservationSubID"`
	ObservationValue             []string  `hl7:"5" json:"ObservationValue"`
	Units                        CE        `hl7:"6" json:"Units"`
	ReferenceRange               string    `hl7:"7" json:"ReferenceRange"`
	AbnormalFlags                []string  `hl7:"8" json:"AbnormalFlags"`
	Probability                  float32   `hl7:"9" json:"Probability"`
	NatureOfAbnormalTest         []string  `hl7:"10" json:"NatureOfAbnormalTest"`
	ResultStatus                 string    `hl7:"11" json:"ResultStatus"`
	DateLastObservedNormalValues time.Time `hl7:"12,longdate" json:"DateLastObservedNormalValues"`
	UserDefinedAccessChecks      string    `hl7:"13" json:"UserDefinedAccessChecks"`
	DateTimeOfObservation        time.Time `hl7:"14,longdate" json:"DateTimeOfObservation"`
	ProducersID                  CE        `hl7:"15" json:"ProducersID"`
	ResponsibleObserver          XCN       `hl7:"16" json:"ResponsibleObserver"`
	ObservationMethod            []CE      `hl7:"17" json:"ObservationMethod"`
	EquipmentInstanceIdentifier  EI        `hl7:"18" json:"EquipmentInstanceIdentifier"`
	DateTimeOfTheAnalysis        time.Time `hl7:"19" json:"DateTimeOfTheAnalysis"`
}
