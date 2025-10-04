package location

import (
	"log"
	"log/slog"

	"fyne.io/fyne/v2/internal/driver/mobile"
)

func GetCurrent() (lat, lon float64, err error) {
	slog.Debug("location.GetCurrent")
	log.Println("LOL location.GetCurrent")

	lat, lon, err = mobile.GetCurrentLocation()
	return lat, lon, err
}
