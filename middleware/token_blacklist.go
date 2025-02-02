package middleware

import (
	"sync"
	"time"
)

var TokenBlacklist = struct {
	sync.RWMutex
	blacklist map[string]time.Time
}{blacklist: make(map[string]time.Time)}

func BlacklistToken(token string, exp time.Time) {
	TokenBlacklist.Lock()
	TokenBlacklist.blacklist[token] = exp
	TokenBlacklist.Unlock()
}

func IsTokenBlacklisted(token string) bool {
	TokenBlacklist.RLock()
	defer TokenBlacklist.RUnlock()

	exp, exists := TokenBlacklist.blacklist[token]
	if !exists {
		return false
	}

	if time.Now().After(exp) {
		TokenBlacklist.Lock()
		delete(TokenBlacklist.blacklist, token)
		TokenBlacklist.Unlock()
		return false
	}

	return true
}
