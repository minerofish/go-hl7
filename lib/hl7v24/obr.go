package hl7v24

import "time"

// HL7 v2.4 - OBR - Observation Reques
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/OBR
type OBR struct {
	SetID                                int       `hl7:"1" json:"SetID"`
	PlacerOrderNumber                    EI        `hl7:"2" json:"PlacerOrderNumber"`
	FillerOrderNumber                    EI        `hl7:"3" json:"FillerOrderNumber"`
	UniversalServiceIdentifier           CE        `hl7:"4" json:"UniversalServiceIdentifier"`
	Priority                             string    `hl7:"5" json:"Priority"`
	RequestedDateTime                    time.Time `hl7:"6" json:"RequestedDateTime"`
	ObservationDateTime                  time.Time `hl7:"7" json:"ObservationDateTime"`
	ObservationEndDateTime               time.Time `hl7:"8" json:"ObservationEndDateTime"`
	CollectionVolume                     CQ        `hl7:"9" json:"CollectionVolume"`
	CollectorIdentifier                  []XCN     `hl7:"10" json:"CollectorIdentifier"`
	SpecimenActionCode                   string    `hl7:"11" json:"SpecimenActionCode"`
	DangerCode                           CE        `hl7:"12" json:"DangerCode"`
	RelevantClinicalInformation          string    `hl7:"13" json:"RelevantClinicalInformation"`
	SpecimenReceivedDateTime             time.Time `hl7:"14" json:"SpecimenReceivedDateTime"`
	SpecimenSource                       SPS       `hl7:"15" json:"SpecimenSource"`
	OrderingProvider                     []XCN     `hl7:"16" json:"OrderingProvider"`
	OrderCallbackPhoneNumber             []XTN     `hl7:"17" json:"OrderCallbackPhoneNumber"`
	PlacerField1                         string    `hl7:"18" json:"PlacerField1"`
	PlacerField2                         string    `hl7:"19" json:"PlacerField2"`
	FillerField1                         string    `hl7:"20" json:"FillerField1"`
	FillerField2                         string    `hl7:"21" json:"FillerField2"`
	ResultsRptStatusChngDateTime         time.Time `hl7:"22" json:"ResultsRptStatusChngDateTime"`
	ChargeToPractice                     MOC       `hl7:"23" json:"ChargeToPractice"`
	DiagnosticServiceSectionID           string    `hl7:"24" json:"DiagnosticServiceSectionID"`
	ResultStatus                         string    `hl7:"25" json:"ResultStatus"`
	ParentResult                         PRL       `hl7:"26" json:"ParentResult"`
	QuantityTiming                       []TQ      `hl7:"27" json:"QuantityTiming"`
	ResultCopiesTo                       []XCN     `hl7:"28" json:"ResultCopiesTo"`
	ParentNumber                         EIP       `hl7:"29" json:"ParentNumber"`
	TransportationMode                   string    `hl7:"30" json:"TransportationMode"`
	ReasonForStudy                       []CE      `hl7:"31" json:"ReasonForStudy"`
	PrincipalResultInterpreter           NDL       `hl7:"32" json:"PrincipalResultInterpreter"`
	AssistantResultInterpreter           []NDL     `hl7:"33" json:"AssistantResultInterpreter"`
	Technician                           []NDL     `hl7:"34" json:"Technician"`
	Transcriptionist                     []NDL     `hl7:"35" json:"Transcriptionist"`
	ScheduledDateTime                    time.Time `hl7:"36" json:"ScheduledDateTime"`
	NumberOfSampleContainers             int       `hl7:"37" json:"NumberOfSampleContainers"`
	TransportLogisticsOfCollectedSample  []CE      `hl7:"38" json:"TransportLogisticsOfCollectedSample"`
	CollectorsComment                    []CE      `hl7:"39" json:"CollectorsComment"`
	TransportArrangementResponsibility   CE        `hl7:"40" json:"TransportArrangementResponsibility"`
	TransportArranged                    string    `hl7:"41" json:"TransportArranged"`
	EscortRequired                       string    `hl7:"42" json:"EscortRequired"`
	PlannedPatientTransportComment       []CE      `hl7:"43" json:"PlannedPatientTransportComment"`
	ProcedureCode                        CE        `hl7:"44" json:"ProcedureCode"`
	ProcedureCodeModifier                []CE      `hl7:"45" json:"ProcedureCodeModifier"`
	PlacerSupplementalServiceInformation []CE      `hl7:"46" json:"PlacerSupplementalServiceInformation"`
	FillerSupplementalSErviceInformation []CE      `hl7:"47" json:"FillerSupplementalSErviceInformation"`
}
