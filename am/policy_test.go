package am

import (
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"io/ioutil"
	"testing"

	mocks "github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/mocks/am"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPolicySetExists(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	common.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("policyset-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := PolicySetExists("Open Banking")

	assert.True(t, b)
}

func TestPolicyExists(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	common.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("policy-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := PolicyExists("AISP Policy")

	assert.True(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)
}
