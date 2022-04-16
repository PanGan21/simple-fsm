package example

import (
	"fmt"

	fsm "github.com/PanGan21/simplefsm"
)

const (
	Off fsm.StateType = "Off"
	On  fsm.StateType = "On"

	SwitchOff fsm.EventType = "SwitchOff"
	SwitchOn  fsm.EventType = "SwitchOn"
)

// OffAction represents the action executed on entering the Off state
type OffAction struct{}

func (a *OffAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	fmt.Println("The light has been switched off")
	return fsm.NoOp
}

// OnAction represents the action executed on entering the Off state
type OnAction struct{}

func (a *OnAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	fmt.Println("The light has been switched on")
	return fsm.NoOp
}

func NewLightSwitchFSM() *fsm.StateMachine {
	return &fsm.StateMachine{
		States: fsm.States{
			fsm.Default: fsm.State{
				Events: fsm.Events{
					SwitchOff: Off,
				},
			},
			Off: fsm.State{
				Action: &OffAction{},
				Events: fsm.Events{
					SwitchOn: On,
				},
			},
			On: fsm.State{
				Action: &OnAction{},
				Events: fsm.Events{
					SwitchOff: Off,
				},
			},
		},
	}
}
