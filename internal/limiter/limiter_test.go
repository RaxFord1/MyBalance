package limiter

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestNewEventLimiter(t *testing.T) {
	got := NewEventLimiter(1, 1)

	if got == nil {
		t.Fatal("NewUserLimiter() = nil, want non-nil")
	}

	log.Println(got)
}

func TestEventLimiter_GetLimiter(t *testing.T) {
	userLimiter := NewEventLimiter(1, 1)
	userID := "testUser"

	limiter := userLimiter.GetLimiter(userID)
	if limiter == nil {
		t.Errorf("GetLimiter() = nil, want non-nil")
	}

	// Ensure that the same limiter instance is returned for the same user ID.
	if got := userLimiter.GetLimiter(userID); got != limiter {
		t.Errorf("GetLimiter() = %v, want %v", got, limiter)
	}
}

func TestEventLimiter_AllowAction(t *testing.T) {
	userLimiter := NewEventLimiter(1, 1)
	userID := "testUser"

	// Allow the first action.
	if err := userLimiter.Allow(userID); err != nil {
		t.Errorf("AllowAction() = false, want true")
	}

	// Allow the second action because of the burst capacity.
	if err := userLimiter.Allow(userID); err == nil {
		t.Errorf("AllowAction() = true, want false")
	}

	// The third action should be disallowed immediately due to rate limiting.
	if err := userLimiter.Allow(userID); err == nil {
		t.Errorf("AllowAction() = true, want false")
	}

	// Wait for a second to allow the rate limiter to replenish its tokens.
	time.Sleep(time.Second)

	// The action should be allowed again after the rate limiter has replenished.
	if err := userLimiter.Allow(userID); err != nil {
		t.Errorf("AllowAction() = false, want true")
	}
}

func TestEventLimiter_MultipleUsers(t *testing.T) {
	userLimiter := NewEventLimiter(1, 1)
	userID1 := "user1"
	userID2 := "user2"

	// Allow actions for the first user.
	if err := userLimiter.Allow(userID1); err != nil {
		t.Errorf("AllowAction() for %s = false, want true", userID1)
	}
	if err := userLimiter.Allow(userID1); err == nil {
		t.Errorf("AllowAction() for %s = true, want false", userID1)
	}
	if err := userLimiter.Allow(userID1); err == nil {
		t.Errorf("AllowAction() for %s = true, want false", userID1)
	}

	// Allow actions for the second user independently.
	if err := userLimiter.Allow(userID2); err != nil {
		t.Errorf("AllowAction() for %s = false, want true", userID2)
	}
	if err := userLimiter.Allow(userID2); err == nil {
		t.Errorf("AllowAction() for %s = true, want false", userID2)
	}
	if err := userLimiter.Allow(userID2); err == nil {
		t.Errorf("AllowAction() for %s = true, want false", userID2)
	}
}

func TestEventLimiter_Concurrency(t *testing.T) {
	// expecting 1 to pass and all other to fail
	// numGoroutines := 10
	// numRequests := 5

	// expecting 2 to pass and all other to fail
	// numGoroutines := 10
	// numRequests := 11

	userLimiter := NewEventLimiter(1, 1)
	userID := "testUser"
	var wg sync.WaitGroup
	numGoroutines := 10
	numRequests := 11
	results := make(chan error, numGoroutines*numRequests)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numRequests; j++ {
				allowed := userLimiter.Allow(userID)
				results <- allowed
				time.Sleep(100 * time.Millisecond) // slight delay to simulate real-world usage
			}
		}()
	}

	wg.Wait()
	close(results)

	var allowedCount, deniedCount int
	for result := range results {
		if result == nil {
			allowedCount++
		} else {
			deniedCount++
		}
	}

	t.Logf("Allowed: %d, Denied: %d", allowedCount, deniedCount)

	if allowedCount == 0 {
		t.Errorf("Expected some allowed actions, got 0")
	}
	if deniedCount == 0 {
		t.Errorf("Expected some denied actions, got 0")
	}
}

func TestEventLimiter_ConcurrencyMultipleUsers(t *testing.T) {
	// for numRequestsPerGoroutine = 11 expecting 2 allowed, bc 11 - 1 (y) + 9(no) + 1 (y)
	// for 1 > numRequestsPerGoroutine < 10 expecting 1 allowed, bc 1 sec would be timeout and other requests would be blocked
	userLimiter := NewEventLimiter(1, 1) // 1 request per second with a burst of 1
	var wg sync.WaitGroup
	numUsers := 5
	numGoroutinesPerUser := 2
	numRequestsPerGoroutine := 11
	results := make(chan struct {
		userID  string
		allowed error
	}, numUsers*numGoroutinesPerUser*numRequestsPerGoroutine)

	for u := 0; u < numUsers; u++ {
		userID := "user" + string(rune('A'+u))
		for g := 0; g < numGoroutinesPerUser; g++ {
			wg.Add(1)
			go func(userID string) {
				defer wg.Done()
				for r := 0; r < numRequestsPerGoroutine; r++ {
					allowed := userLimiter.Allow(userID)
					results <- struct {
						userID  string
						allowed error
					}{userID, allowed}
					time.Sleep(100 * time.Millisecond) // slight delay to simulate real-world usage
				}
			}(userID)
		}
	}

	wg.Wait()
	close(results)

	userResults := make(map[string]struct {
		allowed int
		denied  int
	})
	for result := range results {
		res := userResults[result.userID]
		if result.allowed == nil {
			res.allowed++
		} else {
			res.denied++
		}
		userResults[result.userID] = res
	}

	for userID, res := range userResults {
		t.Logf("User %s - Allowed: %d, Denied: %d", userID, res.allowed, res.denied)
		if res.allowed == 0 {
			t.Errorf("Expected some allowed actions for %s, got 0", userID)
		}
		if res.denied == 0 {
			t.Errorf("Expected some denied actions for %s, got 0", userID)
		}
	}
}
