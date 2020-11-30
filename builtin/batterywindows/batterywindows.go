package batterywindows

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetLevel() string {

	value, err := exec.Command("WMIC", "PATH", "Win32_Battery", "Get", "EstimatedChargeRemaining").Output()
	if err != nil {
		log.Printf("%s", err)
	}

	lines := strings.Split(string(value), "\n")

	return lines[1]
}

func IsPluggedIn() bool {

	// WMIC PATH Win32_Battery Get BatteryStatus
	// For status
	// Other (1)
	// The battery is discharging.
	// Unknown (2)
	// The system has access to AC so no battery is being discharged. However, the battery is not necessarily charging.
	// Fully Charged (3)
	// Low (4)
	// Critical (5)
	// Charging (6)
	// Charging and High (7)
	// Charging and Low (8)
	// Charging and Critical (9)
	// Undefined (10)
	// Partially Charged (11)

	value, err := exec.Command("WMIC", "PATH", "Win32_Battery", "Get", "BatteryStatus").Output()
	if err != nil {
		log.Printf("%s", err)
	}

	lines := strings.Split(string(value), "\n")

	status, err := strconv.Atoi(strings.TrimSpace(lines[1]))
	if err != nil {
		log.Printf("%s", err)
	}

	switch status {
	case 1, 10:
		return false
	default:
		return true
	}

}
