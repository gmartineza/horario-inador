package processor

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/gmartineza/horario-inador/internal/models"
)

// ImageConfig holds the configuration for image generation
type ImageConfig struct {
	Width  int
	Height int
}

// GenerateImage creates a new schedule image
func GenerateImage(schedule *models.Schedule, config ImageConfig) (*image.RGBA, error) {
	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, config.Width, config.Height))

	// Fill with white background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// TODO: Add schedule text rendering here in next step

	return img, nil
}

// SavePNG saves the image to a PNG file
func SavePNG(img *image.RGBA, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
