package securebanking

import (
	"io/ioutil"
	"secure-banking-uk-initializer/pkg/httprest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocks "secure-banking-uk-initializer/pkg/mocks/am"
)

func TestPolicySetExists(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("policyset-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := httprest.PolicySetExists("Open Banking")

	assert.True(t, b)
}

func TestPolicyExists(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("policy-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := httprest.PolicyExists("AISP Policy")

	assert.True(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)
}
