package speedbump

import (
	"strconv"
	"time"

	"github.com/benbjohnson/clock"
)

// PerSecondHasher generates hashes per second. This means you can keep track
// of N request per second.
type PerSecondHasher struct {
	// Clock is the time reference that will be used by the hasher. If it is
	// not provided, the hashing function will use the default time. This can
	// be replaced with a mock clock object for testing.
	Clock clock.Clock
}

// Hash generates the hash for the current period and client.
func (h PerSecondHasher) Hash(id string) string {
	if h.Clock == nil {
		h.Clock = clock.New()
	}

	return id + ":" + strconv.FormatInt(h.Clock.Now().Unix(), 10)
}

// Duration gets the duration of each period.
func (h PerSecondHasher) Duration() time.Duration {
	return time.Second
}

// PerMinuteHasher generates hashes per minute. This means you can keep track
// of N request per minute.
type PerMinuteHasher struct {
	Clock clock.Clock
}

// Hash generates the hash for the current period and client.
func (h PerMinuteHasher) Hash(id string) string {
	if h.Clock == nil {
		h.Clock = clock.New()
	}

	return id + ":" + h.Clock.Now().Format("2006-01-02T15:04")
}

// Duration gets the duration of each period.
func (h PerMinuteHasher) Duration() time.Duration {
	return time.Minute
}

// PerHourHasher generates hashes per hour. This means you can keep track
// of N request per hour.
type PerHourHasher struct {
	Clock clock.Clock
}

// Hash generates the hash for the current period and client.
func (h PerHourHasher) Hash(id string) string {
	if h.Clock == nil {
		h.Clock = clock.New()
	}

	return id + ":" + h.Clock.Now().Format("2006-01-02T15")
}

// Duration gets the duration of each period.
func (h PerHourHasher) Duration() time.Duration {
	return time.Hour
}
