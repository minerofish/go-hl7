package hl7v24

// NTE - Notes and comments segment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NTE
type NTE struct {
	CommentID       int      `hl7:"1" json:"CommentID"`
	SourceOfComment string   `hl7:"2" json:"SourceOfComment"`
	Comment         []string `hl7:"3" json:"Comment"`
	CommentType     string   `hl7:"4" json:"CommentType"`
}
