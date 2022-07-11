package hl7v24

import "time"

type FT1 struct {
	SetID                     int       `hl7:"1,sequence"`
	TransactionID             string    `hl7:"2"`
	TransactionBatchID        string    `hl7:"3"`
	TransactionDate           time.Time `hl7:"4,longdate"`
	TransactionPostingDate    time.Time `hl7:"5,longdate"`
	TransactionType           string    `hl7:"6"`
	TransactionCode           CE        `hl7:"7"`
	TransactionDescription    string    `hl7:"8"`
	TransactionDescriptionAlt string    `hl7:"9"`
	TransactionQuantity       float32   `hl7:"10"`
	TransactionAmountExtended CP        `hl7:"11"`
	TransactionAmountUnit     CP        `hl7:"12"`
	DepartmentCode            CE        `hl7:"13"`
	InsurancePlanID           CE        `hl7:"14"`
	InsuranceAmount           CP        `hl7:"15"`
	AssignedPatientLocaton    PL        `hl7:"16"`
	FeeSchedule               string    `hl7:"17"`
	PatientType               string    `hl7:"18"`
	DiagnosisCodeFT1          []CE      `hl7:"19"`
	PerformedByCode           []XCN     `hl7:"20"`
	OrderedByCode             []XCN     `hl7:"21"`
	UnitCost                  CP        `hl7:"22"`
	FillerOrderNUmber         EI        `hl7:"23"`
	EnteredByCode             []XCN     `hl7:"24"`
	ProcedureCode             CE        `hl7:"25"`
	ProcedureCodeModifier     []CE      `hl7:"26"`
}
