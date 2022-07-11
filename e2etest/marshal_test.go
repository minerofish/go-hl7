package e2e

import (
	"fmt"
	"strings"
	"testing"

	"github.com/DRK-Blutspende-BaWueHe/go-hl7"
	"github.com/DRK-Blutspende-BaWueHe/go-hl7/lib/hl7v23"
	"github.com/stretchr/testify/assert"
)

func TestMarshalMSH(t *testing.T) {
	// Arrange
	encodingType := hl7.EncodingUTF8
	timezoneType := hl7.TimezoneEuropeBerlin
	messageString := "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|0||ER|ER||8859/1|\u000d"

	var err error
	var messageBytes hl7v23.ORM_001
	err = hl7.Unmarshal(
		[]byte(messageString),
		&messageBytes,
		encodingType,
		timezoneType)

	// Act
	var marshalledMessageBytes [][]byte
	marshalledMessageBytes, err = hl7.Marshal(
		messageBytes,
		hl7.StandardFieldSeparator,
		encodingType,
		timezoneType,
		hl7.StandardNotation)

	// Assert
	assert.Nil(t, err)
	assert.Nil(t, marshalledMessageBytes)
	assert.Equal(t, 1, len(marshalledMessageBytes))
	// TODO: fix missing empty section
	assert.Equal(t, messageString, string(marshalledMessageBytes[0]))
}

func TestMarshalORM(t *testing.T) {
	var message hl7v23.ORM_001

	message.MSH.SendingApplication.NamespaceId = "NamespaceId"
	message.MSH.SendingApplication.UniversalId = "UniversalId"
	message.MSH.SendingApplication.UniversalIdType = "UniversalIdType"

	bytes, err := hl7.Marshal(message, hl7.StandardFieldSeparator, hl7.EncodingASCII, hl7.TimezoneEuropeBerlin, hl7.StandardNotation)

	assert.Nil(t, err)
	for _, x := range bytes {
		fmt.Println(string(x))
	}
}

func TestMarshalAnUnmarshal_1(t *testing.T) {
	var filedata string
	filedata = filedata + "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|0||ER|ER||8859/1|\u000d"
	filedata = filedata + "PID|1|a^b~^c|00100M56016||Smith^Harry||19500412|M\u000d"
	//filedata = filedata + "ORC|NW|000218T018||||Not used|^^^^^R||20110926120055\u000d"
	//filedata = filedata + "OBR|1|000218T018||101~102||20110926120000|||||A||||\u000d"

	var message hl7v23.ORM_001
	err := hl7.Unmarshal(
		[]byte(filedata),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)
	fmt.Println("1 ", message.Patient.PatientIdentification.ExternalID[0].Id)
	fmt.Println("2 ", message.Patient.PatientIdentification.ExternalID[0].CheckDigit)

	bytes, err := hl7.Marshal(message, hl7.StandardFieldSeparator, hl7.EncodingASCII, hl7.TimezoneEuropeBerlin, hl7.StandardNotation)

	assert.Nil(t, err)
	ofile := strings.Split(filedata, "\u000d")
	for i, x := range bytes {
		if i < len(ofile) {
			fmt.Println(ofile[i])
		}
		fmt.Println(string(x))
		fmt.Println("-----")
	}

}
