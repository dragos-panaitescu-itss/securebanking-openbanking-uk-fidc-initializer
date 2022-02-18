package am

import (
	"io/ioutil"
	"testing"

	mocks "github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/mocks/am"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServiceIdentityExists(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("serviceaccount-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := ServiceIdentityExists("service_account.ig")

	assert.True(t, b)
}

func TestServiceIdentityNotExists(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("serviceaccount-empty-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := ServiceIdentityExists("service_account.ig")

	assert.False(t, b)
}
