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

func TestWillTemplateFile(t *testing.T) {
	templatefile := `{"key":"{{.Hosts.IdentityPlatformFQDN}}"}`
	expectedString := `{"key":"https://example.com"}`
	tmp, err := os.CreateTemp("", "")
	os.WriteFile(tmp.Name(), []byte(templatefile), 0666)
	cfg := &types.Configuration{}
	cfg.Hosts.IdentityPlatformFQDN = "https://example.com"
	b, err := Template(tmp.Name(), cfg)
	assert.Nil(t, err)
	assert.Equal(t, string(b), expectedString)
}

func TestWillUnmarshalFileWithConditional(t *testing.T) {
	templatefile := `{"key":"{{- if .Hosts.WildcardFQDN -}}{{ .Hosts.WildcardFQDN }}{{- else -}}{{ .Hosts.BaseFQDN }}{{- end }}"}`
	tmp, err := os.CreateTemp("", "")
	os.WriteFile(tmp.Name(), []byte(templatefile), 0666)
	cfg := &types.Configuration{}
	cfg.Hosts.WildcardFQDN = ""
	cfg.Hosts.BaseFQDN = "forgerock.com"
	testJSON := &Test{}
	err = Unmarshal(tmp.Name(), cfg, testJSON)
	assert.Nil(t, err)
	assert.Equal(t, testJSON.Key, cfg.Hosts.BaseFQDN)
}
