package am

import "go.uber.org/zap"

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

func Find(arg string, ob *AmResult, fn func(*Result) string) bool {
	for _, e := range ob.Result {
		if fn(&e) == arg {
			zap.S().Infow("Argument found", "arg", arg)
			return true
		}
	}
	return false
}

func FindIdByName(name string, ob *AmResult, fn func(*Result) string) string {
	for _, e := range ob.Result {
		if fn(&e) == name {
			zap.S().Infow("ID found", "id", e.ID)
			return e.ID
		}
	}
	return ""
}
