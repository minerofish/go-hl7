package hl7v24

// HL7 v2.4 - ORL_O22 - General laboratory order response message
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/ORL_O22
type ORL_O22 struct {
	MSH                    MSH   `hl7:"MSH"`
	MessageAcknowledgement MSA   `hl7:"MSA"`
	Error                  ERR   `hl7:"ERR,optional"`
	NotesAndComments       []NTE `hl7:"NTE,optional"`
	Response               struct {
		Patient struct {
			PatientIdentification PID `hl7:"PID,optional"`
		}
		GeneralOrder []struct {
			Container struct {
				SpecimenAndContainerDetail SAC   `hl7:"SAC,optional"`
				ObservationResult          []OBX `hl7:"OBX,optional"`
			}
			Order []struct {
				CommonOrder        ORC `hl7:"ORC,optional"`
				ObservationRequest struct {
					TheRequest                 OBR   `hl7:"OBR,optional"`
					SpecimenAndContainerDetail []SAC `hl7:"SAC,optional"`
				}
			}
		}
	}
}
