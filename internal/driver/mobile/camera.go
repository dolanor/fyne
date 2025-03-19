package mobile

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

type hasCameraOpen interface {
	ShowCameraOpen(callback func(path string, closer func()), filename string)
}

// ShowCameraOpen loads the native file save dialog and returns the chosen file path via the callback func.
func ShowCameraOpen(callback func(fyne.URIWriteCloser, error), filename string) {
	drv := fyne.CurrentApp().Driver().(*driver)
	if a, ok := drv.app.(hasCameraOpen); ok {
		a.ShowCameraOpen(func(path string, closer func()) {
			if path == "" {
				callback(nil, nil)
				return
			}

			uri, err := storage.ParseURI(path)
			if err != nil {
				callback(nil, err)
				return
			}

			f, err := fileWriterForURI(uri)
			if f != nil {
				f.(*fileSave).done = closer
			}
			callback(f, err)
		}, filename)
	}
}
