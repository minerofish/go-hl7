package hl7v24

import "time"

// HL7 v2.4 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN1
type IN1 struct {
	SetID                          string    `hl7:"1" json:"SetID"`
	PlanID                         CE        `hl7:"2" json:"PlanID"`
	CompanyID                      []CX      `hl7:"3" json:"CompanyID"`
	CompanyName                    []XON     `hl7:"4" json:"CompanyName"`
	CompanyAddress                 []XAD     `hl7:"5" json:"CompanyAddress"`
	ContactPerson                  []XPN     `hl7:"6" json:"ContactPerson"`
	PhoneNumber                    []XTN     `hl7:"7" json:"PhoneNumber"`
	GroupNumber                    string    `hl7:"8" json:"GroupNumber"`
	GroupName                      []XON     `hl7:"9" json:"GroupName"`
	GroupEmployerID                []CX      `hl7:"10" json:"GroupEmployerID"`
	GroupEmployerName              []XON     `hl7:"11" json:"GroupEmployerName"`
	PlanEffectiveDate              time.Time `hl7:"12,shortdate" json:"PlanEffectiveDate"`
	PlanExpirationDate             time.Time `hl7:"13,shortdate" json:"PlanExpirationDate"`
	AuthorizationInformation       AUI       `hl7:"14" json:"AuthorizationInformation"`
	PlanType                       string    `hl7:"15" json:"PlanType"`
	NameOfInsured                  []XPN     `hl7:"16" json:"NameOfInsured"`
	InsuredRelationshipToPatient   CE        `hl7:"17" json:"InsuredRelationshipToPatient"`
	InsuredDateOfBirth             time.Time `hl7:"18,longdate" json:"InsuredDateOfBirth"`
	InsuredAddress                 []XAD     `hl7:"19" json:"InsuredAddress"`
	AssignmentOfBenefits           string    `hl7:"20" json:"AssignmentOfBenefits"`
	CoordinationOfBenefits         string    `hl7:"21" json:"CoordinationOfBenefits"`
	CoordOfBenPriority             string    `hl7:"22" json:"CoordOfBenPriority"`
	NoticeOfAdmissionFlag          string    `hl7:"23" json:"NoticeOfAdmissionFlag"`
	NoticeOfAdmissionDate          time.Time `hl7:"24,shortdate" json:"NoticeOfAdmissionDate"`
	ReportOfEligibilityFlag        string    `hl7:"25" json:"ReportOfEligibilityFlag"`
	ReportOfEligibilityDate        time.Time `hl7:"26,shortdate" json:"ReportOfEligibilityDate"`
	ReleaseInformationCode         string    `hl7:"27" json:"ReleaseInformationCode"`
	PreAdmitCert                   string    `hl7:"28" json:"PreAdmitCert"`
	VerificationDateTime           time.Time `hl7:"29,longdate" json:"VerificationDateTime"`
	VerificationBy                 []XCN     `hl7:"30" json:"VerificationBy"`
	TypeOfAgreementCode            string    `hl7:"31" json:"TypeOfAgreementCode"`
	BillingStatus                  string    `hl7:"32" json:"BillingStatus"`
	LifetimeReserveDays            float32   `hl7:"33" json:"LifetimeReserveDays"`
	DelayBeforeLifetimeReserveDays float32   `hl7:"34" json:"DelayBeforeLifetimeReserveDays"`
	CompanyPlanCode                string    `hl7:"35" json:"CompanyPlanCode"`
	PolicyNumber                   string    `hl7:"36" json:"PolicyNumber"`
	PolicyDeductible               CP        `hl7:"37" json:"PolicyDeductible"`
	PolicyLimitAmount              CP        `hl7:"38" json:"PolicyLimitAmount"`
	PolicyLimitDays                float32   `hl7:"39" json:"PolicyLimitDays"`
	RoomRateSemiPrivate            CP        `hl7:"40" json:"RoomRateSemiPrivate"`
	RoomRatePrivate                CP        `hl7:"41" json:"RoomRatePrivate"`
	InsuredEmploymentStatus        CE        `hl7:"42" json:"InsuredEmploymentStatus"`
	InsuredSex                     string    `hl7:"43" json:"InsuredSex"`
	InsuredEmployerAddress         []XAD     `hl7:"44" json:"InsuredEmployerAddress"`
	VerificationStatus             string    `hl7:"45" json:"VerificationStatus"`
	PriorInsurancePlanID           string    `hl7:"46" json:"PriorInsurancePlanID"`
	CoverageType                   string    `hl7:"47" json:"CoverageType"`
	Handicap                       string    `hl7:"48" json:"Handicap"`
	InsuredIDNumber                []CX      `hl7:"49" json:"InsuredIDNumber"`
}
