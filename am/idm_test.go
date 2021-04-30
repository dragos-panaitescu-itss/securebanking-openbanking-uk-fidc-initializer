package am

import (
	"io/ioutil"
	"testing"

	mocks "github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/mocks/am"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWillReturnAllMissingObjects(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("managed-objects-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	expectedMissing := []string{"abc", "def"}
	allMissing := MissingObjects(expectedMissing)
	assert.Equal(t, expectedMissing, allMissing)
}

func TestWillReturnPartialListOfMissingObjects(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("managed-objects-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	expectedMissing := []string{"abc", "def"}
	allMissing := MissingObjects([]string{"anotherObject", "abc", "def", "api_client"})
	assert.Equal(t, expectedMissing, allMissing)
}

func TestWillReturnNoMissingObjects(t *testing.T) {
	mockRestReaderWriter := &mocks.RestReaderWriter{}
	Client = mockRestReaderWriter
	buffer, _ := ioutil.ReadFile("managed-objects-test.json")
	mockRestReaderWriter.On("Get", mock.Anything, mock.Anything).
		Return(buffer)

	expectedMissing := []string{}
	allMissing := MissingObjects([]string{"anotherObject", "api_client"})
	assert.Equal(t, expectedMissing, allMissing)
}

func TestWillReturnObjectNamesFromPath(t *testing.T) {
	names := ObjectNames("testconfig/")
	expectedNames := []string{"test.user.1", "test1", "test2", "test3"}

	assert.Equal(t, expectedNames, names)
}

func TestManagedObjectDirectoriesExist(t *testing.T) {
	_, err := ioutil.ReadDir("../config/managed-objects/additional")

	assert.Nil(t, err, "The managed object directory additional/ should exist")

	_, err = ioutil.ReadDir("../config/managed-objects/openbanking")

	assert.Nil(t, err, "The managed object config directory config/managed-objects/openbanking/ should exist")
}
