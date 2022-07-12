package hl7v24

// HL7 v2.4 - ORU_R01 - Unsolicited transmission of an observation message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORU_R01
type ORU_R01 struct {
	MSH MSH `hl7:"MSH" json:"MSH"`

	PatientResult []struct {
		Patient struct {
			PatientIdentification PID   `hl7:"PID,optional" json:"PatientIdentification"`
			PatientDemographics   PD1   `hl7:"PD1,optional" json:"PatientDemographics"`
			NextOfKin             NK1   `hl7:"PD1,optional" json:"NextOfKin"`
			NotesAndComments      []NTE `hl7:"NTE,optional" json:"NotesAndComments"`
			Visit                 struct {
				PatientVisit          PV1 `hl7:"PV1,optional" json:"PatientVisit"`
				AdditionalInformation PV2 `hl7:"PV2,optional" json:"AdditionalInformation"`
			}
		}

		OrderObservation []struct {
			CommonOrder        ORC   `hl7:"ORC,optional" json:"CommonOrder"`
			ObservationRequest OBR   `hl7:"OBR,optional" json:"ObservationRequest"`
			NotesAndComments   []NTE `hl7:"NTE,optional" json:"NotesAndComments"`
			ContactData        []CTD `hl7:"CTD,optional" json:"ContactData"`

			Observation []struct {
				Observation                 OBX   `hl7:"OBX,optional" json:"Observation"`
				ObservationNotesAndComments []NTE `hl7:"NTE,optional" json:"ObservationNotesAndComments"`
			}
			FinancialTransaction        []FT1 `hl7:"FT1,optional" json:"FinancialTransaction"`
			ClinicalTrialIdentification []CTI `hl7:"CTI,optional" json:"ClinicalTrialIdentification"`
		}
	}
	ContinuationPointer DSC `hl7:"MSH" json:"ContinuationPointer"`
}
