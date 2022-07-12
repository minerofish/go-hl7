package hl7v24

import "time"

// HL7 v2.4 - IN3 - Insurance Additional Information, Certification
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/IN3
type IN3 struct {
	SetID                              string    `hl7:"1" json:"SetID"`
	CertificationNumber                CX        `hl7:"2" json:"CertificationNumber"`
	CertifiedBy                        []XCN     `hl7:"3" json:"CertifiedBy"`
	CertificationRequired              string    `hl7:"4" json:"CertificationRequired"`
	Penalty                            MOP       `hl7:"5" json:"Penalty"`
	CertificationDateTime              time.Time `hl7:"6,longdate" json:"CertificationDateTime"`
	CertificationModifyDateTime        time.Time `hl7:"7,longdate" json:"CertificationModifyDateTime"`
	Operator                           []XCN     `hl7:"8" json:"Operator"`
	CertificationBeginDate             time.Time `hl7:"9,shortdate" json:"CertificationBeginDate"`
	CertificationEndDate               time.Time `hl7:"10,shortdate" json:"CertificationEndDate"`
	Days                               DTN       `hl7:"11" json:"Days"`
	NonConcurCode                      CE        `hl7:"12" json:"NonConcurCode"`
	NonConcurEffectiveDateTime         time.Time `hl7:"13,longdate" json:"NonConcurEffectiveDateTime"`
	PhysicianReviewer                  []XCN     `hl7:"14" json:"PhysicianReviewer"`
	CertificationContact               string    `hl7:"15" json:"CertificationContact"`
	CertificationContactPhoneNumber    []XTN     `hl7:"16" json:"CertificationContactPhoneNumber"`
	AppealReason                       CE        `hl7:"17" json:"AppealReason"`
	CertificationAgency                CE        `hl7:"18" json:"CertificationAgency"`
	CertificationAgencyPhoneNumber     []XTN     `hl7:"19" json:"CertificationAgencyPhoneNumber"`
	PreCertificationRequired           []PCF     `hl7:"20" json:"PreCertificationRequired"`
	CaseManager                        string    `hl7:"21" json:"CaseManager"`
	SecondOpinionDate                  time.Time `hl7:"22,shortdate" json:"SecondOpinionDate"`
	SecondOpinionStatus                string    `hl7:"23" json:"SecondOpinionStatus"`
	SecondOpinionDocumentationReceived []string  `hl7:"24" json:"SecondOpinionDocumentationReceived"`
	SecondOptionPhysician              []XCN     `hl7:"25" json:"SecondOptionPhysician"`
}
