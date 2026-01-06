package mobile

import (
	"bytes"
	"io"
	"log/slog"

	"fyne.io/fyne/v2"
)

type hasCameraOpen interface {
	ShowCameraOpen(callback func(base64Image string, closer func()), filename string)
}

// ShowCameraOpen open the camera intent and return the image as a base64 data.
func ShowCameraOpen(callback func(r io.Reader, err error), filename string) {
	drv, ok := fyne.CurrentApp().Driver().(*driver)
	if !ok {
		return
	}

	a, ok := drv.app.(hasCameraOpen)
	if !ok {
		return
	}

	a.ShowCameraOpen(func(base64Image string, closer func()) {
		slog.Debug("app show camera open", "base64Image_nil", base64Image == "")
		if base64Image == "" {
			callback(nil, nil)
			return
		}
		// TODO: maybe de-base64 the data here

		buf := bytes.NewBufferString(base64Image)

		slog.Debug("app show camera open: call callback")

		callback(buf, nil)
	}, filename)
}
