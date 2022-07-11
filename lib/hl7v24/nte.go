package hl7v24

// NTE - Notes and comments segment
// https://hl7-definition.caristix.com/v2/HL7v2.4/Segments/NTE
type NTE struct {
	CommentID       int      `hl7:"1"`
	SourceOfcomment string   `hl7:"2"`
	Comment         []string `hl7:"3"`
	CommentType     string   `hl7:"4"`
}
