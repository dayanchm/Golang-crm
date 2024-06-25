package helpers

import (
	"net"
	"net/http"
	"sync"
	"time"
)

var loginAttempts = make(map[string]int)
var blockTimes = make(map[string]time.Time)
var lock = sync.Mutex{}

func getIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func CheckLoginAttempts(r *http.Request) (bool, time.Duration) {
	ip := getIP(r)
	lock.Lock()
	defer lock.Unlock()

	if blockTime, exists := blockTimes[ip]; exists && time.Now().Before(blockTime) {
		remainingTime := time.Until(blockTime)
		return false, remainingTime
	}

	if count, exists := loginAttempts[ip]; exists {
		if count >= 5 && count < 10 {
			blockDuration := time.Minute * 15
			blockTimes[ip] = time.Now().Add(blockDuration)
			return false, blockDuration
		} else if count >= 10 {
			blockDuration := time.Minute * 30
			blockTimes[ip] = time.Now().Add(blockDuration)
			return false, blockDuration
		}
	}
	return true, 0
}

func RecordLoginAttempt(r *http.Request, success bool) {
	ip := getIP(r)
	lock.Lock()
	defer lock.Unlock()
	if success {
		delete(loginAttempts, ip)
		delete(blockTimes, ip)
	} else {
		loginAttempts[ip]++
	}
}
