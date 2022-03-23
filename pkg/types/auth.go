package types

type AuthzRole struct {
	Ref           string      `json:"_ref"`
	RefProperties interface{} `json:"_refProperties,omitempty"`
}

type ServiceUser struct {
	UserName  string      `json:"userName"`
	SN        string      `json:"sn"`
	GivenName string      `json:"givenName"`
	Mail      string      `json:"mail"`
	Password  string      `json:"password"`
	AuthzRole []AuthzRole `json:"authzRoles"`
}
