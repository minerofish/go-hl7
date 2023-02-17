// Standard implementation for the lis2a2 protocol according to standard in every detail,
// will work for most with some tinkering....
package hl7v23

type Observation struct {
	Observation                 OBX   `hl7:"OBX,optional"`
	ObservationNotesAndComments []NTE `hl7:"NTE,optional"`
}

type Detail struct {
	ObservationRequestSegment OBR `hl7:"OBR,optional"`
	RequisitionDetail         RQD `hl7:"RQD,optional"`
	RequisitionDetail1        RQ1 `hl7:"RQ1,optional"`
	PharmacyPrescription      RXO `hl7:"RQ1,optional"`
	DietaryOrders             ODS `hl7:"ODS,optional"`
	DietTrayInstructions      ODT `hl7:"ODT,optional"`

	NotesAndComments []NTE `hl7:"NTE,optional"`
	Diagnosis        DG1   `hl7:"DG1,optional"`

	Observation []Observation
}

type Order struct {
	CommonOrderSegment ORC `hl7:"ORC"`
	Detail             Detail

	ClinicalTrialIdentification []CTI `hl7:"CTI,optional"`
	Billing                     BLG   `hl7:"BLG,optional"`
}

type Insurance struct {
	Insurance             IN1 `hl7:"IN1,optional"`
	AdditionalInformation IN2 `hl7:"IN2,optional"`
	Certification         IN3 `hl7:"IN3,optional"`
}

type PatientVisit struct {
	PatientVisit          PV1 `hl7:"PV1,optional"`
	AdditionalInformation PV2 `hl7:"PV2,optional"`
}

type Patient struct {
	PatientIdentification PID   `hl7:"PID,optional"`
	PatientDemographics   PD1   `hl7:"PD1,optional"`
	NotesAndComments      []NTE `hl7:"NTE,optional"`
}

// ORM_O01 - Order message
// https://hl7-definition.caristix.com/v2/HL7v2.3/TriggerEvents/ORM_O01
type ORM_O01 struct {
	MSH                       MSH `hl7:"MSH"`
	NotesAndComments          NTE `hl7:"NTE,optional"`
	Patient                   Patient
	PatientVisit              PatientVisit
	Insurance                 []Insurance
	Guarantor                 GT1 `hl7:"GT1,optional"`
	PatientAllergyInformation AL1 `hl7:"AL1,optional"`
	Order                     []Order
}
