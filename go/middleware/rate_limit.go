package middleware

import (
	"math"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	replicas = 1
)

type BurstBucket struct {
	replicas     int64 // for future use
	interval     time.Duration
	maxBurst     int64
	currentBurst int64
	lastEntry    time.Time
}

func NewBurstBucket(maxBurst int64, interval time.Duration, replicas int64) *BurstBucket {
	return &BurstBucket{
		replicas:     replicas,
		interval:     interval,
		maxBurst:     maxBurst,
		currentBurst: 0,
		lastEntry:    time.Time{},
	}
}

func (b *BurstBucket) SetReplicas(replicas int64) {
	b.currentBurst = int64(
		math.Floor(float64(b.currentBurst) * (float64(b.replicas) / float64(replicas))))
	b.replicas = replicas
}

func (b *BurstBucket) MaxBurst() int64 {
	return b.maxBurst / b.replicas
}

func (b *BurstBucket) Recoverable() int64 {
	elapsedMs := float64(time.Since(b.lastEntry).Milliseconds())
	recoveryRateMs := float64(b.interval.Milliseconds()) / float64(b.MaxBurst())
	return int64(elapsedMs / recoveryRateMs)
}

func (b *BurstBucket) Clean() {
	currentBurst := b.currentBurst - b.Recoverable()

	if currentBurst < 0 {
		b.currentBurst = 0
	} else {
		b.currentBurst = currentBurst
	}
}

func (b *BurstBucket) CurrentBurst() int64 {
	b.Clean()
	return b.currentBurst
}

func (b *BurstBucket) Add() bool {
	b.Clean()

	// if we are at the limit, we can't add more
	if b.currentBurst >= b.MaxBurst() {
		return false
	}

	// we add an element and update the last time
	b.currentBurst++
	b.lastEntry = time.Now()
	return true
}

type RateLimiter struct {
	maxBurst     int64
	recoveryRate time.Duration
	replicas     int64
	buckets      map[string]*BurstBucket
}

func NewRateLimiter(maxBurst int64, recoveryRate time.Duration, replicas int64) *RateLimiter {
	return &RateLimiter{
		replicas:     replicas,
		maxBurst:     maxBurst,
		recoveryRate: recoveryRate,
		buckets:      make(map[string]*BurstBucket),
	}
}

func (r *RateLimiter) SetReplicas(replicas int64) {
	r.replicas = replicas

	for _, bucket := range r.buckets {
		bucket.SetReplicas(replicas)
	}
}

func (r *RateLimiter) Clean() {
	for key, bucket := range r.buckets {
		if bucket.CurrentBurst() == 0 {
			delete(r.buckets, key)
		}
	}
}

func (r *RateLimiter) Add(key string) bool {
	bucket, ok := r.buckets[key]
	if !ok {
		bucket = NewBurstBucket(r.maxBurst, r.recoveryRate, r.replicas)
		r.buckets[key] = bucket
	}

	return bucket.Add()
}

// endpints that send emails.
func sendsEmail(path string, verifyEmailEnabled bool) bool {
	if verifyEmailEnabled {
		return slices.Contains([]string{
			"/signin/passwordless/email",
			"/user/email/change",
			"/user/email/send-verification-email",
			"/user/password/reset",
			"/signup/email-password",
			"/user/deanonymize",
		}, path)
	}
	return slices.Contains([]string{
		"/signin/passwordless/email",
		"/user/email/change",
		"/user/password/reset",
	}, path)
}

// endpoints that send SMS.
func sendsSMS(path string) bool {
	return slices.Contains([]string{
		"/signin/passwordless/sms",
	}, path)
}

// endpnits that can be brute forced.
func bruteForceProtected(path string) bool {
	return strings.HasPrefix(path, "/signin") ||
		strings.HasSuffix(path, "/verify") ||
		strings.HasSuffix(path, "/otp")
}

// signups.
func isSignup(path string) bool {
	return strings.HasPrefix(path, "/signup")
}

type RateLimiterConfiguration struct {
	IgnorePrefix string

	GlobalLimit    int64
	GlobalInterval time.Duration

	EmailLimit         int64
	EmailInterval      time.Duration
	EmailIsGlobal      bool
	EmailVerifyEnabled bool

	SMSLimit    int64
	SMSInterval time.Duration

	BruteForceLimit    int64
	BruteForceInterval time.Duration

	SignupsLimit    int64
	SignupsInterval time.Duration
}

func RateLimit( //nolint:cyclop,funlen,gocognit
	ignorePrefix string,
	globalLimit int64,
	globalInterval time.Duration,
	emailLimit int64,
	emailInterval time.Duration,
	emailIsGlobal bool,
	emailVerifyEnabled bool,
	smsLimit int64,
	smsInterval time.Duration,
	bruteForceLimit int64,
	bruteForceInterval time.Duration,
	signupsLimit int64,
	signupsInterval time.Duration,
) gin.HandlerFunc {
	lastClean := time.Now()

	perUserRL := NewRateLimiter(globalLimit, globalInterval, replicas)

	var globalEmailRL *BurstBucket
	var perUserEmailRL *RateLimiter
	if emailIsGlobal {
		globalEmailRL = NewBurstBucket(emailLimit, emailInterval, replicas)
	} else {
		perUserEmailRL = NewRateLimiter(emailLimit, emailInterval, replicas)
	}

	globalSMSRL := NewRateLimiter(smsLimit, smsInterval, replicas)
	perUserBruteForceRL := NewRateLimiter(bruteForceLimit, bruteForceInterval, replicas)
	perUserSignupsRL := NewBurstBucket(signupsLimit, signupsInterval, replicas)

	return func(ctx *gin.Context) {
		if time.Since(lastClean) > 1*time.Minute {
			perUserRL.Clean()
			if globalEmailRL != nil {
				globalEmailRL.Clean()
			}
			if perUserEmailRL != nil {
				perUserEmailRL.Clean()
			}
			globalSMSRL.Clean()
			perUserBruteForceRL.Clean()
			perUserSignupsRL.Clean()

			lastClean = time.Now()
		}

		clientIP := ctx.ClientIP()
		if !perUserRL.Add(clientIP) {
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		path := strings.TrimPrefix(ctx.Request.URL.Path, ignorePrefix)

		if sendsEmail(path, emailVerifyEnabled) {
			if globalEmailRL != nil && !globalEmailRL.Add() {
				ctx.AbortWithStatus(http.StatusTooManyRequests)
			}

			if perUserEmailRL != nil && !perUserEmailRL.Add(clientIP) {
				ctx.AbortWithStatus(http.StatusTooManyRequests)
			}
		}

		if sendsSMS(path) {
			if !globalSMSRL.Add(clientIP) {
				ctx.AbortWithStatus(http.StatusTooManyRequests)
			}
		}

		if bruteForceProtected(path) {
			if !perUserBruteForceRL.Add(clientIP) {
				ctx.AbortWithStatus(http.StatusTooManyRequests)
			}
		}

		if isSignup(path) {
			if !perUserSignupsRL.Add() {
				ctx.AbortWithStatus(http.StatusTooManyRequests)
			}
		}

		ctx.Next()
	}
}
