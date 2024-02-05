package generator

type Generator interface {
	Trigger() chan TriggerReason
	Output() chan string
}

type TriggerReason string // For logging, we can use this from a frontend to trigger manual updates to certain sensors
