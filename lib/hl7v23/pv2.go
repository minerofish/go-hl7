package hl7v23

import "time"

// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PV2
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
	ReferralSourceCode               XCN       `hl7:"13"`
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
	PatientChargeAdjustmentCode      string    `hl7:"30"`
	RecurringServicecoce             string    `hl7:"31"`
	BillingMediaCode                 string    `hl7:"32"`
	EpxectedSurgeryDateTime          time.Time `hl7:"33"`
	MilitaryPartnershipCode          string    `hl7:"34"`
	MilitaryNonAvailabilityCode      string    `hl7:"35"`
	NewbornBabyIndicator             string    `hl7:"36"`
	BabyDetainedIndicator            string    `hl7:"37"`
}
