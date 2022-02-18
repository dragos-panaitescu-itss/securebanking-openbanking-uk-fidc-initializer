package am

import (
	"io/ioutil"
	"testing"

	mocks "github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/mocks/am"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindExistingAlphaClient(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("client-check-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := AlphaClientsExist("Doesn't exist")
	assert.False(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)

	b = AlphaClientsExist("ig-client")
	assert.True(t, b)
}
