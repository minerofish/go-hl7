package hl7v23

import "time"

// DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/DG1
type DG1 struct {
	SetID                    int       `hl7:"1, sequence"`
	DiagnosisCodingMethod    string    `hl7:"2"`
	DiagnosisCode            CE        `hl7:"3"`
	DiagnosisDescription     string    `hl7:"4"`
	DiagnosisDateTime        time.Time `hl7:"5,longdate"`
	DiagnosisType            string    `hl7:"6"`
	MajjorDiagnosticCategory CE        `hl7:"7"`
	DiagnosticRelatedGroup   CE        `hl7:"8"`
	DRGApprovalIndicator     string    `hl7:"9"`
	DRGGrouperReviewCode     string    `hl7:"10"`
	OutlierType              CE        `hl7:"11"`
	OutlierDays              int       `hl7:"12"`
	OutlierCost              CP        `hl7:"13"`
	GrouperVersionAndType    string    `hl7:"14"`
	DiagnosisPriority        int       `hl7:"15"`
	DiagnosingClinican       []XCN     `hl7:"16"`
	DiagnosisClassification  string    `hl7:"17"`
	ConfidentialIndicator    string    `hl7:"18"`
	AttestationDateTime      time.Time `hl7:"19,longdate"`
}
