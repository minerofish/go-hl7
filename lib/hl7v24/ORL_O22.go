package hl7v24

// HL7 v2.4 - ORL_O22 - General laboratory order response message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORL_O22
type ORL_O22 struct {
	MSH                    MSH   `hl7:"MSH" json:"MSH"`
	MessageAcknowledgement MSA   `hl7:"MSA" json:"MessageAcknowledgement"`
	Error                  ERR   `hl7:"ERR,optional" json:"Error"`
	NotesAndComments       []NTE `hl7:"NTE,optional" json:"NotesAndComments"`
	Response               struct {
		Patient struct {
			PatientIdentification PID `hl7:"PID,optional" json:"PatientIdentification"`
		}
		GeneralOrder []struct {
			Container struct {
				SpecimenAndContainerDetail SAC   `hl7:"SAC,optional" json:"SpecimenAndContainerDetail"`
				ObservationResult          []OBX `hl7:"OBX,optional" json:"ObservationResult"`
			}
			Order []struct {
				CommonOrder        ORC `hl7:"ORC,optional" json:"CommonOrder"`
				ObservationRequest struct {
					TheRequest                 OBR   `hl7:"OBR,optional" json:"TheRequest"`
					SpecimenAndContainerDetail []SAC `hl7:"SAC,optional" json:"SpecimenAndContainerDetail"`
				}
			}
		}
	}
}
