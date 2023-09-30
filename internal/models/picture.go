package models

import "fmt"

const MaxPictureSizeBytes = 20_000_000

const MaxPicturesPerPost = 10

type ImageType string

const (
	ImageTypeJPEG ImageType = "image/jpeg"
	ImageTypePNG  ImageType = "image/png"
	ImageTypeGIF  ImageType = "image/gif"
)

type Picture struct {
	Value string
	Type  ImageType
	Size  int
}

func StringToImageType(s string) (ImageType, error) {
	switch s {
	case "image/jpeg":
		return ImageTypeJPEG, nil
	case "image/gif":
		return ImageTypeGIF, nil
	case "image/png":
		return ImageTypePNG, nil
	default:
		return "", fmt.Errorf("unknown or unsupported image type: %v", s)
	}
}
