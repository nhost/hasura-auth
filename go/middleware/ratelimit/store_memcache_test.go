package ratelimit_test

import (
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/nhost/hasura-auth/go/middleware/ratelimit"
)

func TestNewMemcacheStore(t *testing.T) {
	t.Parallel()

	cl := memcache.New("localhost:11211")
	store := ratelimit.NewMemcacheStore(cl)

	if e := store.Get("key"); e != 0 {
		t.Errorf("Expected 0, got %d", e)
	}

	if e := store.Increment("key", time.Second); e != 1 {
		t.Errorf("Expected 1, got %d", e)
	}

	if e := store.Get("key"); e != 1 {
		t.Errorf("Expected 1, got %d", e)
	}

	if e := store.Increment("key", time.Second); e != 2 {
		t.Errorf("Expected 1, got %d", e)
	}

	if e := store.Get("key"); e != 2 {
		t.Errorf("Expected 1, got %d", e)
	}

	time.Sleep(time.Second)
	if e := store.Get("key"); e != 0 {
		t.Errorf("Expected 0, got %d", e)
	}
}

func TestRateLimitMemcache(t *testing.T) {
	t.Parallel()

	store := ratelimit.NewInMemoryStore()

	rl := ratelimit.NewSlidingWindow("test", 5, time.Second, store)
	now := time.Now()
	t.Log(1, rl.Allow("key"), time.Since(now))
	t.Log(2, rl.Allow("key"), time.Since(now))
	t.Log(3, rl.Allow("key"), time.Since(now))
	t.Log(4, rl.Allow("key"), time.Since(now))
	t.Log(5, rl.Allow("key"), time.Since(now))
	t.Log(6, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(7, rl.Allow("key"), time.Since(now))
	t.Log(7, rl.Allow("key"), time.Since(now))
	t.Log(7, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(8, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(9, rl.Allow("key"), time.Since(now))
	t.Log(9, rl.Allow("key"), time.Since(now))
	t.Log(9, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(10, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(11, rl.Allow("key"), time.Since(now))
	t.Log(11, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(12, rl.Allow("key"), time.Since(now))
	t.Log(12, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(13, rl.Allow("key"), time.Since(now))
	t.Log(13, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(14, rl.Allow("key"), time.Since(now))
	t.Log(14, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(15, rl.Allow("key"), time.Since(now))
	t.Log(15, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(7, rl.Allow("key"), time.Since(now))
	t.Log(7, rl.Allow("key"), time.Since(now))
	t.Log(7, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(8, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(9, rl.Allow("key"), time.Since(now))
	t.Log(9, rl.Allow("key"), time.Since(now))
	t.Log(9, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(10, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(11, rl.Allow("key"), time.Since(now))
	t.Log(11, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(12, rl.Allow("key"), time.Since(now))
	t.Log(12, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(13, rl.Allow("key"), time.Since(now))
	t.Log(13, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(14, rl.Allow("key"), time.Since(now))
	t.Log(14, rl.Allow("key"), time.Since(now))
	time.Sleep(time.Second / 5)
	t.Log(15, rl.Allow("key"), time.Since(now))
	t.Log(15, rl.Allow("key"), time.Since(now))
}
