package middleware //nolint:testpackage

import (
	"testing"
	"time"
)

func TestBurstBucket(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name   string
		bucket func() *BurstBucket
		added  bool
	}{
		{
			name: "empty",
			bucket: func() *BurstBucket {
				return NewBurstBucket(10, time.Second)
			},
			added: true,
		},
		{
			name: "full",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(10, time.Second)
				bucket.currentBurst = 10
				bucket.lastEntry = time.Now()
				return bucket
			},
			added: false,
		},
		{
			name: "5/s - after 198ms",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(5, time.Second)
				bucket.lastEntry = time.Now().Add(-198 * time.Millisecond)
				bucket.currentBurst = 5
				return bucket
			},
			added: false,
		},
		{
			name: "5/s - after 200ms",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(5, time.Second)
				bucket.lastEntry = time.Now().Add(-200 * time.Millisecond)
				bucket.currentBurst = 5
				return bucket
			},
			added: true,
		},
		{
			name: "60/5h - after 4.9m",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(60, 5*time.Hour)
				bucket.lastEntry = time.Now().Add(-4*time.Minute - 59*time.Second)
				bucket.currentBurst = 60
				return bucket
			},
			added: false,
		},
		{
			name: "60/5h - after 5m",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(60, 5*time.Hour)
				bucket.lastEntry = time.Now().Add(-5 * time.Minute)
				bucket.currentBurst = 60
				return bucket
			},
			added: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			bucket := tc.bucket()
			got := bucket.Add()
			if got != tc.added {
				t.Errorf("Add() = %v; want %v", got, tc.added)
			}
		})
	}
}

func TestRateLimiterAdd(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		rateLimiter func() RateLimiter
		added       bool
	}{
		{
			name: "empty",
			rateLimiter: func() RateLimiter {
				return RateLimiter{
					recoveryRate: time.Second,
					maxBurst:     10,
					buckets:      make(map[string]*BurstBucket),
				}
			},
			added: true,
		},
		{
			name: "full",
			rateLimiter: func() RateLimiter {
				return RateLimiter{
					recoveryRate: time.Second,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": {
							maxBurst:     10,
							currentBurst: 10,
							lastEntry:    time.Now(),
						},
					},
				}
			},
			added: false,
		},
		{
			name: "recovered",
			rateLimiter: func() RateLimiter {
				return RateLimiter{
					recoveryRate: time.Second,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": {
							maxBurst:     10,
							currentBurst: 10,
							lastEntry:    time.Now().Add(-1 * time.Second),
						},
					},
				}
			},
			added: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rateLimiter := tc.rateLimiter()
			got := rateLimiter.Add("key")
			if got != tc.added {
				t.Errorf("Add() = %v; want %v", got, tc.added)
			}
		})
	}
}

func TestRateLimiterClean(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		rateLimiter func() RateLimiter
		length      int
	}{
		{
			name: "empty",
			rateLimiter: func() RateLimiter {
				return RateLimiter{
					recoveryRate: time.Second,
					maxBurst:     10,
					buckets:      make(map[string]*BurstBucket),
				}
			},
			length: 0,
		},
		{
			name: "with one",
			rateLimiter: func() RateLimiter {
				return RateLimiter{
					recoveryRate: time.Second,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": {
							maxBurst:     10,
							currentBurst: 10,
							lastEntry:    time.Now(),
						},
					},
				}
			},
			length: 1,
		},
		{
			name: "with one expired",
			rateLimiter: func() RateLimiter {
				return RateLimiter{
					recoveryRate: time.Second,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": {
							maxBurst:     10,
							currentBurst: 10,
							lastEntry:    time.Now().Add(-1 * time.Minute),
						},
					},
				}
			},
			length: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rateLimiter := tc.rateLimiter()
			rateLimiter.Clean()
			got := len(rateLimiter.buckets)
			if got != tc.length {
				t.Errorf("Add() = %v; want %v", got, tc.length)
			}
		})
	}
}
