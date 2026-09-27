package execution

import (
	"errors"
)

// Atomic execution (basic/single) execution unit.
type AtomicExecution struct {
	name      string
	action    func() error
	condition bool
}

// The bundled execution unit.
type execution struct {
	identifier    string
	atomics       []AtomicExecution
	defalutAction func() error
}

// Create new atomic execution without validation.
func UnsafeAtomicExecution(n string, a func() error, c bool) AtomicExecution {
	return AtomicExecution{
		name:      n,
		action:    a,
		condition: c,
	}
}

// Create a new atomic execution.
func NewAtomicExecution(n string, a func() error, c bool) (AtomicExecution, error) {
	if len(n) == 0 || a == nil {
		return AtomicExecution{}, errors.New("Failed to create a new atomic execution.")
	}
	return UnsafeAtomicExecution(n, a, c), nil
}

// Create a new bundled execution without validation.
func UnsafeExecution(id string, as []AtomicExecution, da func() error) execution {
	return execution{
		identifier:    id,
		atomics:       as,
		defalutAction: da,
	}
}

// Create new bundled execution.
func NewExecution(id string, as []AtomicExecution, da func() error) (execution, error) {
	if len(id) == 0 || len(as) == 0 || da == nil {
		return execution{}, errors.New("Failed to create a new execution, command couldn't be created.")
	}
	return UnsafeExecution(id, as, da), nil
}

// Run single atomic execution instance.
func (e AtomicExecution) Run() (bool, error) {
	if e.condition {
		// Only trigger error if action have err to return.
		err := e.action()
		if err != nil {
			// Return with false if action performed but there was an error.
			return false, err
		}
		// Return with true if action performed and correctly.
		return true, nil
	}
	// Return with false if no action performed.
	return false, nil
}

// Run all atomic instances in sequential manner.
// If breakable passed as true then run actions until a status of an action returns true.
func (e execution) SequentialRun(breakable bool) error {
	broken := false
	for i := 0; i < len(e.atomics); i++ {
		status, err := e.atomics[i].Run()
		if err != nil {
			return err
		}
		if status {
			broken = true
			if breakable {
				break
			}
		}
	}
	if !broken {
		e.defalutAction()
	}
	return nil
}
