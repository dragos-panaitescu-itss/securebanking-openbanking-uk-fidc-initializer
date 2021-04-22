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
	buffer, _ := ioutil.ReadFile("client-check-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := ServiceIdentityExists("ig-client")

	assert.True(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)

	b = ServiceIdentityExists("Doesnt existy")
	assert.False(t, b)
}
