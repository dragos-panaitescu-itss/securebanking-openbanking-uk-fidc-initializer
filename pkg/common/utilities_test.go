package common

import (
	"os"
	"secure-banking-uk-initializer/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	Key string `json:"key"`
}

func TestWillUnmarshalFile(t *testing.T) {
	templatefile := `{"key":"{{.Hosts.IdentityPlatformFQDN}}"}`
	tmp, err := os.CreateTemp("", "")
	os.WriteFile(tmp.Name(), []byte(templatefile), 0666)
	cfg := &types.Configuration{}
	cfg.Hosts.IdentityPlatformFQDN = "https://example.com"
	testJSON := &Test{}
	err = Unmarshal(tmp.Name(), cfg, testJSON)
	assert.Nil(t, err)
	assert.Equal(t, testJSON.Key, cfg.Hosts.IdentityPlatformFQDN)
}
