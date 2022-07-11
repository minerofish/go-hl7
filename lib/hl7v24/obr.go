package hl7v24

import "time"

// HL7 v2.4 - OBR - Observation Reques
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OBR
type OBR struct {
	ObservationRequest                   string    `hl7:"1"`
	PlacerOrderNumber                    EI        `hl7:"2"`
	FillerOrderNumber                    EI        `hl7:"3"`
	UniversalServiceIdentifier           CE        `hl7:"4"`
	Priority                             string    `hl7:"5"`
	RequestedDateTime                    time.Time `hl7:"6"`
	ObservationDateTime                  time.Time `hl7:"7"`
	ObservationEndDateTime               time.Time `hl7:"8"`
	CollectionVolume                     CQ        `hl7:"9"`
	CollectorIdentifier                  []XCN     `hl7:"10"`
	SpecimenActionCode                   string    `hl7:"11"`
	DangerCode                           CE        `hl7:"12"`
	RelevantClinicalInformation          string    `hl7:"13"`
	SpecimenReceivedDateTime             time.Time `hl7:"14"`
	SpecimenSource                       SPS       `hl7:"15"`
	OrderingProvider                     []XCN     `hl7:"16"`
	OrderCallbackPhoneNumber             []XTN     `hl7:"17"`
	PlacerField1                         string    `hl7:"18"`
	PlacerField2                         string    `hl7:"19"`
	FillerField1                         string    `hl7:"20"`
	FillerField2                         string    `hl7:"21"`
	ResultsRptStatusChngDateTime         time.Time `hl7:"22"`
	ChargeToPractice                     MOC       `hl7:"23"`
	DiagnosticServiceSectionID           string    `hl7:"24"`
	ResultStatus                         string    `hl7:"25"`
	ParentResult                         PRL       `hl7:"26"`
	QuantityTiming                       []TQ      `hl7:"27"`
	ResultCopiesTo                       []XCN     `hl7:"28"`
	ParentNumber                         EIP       `hl7:"29"`
	TransportationMode                   string    `hl7:"30"`
	ReasonForStudy                       []CE      `hl7:"31"`
	PrincipalResultInterpreter           NDL       `hl7:"32"`
	AssistantResultInterpreter           []NDL     `hl7:"33"`
	Technician                           []NDL     `hl7:"34"`
	Transcriptionist                     []NDL     `hl7:"35"`
	ScheduledDateTime                    time.Time `hl7:"36"`
	NumberOfSampleContainers             int       `hl7:"37"`
	TransportLogisticsOfCollectedSample  []CE      `hl7:"38"`
	CollectorsComment                    []CE      `hl7:"39"`
	TransportArrangementResponsibility   CE        `hl7:"40"`
	TransportArranged                    string    `hl7:"41"`
	EscortRequired                       string    `hl7:"42"`
	PlannedPatientTransportComment       []CE      `hl7:"43"`
	ProcedureCode                        CE        `hl7:"44"`
	ProcedureCodeModifier                []CE      `hl7:"45"`
	PlacerSupplementalServiceInformation []CE      `hl7:"46"`
	FillerSupplementalSErviceInformation []CE      `hl7:"47"`
}
