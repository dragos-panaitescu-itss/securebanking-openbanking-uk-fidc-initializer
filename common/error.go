package common

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// RestError for use with goresty SetError client method
//    Allow us to marshal client and server error responses
type RestError struct {
	Message string
	Code    int
	Detail  interface{}
}

// RaiseForStatus will exit if go resty returns an error in STRICT mode,
//    Be it client error, server error or other. Turning off
//    STRICT mode will simply warn of client/server errors.
func RaiseForStatus(err error, restError interface{}) {
	if err != nil {
		zap.S().Fatalw("Goresty has thrown an error when attempting to send", "error", err)
	}

	if restError != nil {
		strict := viper.GetBool("STRICT")
		if strict {
			zap.S().Fatalw("Goresty has sent the request but the status is > 399", "error", restError)
		}
		zap.S().Warnw("Goresty has sent the request but the status is > 399", "error", restError)
	}
}
