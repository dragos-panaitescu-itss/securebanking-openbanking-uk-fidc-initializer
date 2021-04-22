package am

type (
	ActionValues struct {
		GET bool `json:"GET"`
	}
	Subject struct {
		Type       string `json:"type"`
		ClaimName  string `json:"claimName"`
		ClaimValue string `json:"claimValue"`
	}
	Condition struct {
		Type     string `json:"type"`
		ScriptID string `json:"scriptId"`
	}
)

// OpenBankingPolicySet model
type OpenBankingPolicySet struct {
	Name              string   `json:"name"`
	DisplayName       string   `json:"displayName"`
	Description       string   `json:"description"`
	ResourceTypeUuids []string `json:"resourceTypeUuids"`
	Realm             string   `json:"realm"`
	ApplicationType   string   `json:"applicationType"`
	Conditions        []string `json:"conditions"`
	Subjects          []string `json:"subjects"`
}

// CreatePolicy AISP or PISP policy model
type CreatePolicy struct {
	Name             string       `json:"name"`
	Active           bool         `json:"active"`
	Description      string       `json:"description"`
	Resources        []string     `json:"resources"`
	ApplicationName  string       `json:"applicationName"`
	ActionValues     ActionValues `json:"actionValues"`
	Subject          Subject      `json:"subject"`
	Condition        Condition    `json:"condition"`
	ResourceTypeUUID string       `json:"resourceTypeUuid"`
}

// EngineOAuth2Client model
type EngineOAuth2Client struct {
	CoreOAuth2ClientConfig struct {
		Userpassword                 string `json:"userpassword"`
		LoopbackInterfaceRedirection struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"loopbackInterfaceRedirection"`
		DefaultScopes struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"defaultScopes"`
		RefreshTokenLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"refreshTokenLifetime"`
		Scopes struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"scopes"`
		Status struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"status"`
		AccessTokenLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"accessTokenLifetime"`
		RedirectionUris struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"redirectionUris"`
		ClientName struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"clientName"`
		ClientType struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"clientType"`
		AuthorizationCodeLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"authorizationCodeLifetime"`
	} `json:"coreOAuth2ClientConfig"`

	AdvancedOAuth2ClientConfig struct {
		Descriptions struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"descriptions"`
		RequestUris struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"requestUris"`
		LogoURI struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"logoUri"`
		SubjectType struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"subjectType"`
		ClientURI struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"clientUri"`
		Name struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"name"`
		Contacts struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"contacts"`
		ResponseTypes struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"responseTypes"`
		UpdateAccessToken struct {
			Inherited bool `json:"inherited"`
		} `json:"updateAccessToken"`
		MixUpMitigation struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"mixUpMitigation"`
		JavascriptOrigins struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"javascriptOrigins"`
		PolicyURI struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"policyUri"`
		SectorIdentifierURI struct {
			Inherited bool `json:"inherited"`
		} `json:"sectorIdentifierUri"`
		TokenEndpointAuthMethod struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenEndpointAuthMethod"`
		IsConsentImplied struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"isConsentImplied"`
		GrantTypes struct {
			Inherited bool     `json:"inherited"`
			Value     []string `json:"value"`
		} `json:"grantTypes"`
	} `json:"advancedOAuth2ClientConfig"`

	SignEncOAuth2ClientConfig struct {
		TokenEndpointAuthSigningAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenEndpointAuthSigningAlgorithm"`
		IDTokenEncryptionEnabled struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"idTokenEncryptionEnabled"`
		TokenIntrospectionEncryptedResponseEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionEncryptedResponseEncryptionAlgorithm"`
		RequestParameterSignedAlg struct {
			Inherited bool `json:"inherited"`
		} `json:"requestParameterSignedAlg"`
		ClientJwtPublicKey struct {
			Inherited bool `json:"inherited"`
		} `json:"clientJwtPublicKey"`
		IDTokenPublicEncryptionKey struct {
			Inherited bool `json:"inherited"`
		} `json:"idTokenPublicEncryptionKey"`
		MTLSSubjectDN struct {
			Inherited bool `json:"inherited"`
		} `json:"mTLSSubjectDN"`
		UserinfoResponseFormat struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"userinfoResponseFormat"`
		MTLSCertificateBoundAccessTokens struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"mTLSCertificateBoundAccessTokens"`
		PublicKeyLocation struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"publicKeyLocation"`
		TokenIntrospectionResponseFormat struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionResponseFormat"`
		JwkStoreCacheMissCacheTime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"jwkStoreCacheMissCacheTime"`
		RequestParameterEncryptedEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"requestParameterEncryptedEncryptionAlgorithm"`
		UserinfoSignedResponseAlg struct {
			Inherited bool `json:"inherited"`
		} `json:"userinfoSignedResponseAlg"`
		IDTokenEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"idTokenEncryptionAlgorithm"`
		RequestParameterEncryptedAlg struct {
			Inherited bool `json:"inherited"`
		} `json:"requestParameterEncryptedAlg"`
		MTLSTrustedCert struct {
			Inherited bool `json:"inherited"`
		} `json:"mTLSTrustedCert"`
		JwkSet struct {
			Inherited bool `json:"inherited"`
		} `json:"jwkSet"`
		IDTokenEncryptionMethod struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"idTokenEncryptionMethod"`
		JwksCacheTimeout struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"jwksCacheTimeout"`
		UserinfoEncryptedResponseAlg struct {
			Inherited bool `json:"inherited"`
		} `json:"userinfoEncryptedResponseAlg"`
		IDTokenSignedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"idTokenSignedResponseAlg"`
		JwksURI struct {
			Inherited bool `json:"inherited"`
		} `json:"jwksUri"`
		TokenIntrospectionSignedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionSignedResponseAlg"`
		UserinfoEncryptedResponseEncryptionAlgorithm struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"userinfoEncryptedResponseEncryptionAlgorithm"`
		TokenIntrospectionEncryptedResponseAlg struct {
			Inherited bool   `json:"inherited"`
			Value     string `json:"value"`
		} `json:"tokenIntrospectionEncryptedResponseAlg"`
	} `json:"signEncOAuth2ClientConfig"`

	CoreOpenIDClientConfig struct {
		Claims struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"claims"`
		ClientSessionURI struct {
			Inherited bool `json:"inherited"`
		} `json:"clientSessionUri"`
		DefaultAcrValues struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"defaultAcrValues"`
		JwtTokenLifetime struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"jwtTokenLifetime"`
		DefaultMaxAgeEnabled struct {
			Inherited bool `json:"inherited"`
			Value     bool `json:"value"`
		} `json:"defaultMaxAgeEnabled"`
		DefaultMaxAge struct {
			Inherited bool `json:"inherited"`
			Value     int  `json:"value"`
		} `json:"defaultMaxAge"`
		PostLogoutRedirectURI struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"postLogoutRedirectUri"`
	} `json:"coreOpenIDClientConfig"`

	CoreUmaClientConfig struct {
		ClaimsRedirectionUris struct {
			Inherited bool          `json:"inherited"`
			Value     []interface{} `json:"value"`
		} `json:"claimsRedirectionUris"`
	} `json:"coreUmaClientConfig"`
}
