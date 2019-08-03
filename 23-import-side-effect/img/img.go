package img

import (
	"image"
	"image/jpeg"
	_ "image/png" // for png decode
	"log"
	"os"
)

// LoadImage from disk file
func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	image, format, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	log.Printf("%s is a %s image\n", path, format)
	return image, err
}

// ToJpeg converts any image to jpeg format
func ToJpeg(path string, image image.Image) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(
		out,
		image,
		&jpeg.Options{Quality: 30},
	)
}
