package mobile

import (
	"errors"
	"log"

	"fyne.io/fyne/v2"
)

type hasGetCurrentLocation interface {
	GetCurrentLocation() (lat, lon float64, err error)
}

func GetCurrentLocation() (lat, lon float64, err error) {
	drv, ok := fyne.CurrentApp().Driver().(*driver)
	if !ok {
		return 0, 0, errors.New("the app has not a mobile driver")
	}
	log.Printf("driver: %T", drv)

	log.Printf("hasGetCurrentLocation: %T", drv.app)
	a, ok := drv.app.(hasGetCurrentLocation)
	if !ok {
		return 0, 0, errors.New("the app does not handle location")
	}

	lat, lon, err = a.GetCurrentLocation()

	return lat, lon, err
}

type hasStartUpdatingLocation interface {
	StartUpdatingLocation() error
}

func StartUpdatingLocation() error {
	drv, ok := fyne.CurrentApp().Driver().(*driver)
	if !ok {
		return errors.New("the app has not a mobile driver")
	}
	log.Printf("driver: %T", drv)

	log.Printf("hasStartUpdatingLocation: %T", drv.app)
	a, ok := drv.app.(hasStartUpdatingLocation)
	if !ok {
		return errors.New("the app does not handle location")
	}

	a.StartUpdatingLocation()

	return nil
}
