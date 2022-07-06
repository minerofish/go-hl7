package e2e

import (
	"testing"

	"github.com/DRK-Blutspende-BaWueHe/go-hl7/hl7"
	"github.com/DRK-Blutspende-BaWueHe/go-hl7/lib/hl7v23"
	"github.com/stretchr/testify/assert"
)

func TestReadMinimalMessage(t *testing.T) {
	fileData := ""
	fileData = fileData + "MSH|^~\\&|HL7_Host|HL7_Office|CIT|LAB|20110926125155||ORM^O01|20110926125155|P|2.3|||ER|ER||8859/1|<\r"
	fileData = fileData + "PID|1||00100M56016||Smith^Harry||19500412|M\r"
	fileData = fileData + "ORC|NW|000218T018|||||^^^^^R||20110926120055\r"
	fileData = fileData + "OBR|1|000218T018||101~102||20110926120000|||||A||||\r"

	var message hl7v23.ORM_001
	err := hl7.Unmarshal([]byte(fileData), &message,
		hl7.EncodingUTF8, hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)

}
