package securebanking

import (
	"io/ioutil"
	"secure-banking-uk-initializer/pkg/httprest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocks "secure-banking-uk-initializer/pkg/mocks/am"
)

func TestFindSoftwarePublisherAgent(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("oauth2-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := softwarePublisherAgentExists("OBRI")
	assert.True(t, b)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)

	b = softwarePublisherAgentExists("test-publisher")
	assert.True(t, b)
}

func TestFindRemoteConsent(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("remote-consent-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := remoteConsentExists("secure-open-banking-rcs")

	assert.True(t, b)
}

func TestReturnScriptId(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("script-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	s := httprest.GetScriptIdByName("test script")

	assert.Equal(t, "123", s)
	mockRestReaderWriter.AssertCalled(t, "Get", mock.Anything, mock.Anything)

	s = httprest.GetScriptIdByName("Doesnt existy")
	assert.Equal(t, "", s)
}

func TestFindOAuth2Provider(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("oauth2provider-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	b := oauth2ProviderExists("oauth-oidc")

	assert.True(t, b)
}
