package hl7v24

// HL7 v2.4 - CTD - Contact Data
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CTD
type CTD struct {
	ContactRole                     []CE  `hl7:"1"`
	ContactName                     []XPN `hl7:"2"`
	ContactAddress                  []XAD `hl7:"3"`
	ContactLocation                 PL    `hl7:"4"`
	ContactCommunicationInformation []XTN `hl7:"5"`
	PreferredMethodOfContact        CE    `hl7:"6"`
	ContactIdentifier               []PI  `hl7:"7"`
}
