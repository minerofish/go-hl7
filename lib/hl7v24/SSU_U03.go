package hl7v24

// HL7 v2.4 - SSU_U03 - Specimen status update
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/SSU_U03
type SSU_U03 struct {
	MSH               MSH `hl7:"MSH" json:"MSH"`
	EquipmentDetail   EQU `hl7:"EQU" json:"EquipmentDetail"`
	SpecimenContainer []struct {
		SpecimenContainerDetail SAC `hl7:"SAC" json:"SpecimenContainerDetail"`
		ObservationResult       OBX `hl7:"OBX" json:"ObservationResult"`
	}
	Role ROL `hl7:"ROL,optional" json:"Role"`
}
