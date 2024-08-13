package middleware //nolint:testpackage

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBurstBucket(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name           string
		bucket         func() *BurstBucket
		added          bool
		expectedBucket *BurstBucket
	}{
		{
			name: "empty",
			bucket: func() *BurstBucket {
				return NewBurstBucket(10, time.Second, 1)
			},
			added: true,
			expectedBucket: &BurstBucket{
				replicas:     1,
				interval:     time.Second,
				maxBurst:     10,
				currentBurst: 1,
				lastEntry:    time.Time{},
			},
		},
		{
			name: "full",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(10, time.Second, 1)
				bucket.currentBurst = 10
				bucket.lastEntry = time.Now()
				return bucket
			},
			added: false,
			expectedBucket: &BurstBucket{
				replicas:     1,
				interval:     time.Second,
				maxBurst:     10,
				currentBurst: 10,
				lastEntry:    time.Time{},
			},
		},
		{
			name: "5/s - after 198ms",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(5, time.Second, 1)
				bucket.lastEntry = time.Now().Add(-198 * time.Millisecond)
				bucket.currentBurst = 5
				return bucket
			},
			added: false,
			expectedBucket: &BurstBucket{
				replicas:     1,
				interval:     time.Second,
				maxBurst:     5,
				currentBurst: 5,
				lastEntry:    time.Time{},
			},
		},
		{
			name: "5/s - after 200ms",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(5, time.Second, 1)
				bucket.lastEntry = time.Now().Add(-200 * time.Millisecond)
				bucket.currentBurst = 5
				return bucket
			},
			added: true,
			expectedBucket: &BurstBucket{
				replicas:     1,
				interval:     time.Second,
				maxBurst:     5,
				currentBurst: 5,
				lastEntry:    time.Time{},
			},
		},
		{
			name: "60/5h - after 4.9m",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(60, 5*time.Hour, 1)
				bucket.lastEntry = time.Now().Add(-4*time.Minute - 59*time.Second)
				bucket.currentBurst = 60
				return bucket
			},
			expectedBucket: &BurstBucket{
				replicas:     1,
				interval:     5 * time.Hour,
				maxBurst:     60,
				currentBurst: 60,
				lastEntry:    time.Time{},
			},
			added: false,
		},
		{
			name: "60/5h - after 5m",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(60, 5*time.Hour, 1)
				bucket.lastEntry = time.Now().Add(-5 * time.Minute)
				bucket.currentBurst = 60
				return bucket
			},
			added: true,
			expectedBucket: &BurstBucket{
				replicas:     1,
				interval:     5 * time.Hour,
				maxBurst:     60,
				currentBurst: 60,
				lastEntry:    time.Time{},
			},
		},
		{
			name: "full from 1 to 2 replicas",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(10, time.Second, 1)
				bucket.currentBurst = 10
				bucket.lastEntry = time.Now()

				bucket.SetReplicas(2)

				return bucket
			},
			added: false,
			expectedBucket: &BurstBucket{
				replicas:     2,
				interval:     time.Second,
				maxBurst:     10,
				currentBurst: 5,
				lastEntry:    time.Time{},
			},
		},
		{
			name: "full from 2 to 1 replicas",
			bucket: func() *BurstBucket {
				bucket := NewBurstBucket(10, time.Second, 2)
				bucket.currentBurst = 5
				bucket.lastEntry = time.Now()

				bucket.SetReplicas(1)

				return bucket
			},
			added: false,
			expectedBucket: &BurstBucket{
				replicas:     1,
				interval:     time.Second,
				maxBurst:     10,
				currentBurst: 10,
				lastEntry:    time.Time{},
			},
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

			//nolint:exhaustruct
			opts := []cmp.Option{
				cmp.AllowUnexported(BurstBucket{}),
				cmpopts.IgnoreFields(BurstBucket{}, "lastEntry"),
			}
			if diff := cmp.Diff(tc.expectedBucket, bucket, opts...); diff != "" {
				t.Errorf("Add() mismatch (-want +got):\n%s", diff)
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
					replicas:     1,
					recoveryRate: time.Minute,
					maxBurst:     10,
					buckets:      make(map[string]*BurstBucket),
				}
			},
			added: true,
		},
		{
			name: "full",
			rateLimiter: func() RateLimiter {
				bucket := NewBurstBucket(10, time.Minute, 1)
				bucket.lastEntry = time.Now()
				bucket.currentBurst = 10

				return RateLimiter{
					replicas:     1,
					recoveryRate: time.Minute,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": bucket,
					},
				}
			},
			added: false,
		},
		{
			name: "recovered",
			rateLimiter: func() RateLimiter {
				bucket := NewBurstBucket(10, time.Minute, 1)
				bucket.lastEntry = time.Now().Add(-1 * time.Minute)
				bucket.currentBurst = 10

				return RateLimiter{
					replicas:     1,
					recoveryRate: time.Minute,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": bucket,
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

func TestRateLimiterClean(t *testing.T) { //nolint:tparallel,paralleltest
	cases := []struct {
		name        string
		rateLimiter func() RateLimiter
		length      int
	}{
		{
			name: "empty",
			rateLimiter: func() RateLimiter {
				return RateLimiter{
					replicas:     1,
					recoveryRate: time.Minute,
					maxBurst:     10,
					buckets:      make(map[string]*BurstBucket),
				}
			},
			length: 0,
		},
		{
			name: "with one",
			rateLimiter: func() RateLimiter {
				bucket := NewBurstBucket(10, time.Minute, 1)
				bucket.lastEntry = time.Now()
				bucket.currentBurst = 10

				return RateLimiter{
					replicas:     1,
					recoveryRate: time.Minute,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": bucket,
					},
				}
			},
			length: 1,
		},
		{
			name: "with one expired",
			rateLimiter: func() RateLimiter {
				bucket := NewBurstBucket(10, time.Minute, 1)
				bucket.lastEntry = time.Now().Add(-1 * time.Minute)
				bucket.currentBurst = 10

				return RateLimiter{
					replicas:     1,
					recoveryRate: time.Minute,
					maxBurst:     10,
					buckets: map[string]*BurstBucket{
						"key": bucket,
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
				t.Errorf("len() = %v; want %v", got, tc.length)
			}
		})
	}
}
