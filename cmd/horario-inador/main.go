package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gmartineza/horario-inador/internal/models"
	"github.com/gmartineza/horario-inador/internal/processor"
)

func main() {
	// Command line flags
	inputFile := flag.String("input", "", "Input JSON schedule file")
	outputFile := flag.String("output", "schedule.png", "Output PNG file")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Please provide an input file with -input flag")
		os.Exit(1)
	}

	// Read and parse JSON file
	data, err := os.ReadFile(*inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	var schedule models.Schedule
	if err := json.Unmarshal(data, &schedule); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Generate image with default config
	config := processor.DefaultConfig()
	img, err := processor.GenerateImage(schedule, config)
	if err != nil {
		fmt.Printf("Error generating image: %v\n", err)
		os.Exit(1)
	}

	// Save image
	if err := processor.SavePNG(img, *outputFile); err != nil {
		fmt.Printf("Error saving image: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Schedule wallpaper generated: %s\n", *outputFile)
}
