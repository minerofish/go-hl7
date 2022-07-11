package hl7v23

import "time"

// HL7 v2.3 - IN1 - Insurance
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/IN1
type IN1 struct {
	SetID                         string    `hl7:"1"`
	PlanID                        CE        `hl7:"2"`
	CompanyID                     []CX      `hl7:"3"`
	CompanyName                   []XON     `hl7:"4"`
	CompanyAddress                []XAD     `hl7:"5"`
	ContaactPerson                []XPN     `hl7:"6"`
	PhoneNumber                   []XTN     `hl7:"7"`
	GroupNumber                   string    `hl7:"8"`
	GroupName                     []XON     `hl7:"9"`
	GroupEmployerID               []CX      `hl7:"10"`
	GroupEmployerName             []XON     `hl7:"11"`
	PlanEffectiveDate             time.Time `hl7:"12,shortdate"`
	PlanExpirationDate            time.Time `hl7:"13,shortdate"`
	AutorizationInformation       CM_AUI    `hl7:"14"`
	PlanType                      string    `hl7:"15"`
	NameOfInsured                 []XPN     `hl7:"16"`
	InsuredRelationshipToPatient  string    `hl7:"17"`
	InsuredDateOfBirth            time.Time `hl7:"18,shortdate"`
	InsuredAddress                []XAD     `hl7:"19"`
	AsiggnmentOfBenefits          string    `hl7:"20"`
	CooridnationOfBenefits        string    `hl7:"21"`
	CoordiantionOfBenefitsPrio    string    `hl7:"22"`
	NoticeOfAdmissionCode         string    `hl7:"23"`
	NoticeOfAdmissionDate         time.Time `hl7:"24,shortdate"`
	ReportOfEigibilityCode        string    `hl7:"25"`
	ReportOfEigibilityDate        time.Time `hl7:"26,shortdate"`
	ReleaseInformationCode        string    `hl7:"27"`
	PreAdmitCert                  string    `hl7:"28"`
	VerificationDateTime          time.Time `hl7:"29,longdate"`
	VerificationBy                XCN       `hl7:"30"`
	TypeOfAgreementCode           string    `hl7:"31"`
	BillingStatus                 string    `hl7:"32"`
	LifetimeReserveDays           float32   `hl7:"33"`
	DelayBeforeLifetimeReerveDays float32   `hl7:"34"`
	CompanyPlanCode               string    `hl7:"35"`
	PolicyNumber                  string    `hl7:"36"`
	PolicyDeductible              CP        `hl7:"37"`
	PolicyLimitAmount             CP        `hl7:"38"`
	PolicyLimitDays               float32   `hl7:"39"`
	RoomRateSemiPrivate           CP        `hl7:"40"`
	RommRatePrivate               CP        `hl7:"41"`
	InsuredEmploymentStatus       CE        `hl7:"42"`
	InsuredSex                    string    `hl7:"43"`
	InsuredEmployerAddress        []XAD     `hl7:"44"`
	VerificationStatus            string    `hl7:"45"`
	PriorInsurancePlanID          string    `hl7:"46"`
	CoverageType                  string    `hl7:"47"`
	Handicap                      string    `hl7:"48"`
	InsuredIDNumber               []CX      `hl7:"49"`
}
