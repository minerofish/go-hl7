package hl7v23

// CTI - Clinical Trial Identification
// https://hl7-definition.caristix.com/v2/HL7v2.3/Segments/CTI
type CTI struct {
	SponsorStudyID          EI `hl7:"1"`
	StudyPhaseIdentifier    CE `hl7:"2"`
	StudyScheduledTimePoint CE `hl7:"3"`
}
