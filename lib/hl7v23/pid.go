package hl7v23

// PID - Patient Identification
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/PID
type PID struct {
	ID         int  `hl7:"1"`
	ExternalID []CX `hl7:"2"`
	/*
		ExternalID              []CX23    `hl7:"2"`
		InternalID              []CX23    `hl7:"3"`
		AlternateID             []CX23    `hl7:"4"`
		Name                    XPN23     `hl7:"5"`
		MothersMaidenName       XPN23     `hl7:"6"`
		DOB                     time.Time `hl7:"7"`
		Sex                     string    `hl7:"8"`
		Alias                   []XPN23   `hl7:"9"`
		Race                    string    `hl7:"10"`
		Address                 []XAD23   `hl7:"11"`
		CountryCode             string    `hl7:"12"`
		PhoneNumber             []XTN23   `hl7:"13"`
		PhoneNumberBusiness     []XTN23   `hl7:"14"`
		PrimaryLanguage         CE23      `hl7:"15"`
		MaritalStatus           string    `hl7:"16"`
		Religion                string    `hl7:"17"`
		PatientAccountNumber    CX23      `hl7:"18"`
		SSNNumber               string    `hl7:"19"`
		DriversLicenseNumber    DLN23     `hl7:"20"`
		MothersIdentifier       CX23      `hl7:"21"`
		EthnicGroup             string    `hl7:"22"`
		BirthPlace              string    `hl7:"23"`
		MultipleBirthIndicator  string    `hl7:"24"`
		BirthOrder              int       `hl7:"25"`
		Citizenship             []string  `hl7:"26"`
		VeteransMilitaryStatus  CE23      `hl7:"27"`
		NationalityCode         CE23      `hl7:"28"`
		PatientDeathDateAndTime time.Time `hl7:"29"`
		PatientDeathIndicator   string    `hl7:"30"`
	*/
}
