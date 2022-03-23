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
