package hl7v24

import "time"

// HL7 v2.4 - GT1 - Guarantor
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/GT1
type GT1 struct {
	SetId                         string    `hl7:"1" json:"SetId"`
	Number                        []CX      `hl7:"2" json:"Number"`
	Name                          []XPN     `hl7:"3" json:"Name"`
	SpouseName                    []XPN     `hl7:"4" json:"SpouseName"`
	Address                       []XAD     `hl7:"5" json:"Address"`
	PhoneNumberHome               []XTN     `hl7:"6" json:"PhoneNumberHome"`
	PhoneNumberBusiness           []XTN     `hl7:"7" json:"PhoneNumberBusiness"`
	DateTimeOfBirth               time.Time `hl7:"8,longdate" json:"DateTimeOfBirth"`
	Sex                           string    `hl7:"9" json:"Sex"`
	Type                          string    `hl7:"10" json:"Type"`
	Relationship                  CE        `hl7:"11" json:"Relationship"`
	SSN                           string    `hl7:"12" json:"SSN"`
	DateBegin                     time.Time `hl7:"13,shortdate" json:"DateBegin"`
	DateEnd                       time.Time `hl7:"14,shortdate" json:"DateEnd"`
	Priority                      float32   `hl7:"15" json:"Priority"`
	EmployerName                  []XPN     `hl7:"16" json:"EmployerName"`
	EmployerAddress               []XAD     `hl7:"17" json:"EmployerAddress"`
	EmployerPhoneNumber           []XTN     `hl7:"18" json:"EmployerPhoneNumber"`
	EmployeeIDNumber              []CX      `hl7:"19" json:"EmployeeIDNumber"`
	EmploymentStatus              string    `hl7:"20" json:"EmploymentStatus"`
	OrganizationName              []XON     `hl7:"21" json:"OrganizationName"`
	BillingHoldFlag               string    `hl7:"22" json:"BillingHoldFlag"`
	CreditRatingCode              CE        `hl7:"23" json:"CreditRatingCode"`
	DeathDateAndTime              time.Time `hl7:"24,longdate" json:"DeathDateAndTime"`
	DeathFlag                     string    `hl7:"25" json:"DeathFlag"`
	ChargeAdjustmentCode          CE        `hl7:"26" json:"ChargeAdjustmentCode"`
	HouseholdAnnualIncome         CP        `hl7:"27" json:"HouseholdAnnualIncome"`
	HouseholdSize                 float32   `hl7:"28" json:"HouseholdSize"`
	EmployerIDNumber              []CX      `hl7:"29" json:"EmployerIDNumber"`
	MaritalStatusCode             CE        `hl7:"30" json:"MaritalStatusCode"`
	HireEffectiveDate             time.Time `hl7:"31,shortdate" json:"HireEffectiveDate"`
	EmploymentStopDAte            time.Time `hl7:"32,shortdate" json:"EmploymentStopDAte"`
	LivingDependency              string    `hl7:"33" json:"LivingDependency"`
	AmbulatoryStatusCode          []string  `hl7:"34" json:"AmbulatoryStatusCode"`
	Citizenship                   []CE      `hl7:"35" json:"Citizenship"`
	PrimaryLanguage               CE        `hl7:"36" json:"PrimaryLanguage"`
	LivingArrangement             string    `hl7:"37" json:"LivingArrangement"`
	PublicityIndicator            CE        `hl7:"38" json:"PublicityIndicator"`
	ProtectionIndicator           string    `hl7:"39" json:"ProtectionIndicator"`
	StudentIndicator              string    `hl7:"40" json:"StudentIndicator"`
	Religion                      CE        `hl7:"41" json:"Religion"`
	MothersMaidenName             []XPN     `hl7:"42" json:"MothersMaidenName"`
	Nationality                   CE        `hl7:"43" json:"Nationality"`
	EthnicGroup                   []CE      `hl7:"44" json:"EthnicGroup"`
	ContactPersonsName            []XPN     `hl7:"45" json:"ContactPersonsName"`
	ContactPersonsTelephoneNumber []XTN     `hl7:"46" json:"ContactPersonsTelephoneNumber"`
	ContactReason                 CE        `hl7:"47" json:"ContactReason"`
	ContactRelationshipCode       string    `hl7:"48" json:"ContactRelationshipCode"`
	JobTitle                      string    `hl7:"49" json:"JobTitle"`
	JobCode                       JCC       `hl7:"50" json:"JobCode"`
	EmployersOrganizationName     []XON     `hl7:"51" json:"EmployersOrganizationName"`
	Handicap                      string    `hl7:"52" json:"Handicap"`
	JobStatus                     string    `hl7:"53" json:"JobStatus"`
	FinancialClass                FC        `hl7:"54" json:"FinancialClass"`
	GuarantorRace                 []CE      `hl7:"55" json:"GuarantorRace"`
}
