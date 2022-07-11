package hl7v24

// HL7 v2.4 - OUL_R21 - Unsolicited laboratory observation
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/OUL_R21
type OUL_R21 struct {
	MSH              MSH   `hl7:"MSH"`
	NotesAndComments []NTE `hl7:"NTE,optional"`
	Patient          struct {
		PatientIdentification PID   `hl7:"PID,optional"`
		PatientDemographics   PD1   `hl7:"PD1,optional"`
		NotesAndComments      []NTE `hl7:"NTE,optional"`
	}
	Visit struct {
		PatientVisit          PV1 `hl7:"PV1,optional"`
		AdditionalInformation PV2 `hl7:"PV2,optional"`
	}
	OrderObservation []struct {
		Container struct {
			SpecimenAndContainerDetail SAC   `hl7:"SAC,optional"`
			SubstanceIdentifier        SID   `hl7:"SID,optional"`
			ObservationResult          []OBX `hl7:"OBX,optional"`
		}
		CommonOrder        ORC   `hl7:"ORC,optional"`
		ObservationRequest OBR   `hl7:"OBR,optional"`
		NotesAndComments   []NTE `hl7:"NTE,optional"`
		Observation        []struct {
			Observation         OBX   `hl7:"OBX,optional"`
			TestCodeDetail      TCD   `hl7:"TCD,optional"`
			SubstanceIdentifier SID   `hl7:"SID,optional"`
			NotesAndComments    []NTE `hl7:"NTE,optional"`
		}
		ClinicalTrialIdentification []CTI `hl7:"CTI,optional"`
	}
	ContinuationPointer DSC `hl7:"MSH"`
}
