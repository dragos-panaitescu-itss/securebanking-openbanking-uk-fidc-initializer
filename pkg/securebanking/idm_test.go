package securebanking

import (
	"io/ioutil"
	"secure-banking-uk-initializer/pkg/httprest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mocks "secure-banking-uk-initializer/pkg/mocks/am"
)

func TestWillReturnAllMissingObjects(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("managed-objects-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	expectedMissing := []string{"abc", "def"}
	allMissing := missingObjects(expectedMissing)
	assert.Equal(t, expectedMissing, allMissing)
}

func TestWillReturnPartialListOfMissingObjects(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("managed-objects-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	expectedMissing := []string{"abc", "def"}
	allMissing := missingObjects([]string{"anotherObject", "abc", "def", "api_client"})
	assert.Equal(t, expectedMissing, allMissing)
}

func TestWillReturnNoMissingObjects(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	httprest.Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("managed-objects-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	expectedMissing := []string{}
	allMissing := missingObjects([]string{"anotherObject", "api_client"})
	assert.Equal(t, expectedMissing, allMissing)
}

func TestWillReturnObjectNamesFromPath(t *testing.T) {
	names := objectNames("testconfig/")
	expectedNames := []string{"test.user.1", "test1", "test2", "test3"}

	assert.Equal(t, expectedNames, names)
}

func TestManagedObjectDirectoriesExist(t *testing.T) {
	_, err := ioutil.ReadDir("../config/defaults/managed-objects/openbanking")

	assert.Nil(t, err, "The managed object config directory config/managed-objects/openbanking/ should exist")
}
