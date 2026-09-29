package app

import (
	"testing"
	"time"
)

func TestGetBestLocation(t *testing.T) {
	locs := []location{
		{
			Accuracy:                      24,
			Latitude:                      48.58385441549495,
			LocationMonotonicElapsedNanos: 396638846287797,
			Longitude:                     7.745837938701012,
			MonotonicElapsedNanos:         396643954254408,
			Provider:                      "passive",
		},
		{
			Accuracy:                      100,
			Latitude:                      48.5837005,
			LocationMonotonicElapsedNanos: 396597031858750,
			Longitude:                     7.7450963,
			MonotonicElapsedNanos:         396643954254408,
			Provider:                      "network",
		},
		{
			Accuracy:                      36,
			Latitude:                      48.5838544,
			LocationMonotonicElapsedNanos: 396638846287797,
			Longitude:                     7.7458379,
			MonotonicElapsedNanos:         396643954254408,
			Provider:                      "fused",
		},
		{
			Accuracy:                      24,
			Latitude:                      48.58385441549495,
			LocationMonotonicElapsedNanos: 396638846287797,
			Longitude:                     7.745837938701012,
			MonotonicElapsedNanos:         396643954254408,
			Provider:                      "gps",
		},
	}

	//FIXME
	want := location{
		Accuracy:                      100,
		Latitude:                      48.5837005,
		LocationMonotonicElapsedNanos: 396597031858750,
		Longitude:                     7.7450963,
		MonotonicElapsedNanos:         396643954254408,
		Provider:                      "network",
	}

	oldBest := location{
		Accuracy:                      24,
		Latitude:                      48.5838544,
		LocationMonotonicElapsedNanos: 396638846287797,
		Longitude:                     7.7458379,
		MonotonicElapsedNanos:         396643954254408,
		Provider:                      "fused",
	}
	now := time.Unix(0, 396638846287797)
	got := getBestLocation(locs, oldBest, now)
	if got != want {
		t.Fatalf("wrong location chosen\n\twant: %+v\n\t got: %+v", want, got)
	}
}
