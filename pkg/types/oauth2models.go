package types

type (
	InheritedValueString struct {
		Inherited bool   `json:"inherited"`
		Value     string `json:"value"`
	}
	InheritedValueInt struct {
		Inherited bool `json:"inherited"`
		Value     int  `json:"value"`
	}
	InheritedValueBool struct {
		Inherited bool `json:"inherited"`
		Value     bool `json:"value"`
	}
	JwkSet struct {
		Inherited bool   `json:"inherited"`
		Value     string `json:"value"`
	}
	Type struct {
		ID         string `json:"_id"`
		Name       string `json:"name"`
		Collection bool   `json:"collection"`
	}

	CoreOAuth2Config struct {
		UsePolicyEngineForScope           bool `json:"usePolicyEngineForScope"`
		MacaroonTokensEnabled             bool `json:"macaroonTokensEnabled"`
		StatelessTokensEnabled            bool `json:"statelessTokensEnabled"`
		CodeLifetime                      int  `json:"codeLifetime"`
		IssueRefreshTokenOnRefreshedToken bool `json:"issueRefreshTokenOnRefreshedToken"`
		RefreshTokenLifetime              int  `json:"refreshTokenLifetime"`
		AccessTokenLifetime               int  `json:"accessTokenLifetime"`
		IssueRefreshToken                 bool `json:"issueRefreshToken"`
	}

	AdvancedOAuth2Config struct {
		TLSClientCertificateHeaderFormat        string        `json:"tlsClientCertificateHeaderFormat"`
		SupportedSubjectTypes                   []string      `json:"supportedSubjectTypes"`
		DefaultScopes                           []interface{} `json:"defaultScopes"`
		MacaroonTokenFormat                     string        `json:"macaroonTokenFormat"`
		CodeVerifierEnforced                    string        `json:"codeVerifierEnforced"`
		GrantTypes                              []string      `json:"grantTypes"`
		AuthenticationAttributes                []string      `json:"authenticationAttributes"`
		TokenSigningAlgorithm                   string        `json:"tokenSigningAlgorithm"`
		TokenEncryptionEnabled                  bool          `json:"tokenEncryptionEnabled"`
		HashSalt                                string        `json:"hashSalt"`
		ModuleMessageEnabledInPasswordGrant     bool          `json:"moduleMessageEnabledInPasswordGrant"`
		TLSCertificateBoundAccessTokensEnabled  bool          `json:"tlsCertificateBoundAccessTokensEnabled"`
		DisplayNameAttribute                    string        `json:"displayNameAttribute"`
		SupportedScopes                         []string      `json:"supportedScopes"`
		ResponseTypeClasses                     []string      `json:"responseTypeClasses"`
		TokenCompressionEnabled                 bool          `json:"tokenCompressionEnabled"`
		AllowedAudienceValues                   []interface{} `json:"allowedAudienceValues"`
		TLSCertificateRevocationCheckingEnabled bool          `json:"tlsCertificateRevocationCheckingEnabled"`
	}

	CoreOIDCConfig struct {
		SupportedClaims                      []string `json:"supportedClaims"`
		JwtTokenLifetime                     int      `json:"jwtTokenLifetime"`
		SupportedIDTokenEncryptionAlgorithms []string `json:"supportedIDTokenEncryptionAlgorithms"`
		SupportedIDTokenEncryptionMethods    []string `json:"supportedIDTokenEncryptionMethods"`
		SupportedIDTokenSigningAlgorithms    []string `json:"supportedIDTokenSigningAlgorithms"`
	}

	LoaMapping struct {
		UrnOpenbankingPsd2Sca string `json:"urn:openbanking:psd2:sca"`
		UrnOpenbankingPsd2Ca  string `json:"urn:openbanking:psd2:ca"`
	}
	AmrMappings struct {
	}

	AdvancedOIDCConfig struct {
		SupportedTokenIntrospectionResponseSigningAlgorithms    []string      `json:"supportedTokenIntrospectionResponseSigningAlgorithms"`
		SupportedRequestParameterSigningAlgorithms              []string      `json:"supportedRequestParameterSigningAlgorithms"`
		IDTokenInfoClientAuthenticationEnabled                  bool          `json:"idTokenInfoClientAuthenticationEnabled"`
		AlwaysAddClaimsToToken                                  bool          `json:"alwaysAddClaimsToToken"`
		LoaMapping                                              LoaMapping    `json:"loaMapping"`
		SupportedTokenEndpointAuthenticationSigningAlgorithms   []string      `json:"supportedTokenEndpointAuthenticationSigningAlgorithms"`
		SupportedRequestParameterEncryptionAlgorithms           []string      `json:"supportedRequestParameterEncryptionAlgorithms"`
		AuthorisedOpenIDConnectSSOClients                       []interface{} `json:"authorisedOpenIdConnectSSOClients"`
		StoreOpsTokens                                          bool          `json:"storeOpsTokens"`
		SupportedRequestParameterEncryptionEnc                  []string      `json:"supportedRequestParameterEncryptionEnc"`
		DefaultACR                                              []interface{} `json:"defaultACR"`
		AmrMappings                                             AmrMappings   `json:"amrMappings"`
		ClaimsParameterSupported                                bool          `json:"claimsParameterSupported"`
		SupportedUserInfoEncryptionAlgorithms                   []string      `json:"supportedUserInfoEncryptionAlgorithms"`
		SupportedUserInfoEncryptionEnc                          []string      `json:"supportedUserInfoEncryptionEnc"`
		SupportedUserInfoSigningAlgorithms                      []string      `json:"supportedUserInfoSigningAlgorithms"`
		SupportedTokenIntrospectionResponseEncryptionEnc        []string      `json:"supportedTokenIntrospectionResponseEncryptionEnc"`
		SupportedTokenIntrospectionResponseEncryptionAlgorithms []string      `json:"supportedTokenIntrospectionResponseEncryptionAlgorithms"`
		AuthorisedIdmDelegationClients                          []interface{} `json:"authorisedIdmDelegationClients"`
	}

	ClientDynamicRegistrationConfig struct {
		DynamicClientRegistrationSoftwareStatementRequired bool          `json:"dynamicClientRegistrationSoftwareStatementRequired"`
		GenerateRegistrationAccessTokens                   bool          `json:"generateRegistrationAccessTokens"`
		RequiredSoftwareStatementAttestedAttributes        []interface{} `json:"requiredSoftwareStatementAttestedAttributes"`
		DynamicClientRegistrationScope                     string        `json:"dynamicClientRegistrationScope"`
		AllowDynamicRegistration                           bool          `json:"allowDynamicRegistration"`
	}

	CibaConfig struct {
		SupportedCibaSigningAlgorithms []string `json:"supportedCibaSigningAlgorithms"`
		CibaAuthReqIDLifetime          int      `json:"cibaAuthReqIdLifetime"`
		CibaMinimumPollingInterval     int      `json:"cibaMinimumPollingInterval"`
	}

	Consent struct {
		ClientsCanSkipConsent                    bool     `json:"clientsCanSkipConsent"`
		SupportedRcsRequestSigningAlgorithms     []string `json:"supportedRcsRequestSigningAlgorithms"`
		SupportedRcsRequestEncryptionMethods     []string `json:"supportedRcsRequestEncryptionMethods"`
		SupportedRcsRequestEncryptionAlgorithms  []string `json:"supportedRcsRequestEncryptionAlgorithms"`
		SupportedRcsResponseSigningAlgorithms    []string `json:"supportedRcsResponseSigningAlgorithms"`
		EnableRemoteConsent                      bool     `json:"enableRemoteConsent"`
		SupportedRcsResponseEncryptionAlgorithms []string `json:"supportedRcsResponseEncryptionAlgorithms"`
		SupportedRcsResponseEncryptionMethods    []string `json:"supportedRcsResponseEncryptionMethods"`
		RemoteConsentServiceID                   string   `json:"remoteConsentServiceId"`
	}

	PluginsConfig struct {
		AccessTokenModificationScript           string `json:"accessTokenModificationScript"`
		AccessTokenEnricherClass                string `json:"accessTokenEnricherClass"`
		AccessTokenModificationPluginType       string `json:"accessTokenModificationPluginType"`
		AccessTokenModifierClass                string `json:"accessTokenModifierClass"`
		AuthorizeEndpointDataProviderClass      string `json:"authorizeEndpointDataProviderClass"`
		AuthorizeEndpointDataProviderPluginType string `json:"authorizeEndpointDataProviderPluginType"`
		OidcClaimsScript                        string `json:"oidcClaimsScript"`
		OidcClaimsClass                         string `json:"oidcClaimsClass"`
		OidcClaimsPluginType                    string `json:"oidcClaimsPluginType"`
		EvaluateScopeClass                      string `json:"evaluateScopeClass"`
		EvaluateScopePluginType                 string `json:"evaluateScopePluginType"`
		ValidateScopeClass                      string `json:"validateScopeClass"`
		ValidateScopePluginType                 string `json:"validateScopePluginType"`
	}

	DeviceCodeConfig struct {
		DevicePollInterval int `json:"devicePollInterval"`
		DeviceCodeLifetime int `json:"deviceCodeLifetime"`
	}
)

// RemoteConsent struct for configuring the remote consent service
type RemoteConsent struct {
	Userpassword                             interface{}          `json:"userpassword"`
	RemoteConsentRequestEncryptionAlgorithm  InheritedValueString `json:"remoteConsentRequestEncryptionAlgorithm"`
	PublicKeyLocation                        InheritedValueString `json:"publicKeyLocation"`
	JwksCacheTimeout                         InheritedValueInt    `json:"jwksCacheTimeout"`
	RemoteConsentResponseSigningAlg          InheritedValueString `json:"remoteConsentResponseSigningAlg"`
	RemoteConsentRequestSigningAlgorithm     InheritedValueString `json:"remoteConsentRequestSigningAlgorithm"`
	JwkSet                                   JwkSet               `json:"jwkSet"`
	JwkStoreCacheMissCacheTime               InheritedValueInt    `json:"jwkStoreCacheMissCacheTime"`
	RemoteConsentResponseEncryptionMethod    InheritedValueString `json:"remoteConsentResponseEncryptionMethod"`
	RemoteConsentRedirectURL                 InheritedValueString `json:"remoteConsentRedirectUrl"`
	RemoteConsentRequestEncryptionEnabled    InheritedValueBool   `json:"remoteConsentRequestEncryptionEnabled"`
	RemoteConsentRequestEncryptionMethod     InheritedValueString `json:"remoteConsentRequestEncryptionMethod"`
	RemoteConsentResponseEncryptionAlgorithm InheritedValueString `json:"remoteConsentResponseEncryptionAlgorithm"`
	RequestTimeLimit                         InheritedValueInt    `json:"requestTimeLimit"`
	JwksURI                                  InheritedValueString `json:"jwksUri"`
	Type                                     Type                 `json:"_type"`
}

type PublisherAgent struct {
	Userpassword                      interface{}          `json:"userpassword"`
	PublicKeyLocation                 InheritedValueString `json:"publicKeyLocation"`
	JwksCacheTimeout                  InheritedValueInt    `json:"jwksCacheTimeout"`
	SoftwareStatementSigningAlgorithm InheritedValueString `json:"softwareStatementSigningAlgorithm"`
	JwkSet                            JwkSet               `json:"jwkSet"`
	Issuer                            InheritedValueString `json:"issuer"`
	JwkStoreCacheMissCacheTime        InheritedValueInt    `json:"jwkStoreCacheMissCacheTime"`
	JwksURI                           InheritedValueString `json:"jwksUri"`
}

type OAuth2Provider struct {
	CoreOAuth2Config                CoreOAuth2Config                `json:"coreOAuth2Config"`
	AdvancedOAuth2Config            AdvancedOAuth2Config            `json:"advancedOAuth2Config"`
	CoreOIDCConfig                  CoreOIDCConfig                  `json:"coreOIDCConfig"`
	AdvancedOIDCConfig              AdvancedOIDCConfig              `json:"advancedOIDCConfig"`
	ClientDynamicRegistrationConfig ClientDynamicRegistrationConfig `json:"clientDynamicRegistrationConfig"`
	CibaConfig                      CibaConfig                      `json:"cibaConfig"`
	Consent                         Consent                         `json:"consent"`
	DeviceCodeConfig                DeviceCodeConfig                `json:"deviceCodeConfig"`
	PluginsConfig                   PluginsConfig                   `json:"pluginsConfig"`
	ID                              string                          `json:"_id"`
	Type                            Type                            `json:"_type"`
}
