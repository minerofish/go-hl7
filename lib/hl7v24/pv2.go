package hl7v24

import "time"

// HL7 v2.4 - PV2 - Patient visit - additional information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PV2
type PV2 struct {
	PriorPendigLocation              PL        `hl7:"1"`
	AccomodationCode                 CE        `hl7:"2"`
	AdmitReason                      CE        `hl7:"3"`
	TransferReason                   CE        `hl7:"4"`
	PatientValuables                 string    `hl7:"5"`
	PatientValuablesLocation         string    `hl7:"6"`
	VisitUserCode                    string    `hl7:"7"`
	ExpectedAdmitDate                time.Time `hl7:"8"`
	ExpectedDischargeDate            time.Time `hl7:"9"`
	EstimatedLengthOfInpatientStay   float32   `hl7:"10"`
	ActualLengthOfImpatientStay      float32   `hl7:"11"`
	VisitDescription                 string    `hl7:"12"`
	ReferralSourceCode               []XCN     `hl7:"13"`
	PreviousServiceDate              time.Time `hl7:"14"`
	EmploymentIllnesRelatedIndicator string    `hl7:"15"`
	PurgeStatusCode                  string    `hl7:"16"`
	PurgeStatusDate                  time.Time `hl7:"17"`
	SpecialProgramCode               string    `hl7:"18"`
	RetentionIndicator               string    `hl7:"19"`
	ExpectedNumberOfInsurancePlans   float32   `hl7:"20"`
	VisitPublicitycode               string    `hl7:"21"`
	VisitProtectionIndicator         string    `hl7:"22"`
	ClinicOrganizationName           XON       `hl7:"23"`
	PatientStatusCode                string    `hl7:"24"`
	VisitPriorityCode                string    `hl7:"25"`
	PreviousTreatmentCode            string    `hl7:"26"`
	ExpectedDischargeDisposition     string    `hl7:"27"`
	SignatureOnFileDate              time.Time `hl7:"28"`
	FirstSimilarIllnessDate          time.Time `hl7:"29"`
	PatientChargeAdjustmentCode      CE        `hl7:"30"`
	RecurringServicecoce             string    `hl7:"31"`
	BillingMediaCode                 string    `hl7:"32"`
	EpxectedSurgeryDateTime          time.Time `hl7:"33"`
	MilitaryPartnershipCode          string    `hl7:"34"`
	MilitaryNonAvailabilityCode      string    `hl7:"35"`
	NewbornBabyIndicator             string    `hl7:"36"`
	BabyDetainedIndicator            string    `hl7:"37"`
	ModeOfArrivalCode                CE        `hl7:"38"`
	RecreationalDrugUseCode          []CE      `hl7:"39"`
	AdmissionLevelOfCareCode         CE        `hl7:"40"`
	PrecautionCode                   []CE      `hl7:"41"`
	PatientConditionCode             CE        `hl7:"42"`
	LivingWillCode                   string    `hl7:"43"`
	OrganDonorCode                   string    `hl7:"44"`
	AdvanceDirectiveCode             []CE      `hl7:"45"`
	PatientStatusEffectiveDate       time.Time `hl7:"46,shortdate"`
	ExpectedLOAReturnDateTime        time.Time `hl7:"47,longdate"`
}
