package processor

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gmartineza/horario-inador/internal/models"
)

// ImageConfig holds the configuration for image generation
type ImageConfig struct {
	Width      int
	Height     int
	Padding    int
	FontSize   float64
	LineHeight float64
}

// DefaultConfig returns a default configuration for phone wallpapers
func DefaultConfig() ImageConfig {
	return ImageConfig{
		Width:      1080,
		Height:     1920,
		Padding:    40,
		FontSize:   20,
		LineHeight: 30,
	}
}

// GenerateImage creates a new schedule image
func GenerateImage(schedule models.Schedule, config ImageConfig) (*image.RGBA, error) {
	// Create a new RGBA image
	img := image.NewRGBA(image.Rect(0, 0, config.Width, config.Height))

	// Fill with white background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Initialize font
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %v", err)
	}

	// Create font context
	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFont(font)
	ctx.SetFontSize(config.FontSize)
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	ctx.SetSrc(image.NewUniform(color.Black))

	// Start drawing position
	x := config.Padding
	y := config.Padding + int(config.FontSize)

	// Draw title
	y += drawText(ctx, "Horario", x, y, config)
	y += int(config.LineHeight)

	// Sort days of the week
	days := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes"}

	// Draw schedule by day
	for _, day := range days {
		// Draw day header
		y += drawText(ctx, fmt.Sprintf("=== %s ===", day), x, y, config)
		y += int(config.LineHeight / 2)

		// Find classes for this day
		for _, course := range schedule {
			for _, session := range course.Schedule {
				if session.Day == day {
					// Draw course info
					courseInfo := fmt.Sprintf("%s - %s", session.TimeSlot, course.Name)
					y += drawText(ctx, courseInfo, x+20, y, config)
					y += drawText(ctx, fmt.Sprintf("Aula: %s", session.Room), x+40, y, config)
					y += int(config.LineHeight)
				}
			}
		}
		y += int(config.LineHeight)
	}

	return img, nil
}

// drawText draws text at the specified position and returns the height used
func drawText(ctx *freetype.Context, text string, x, y int, config ImageConfig) int {
	pt := freetype.Pt(x, y)
	_, err := ctx.DrawString(text, pt)
	if err != nil {
		return int(ctx.PointToFixed(config.FontSize).Round())
	}
	return int(ctx.PointToFixed(config.FontSize).Round())
}

// SavePNG saves the image to a PNG file
func SavePNG(img *image.RGBA, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	return nil
}
