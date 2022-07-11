package hl7v24

// HL7 v2.4 - SSR_U04 - Specimen status request
// https://hl7-definition.caristix.com/v2/HL7v2.4/TriggerEvents/SSR_U04
type SSR_U04 struct {
	MSH                     MSH `hl7:"MSH"`
	EquipmentDetail         EQU `hl7:"EQU"`
	SpecimenContainerDetail SAC `hl7:"SAC"`
	Role                    ROL `hl7:"ROL,optional"`
}
