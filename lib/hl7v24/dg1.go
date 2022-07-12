package hl7v24

import "time"

// HL7 v2.4 - DG1 - Diagnosis
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/DG1
type DG1 struct {
	SetID                   int       `hl7:"1, sequence" json:"SetID "`
	DiagnosisCodingMethod   string    `hl7:"2" json:"DiagnosisCodingMethod"`
	DiagnosisCode           CE        `hl7:"3" json:"DiagnosisCode"`
	DiagnosisDescription    string    `hl7:"4" json:"DiagnosisDescription"`
	DiagnosisDateTime       time.Time `hl7:"5,longdate" json:"DiagnosisDateTime"`
	DiagnosisType           string    `hl7:"6" json:"DiagnosisType"`
	MajorDiagnosticCategory CE        `hl7:"7" json:"MajorDiagnosticCategory"`
	DiagnosticRelatedGroup  CE        `hl7:"8" json:"DiagnosticRelatedGroup"`
	DRGApprovalIndicator    string    `hl7:"9" json:"DRGApprovalIndicator"`
	DRGGrouperReviewCode    string    `hl7:"10" json:"DRGGrouperReviewCode"`
	OutlierType             CE        `hl7:"11" json:"OutlierType"`
	OutlierDays             int       `hl7:"12" json:"OutlierDays"`
	OutlierCost             CP        `hl7:"13" json:"OutlierCost"`
	GrouperVersionAndType   string    `hl7:"14" json:"GrouperVersionAndType"`
	DiagnosisPriority       int       `hl7:"15" json:"DiagnosisPriority"`
	DiagnosingClinician     []XCN     `hl7:"16" json:"DiagnosingClinician"`
	DiagnosisClassification string    `hl7:"17" json:"DiagnosisClassification"`
	ConfidentialIndicator   string    `hl7:"18" json:"ConfidentialIndicator"`
	AttestationDateTime     time.Time `hl7:"19,longdate" json:"AttestationDateTime"`
}
