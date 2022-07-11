package hl7v23

import "time"

//  PV1 - Patient visit
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV1
type PV1 struct {
	SetID                   string      `hl7:"1"`
	PatientClass            string      `hl7:"2"`
	AssignedPatientLocation PL          `hl7:"3"`
	AdmissionType           string      `hl7:"4"`
	PreadmitNumber          CX          `hl7:"5"`
	PriorPatientLocation    PL          `hl7:"6"`
	AttendingDoctor         []XCN       `hl7:"7"`
	ReferringDoctor         []XCN       `hl7:"8"`
	ConsultingDoctor        []XCN       `hl7:"9"`
	HospitalService         string      `hl7:"10"`
	TemporaryLocation       PL          `hl7:"11"`
	PreadmitTestIndicator   string      `hl7:"12"`
	ReadmissionIndicator    string      `hl7:"13"`
	AdmitSource             string      `hl7:"14"`
	AmbulatoryStatus        []string    `hl7:"15"`
	VIPIndicator            string      `hl7:"16"`
	AdmittingDoctor         []XCN       `hl7:"17"`
	PatientType             string      `hl7:"18"`
	VisitNumber             CX          `hl7:"19"`
	FinancialClass          FC          `hl7:"20"`
	ChargePriceIndicator    string      `hl7:"21"`
	CourtesyCode            string      `hl7:"22"`
	CreditRating            string      `hl7:"23"`
	ContractCode            []string    `hl7:"24"`
	ContractEffectiveDate   []time.Time `hl7:"25,shortdate"`
	ContractAmount          []float32   `hl7:"26"`
	ContractPeriod          []float32   `hl7:"27"`
	InterestCode            string      `hl7:"28"`
	TransferToBadDebtCode   string      `hl7:"29"`
	TransferToBadDebtDate   time.Time   `hl7:"30,shortdate"`
	BadDebtAgencyCode       string      `hl7:"31"`
	BadDebtTransferAmount   float32     `hl7:"32"`
	BadDebtRecoveryAmount   float32     `hl7:"33"`
	DeleteAccountIndicator  string      `hl7:"34"`
	DeleteAccountDate       time.Time   `hl7:"35"`
	DischargeDisposition    string      `hl7:"36"`
	DischardedToLocation    CM_DLD      `hl7:"37"`
	DietType                string      `hl7:"38"`
	ServicingFacility       string      `hl7:"39"`
	BedStatus               string      `hl7:"40"`
	AccountStatus           string      `hl7:"41"`
	PendingLocation         PL          `hl7:"42"`
	PriorTemporaryLocation  PL          `hl7:"43"`
	AdmitDateTime           time.Time   `hl7:"44"`
	DischarcheDateTime      time.Time   `hl7:"45"`
	CurrentPatientBalance   float32     `hl7:"46"`
	TotalCharges            float32     `hl7:"47"`
	TotalAdjustments        float32     `hl7:"48"`
	TotalPayments           float32     `hl7:"49"`
	AlternateVisitID        CX          `hl7:"50"`
	VisitIndicator1         string      `hl7:"51"`
	OhterHealthcareProvider []XCN       `hl7:"52"`
}
