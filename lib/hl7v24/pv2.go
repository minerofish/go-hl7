package hl7v24

import "time"

// HL7 v2.4 - PV2 - Patient visit - additional information
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/PV2
type PV2 struct {
	PriorPendingLocation              PL        `hl7:"1" json:"PriorPendingLocation"`
	AccommodationCode                 CE        `hl7:"2" json:"AccommodationCode"`
	AdmitReason                       CE        `hl7:"3" json:"AdmitReason"`
	TransferReason                    CE        `hl7:"4" json:"TransferReason"`
	PatientValuables                  string    `hl7:"5" json:"PatientValuables"`
	PatientValuablesLocation          string    `hl7:"6" json:"PatientValuablesLocation"`
	VisitUserCode                     string    `hl7:"7" json:"VisitUserCode"`
	ExpectedAdmitDate                 time.Time `hl7:"8" json:"ExpectedAdmitDate"`
	ExpectedDischargeDate             time.Time `hl7:"9" json:"ExpectedDischargeDate"`
	EstimatedLengthOfInpatientStay    float32   `hl7:"10" json:"EstimatedLengthOfInpatientStay"`
	ActualLengthOfImpatientStay       float32   `hl7:"11" json:"ActualLengthOfImpatientStay"`
	VisitDescription                  string    `hl7:"12" json:"VisitDescription"`
	ReferralSourceCode                []XCN     `hl7:"13" json:"ReferralSourceCode"`
	PreviousServiceDate               time.Time `hl7:"14" json:"PreviousServiceDate"`
	EmploymentIllnessRelatedIndicator string    `hl7:"15" json:"EmploymentIllnessRelatedIndicator"`
	PurgeStatusCode                   string    `hl7:"16" json:"PurgeStatusCode"`
	PurgeStatusDate                   time.Time `hl7:"17" json:"PurgeStatusDate"`
	SpecialProgramCode                string    `hl7:"18" json:"SpecialProgramCode"`
	RetentionIndicator                string    `hl7:"19" json:"RetentionIndicator"`
	ExpectedNumberOfInsurancePlans    float32   `hl7:"20" json:"ExpectedNumberOfInsurancePlans"`
	VisitPublicityCode                string    `hl7:"21" json:"VisitPublicityCode"`
	VisitProtectionIndicator          string    `hl7:"22" json:"VisitProtectionIndicator"`
	ClinicOrganizationName            XON       `hl7:"23" json:"ClinicOrganizationName"`
	PatientStatusCode                 string    `hl7:"24" json:"PatientStatusCode"`
	VisitPriorityCode                 string    `hl7:"25" json:"VisitPriorityCode"`
	PreviousTreatmentCode             string    `hl7:"26" json:"PreviousTreatmentCode"`
	ExpectedDischargeDisposition      string    `hl7:"27" json:"ExpectedDischargeDisposition"`
	SignatureOnFileDate               time.Time `hl7:"28" json:"SignatureOnFileDate"`
	FirstSimilarIllnessDate           time.Time `hl7:"29" json:"FirstSimilarIllnessDate"`
	PatientChargeAdjustmentCode       CE        `hl7:"30" json:"PatientChargeAdjustmentCode"`
	RecurringServiceCode              string    `hl7:"31" json:"RecurringServiceCode"`
	BillingMediaCode                  string    `hl7:"32" json:"BillingMediaCode"`
	ExpectedSurgeryDateTime           time.Time `hl7:"33" json:"ExpectedSurgeryDateTime"`
	MilitaryPartnershipCode           string    `hl7:"34" json:"MilitaryPartnershipCode"`
	MilitaryNonAvailabilityCode       string    `hl7:"35" json:"MilitaryNonAvailabilityCode"`
	NewbornBabyIndicator              string    `hl7:"36" json:"NewbornBabyIndicator"`
	BabyDetainedIndicator             string    `hl7:"37" json:"BabyDetainedIndicator"`
	ModeOfArrivalCode                 CE        `hl7:"38" json:"ModeOfArrivalCode"`
	RecreationalDrugUseCode           []CE      `hl7:"39" json:"RecreationalDrugUseCode"`
	AdmissionLevelOfCareCode          CE        `hl7:"40" json:"AdmissionLevelOfCareCode"`
	PrecautionCode                    []CE      `hl7:"41" json:"PrecautionCode"`
	PatientConditionCode              CE        `hl7:"42" json:"PatientConditionCode"`
	LivingWillCode                    string    `hl7:"43" json:"LivingWillCode"`
	OrganDonorCode                    string    `hl7:"44" json:"OrganDonorCode"`
	AdvanceDirectiveCode              []CE      `hl7:"45" json:"AdvanceDirectiveCode"`
	PatientStatusEffectiveDate        time.Time `hl7:"46,shortdate" json:"PatientStatusEffectiveDate"`
	ExpectedLOAReturnDateTime         time.Time `hl7:"47,longdate" json:"ExpectedLOAReturnDateTime"`
}
