package middleware

import (
	"net/http"
<<<<<<< HEAD
	"strconv"
=======
>>>>>>> dev
	"time"
)

const (
	errCapacityExceeded = "Server capacity exceeded."
	errTimedOut         = "Timed out while waiting for a pending request to complete."
	errContextCanceled  = "Context was canceled."
)

var (
	defaultBacklogTimeout = time.Second * 60
)

<<<<<<< HEAD
// ThrottleOpts represents a set of throttling options.
type ThrottleOpts struct {
	Limit          int
	BacklogLimit   int
	BacklogTimeout time.Duration
	RetryAfterFn   func(ctxDone bool) time.Duration
}

=======
>>>>>>> dev
// Throttle is a middleware that limits number of currently processed requests
// at a time across all users. Note: Throttle is not a rate-limiter per user,
// instead it just puts a ceiling on the number of currentl in-flight requests
// being processed from the point from where the Throttle middleware is mounted.
func Throttle(limit int) func(http.Handler) http.Handler {
<<<<<<< HEAD
	return ThrottleWithOpts(ThrottleOpts{Limit: limit, BacklogTimeout: defaultBacklogTimeout})
=======
	return ThrottleBacklog(limit, 0, defaultBacklogTimeout)
>>>>>>> dev
}

// ThrottleBacklog is a middleware that limits number of currently processed
// requests at a time and provides a backlog for holding a finite number of
// pending requests.
func ThrottleBacklog(limit int, backlogLimit int, backlogTimeout time.Duration) func(http.Handler) http.Handler {
<<<<<<< HEAD
	return ThrottleWithOpts(ThrottleOpts{Limit: limit, BacklogLimit: backlogLimit, BacklogTimeout: backlogTimeout})
}

// ThrottleWithOpts is a middleware that limits number of currently processed requests using passed ThrottleOpts.
func ThrottleWithOpts(opts ThrottleOpts) func(http.Handler) http.Handler {
	if opts.Limit < 1 {
		panic("chi/middleware: Throttle expects limit > 0")
	}

	if opts.BacklogLimit < 0 {
=======
	if limit < 1 {
		panic("chi/middleware: Throttle expects limit > 0")
	}

	if backlogLimit < 0 {
>>>>>>> dev
		panic("chi/middleware: Throttle expects backlogLimit to be positive")
	}

	t := throttler{
<<<<<<< HEAD
		tokens:         make(chan token, opts.Limit),
		backlogTokens:  make(chan token, opts.Limit+opts.BacklogLimit),
		backlogTimeout: opts.BacklogTimeout,
		retryAfterFn:   opts.RetryAfterFn,
	}

	// Filling tokens.
	for i := 0; i < opts.Limit+opts.BacklogLimit; i++ {
		if i < opts.Limit {
=======
		tokens:         make(chan token, limit),
		backlogTokens:  make(chan token, limit+backlogLimit),
		backlogTimeout: backlogTimeout,
	}

	// Filling tokens.
	for i := 0; i < limit+backlogLimit; i++ {
		if i < limit {
>>>>>>> dev
			t.tokens <- token{}
		}
		t.backlogTokens <- token{}
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			select {

			case <-ctx.Done():
<<<<<<< HEAD
				t.setRetryAfterHeaderIfNeeded(w, true)
=======
>>>>>>> dev
				http.Error(w, errContextCanceled, http.StatusServiceUnavailable)
				return

			case btok := <-t.backlogTokens:
				timer := time.NewTimer(t.backlogTimeout)

				defer func() {
					t.backlogTokens <- btok
				}()

				select {
				case <-timer.C:
<<<<<<< HEAD
					t.setRetryAfterHeaderIfNeeded(w, false)
=======
>>>>>>> dev
					http.Error(w, errTimedOut, http.StatusServiceUnavailable)
					return
				case <-ctx.Done():
					timer.Stop()
<<<<<<< HEAD
					t.setRetryAfterHeaderIfNeeded(w, true)
=======
>>>>>>> dev
					http.Error(w, errContextCanceled, http.StatusServiceUnavailable)
					return
				case tok := <-t.tokens:
					defer func() {
						timer.Stop()
						t.tokens <- tok
					}()
					next.ServeHTTP(w, r)
				}
				return

			default:
<<<<<<< HEAD
				t.setRetryAfterHeaderIfNeeded(w, false)
=======
>>>>>>> dev
				http.Error(w, errCapacityExceeded, http.StatusServiceUnavailable)
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}

// token represents a request that is being processed.
type token struct{}

// throttler limits number of currently processed requests at a time.
type throttler struct {
	tokens         chan token
	backlogTokens  chan token
	backlogTimeout time.Duration
<<<<<<< HEAD
	retryAfterFn   func(ctxDone bool) time.Duration
}

// setRetryAfterHeaderIfNeeded sets Retry-After HTTP header if corresponding retryAfterFn option of throttler is initialized.
func (t throttler) setRetryAfterHeaderIfNeeded(w http.ResponseWriter, ctxDone bool) {
	if t.retryAfterFn == nil {
		return
	}
	w.Header().Set("Retry-After", strconv.Itoa(int(t.retryAfterFn(ctxDone).Seconds())))
=======
>>>>>>> dev
}
