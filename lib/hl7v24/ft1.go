package hl7v24

import "time"

type FT1 struct {
	SetID                     int       `hl7:"1,sequence" json:"SetID"`
	TransactionID             string    `hl7:"2" json:"TransactionID"`
	TransactionBatchID        string    `hl7:"3" json:"TransactionBatchID"`
	TransactionDate           time.Time `hl7:"4,longdate" json:"TransactionDate"`
	TransactionPostingDate    time.Time `hl7:"5,longdate" json:"TransactionPostingDate"`
	TransactionType           string    `hl7:"6" json:"TransactionType"`
	TransactionCode           CE        `hl7:"7" json:"TransactionCode"`
	TransactionDescription    string    `hl7:"8" json:"TransactionDescription"`
	TransactionDescriptionAlt string    `hl7:"9" json:"TransactionDescriptionAlt"`
	TransactionQuantity       float32   `hl7:"10" json:"TransactionQuantity"`
	TransactionAmountExtended CP        `hl7:"11" json:"TransactionAmountExtended"`
	TransactionAmountUnit     CP        `hl7:"12" json:"TransactionAmountUnit"`
	DepartmentCode            CE        `hl7:"13" json:"DepartmentCode"`
	InsurancePlanID           CE        `hl7:"14" json:"InsurancePlanID"`
	InsuranceAmount           CP        `hl7:"15" json:"InsuranceAmount"`
	AssignedPatientLocaton    PL        `hl7:"16" json:"AssignedPatientLocaton"`
	FeeSchedule               string    `hl7:"17" json:"FeeSchedule"`
	PatientType               string    `hl7:"18" json:"PatientType"`
	DiagnosisCodeFT1          []CE      `hl7:"19" json:"DiagnosisCodeFT1"`
	PerformedByCode           []XCN     `hl7:"20" json:"PerformedByCode"`
	OrderedByCode             []XCN     `hl7:"21" json:"OrderedByCode"`
	UnitCost                  CP        `hl7:"22" json:"UnitCost"`
	FillerOrderNUmber         EI        `hl7:"23" json:"FillerOrderNUmber"`
	EnteredByCode             []XCN     `hl7:"24" json:"EnteredByCode"`
	ProcedureCode             CE        `hl7:"25" json:"ProcedureCode"`
	ProcedureCodeModifier     []CE      `hl7:"26" json:"ProcedureCodeModifier"`
}
