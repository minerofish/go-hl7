package hl7v24

import "time"

// HL7 v2.4 - NK1 - Next of kin / associated parties
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NK1
type NK1 struct {
	SetID                                    int       `hl7:"1" json:"SetID"`
	Name                                     []XPN     `hl7:"2" json:"Name"`
	Relationship                             CE        `hl7:"3" json:"Relationship"`
	Address                                  []XAD     `hl7:"4" json:"Address"`
	PhoneNumber                              []XTN     `hl7:"5" json:"PhoneNumber"`
	BusinessPhoneNumber                      []XTN     `hl7:"6" json:"BusinessPhoneNumber"`
	ContactRole                              CE        `hl7:"7" json:"ContactRole"`
	StartDate                                time.Time `hl7:"8" json:"StartDate"`
	EndDate                                  time.Time `hl7:"9" json:"EndDate"`
	NextOfKinAssociatedPartiesJobTitle       string    `hl7:"10" json:"NextOfKinAssociatedPartiesJobTitle"`
	NextOfKinAssociatedPartiesJobCode        CE        `hl7:"11" json:"NextOfKinAssociatedPartiesJobCode"`
	NextOfKinAssociatedPartiesEmployeeNumber CX        `hl7:"12" json:"NextOfKinAssociatedPartiesEmployeeNumber"`
	OrganizationName                         []XON     `hl7:"13" json:"OrganizationName"`
	MaritalStatus                            CE        `hl7:"14" json:"MaritalStatus"`
	AdministrativeSex                        string    `hl7:"15" json:"AdministrativeSex"`
	DOB                                      time.Time `hl7:"16" json:"DOB"`
	LivingDependency                         []string  `hl7:"17" json:"LivingDependency"`
	AmbulatoryStatus                         []string  `hl7:"18" json:"AmbulatoryStatus"`
	Citzenship                               []CE      `hl7:"19" json:"Citzenship"`
	PrimaryLanguage                          CE        `hl7:"20" json:"PrimaryLanguage"`
	LivingArrangement                        string    `hl7:"21" json:"LivingArrangement"`
	PublicityCode                            CE        `hl7:"22" json:"PublicityCode"`
	ProtectionIndicator                      string    `hl7:"23" json:"ProtectionIndicator"`
	StudentIndicator                         string    `hl7:"24" json:"StudentIndicator"`
	Religion                                 CE        `hl7:"25" json:"Religion"`
	MothersMaidenName                        []XPN     `hl7:"26" json:"MothersMaidenName"`
	Nationality                              CE        `hl7:"27" json:"Nationality"`
	EthnicGroup                              []CE      `hl7:"28" json:"EthnicGroup"`
	ContactReason                            []CE      `hl7:"29" json:"ContactReason"`
	ContactPersonsName                       []XPN     `hl7:"30" json:"ContactPersonsName"`
	ContactPersonsTelephoneNumber            []XTN     `hl7:"31" json:"ContactPersonsTelephoneNumber"`
	NextOfKin                                []CX      `hl7:"32" json:"NextOfKin"`
	JobStatus                                string    `hl7:"34" json:"JobStatus"`
	Race                                     []CE      `hl7:"35" json:"Race"`
	Handicap                                 string    `hl7:"36" json:"Handicap"`
	ContactPersonSocialSecurityNumber        string    `hl7:"37" json:"ContactPersonSocialSecurityNumber"`
}
