package worker

import (
	"errors"
	"strings"
)

// errorContainsMessage checks if the message of err1 is contained within any error in the chain of err2.
func errorContainsMessage(err1, err2 error) bool {
	targetMessage := err1.Error()
	for {
		if strings.Contains(err2.Error(), targetMessage) {
			return true
		}
		// Unwrap the next error in the chain
		unwrappedErr := errors.Unwrap(err2)
		if unwrappedErr == nil {
			break
		}
		err2 = unwrappedErr
	}
	return false
}
