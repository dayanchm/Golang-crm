package helpers

import (
	"net/http"
	"sync"
	"time"
)

var loginAttempts = make(map[string]int)
var blockTimes = make(map[string]time.Time)
var lock = sync.Mutex{}

func CheckLoginAttempts(r *http.Request) (bool, time.Duration) {
	ip := r.RemoteAddr
	lock.Lock()
	defer lock.Unlock()

	if blockTime, exists := blockTimes[ip]; exists && time.Now().Before(blockTime) {
		remainingTime := time.Until(blockTime)
		return false, remainingTime
	}
	if count, exists := loginAttempts[ip]; exists && count >= 5 {
		blockDuration := time.Minute * time.Duration(15<<uint((count-5)/5))
		blockTimes[ip] = time.Now().Add(blockDuration)
		return false, blockDuration
	}
	return true, 0
}

func RecordLoginAttempt(r *http.Request, success bool) {
	ip := r.RemoteAddr
	lock.Lock()
	defer lock.Unlock()
	if success {
		delete(loginAttempts, ip)
		delete(blockTimes, ip)
	} else {
		loginAttempts[ip]++
	}
}
