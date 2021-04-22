package am

import (
	"io/ioutil"
	"testing"

	mocks "github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/mocks/am"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMangedObjectExists(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("managed-objects-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := ManagedObjectExists("api_client")
	assert.True(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)

	b = ManagedObjectExists("xyz")
	assert.False(t, b)
}
