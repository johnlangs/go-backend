package utils

import (
	"context"
	"errors"
	"sync"
	"time"
)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consectiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		m.RLock()

		d := consectiveFailures - int(failureThreshold)

		if d >= 0 {
			// Exponential back-off calling
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}
	
		m.RUnlock()

		response, err := circuit(ctx)

		m.Lock()
		defer m.Unlock()

		lastAttempt = time.Now()

		if err != nil {
			consectiveFailures++
			return response, err
		}

		consectiveFailures = 0
		
		return response, err
	}
}
