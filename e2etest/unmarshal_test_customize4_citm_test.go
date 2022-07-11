package e2e

import (
	"testing"

	"github.com/DRK-Blutspende-BaWueHe/go-hl7"
	"github.com/DRK-Blutspende-BaWueHe/go-hl7/lib/hl7v23"
	"github.com/DRK-Blutspende-BaWueHe/go-hl7/lib/hl7v24"
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

func TestCit_OUL_R21(t *testing.T) {
	var filedata string
	filedata = filedata + "MSH|^~\\&|Roche Diagnostics|cITm 1.10.02.0572|DRK FFM||20220711130056||OUL^R21|107737129|P|2.4|||NE|NE||UNICODE UTF-8<13>PID|1|?|||^|||\u000d"
	filedata = filedata + "SAC|||AA4F1A1A6J\u000d"
	filedata = filedata + "ORC|RE|AA4F1A1A6J|||||^^^^^?||20220709195656\u000d"
	filedata = filedata + "OBR|1|AA4F1A1A6J||ALL||20220711075520|||||||||1^^^^^^P||||||||||||^^^^^?|||||||\u000d"
	filedata = filedata + "OBX|1||AHBC2-R||-1\\S\\2.14|||N|||F|||20220711122749|^^^CCM2-c8k-5-1859-10#e801#2#2|bmserv\\S\\SYSTEM^System||CCM2-c8k-5-1859-10#e801#2#2|20220711122749\u000d"
	filedata = filedata + "TCD|1|AHBC2-R|1||||||^|\u000d"

	var message hl7v24.OUL_R21
	err := hl7.Unmarshal(
		[]byte(filedata),
		&message,
		hl7.EncodingUTF8,
		hl7.TimezoneEuropeBerlin)

	assert.Nil(t, err)
	assert.Equal(t, "AA4F1A1A6J", message.OrderObservation[0].Container.SpecimenAndContainerDetail.ContainerIdentifier.EntityIdentifier)
}
