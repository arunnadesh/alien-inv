package models

import "fmt"

// InvasionError custom error
type InvasionError struct {
	Reason string // Error reason
}

const (
	ErrAlienDeadOrTrapped = "alien-dead-or-trapped"
	ErrWorldDestroyed     = "world-destroyed"
	ErrNoAttack           = "no-attacks-made"
)

// Error string
func (err *InvasionError) Error() string {
	return fmt.Sprintf("Invasion error with reason: %s", err.Reason)
}
