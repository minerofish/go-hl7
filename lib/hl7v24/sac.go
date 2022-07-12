package hl7v24

import "time"

// HL7 v2.4 - SAC - Specimen and container detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SAC
type SAC struct {
	ExternalAccessionIdentifier       EI        `hl7:"1" json:"ExternalAccessionIdentifier"`
	AccessionIdentifier               EI        `hl7:"2" json:"AccessionIdentifier"`
	ContainerIdentifier               EI        `hl7:"3" json:"ContainerIdentifier"`
	Primary                           EI        `hl7:"4" json:"Primary"`
	EquipmentContainerIdentifier      EI        `hl7:"5" json:"EquipmentContainerIdentifier"`
	SpecimenSource                    SPS       `hl7:"6" json:"SpecimenSource"`
	RegistrationDateTime              time.Time `hl7:"7,longdate" json:"RegistrationDateTime"`
	ContainerStatus                   CE        `hl7:"8" json:"ContainerStatus"`
	CarrierType                       CE        `hl7:"9" json:"CarrierType"`
	CarrierIdentifier                 EI        `hl7:"10" json:"CarrierIdentifier"`
	PositionInCarrier                 NA        `hl7:"11" json:"PositionInCarrier"`
	TrayTypeSAC                       CE        `hl7:"12" json:"TrayTypeSAC"`
	TrayIdentifier                    EI        `hl7:"13" json:"TrayIdentifier"`
	PositionInTray                    NA        `hl7:"14" json:"PositionInTray"`
	Location                          []CE      `hl7:"15" json:"Location"`
	ContainerHeight                   float32   `hl7:"16" json:"ContainerHeight"`
	ContainerDiameter                 float32   `hl7:"17" json:"ContainerDiameter"`
	BarrierDelta                      float32   `hl7:"18" json:"BarrierDelta"`
	BottomDelta                       float32   `hl7:"19" json:"BottomDelta"`
	ContainerHeightDiameterDeltaUnits CE        `hl7:"20" json:"ContainerHeightDiameterDeltaUnits"`
	ContainerVolume                   float32   `hl7:"21" json:"ContainerVolume"`
	AvailableVolume                   float32   `hl7:"22" json:"AvailableVolume"`
	InitialSpecimenVolume             float32   `hl7:"23" json:"InitialSpecimenVolume"`
	VolumeUnits                       CE        `hl7:"24" json:"VolumeUnits"`
	SeparatorType                     CE        `hl7:"25" json:"SeparatorType"`
	CapType                           CE        `hl7:"26" json:"CapType"`
	Additive                          []CE      `hl7:"27" json:"Additive"`
	SpecimenComponent                 CE        `hl7:"28" json:"SpecimenComponent"`
	DilutionFactor                    SN        `hl7:"29" json:"DilutionFactor"`
	Treatment                         CE        `hl7:"30" json:"Treatment"`
	Temperature                       SN        `hl7:"31" json:"Temperature"`
	HemolysisIndex                    float32   `hl7:"32" json:"HemolysisIndex"`
	HomolysisIndexUnits               CE        `hl7:"33" json:"HomolysisIndexUnits"`
	LipemiaIndex                      float32   `hl7:"34" json:"LipemiaIndex"`
	LipemiaIndexUnits                 CE        `hl7:"35" json:"LipemiaIndexUnits"`
	IcterusIndex                      float32   `hl7:"26" json:"IcterusIndex"`
	IcterusIndexUnits                 CE        `hl7:"37" json:"IcterusIndexUnits"`
	FibrinIndex                       float32   `hl7:"38" json:"FibrinIndex"`
	FibrinIndexUnits                  CE        `hl7:"39" json:"FibrinIndexUnits"`
	SystemInducedContaminants         []CE      `hl7:"40" json:"SystemInducedContaminants"`
	DrugInterference                  []CE      `hl7:"41" json:"DrugInterference"`
	ArtificialBlood                   CE        `hl7:"42" json:"ArtificialBlood"`
	SpecialHandlingConsiderations     []CE      `hl7:"43" json:"SpecialHandlingConsiderations"`
	OtherEnvironmentalFactors         []CE      `hl7:"44" json:"OtherEnvironmentalFactors"`
}
