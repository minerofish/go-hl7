package hl7v23

import "time"

type IN3 struct {
	SetID                              string    `hl7:"1"`
	CertificationNumber                CX        `hl7:"2"`
	CertifiedBy                        []XCN     `hl7:"3"`
	CertificationRequired              string    `hl7:"4"`
	Penalty                            CM_PEN    `hl7:"5"`
	CertifcationDateTime               time.Time `hl7:"6,longdate"`
	CertificationModifyDateTime        time.Time `hl7:"7,longdate"`
	Operator                           []XCN     `hl7:"8"`
	CertificationBeginDate             time.Time `hl7:"9,shortdate"`
	CertificationEndDate               time.Time `hl7:"10,shortdate"`
	Days                               CM_DTN    `hl7:"11"`
	NonconcurCode                      CE        `hl7:"12"`
	NonConcurEffectiveDateTime         time.Time `hl7:"13,longdate"`
	PhysicianReviewer                  []XCN     `hl7:"14"`
	CertificationContact               string    `hl7:"15"`
	CertificationContactPhoneNumber    []XTN     `hl7:"16"`
	AppealReason                       CE        `hl7:"17"`
	CertificationAgency                CE        `hl7:"18"`
	CertificationAgencyPhoneNumber     []XTN     `hl7:"19"`
	PreCertificationRequired           []CM_PCF  `hl7:"20"`
	CaseManager                        string    `hl7:"21"`
	SecondOpinionDate                  time.Time `hl7:"22,longdate"`
	SecondOpinionStatus                string    `hl7:"23"`
	SecondOpinionDocumentationReceived []string  `hl7:"24"`
	SecondOptionPhysican               []XCN     `hl7:"25"`
}
