package backlight

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

var (
	acpiBacklight = regexp.MustCompile(`acpi_video[0-9]+$`)
)

const backlightsDirectory = "/sys/class/backlight"

// Backlight defines a backlight, being the maximum brightness, and functions for getting and setting brightness
type Backlight struct {
	MaxBrightness int
	SetBrightness func(string)
	GetBrightness func() string
	IsOn          func() bool
}

type backlight struct {
	identifier string
	fallback   bool
}

// PopulateBacklights finds and generates functions for all system backlights
func PopulateBacklights() (BacklightOutput []Backlight) {

	backlights, err := identifyBacklights()

	if err != nil {
		fmt.Println("Error identifying backlights")
		fmt.Println(err)
	}

	for _, backlight := range backlights {
		back := Backlight{}
		back.MaxBrightness, _ = backlight.getMaxBrightness()
		back.SetBrightness = backlight.setBrightness
		back.GetBrightness = backlight.getBrightness
		back.IsOn = backlight.IsOn

		BacklightOutput = append(BacklightOutput, back)
	}

	return BacklightOutput
}

func (b *backlight) currentBrightnessPath() string {
	return filepath.Join(backlightsDirectory, b.identifier, "brightness")
}

func (b *backlight) maxBrightnessPath() string {
	return filepath.Join(backlightsDirectory, b.identifier, "max_brightness")
}

func (b *backlight) dpmsPath() string {
	return filepath.Join(backlightsDirectory, b.identifier, "device/dpms")
}

func (b *backlight) setBrightness(brightness string) {
	f, err := os.OpenFile(b.currentBrightnessPath(), os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(brightness)
	if err != nil {
		log.Println(err)
	}
	return
}

func (b *backlight) getBrightness() string {
	v, err := ioutil.ReadFile(b.currentBrightnessPath())
	if err != nil {
		fmt.Printf("failed to read current brightness of %s: %s", b.identifier, err)
		return ""
	}
	val := string(bytes.TrimSpace(v))
	if err != nil {
		fmt.Printf("failed to parse current brightness of %s: %s", b.identifier, err)
		return ""
	}
	return val
}

func (b *backlight) getMaxBrightness() (int, error) {
	v, err := ioutil.ReadFile(b.maxBrightnessPath())
	if err != nil {
		return 0, fmt.Errorf("failed to read maximum brightness of %s: %s", b.identifier, err)
	}
	val, err := strconv.Atoi(string(bytes.TrimSpace(v)))
	if err != nil {
		return 0, fmt.Errorf("failed to parse maximum brightness of %s: %s", b.identifier, err)
	}
	return val, err
}

func (b *backlight) IsOn() bool {
	v, err := ioutil.ReadFile(b.dpmsPath())
	if err != nil {
		log.Println("failed to read DPMS")
		return false
	}
	if string(v) == "On\n" {
		return true
	}
	return false
}

func MinredToKelvin(minred float64) float64 {
	return 1000000.0 / minred
}

func KelvinToMinred(kelvin float64) float64 {
	return 1000000.0 / kelvin
}

func identifyBacklights() ([]*backlight, error) {
	var backlights []*backlight

	// Locate backlight interfaces. Issues are commonly reported when the ACPI
	// backlight interface is used instead of other interfaces present on the
	// system. As a result, ACPI interfaces are only used when no other
	// interfaces are available.
	err := filepath.Walk(backlightsDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if path == backlightsDirectory {
			return nil
		}

		backlight := &backlight{
			identifier: filepath.Base(path),
			fallback:   acpiBacklight.MatchString(path),
		}
		backlights = append(backlights, backlight)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list backlight interfaces: %s", err)
	}

	return backlights, nil
}
