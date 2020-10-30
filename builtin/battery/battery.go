package battery

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
	acpiBattery = regexp.MustCompile(`BAT[0-9]+$`)
)

const batteryDirectory = "/sys/class/power_supply"

type Battery struct {
	MaxCapacity int
	GetCharge   func() string
	IsPluggedIn func() bool
}

type battery struct {
	identifier string
}

// PopulateBatteries finds and generates functions for all system batteries
func PopulateBatteries() (BatteryOutput []Battery) {

	batteries, err := identifyBatteries()

	if err != nil {
		fmt.Println("Error identifying batteries")
		fmt.Println(err)
	}

	for _, battery := range batteries {
		bat := Battery{}
		bat.MaxCapacity, _ = battery.getMaxCapacity()
		bat.GetCharge = battery.getCharge
		bat.IsPluggedIn = battery.IsPluggedIn

		BatteryOutput = append(BatteryOutput, bat)
	}

	return BatteryOutput
}

func (b *battery) currentChargePath() string {
	return filepath.Join(batteryDirectory, b.identifier, "charge_now")
}

func (b *battery) maxCapacityPath() string {
	return filepath.Join(batteryDirectory, b.identifier, "charge_full")
}

func (b *battery) statusPath() string {
	return filepath.Join(batteryDirectory, b.identifier, "status")
}

func (b *battery) getCharge() string {
	v, err := ioutil.ReadFile(b.currentChargePath())
	if err != nil {
		fmt.Printf("failed to read current charge of %s: %s", b.identifier, err)
		return ""
	}
	val := string(bytes.TrimSpace(v))
	if err != nil {
		fmt.Printf("failed to parse current charge of %s: %s", b.identifier, err)
		return ""
	}
	return val
}

func (b *battery) getMaxCapacity() (int, error) {
	v, err := ioutil.ReadFile(b.maxCapacityPath())
	if err != nil {
		return 0, fmt.Errorf("failed to read maximum capacity of %s: %s", b.identifier, err)
	}
	val, err := strconv.Atoi(string(bytes.TrimSpace(v)))
	if err != nil {
		return 0, fmt.Errorf("failed to parse maximum capacity of %s: %s", b.identifier, err)
	}
	return val, err
}

func (b *battery) IsPluggedIn() bool {
	v, err := ioutil.ReadFile(b.statusPath())
	if err != nil {
		log.Println("failed to read status")
		return false
	}
	if string(v) == "Discharging\n" {
		return false
	}
	return true
}

func identifyBatteries() ([]*battery, error) {
	var batteries []*battery

	// Locate backlight interfaces. Issues are commonly reported when the ACPI
	// backlight interface is used instead of other interfaces present on the
	// system. As a result, ACPI interfaces are only used when no other
	// interfaces are available.
	err := filepath.Walk(batteryDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if path == batteryDirectory {
			return nil
		}

		if acpiBattery.MatchString(path) {
			battery := &battery{
				identifier: filepath.Base(path),
			}
			batteries = append(batteries, battery)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list battery interfaces: %s", err)
	}

	return batteries, nil
}
