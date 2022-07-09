package e2e

import (
	"testing"

	"github.com/DRK-Blutspende-BaWueHe/go-hl7"
	"github.com/DRK-Blutspende-BaWueHe/go-hl7/lib/hl7v23"
	"github.com/stretchr/testify/assert"
)

// This test will fail, because Roche is unable to read documents... OBR can not be repeated
func Test_cITm_Result1(t *testing.T) {
	var filedata string
	filedata = filedata + "MSH|^~\\&|||||20110927155013||ORU^R01|68516|P|2.3|||NE|NE||8859/1\u000d"
	filedata = filedata + "PID|1||4637463G66||Smith^John||19630101|M\u000d"
	filedata = filedata + "NTE||L|1st·comment·on·patient·/·sample·20020604101\u000d"
	filedata = filedata + "NTE||L|2nd·comment·on·patient·/·sample\u000d"
	filedata = filedata + "ORC|RE|20020604101|||||^^^^^R||20110927150630\u000d"
	filedata = filedata + "OBR|1|20020604101||ALL||20110927150629|||||||||S1^^^^^^P||||||||||||^^^^^|||||||\u000d"
	filedata = filedata + "OBX|1||21||101,0|mmol/L||N|||F||23~N|20110927154723|^^^COBAS8K.200|System\u000d"
	//filedata = filedata + "TCD|1|21|Dec\u000d"
	filedata = filedata + "NTE|||L|R|G\u000d"
	filedata = filedata + "OBX|2||82||5,7|mmol/L||H|||F||23~N|20110927154733|^^^COBAS8K.200|System\u000d"
	//filedata = filedata + "TCD|1|82|1\u000d"
	filedata = filedata + "NTE|||L|R|G\u000d"
	filedata = filedata + "OBX|3||89||162,0|mmol/L||H|||F||23~N|20110927154734|^^^COBAS8K.200|System\u000d"
	//filedata = filedata + "TCD|1|89|Inc\u000d"
	filedata = filedata + "NTE||L|Comment·on·test·code·89\u000d"
	filedata = filedata + "NTE|||L|R|G\u000d"
	var message hl7v23.ORU_R01
	err := hl7.Unmarshal(
		[]byte(filedata),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(message.PatientResult))

	assert.Equal(t, "2nd·comment·on·patient·/·sample", message.PatientResult[0].Patient.NotesAndComments[1].Comment)

	assert.Equal(t, 1, len(message.PatientResult[0].OrderObservation))
	assert.Equal(t, "RE", message.PatientResult[0].OrderObservation[0].CommonOrder.OrderControl)
	assert.Equal(t, "ALL", message.PatientResult[0].OrderObservation[0].ObservationRequest.UniversalServiceIdentifier.Identifier)

	assert.Equal(t, 3, len(message.PatientResult[0].OrderObservation[0].Observation))
}
