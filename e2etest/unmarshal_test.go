package e2e

import (
	"testing"

	"github.com/DRK-Blutspende-BaWueHe/go-hl7/hl7"
	"github.com/DRK-Blutspende-BaWueHe/go-hl7/lib/hl7v23"
	"github.com/stretchr/testify/assert"
)

func Test_Parse_MSH_Record(t *testing.T) {
	fileData := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1|<\r`

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(fileData),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)
	assert.NotNil(t, message.MSH)
	assert.Equal(t, "MSH", message.MSH.FieldSeparator)
	assert.Equal(t, "^~\\&", message.MSH.EncodingCharacters)
	//assert.Equal(t, "", message.MSH.SendingApplication)
	//assert.Equal(t, "", message.MSH.SendingFacility)
	//assert.Equal(t, "", message.MSH.ReceivingApplication)
	assert.Equal(t, "LAB", message.MSH.ReceivingFacility)
	assert.Equal(t, "2011-09-26 10:51:55 +0000 UTC", message.MSH.DateTimeOfMessage.String())
	assert.Equal(t, "", message.MSH.Security)
	assert.Equal(t, "ORM", message.MSH.MessageType)
	assert.Equal(t, "O01", message.MSH.MessageTriggerEvent)
	assert.Equal(t, "20110926125155", message.MSH.MessageControlID)
	assert.Equal(t, "P", message.MSH.ProccessingID)
	assert.Equal(t, "2.3", message.MSH.VersionID)
	assert.Equal(t, 0, message.MSH.SequenceNumber)
	assert.Equal(t, "", message.MSH.ContinuationPointer)
	assert.Equal(t, "ER", message.MSH.AcceptAcknowledgementType)
	assert.Equal(t, "ER", message.MSH.ApplicationAcknowledgementType)
	assert.Equal(t, "", message.MSH.CountryCode)
	//assert.Equal(t, "", message.MSH.CharacterSet)
	//assert.Equal(t, "8859/1", message.MSH.PrincipalLanguageOfMessage)
}

func Test_Parse_PID_Record(t *testing.T) {
	fileData := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1|<\r
	PID|1||00100M56016||Smith^Harry||19500412|M\r`

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(fileData),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)

	// TODO: add PID asserts here
	// assert.NotNil(t, message.Patient)
}

func Test_Parse_ORC_Record(t *testing.T) {
	fileData := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1|<\r
	PID|1||00100M56016||Smith^Harry||19500412|M\r
	ORC|NW|000218T018|||||^^^^^R||20110926120055\r`

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(fileData),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)

	// TODO: add ORC asserts here
}

func Test_Parse_OBR_Record(t *testing.T) {
	fileData := `MSH|^~\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1|<\r
	PID|1||00100M56016||Smith^Harry||19500412|M\r
	ORC|NW|000218T018|||||^^^^^R||20110926120055\r
	OBR|1|000218T018||101~102||20110926120000|||||A||||\r`

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(fileData),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)

	// TODO: add OBR asserts here
}
