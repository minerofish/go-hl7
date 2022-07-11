package hl7v24

import "time"

// HL7 v2.4 - SAC - Specimen and container detail
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/SAC
type SAC struct {
	ExternalAccesionIDentifier        EI        `hl7:"1"`
	AccesionIdntifier                 EI        `hl7:"2"`
	ContainerIdentifier               EI        `hl7:"3"`
	Primary                           EI        `hl7:"4"`
	EquipmentContainerIDentifier      EI        `hl7:"5"`
	SpecimenSource                    SPS       `hl7:"6"`
	RegistrationDateTime              time.Time `hl7:"7,longdate"`
	ContainerStatus                   CE        `hl7:"8"`
	CarrierType                       CE        `hl7:"9"`
	CarrierIdentifier                 EI        `hl7:"10"`
	PositionInCarrier                 NA        `hl7:"11"`
	TrayTypeSAC                       CE        `hl7:"12"`
	TrayIDentifier                    EI        `hl7:"13"`
	PositionInTray                    NA        `hl7:"14"`
	Location                          []CE      `hl7:"15"`
	ContainerHeight                   float32   `hl7:"16"`
	ContainerDiameter                 float32   `hl7:"17"`
	BarrierDelta                      float32   `hl7:"18"`
	BottomDelta                       float32   `hl7:"19"`
	ContainerHeightDiameterDeltaUnits CE        `hl7:"20"`
	ContainerVolume                   float32   `hl7:"21"`
	AvailableVolume                   float32   `hl7:"22"`
	InitialSpecimenVolume             float32   `hl7:"23"`
	VolumeUnits                       CE        `hl7:"24"`
	SeparatorType                     CE        `hl7:"25"`
	CapType                           CE        `hl7:"26"`
	Additive                          []CE      `hl7:"27"`
	SpecimenComponent                 CE        `hl7:"28"`
	DilutionFactor                    SN        `hl7:"29"`
	Treatment                         CE        `hl7:"30"`
	Temperature                       SN        `hl7:"31"`
	HemolysisIndex                    float32   `hl7:"32"`
	HomolysisIndexUnits               CE        `hl7:"33"`
	LipemiaIndex                      float32   `hl7:"34"`
	LipemiaIndexUnits                 CE        `hl7:"35"`
	IcterusIndex                      float32   `hl7:"26"`
	IcterusIndexUnits                 CE        `hl7:"37"`
	FibrinIndex                       float32   `hl7:"38"`
	FibrinIndexUnits                  CE        `hl7:"39"`
	SystemInducedContaiminants        []CE      `hl7:"40"`
	DrugInterference                  []CE      `hl7:"41"`
	ArtificialBlood                   CE        `hl7:"42"`
	SpecialHandlingConsiderations     []CE      `hl7:"43"`
	OtherEnvironmentalFactors         []CE      `hl7:"44"`
}
