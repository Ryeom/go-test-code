package util

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type TokenBucket struct {
	capacity      int        // 버킷에 저장가능한 최대토큰 개수
	tokens        int        // 버킷의 현재토큰수
	refillRate    int        // 초당 추가할 토큰 수
	mu            sync.Mutex
	lastTimestamp time.Time
}

func NewTokenBucket(capacity, refillRate int) *TokenBucket {
	return &TokenBucket{
		capacity:      capacity,
		tokens:        capacity,
		refillRate:    refillRate,
		lastTimestamp: time.Now(),
	}
}

func (tb *TokenBucket) AllowRequest(tokens int) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refillTokens()

	if tokens <= tb.tokens {
		tb.tokens -= tokens
		return true
	}

	return false
}

func (tb *TokenBucket) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(tb.lastTimestamp)
	tb.tokens += int(elapsed.Seconds()) * tb.refillRate

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}

	tb.lastTimestamp = now
}

func TestTb(t *testing.T) {
	//10개의 토큰 용량과 초당 2개의 토큰 리필 속도를 가진 토큰 버킷 생성
	tb := NewTokenBucket(10, 2)

	// 각각 2개의 토큰으로 15개의 요청을 만듦
	for i := 0; i < 15; i++ {
		if tb.AllowRequest(2) {
			fmt.Println("Request", i+1, "succeeded")
		} else {
			fmt.Println("Request", i+1, "failed")
		}

		time.Sleep(500 * time.Millisecond)
	}
}
