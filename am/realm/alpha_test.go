package realm

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/am"
	mocks "github.com/secureBankingAcceleratorToolkit/securebanking-openbanking-uk-fidc-initialiszer/mocks/am"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindExistingAlphaClient(t *testing.T) {
	mockResultFn := func(args mock.Arguments) {
		ob := args.Get(2)
		buffer, _ := ioutil.ReadFile("client-check-test.json")
		_ = json.Unmarshal(buffer, &ob)
	}

	mockRestReaderWriter := &mocks.RestReaderWriter{}
	am.Client = mockRestReaderWriter
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything, mock.Anything).
		Run(mockResultFn)

	b := AlphaClientsExist("Doesnt existy")
	assert.False(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything, mock.Anything)

	b = AlphaClientsExist("ig-client")
	assert.True(t, b)
}

func TestMangedObjectExists(t *testing.T) {
	mockResultFn := func(args mock.Arguments) {
		ob := args.Get(2)
		buffer, _ := ioutil.ReadFile("managed-objects-test.json")
		_ = json.Unmarshal(buffer, &ob)
	}

	mockRestReaderWriter := &mocks.RestReaderWriter{}
	am.Client = mockRestReaderWriter
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything, mock.Anything).
		Run(mockResultFn)

	b := ManagedObjectExists("api_client")
	assert.True(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything, mock.Anything)

	b = ManagedObjectExists("xyz")
	assert.False(t, b)
}
