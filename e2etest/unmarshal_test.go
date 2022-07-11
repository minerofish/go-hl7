package e2e

import (
	"testing"

	"github.com/DRK-Blutspende-BaWueHe/go-hl7/hl7"
	"github.com/stretchr/testify/assert"
)

func TestMessageIdentification(t *testing.T) {
	data := "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3"

	messageType, protocolVersion, err := hl7.IdentifyMessage(
		[]byte(data),
		hl7.EncodingUTF8,
	)

	assert.Nil(t, err)
	assert.Equal(t, "ORM^O01", messageType)
	assert.Equal(t, "2.3", protocolVersion)
}

/*
// Test_Parse_MSH_Segment, this test has only one line
func Test_Parse_MSH_Segment(t *testing.T) {
	fileData := fmt.Sprintf("MSH|^~\\&|HL7_Host^b^c|HL7_Office^^Xyz|CIT^^|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1~second_element|<\u000d")

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(fileData),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)
	assert.NotNil(t, message.MSH)
	assert.Equal(t, "|", message.MSH.FieldSeparator)
	assert.Equal(t, "^~\\&", message.MSH.EncodingCharacters)
	assert.Equal(t, "HL7_Host", message.MSH.SendingApplication.NamespaceId)
	assert.Equal(t, "b", message.MSH.SendingApplication.UniversalId)
	assert.Equal(t, "c", message.MSH.SendingApplication.UniversalIdType)
	assert.Equal(t, "HL7_Office", message.MSH.SendingFacility.NamespaceId)
	assert.Equal(t, "CIT", message.MSH.ReceivingApplication.NamespaceId)
	assert.Equal(t, "LAB", message.MSH.ReceivingFacility.NamespaceId)
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
	assert.Equal(t, "8859/1", message.MSH.CharacterSet[0])
	assert.Equal(t, "second_element", message.MSH.CharacterSet[1])
	assert.Equal(t, "<", message.MSH.PrincipalLanguageOfMessage.Identifier)
}

// Run a Testorder provided by Roche cITM but its some standard with each record once
func Test_Order_ORM_generic1(t *testing.T) {
	var filedata string
	filedata = filedata + "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1|\u000d"
	filedata = filedata + "PID|1|a^b~^c|00100M56016||Smith^Harry||19500412|M\u000d"
	filedata = filedata + "ORC|NW|000218T018||||Not used|^^^^^R||20110926120055\u000d"
	filedata = filedata + "OBR|1|000218T018||101~102||20110926120000|||||A||||\u000d"

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(filedata),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)

	assert.NotNil(t, message.Patient)
	assert.NotNil(t, message.Patient.PatientIdentification)
	assert.Equal(t, 1, message.Patient.PatientIdentification.ID)
	assert.Equal(t, "a", message.Patient.PatientIdentification.ExternalID[0].Id)
	assert.Equal(t, "b", message.Patient.PatientIdentification.ExternalID[0].CheckDigit)
	assert.Equal(t, "", message.Patient.PatientIdentification.ExternalID[0].AssigningAuthority)
	assert.Equal(t, "", message.Patient.PatientIdentification.ExternalID[1].Id)
	assert.Equal(t, "c", message.Patient.PatientIdentification.ExternalID[1].CheckDigit)
	assert.Equal(t, "", message.Patient.PatientIdentification.ExternalID[1].AssigningAuthority)
	assert.Equal(t, "a", message.Patient.PatientIdentification.ExternalID[0].Id)
	assert.Equal(t, "b", message.Patient.PatientIdentification.ExternalID[0].CheckDigit)
	assert.Equal(t, "", message.Patient.PatientIdentification.ExternalID[1].Id)
	assert.Equal(t, "c", message.Patient.PatientIdentification.ExternalID[1].CheckDigit)
	assert.Equal(t, "00100M56016", message.Patient.PatientIdentification.InternalID[0].Id)
	assert.Equal(t, "", message.Patient.PatientIdentification.AlternateID[0].Id)
	assert.Equal(t, "Smith", message.Patient.PatientIdentification.Name.FamilyName)
	assert.Equal(t, "Harry", message.Patient.PatientIdentification.Name.GivenName)
	assert.Equal(t, "", message.Patient.PatientIdentification.Name.MiddleInitialOrName)
	assert.Equal(t, "", message.Patient.PatientIdentification.MothersMaidenName.FamilyName)
	assert.Equal(t, "1950-04-12 00:00:00 +0100 CET", message.Patient.PatientIdentification.DOB.String())
	assert.Equal(t, "M", message.Patient.PatientIdentification.Sex)

	assert.Equal(t, 1, len(message.Order))

	assert.Equal(t, "NW", message.Order[0].CommondOrderSegment.OrderControl)
	assert.Equal(t, "000218T018", message.Order[0].CommondOrderSegment.PlacerOrderNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.PlacerOrderNumber.NamespaceId)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.PlacerOrderNumber.UniversalId)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.PlacerOrderNumber.UniversalIdType)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.FillerOrderNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.PlacerGroupNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.OrderStatus)
	assert.Equal(t, "Not used", message.Order[0].CommondOrderSegment.ResponseFlag)
	assert.Equal(t, "R", message.Order[0].CommondOrderSegment.QuantityTiming.Priority)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.QuantityTiming.Duration)
	assert.Equal(t, "R", message.Order[0].CommondOrderSegment.QuantityTiming.Priority)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.QuantityTiming.Condition)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.ParentOrder.ParentsPlacerOrderNumber)
	assert.Equal(t, "", message.Order[0].CommondOrderSegment.ParentOrder.ParentsFillerOrderNumber)
	assert.Equal(t, "2011-09-26 10:00:55 +0000 UTC", message.Order[0].CommondOrderSegment.DateTimeOfTransaction.String())

	assert.Nil(t, err)
	assert.NotNil(t, message.Order[0].Detail.ObservationRequestSegment)
	assert.Equal(t, "1", message.Order[0].Detail.ObservationRequestSegment.ObservationRequest)
	assert.Equal(t, "000218T018", message.Order[0].Detail.ObservationRequestSegment.PlacerOrderNumber.EntityIdentifier)
	assert.Equal(t, "", message.Order[0].Detail.ObservationRequestSegment.PlacerOrderNumber.NamespaceId)
	assert.Equal(t, "", message.Order[0].Detail.ObservationRequestSegment.PlacerOrderNumber.UniversalId)
	assert.Equal(t, "", message.Order[0].Detail.ObservationRequestSegment.PlacerOrderNumber.UniversalIdType)
	assert.Equal(t, "", message.Order[0].Detail.ObservationRequestSegment.FillerOrderNumber.EntityIdentifier)
	assert.Equal(t, "101", message.Order[0].Detail.ObservationRequestSegment.UniversalServiceIdentifier.Identifier)
	assert.Equal(t, "", message.Order[0].Detail.ObservationRequestSegment.UniversalServiceIdentifier.Text)
	assert.Equal(t, "", message.Order[0].Detail.ObservationRequestSegment.Priority)
	assert.Equal(t, "2011-09-26 10:00:00 +0000 UTC", message.Order[0].Detail.ObservationRequestSegment.RequestedDateTime.String())
	assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", message.Order[0].Detail.ObservationRequestSegment.ObservationDateTime.String())
	assert.Equal(t, "A", message.Order[0].Detail.ObservationRequestSegment.SpecimenActionCode)

}
*/
