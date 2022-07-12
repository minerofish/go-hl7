package hl7v24

// HL7 v2.4 - CTD - Contact Data
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/CTD
type CTD struct {
	ContactRole                     []CE  `hl7:"1" json:"ContactRole"`
	ContactName                     []XPN `hl7:"2" json:"Contactname"`
	ContactAddress                  []XAD `hl7:"3" json:"ContactAddress"`
	ContactLocation                 PL    `hl7:"4" json:"ContactLocation"`
	ContactCommunicationInformation []XTN `hl7:"5" json:"contactCommunicationInformation"`
	PreferredMethodOfContact        CE    `hl7:"6" json:"PreferredMethodOfContact"`
	ContactIdentifier               []PI  `hl7:"7" json:"ContactIdentifier"`
}
