package example

import (
	"testing"

	fsm "github.com/PanGan21/simplefsm"
)

func TestLightSwitchStateMachine(t *testing.T) {
	// Create a new instance of the light switch state machine.
	lightSwitchFSM := NewLightSwitchFSM()

	// See the initial "off" state im the state machine
	err := lightSwitchFSM.SendEvent(SwitchOff, nil)
	if err != nil {
		t.Errorf("Couldn't set the initial state of the state machine, err: %v", err)
	}

	// Send the switch-off event again and expect the state machine to return an error.
	err = lightSwitchFSM.SendEvent(SwitchOff, nil)
	if err != fsm.ErrEventRejected {
		t.Errorf("Expected the event rejected error, got nil")
	}

	// Send the switch-on event and expect th e state machine to transition to the
	// "on" state
	err = lightSwitchFSM.SendEvent(SwitchOn, nil)
	if err != nil {
		t.Errorf("Couldn't switch the light on, err: %v", err)
	}

	// Send the switch-on event again and expect the state machine to return an error.
	err = lightSwitchFSM.SendEvent(SwitchOn, nil)
	if err != fsm.ErrEventRejected {
		t.Errorf("Expected the event rejected error, got nil")
	}

	// Send the switch-off event and expect the state machine to transition back
	// to the "off" state.
	err = lightSwitchFSM.SendEvent(SwitchOff, nil)
	if err != nil {
		t.Errorf("Couldn't switch the light off, err: %v", err)
	}

}
