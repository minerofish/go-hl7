package hl7v24

import "time"

type NK1 struct {
	SetId                                    int       `hl7:"1"`
	Name                                     []XPN     `hl7:"2"`
	RelationShip                             CE        `hl7:"3"`
	Address                                  []XAD     `hl7:"4"`
	PhoneNumber                              []XTN     `hl7:"5"`
	BusinessPhoneNumber                      []XTN     `hl7:"6"`
	ContactRole                              CE        `hl7:"7"`
	StartDate                                time.Time `hl7:"8"`
	EndDate                                  time.Time `hl7:"9"`
	NextOfKinAssociatedPartiesJobTitle       string    `hl7:"10"`
	NextOfKinAssociatedPartiesJobCode        CE        `hl7:"11"`
	NextOfKinAssociatedPartiesEmployeeNumber CX        `hl7:"12"`
	OrganizationName                         []XON     `hl7:"13"`
	MaritalStatus                            CE        `hl7:"14"`
	AdministrativeSex                        string    `hl7:"15"`
	DOB                                      time.Time `hl7:"16"`
	LivingDependency                         []string  `hl7:"17"`
	AmbulatoryStatus                         []string  `hl7:"18"`
	Citzenship                               []CE      `hl7:"19"`
	PrimaryLanguage                          CE        `hl7:"20"`
	LivingArrangement                        string    `hl7:"21"`
	PublicityCode                            CE        `hl7:"22"`
	ProtectionIndeicator                     string    `hl7:"23"`
	StudentIndicator                         string    `hl7:"24"`
	Religion                                 CE        `hl7:"25"`
	MothersMaidenName                        []XPN     `hl7:"26"`
	Nationality                              CE        `hl7:"27"`
	EthnicGroup                              []CE      `hl7:"28"`
	ContactReason                            []CE      `hl7:"29"`
	ContactPersonsName                       []XPN     `hl7:"30"`
	ContactPersonsTelephoneNumber            []XTN     `hl7:"31"`
	NextOfKin                                []CX      `hl7:"32"`
	JobStatus                                string    `hl7:"34"`
	Race                                     []CE      `hl7:"35"`
	Handicap                                 string    `hl7:"36"`
	ContactPersonSocialSecurityNumber        string    `hl7:"37"`
}
