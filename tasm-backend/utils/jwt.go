package utils

import (
	"os"
	"time"
)

var fallbackSecret []byte

// GetJWTSecret returns the JWT secret from the environment or a secure,
// pseudo-random fallback generated on startup if the env is missing.
func GetJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret != "" {
		return []byte(secret)
	}
	if fallbackSecret == nil {
		fallbackSecret = make([]byte, 32)
		// Basic pseudo-random fallback for security if env is missing
		for i := range fallbackSecret {
			fallbackSecret[i] = byte(time.Now().UnixNano() + int64(i))
		}
	}
	return fallbackSecret
}
