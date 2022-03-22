package types

// OAuth2Client model
type OAuth2Client struct {
	CoreOAuth2ClientConfig struct {
		Agentgroup string `json:"agentgroup"`
		Status     struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"status"`
		Userpassword string `json:"userpassword,omitempty"`
		ClientType   struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"clientType"`
		LoopbackInterfaceRedirection struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"loopbackInterfaceRedirection"`
		RedirectionUris struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"redirectionUris"`
		Scopes struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"scopes"`
		DefaultScopes struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"defaultScopes"`
		ClientName struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"clientName"`
		AuthorizationCodeLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"authorizationCodeLifetime"`
		RefreshTokenLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"refreshTokenLifetime"`
		AccessTokenLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"accessTokenLifetime"`
	} `json:"coreOAuth2ClientConfig"`

	AdvancedOAuth2ClientConfig struct {
		Name struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"name"`
		Descriptions struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"descriptions"`
		RequestUris struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"requestUris"`
		ResponseTypes struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"responseTypes"`
		GrantTypes struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"grantTypes"`
		Contacts struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"contacts"`
		TokenEndpointAuthMethod struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenEndpointAuthMethod"`
		SectorIdentifierURI struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"sectorIdentifierUri"`
		SubjectType struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"subjectType"`
		UpdateAccessToken struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"updateAccessToken"`
		ClientURI struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"clientUri"`
		LogoURI struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"logoUri"`
		PolicyURI struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"policyUri"`
		IsConsentImplied struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"isConsentImplied"`
		MixUpMitigation struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"mixUpMitigation"`
	} `json:"advancedOAuth2ClientConfig"`

	CoreOpenIDClientConfig struct {
		Claims struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"claims"`
		PostLogoutRedirectURI struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"postLogoutRedirectUri"`
		ClientSessionURI struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"clientSessionUri"`
		DefaultMaxAge struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"defaultMaxAge"`
		DefaultMaxAgeEnabled struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"defaultMaxAgeEnabled"`
		DefaultAcrValues struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"defaultAcrValues"`
		JwtTokenLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"jwtTokenLifetime"`
	} `json:"coreOpenIDClientConfig"`

	SignEncOAuth2ClientConfig struct {
		JwksURI struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"jwksUri"`
		JwksCacheTimeout struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"jwksCacheTimeout"`
		JwkStoreCacheMissCacheTime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"jwkStoreCacheMissCacheTime"`
		TokenEndpointAuthSigningAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenEndpointAuthSigningAlgorithm"`
		JwkSet struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"jwkSet"`
		IDTokenSignedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"idTokenSignedResponseAlg"`
		IDTokenEncryptionEnabled struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"idTokenEncryptionEnabled"`
		IDTokenEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"idTokenEncryptionAlgorithm"`
		IDTokenEncryptionMethod struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"idTokenEncryptionMethod"`
		IDTokenPublicEncryptionKey struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"idTokenPublicEncryptionKey"`
		ClientJwtPublicKey struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"clientJwtPublicKey"`
		MTLSTrustedCert struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"mTLSTrustedCert"`
		MTLSSubjectDN struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"mTLSSubjectDN"`
		MTLSCertificateBoundAccessTokens struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"mTLSCertificateBoundAccessTokens"`
		PublicKeyLocation struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"publicKeyLocation"`
		UserinfoResponseFormat struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"userinfoResponseFormat"`
		UserinfoSignedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"userinfoSignedResponseAlg"`
		UserinfoEncryptedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"userinfoEncryptedResponseAlg"`
		UserinfoEncryptedResponseEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"userinfoEncryptedResponseEncryptionAlgorithm"`
		RequestParameterSignedAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"requestParameterSignedAlg"`
		RequestParameterEncryptedAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"requestParameterEncryptedAlg"`
		RequestParameterEncryptedEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"requestParameterEncryptedEncryptionAlgorithm"`
		TokenIntrospectionResponseFormat struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionResponseFormat"`
		TokenIntrospectionSignedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionSignedResponseAlg"`
		TokenIntrospectionEncryptedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionEncryptedResponseAlg"`
		TokenIntrospectionEncryptedResponseEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionEncryptedResponseEncryptionAlgorithm"`
	} `json:"signEncOAuth2ClientConfig"`

	CoreUmaClientConfig struct {
		ClaimsRedirectionUris struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"claimsRedirectionUris"`
	} `json:"coreUmaClientConfig"`
}

// PolicyAgent model
type PolicyAgent struct {
	Userpassword         string               `json:"userpassword"`
	IgTokenIntrospection IgTokenIntrospection `json:"igTokenIntrospection"`
}

type IgTokenIntrospection struct {
	Value     string `json:"value"`
	Inherited bool   `json:"inherited"`
}
