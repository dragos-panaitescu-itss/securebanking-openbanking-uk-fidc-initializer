package types

// ServerInfo necessary fields returned from platform
type ServerInfo struct {
	CookieName   string `json:"cookieName"`
	SecureCookie bool   `json:"secureCookie"`
}

type AmResult struct {
	Result                  []Result    `json:"result"`
	Resultcount             int         `json:"resultCount"`
	Pagedresultscookie      interface{} `json:"pagedResultsCookie"`
	Totalpagedresultspolicy string      `json:"totalPagedResultsPolicy"`
	Totalpagedresults       int         `json:"totalPagedResults"`
	Remainingpagedresults   int         `json:"remainingPagedResults"`
}

type Result struct {
	ID       string `json:"_id"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
}

type ResultFilter struct {
	Result      []Result `json:"result"`
	ResultCount int      `json:"resultCount,omitempty"`
}

type Source struct {
	ExtensionClassName string `json:"extensionClassName"`
	Source             string `json:"source"`
	FixedValue         string `json:"fixedValue"`
	ContextPath        string `json:"contextPath"`
	ID                 string `json:"_id"`
	Type               Type   `json:"_type"`
}
