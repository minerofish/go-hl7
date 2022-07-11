package hl7v23

import "time"

type GT1 struct {
	SetId                         string    `hl7:"1"`
	Number                        CX        `hl7:"2"`
	Name                          []XPN     `hl7:"3"`
	SpouseName                    []XPN     `hl7:"4"`
	Address                       []XAD     `hl7:"5"`
	PhoneNumberHome               []XTN     `hl7:"6"`
	PhoneNumberBusiness           []XTN     `hl7:"7"`
	DateTimeOfBirth               time.Time `hl7:"8,shortdate"`
	Sex                           string    `hl7:"9"`
	Type                          string    `hl7:"10"`
	Relationship                  string    `hl7:"11"`
	SSN                           string    `hl7:"12"`
	DateBegin                     time.Time `hl7:"13,shortdate"`
	DateEnd                       time.Time `hl7:"14,shortdate"`
	Priority                      float32   `hl7:"15"`
	Employer                      []XPN     `hl7:"16"`
	EmployerAddress               []XAD     `hl7:"17"`
	EmlpoyerPhoneNumber           []XTN     `hl7:"18"`
	EmployeeID                    CX        `hl7:"19"`
	EmploymentStatus              string    `hl7:"20"`
	Organization                  []XON     `hl7:"21"`
	BillingHoldFlag               string    `hl7:"22"`
	CreditRatingCode              CE        `hl7:"23"`
	DeathDateAndTime              time.Time `hl7:"24,longdate"`
	DeathFlag                     string    `hl7:"25"`
	ChargeAdjustmentCode          CE        `hl7:"26"`
	HouseholdAnnualIncum          CP        `hl7:"27"`
	HouseholdSize                 float32   `hl7:"28"`
	EmployerIDNumber              []CX      `hl7:"29"`
	MaritalStatusCode             string    `hl7:"30"`
	HireEffectiveDate             time.Time `hl7:"31,shortdate"`
	EmploymentStopDAte            time.Time `hl7:"32,shortdate"`
	LivingDepndency               string    `hl7:"33"`
	AmbulatoryStatusCode          string    `hl7:"34"`
	Citizenship                   string    `hl7:"35"`
	PrimaryLanguage               CE        `hl7:"36"`
	LivingArrangement             string    `hl7:"37"`
	PublicityIndicator            CE        `hl7:"38"`
	ProtectionIndicator           string    `hl7:"39"`
	StudentIndicator              string    `hl7:"40"`
	Religion                      string    `hl7:"41"`
	MothersMaidenName             XPN       `hl7:"42"`
	NationalityCode               CE        `hl7:"43"`
	EthnicGroup                   string    `hl7:"44"`
	ContactPersonsName            []XPN     `hl7:"45"`
	ContactPersonsTelephoneNumber []XTN     `hl7:"46"`
	ContactReason                 CE        `hl7:"47"`
	ContactRelationshipCode       string    `hl7:"48"`
	JobTitle                      string    `hl7:"49"`
	JobCode                       JCC       `hl7:"50"`
	EmployersOrganizationName     []XON     `hl7:"51"`
	Handicap                      string    `hl7:"52"`
	JobStatus                     string    `hl7:"53"`
	FinancialClass                FC        `hl7:"54"`
	GuarantorRace                 string    `hl7:"55"`
}
