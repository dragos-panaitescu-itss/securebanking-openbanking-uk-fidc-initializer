package platform

import (
	"net/http"
	"net/url"

	"go.uber.org/zap"

	"github.com/go-resty/resty/v2"
	"github.com/secureBankingAccessToolkit/securebanking-openbanking-uk-fidc-initialiszer/common"
	"github.com/spf13/viper"
)

// AdminToken returned by IDM
type AdminToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// Session containing the access token and cookie
type Session struct {
	authCode  string
	AuthToken AdminToken
	Cookie    *http.Cookie
}

var client = resty.New().SetRedirectPolicy(resty.NoRedirectPolicy()).SetError(common.RestError{})

// Authenticate user against platform, returns the iPlanetDomainPro cookie and access token
func (s *Session) Authenticate() (*http.Cookie, string) {
	s.GetIDMAdminAuthCode()
	s.GetIDMAdminToken()
	return s.Cookie, s.AuthToken.AccessToken
}

// FromAmAdminSession - get a session token from AM for authentication
//    returns the Session object with embedded session cookie
func FromAmAdminSession() *Session {
	zap.L().Debug("Getting an admin session from AM")
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/json/realms/root/authenticate"
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-OpenAM-Username", "amadmin").
		SetHeader("X-OpenAM-Password", viper.GetString("OPEN_AM_PASSWORD")).
		Post(path)
	common.RaiseForStatus(err, resp.Error())

	var cookieValue string = ""
	for _, cookie := range resp.Cookies() {
		zap.S().Debugw("Cookie found", "cookie", cookie)
		if cookie.Name == "iPlanetDirectoryPro" {
			cookieValue = cookie.Value
		}
	}
	if cookieValue == "" {
		zap.S().Fatalw("Cannot find iPlanetDirectoryPro cookie",
			"statusCode", resp.StatusCode(),
			"advice", `Calling this method twice might result in this error,
				 AM will not react well to successive session requests`,
			"error", resp.Error())
	}
	iPlanetDirCookie := &http.Cookie{
		Name:     "iPlanetDirectoryPro",
		Value:    cookieValue,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Domain:   viper.GetString("IAM_FQDN"),
	}
	s := &Session{}
	s.Cookie = iPlanetDirCookie
	zap.S().Infow("New AM session created", "cookie", s.Cookie)
	return s
}

// GetIDMAdminAuthCode - get auth code from IDM.
// 		Redirects should be disabled, we expect a 302 status code here
func (s *Session) GetIDMAdminAuthCode() {
	zap.L().Debug("Getting IDM admin auth code")
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/oauth2/authorize"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetQueryParams(map[string]string{
			"redirect_uri":          "https://" + viper.GetString("IAM_FQDN") + "/platform/appAuthHelperRedirect.html",
			"client_id":             "idmAdminClient",
			"response_type":         "code",
			"scope":                 "fr:idm:*",
			"code_challenge":        "gX2yL78GGlz3QHsQZKPf96twOmUBKxn1-IXPd5_EHdA",
			"code_challenge_method": "S256",
		}).
		SetCookie(s.Cookie).
		Get(path)
	if err != nil {
		if resp.StatusCode() != 302 {
			zap.S().Fatalw("Expecting 302 status code when getting auth code from IDM",
				"statusCode", resp.StatusCode(),
				"advice", "the idmAdminClient must exist in the root realm and redirects must be turned off",
				"error", resp.Error())
		}
	}
	v, err := url.ParseQuery(resp.Header().Get("Location"))
	if err != nil {
		zap.S().Fatalw("Error parsing location header", "statusCode", resp.StatusCode(), "error", err)
	}
	zap.S().Debugw("Got Location header from IDM", "Location", v)
	authCode := v["https://"+viper.GetString("IAM_FQDN")+"/platform/appAuthHelperRedirect.html?code"][0]
	s.authCode = authCode
}

// GetIDMAdminToken - get admin token from IDM
func (s *Session) GetIDMAdminToken() {
	zap.L().Debug("Getting admin token")
	token := &AdminToken{}
	path := "https://" + viper.GetString("IAM_FQDN") + "/am/oauth2/access_token"
	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(token).
		SetCookie(s.Cookie).
		SetFormData(map[string]string{
			"grant_type":    "authorization_code",
			"redirect_uri":  "https://" + viper.GetString("IAM_FQDN") + "/platform/appAuthHelperRedirect.html",
			"client_id":     "idmAdminClient",
			"code":          s.authCode,
			"code_verifier": "codeverifier",
		}).
		Post(path)
	common.RaiseForStatus(err, resp.Error())
	s.AuthToken = *token
}
