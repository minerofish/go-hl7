package hl7v24

import "time"

// PV1 - Patient visit
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV1
type PV1 struct {
	SetID                   string      `hl7:"1" json:"SetID"`
	PatientClass            string      `hl7:"2" json:"PatientClass"`
	AssignedPatientLocation PL          `hl7:"3" json:"AssignedPatientLocation"`
	AdmissionType           string      `hl7:"4" json:"AdmissionType"`
	PreadmitNumber          CX          `hl7:"5" json:"PreadmitNumber"`
	PriorPatientLocation    PL          `hl7:"6" json:"PriorPatientLocation"`
	AttendingDoctor         []XCN       `hl7:"7" json:"AttendingDoctor"`
	ReferringDoctor         []XCN       `hl7:"8" json:"ReferringDoctor"`
	ConsultingDoctor        []XCN       `hl7:"9" json:"ConsultingDoctor"`
	HospitalService         string      `hl7:"10" json:"HospitalService"`
	TemporaryLocation       PL          `hl7:"11" json:"TemporaryLocation"`
	PreadmitTestIndicator   string      `hl7:"12" json:"PreadmitTestIndicator"`
	ReadmissionIndicator    string      `hl7:"13" json:"ReadmissionIndicator"`
	AdmitSource             string      `hl7:"14" json:"AdmitSource"`
	AmbulatoryStatus        []string    `hl7:"15" json:"AmbulatoryStatus"`
	VIPIndicator            string      `hl7:"16" json:"VIPIndicator"`
	AdmittingDoctor         []XCN       `hl7:"17" json:"AdmittingDoctor"`
	PatientType             string      `hl7:"18" json:"PatientType"`
	VisitNumber             CX          `hl7:"19" json:"VisitNumber"`
	FinancialClass          []FC        `hl7:"20" json:"FinancialClass"`
	ChargePriceIndicator    string      `hl7:"21" json:"ChargePriceIndicator"`
	CourtesyCode            string      `hl7:"22" json:"CourtesyCode"`
	CreditRating            string      `hl7:"23" json:"CreditRating"`
	ContractCode            []string    `hl7:"24" json:"ContractCode"`
	ContractEffectiveDate   []time.Time `hl7:"25,shortdate" json:"ContractEffectiveDate"`
	ContractAmount          []float32   `hl7:"26" json:"ContractAmount"`
	ContractPeriod          []float32   `hl7:"27" json:"ContractPeriod"`
	InterestCode            string      `hl7:"28" json:"InterestCode"`
	TransferToBadDebtCode   string      `hl7:"29" json:"TransferToBadDebtCode"`
	TransferToBadDebtDate   time.Time   `hl7:"30,shortdate" json:"TransferToBadDebtDate"`
	BadDebtAgencyCode       string      `hl7:"31" json:"BadDebtAgencyCode"`
	BadDebtTransferAmount   float32     `hl7:"32" json:"BadDebtTransferAmount"`
	BadDebtRecoveryAmount   float32     `hl7:"33" json:"BadDebtRecoveryAmount"`
	DeleteAccountIndicator  string      `hl7:"34" json:"DeleteAccountIndicator"`
	DeleteAccountDate       time.Time   `hl7:"35,shortdate" json:"DeleteAccountDate"`
	DischargeDisposition    string      `hl7:"36" json:"DischargeDisposition"`
	DischargedToLocation    DLD         `hl7:"37" json:"DischargedToLocation"`
	DietType                string      `hl7:"38" json:"DietType"`
	ServicingFacility       string      `hl7:"39" json:"ServicingFacility"`
	BedStatus               string      `hl7:"40" json:"BedStatus"`
	AccountStatus           string      `hl7:"41" json:"AccountStatus"`
	PendingLocation         PL          `hl7:"42" json:"PendingLocation"`
	PriorTemporaryLocation  PL          `hl7:"43" json:"PriorTemporaryLocation"`
	AdmitDateTime           time.Time   `hl7:"44,longdate" json:"AdmitDateTime"`
	DischargeDateTime       []time.Time `hl7:"45,longdate" json:"DischargeDateTime"`
	CurrentPatientBalance   float32     `hl7:"46" json:"CurrentPatientBalance"`
	TotalCharges            float32     `hl7:"47" json:"TotalCharges"`
	TotalAdjustments        float32     `hl7:"48" json:"TotalAdjustments"`
	TotalPayments           float32     `hl7:"49" json:"TotalPayments"`
	AlternateVisitID        CX          `hl7:"50" json:"AlternateVisitID"`
	VisitIndicator1         string      `hl7:"51" json:"VisitIndicator1"`
	OtherHealthcareProvider []XCN       `hl7:"52" json:"OtherHealthcareProvider"`
}
