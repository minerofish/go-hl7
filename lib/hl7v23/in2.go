package hl7v23

import "time"

// IN2 - Insurance additional info
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/IN2
type IN2 struct {
	InsuredEpmloyeeID                    []CX        `hl7:"1"`
	InsuredSSN                           string      `hl7:"2"`
	InsuredEmployerName                  []XCN       `hl7:"3"`
	EmployerInformationData              string      `hl7:"4"`
	MailClaimParty                       []string    `hl7:"5"`
	MediareHealthInsCardNumber           string      `hl7:"6"`
	MedicalCaseName                      []XPN       `hl7:"7"`
	MedicalCaseNumber                    string      `hl7:"8"`
	ChampusSponsorName                   []XPN       `hl7:"9"`
	ChampusID                            string      `hl7:"10"`
	DependentOfChampusRecipient          CE          `hl7:"11"`
	ChampusOrganization                  string      `hl7:"12"`
	ChampusStation                       string      `hl7:"13"`
	ChampusService                       string      `hl7:"14"`
	ChampusRank                          string      `hl7:"15"`
	ChampusStatus                        string      `hl7:"16"`
	ChampusRetireDate                    time.Time   `hl7:"17,shortdate"`
	ChampusNonAvailCertOnFile            string      `hl7:"18"`
	BabyCoverage                         string      `hl7:"19"`
	CombineBabyBill                      string      `hl7:"20"`
	BloodDeductible                      string      `hl7:"21"`
	SpecialCoverageApprovalName          []XPN       `hl7:"22"`
	SpecialCoverageApprovalTitle         string      `hl7:"23"`
	NonCoveredInsuranceCode              []string    `hl7:"24"`
	PayorID                              []CX        `hl7:"25"`
	PayorSubscriberID                    []CX        `hl7:"26"`
	EligibilitySource                    string      `hl7:"27"`
	RoomCoverageType                     []CM_RMC    `hl7:"28"`
	PolicyTypeAmount                     []CM_PTA    `hl7:"29"`
	DailyDeductible                      CM_DDI      `hl7:"30"`
	LivingDependency                     string      `hl7:"31"`
	AmbulatoryStatus                     string      `hl7:"32"`
	Citizenship                          string      `hl7:"33"`
	PrimaryLanguage                      CE          `hl7:"34"`
	LivingArrangement                    string      `hl7:"35"`
	PublicityIndicator                   CE          `hl7:"36"`
	ProtectionIndicator                  string      `hl7:"37"`
	StudentIndicator                     string      `hl7:"38"`
	Religion                             string      `hl7:"39"`
	MothersMaidenName                    XPN         `hl7:"40"`
	NationalityCode                      CE          `hl7:"41"`
	EthnicGroup                          string      `hl7:"42"`
	MaritalStatus                        []string    `hl7:"43"`
	EmploymentStartDate                  time.Time   `hl7:"44,shortdate"`
	EmploymentStopDate                   time.Time   `hl7:"45,shortdate"`
	JobTitle                             string      `hl7:"46"`
	JobCode                              JCC         `hl7:"47"`
	JobStatus                            string      `hl7:"48"`
	EmployerContactpersonName            []XPN       `hl7:"49"`
	EmployerContactPersonPhone           []XTN       `hl7:"50"`
	EmployerContactReason                string      `hl7:"51"`
	InsuredContactPersonName             []XPN       `hl7:"52"`
	InsuredContactPersonTelephone        []XTN       `hl7:"53"`
	InsuredContactPersonReason           []string    `hl7:"54"`
	RelationshipToThePatientStartDate    time.Time   `hl7:"55,shortdate"`
	RelationshipToThePatientStopDate     []time.Time `hl7:"56,shortdate"`
	InsuranceCompanyContactReason        string      `hl7:"57"`
	InsuranceCompanyContactPhone         XTN         `hl7:"58"`
	PolicyScope                          string      `hl7:"59"`
	PolicySource                         string      `hl7:"60"`
	PatientMemberNumber                  CX          `hl7:"61"`
	GuarantorRelationshipToInsured       string      `hl7:"62"`
	InsuredsTelephoneNumberHome          []XTN       `hl7:"63"`
	InsuredEmployerTelephoneNumber       []XTN       `hl7:"64"`
	MilitaryHandicappedProgram           CE          `hl7:"65"`
	SuspendFalg                          string      `hl7:"66"`
	CopayLimitFlag                       string      `hl7:"67"`
	StoplossLimitFlag                    string      `hl7:"68"`
	InsuredOrgannizationNameAndID        []XON       `hl7:"69"`
	InsuredEmployerOrganizationNameAndID []XON       `hl7:"70"`
	Race                                 string      `hl7:"71"`
	PatientRelationshipToInsured         CE          `hl7:"72"`
}
