package hl7v24

// HL7 v2.4 - OUL_R21 - Unsolicited laboratory observation
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/OUL_R21
type OUL_R21 struct {
	MSH              MSH   `hl7:"MSH" json:"MSH"`
	NotesAndComments []NTE `hl7:"NTE,optional" json:"NotesAndComments"`
	Patient          struct {
		PatientIdentification PID   `hl7:"PID,optional" json:"PatientIdentification"`
		PatientDemographics   PD1   `hl7:"PD1,optional" json:"PatientDemographics"`
		NotesAndComments      []NTE `hl7:"NTE,optional" json:"NotesAndComments"`
	}
	Visit struct {
		PatientVisit          PV1 `hl7:"PV1,optional" json:"PatientVisit"`
		AdditionalInformation PV2 `hl7:"PV2,optional" json:"AdditionalInformation"`
	}
	OrderObservation []struct {
		Container struct {
			SpecimenAndContainerDetail SAC   `hl7:"SAC,optional" json:"SpecimenAndContainerDetail"`
			SubstanceIdentifier        SID   `hl7:"SID,optional" json:"SubstanceIdentifier"`
			ObservationResult          []OBX `hl7:"OBX,optional" json:"ObservationResult"`
		}
		CommonOrder        ORC   `hl7:"ORC,optional" json:"CommonOrder"`
		ObservationRequest OBR   `hl7:"OBR,optional" json:"ObservationRequest"`
		NotesAndComments   []NTE `hl7:"NTE,optional" json:"NotesAndComments"`
		Observation        []struct {
			Observation         OBX   `hl7:"OBX,optional" json:"Observation"`
			TestCodeDetail      TCD   `hl7:"TCD,optional" json:"TestCodeDetail"`
			SubstanceIdentifier SID   `hl7:"SID,optional" json:"SubstanceIdentifier"`
			NotesAndComments    []NTE `hl7:"NTE,optional" json:"NotesAndComments"`
		}
		ClinicalTrialIdentification []CTI `hl7:"CTI,optional" json:"ClinicalTrialIdentification"`
	}
	ContinuationPointer DSC `hl7:"MSH" json:"ContinuationPointer"`
}
