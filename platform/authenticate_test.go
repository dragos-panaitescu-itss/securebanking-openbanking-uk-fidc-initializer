package platform

import (
	"log"
	"strconv"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConsumer(t *testing.T) {
	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "initializer",
		Provider: "am",
		Host:     "localhost",
	}
	defer pact.Teardown()

	var test = func() (err error) {
		viper.SetDefault("IAM_FQDN", "localhost:"+strconv.Itoa(pact.Server.Port))
		viper.SetDefault("SCHEME", "http")
		cookie := GetCookieNameFromAm()

		assert.Equal(t, "iPlanetDirectory", cookie)
		return nil
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given("Server info exists").
		UponReceiving("A request to get Server info").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/am/json/serverinfo/*"),
			Headers: dsl.MapMatcher{"Accept": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.MapMatcher{"cookieName": dsl.String("iPlanetDirectory")},
		})

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
