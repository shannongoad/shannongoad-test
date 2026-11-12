package service

import "time"

// Config holds service defaults.
type Config struct {
	Timeout time.Duration
	Retries int
}
const DefaultTimeout = 30 * time.Second
const DefaultRetries = 3

// NewDefaultConfig returns defaults.
func NewDefaultConfig() Config {
	return Config{
		Timeout: DefaultTimeout,
		Retries: DefaultRetries,
	}
}
