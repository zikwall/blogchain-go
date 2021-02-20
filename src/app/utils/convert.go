package utils

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
)

func ToPng(imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	var img image.Image
	var err error

	switch contentType {
	case "image/png":
	case "image/gif":
		img, err = gif.Decode(bytes.NewReader(imageBytes))
	case "image/jpeg":
		img, err = jpeg.Decode(bytes.NewReader(imageBytes))
	}

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	if err := png.Encode(buf, img); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
