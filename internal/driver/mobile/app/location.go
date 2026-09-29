package app

import (
	"log/slog"
	"time"
)

type location struct {
	Provider  string
	Latitude  float64
	Longitude float64

	LocationMonotonicElapsedNanos int64
	MonotonicElapsedNanos         int64
	Accuracy                      float64

	SystemElapsed int64
	NowElapsed    int64
}

func getBestLocation(locations []location, oldBest location, now time.Time) location {

	isStale := func(fixMonotonicElapsedNanos, previousFixMonotonicElapsedNanos int64) bool {
		const staleFix = 2 * time.Minute

		diff := fixMonotonicElapsedNanos - previousFixMonotonicElapsedNanos //fixTimestamp.Before(now.Add(-staleFix))
		slog.Info("isStale", "new", fixMonotonicElapsedNanos, "previous", previousFixMonotonicElapsedNanos)
		return diff > staleFix.Nanoseconds()
	}

	best := oldBest

	for _, loc := range locations {
		stale := isStale(loc.LocationMonotonicElapsedNanos, best.LocationMonotonicElapsedNanos)
		if stale {
			best = loc
		}

		betterAccuracy := best.Accuracy < loc.Accuracy
		if betterAccuracy {
			best = loc
		}

		slog.Info("location", "provider", loc.Provider, "acc", loc.Accuracy, "stale", stale, "better_acc", betterAccuracy, "now", now, "bestNow", best.NowElapsed)
	}

	return best
}
