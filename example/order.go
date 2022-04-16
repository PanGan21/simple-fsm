package example

import (
	"errors"
	"fmt"
	"strings"

	fsm "github.com/PanGan21/simplefsm"
)

const (
	CreatingOrder     fsm.StateType = "CreatingOrder"
	OrderFailed       fsm.StateType = "OrderFailed"
	OrderPlaced       fsm.StateType = "OrderPlaced"
	ChargingCard      fsm.StateType = "ChargingCard"
	TransactionFailed fsm.StateType = "TransactionFailed"
	OrderShipped      fsm.StateType = "OrderShipped"

	CreateOrder     fsm.EventType = "CreateOrder"
	FailOrder       fsm.EventType = "FailOrder"
	PlaceOrder      fsm.EventType = "PlaceOrder"
	ChargeCard      fsm.EventType = "ChargeCard"
	FailTransaction fsm.EventType = "FailTransaction"
	ShipOrder       fsm.EventType = "ShipOrder"
)

type OrderCreationContext struct {
	items []string
	err   error
}

func (c *OrderCreationContext) Strinh() string {
	return fmt.Sprintf("OrderCreationContext [ items: %s, err: %v ]", strings.Join(c.items, ","), c.err)
}

type OrderShipmentContext struct {
	cardNumber string
	address    string
	err        error
}

func (c *OrderShipmentContext) String() string {
	return fmt.Sprintf("OrderShipmentContext [ cardNumber: %s, address: %s, err: %v ]",
		c.cardNumber, c.address, c.err)
}

type CreatingOrderAction struct{}

func (a *CreatingOrderAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Validating, order:", order)
	if len(order.items) == 0 {
		order.err = errors.New("Insufficient number of items in order")
		return FailOrder
	}
	return PlaceOrder
}

type OrderFailedAction struct{}

func (a *OrderFailedAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Order failed, err:", order.err)
	return fsm.NoOp
}

type OrderPlacedAction struct{}

func (a *OrderPlacedAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	order := eventCtx.(*OrderCreationContext)
	fmt.Println("Order placed, items:", order.items)
	return fsm.NoOp
}

type ChargingCardAction struct{}

func (a *ChargingCardAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Validating card, shipment:", shipment)
	if shipment.cardNumber == "" {
		shipment.err = errors.New("Card number is invalid")
		return FailTransaction
	}
	return ShipOrder
}

type TransactionFailedAction struct{}

func (a *TransactionFailedAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Transaction failed, err:", shipment.err)
	return fsm.NoOp
}

type OrderShippedAction struct{}

func (a *OrderShippedAction) Execute(eventCtx fsm.EventContext) fsm.EventType {
	shipment := eventCtx.(*OrderShipmentContext)
	fmt.Println("Order shipped, address:", shipment.address)
	return fsm.NoOp
}

func newOrderFSM() *fsm.StateMachine {
	return &fsm.StateMachine{
		States: fsm.States{
			fsm.Default: fsm.State{
				Events: fsm.Events{
					CreateOrder: CreatingOrder,
				},
			},
			CreatingOrder: fsm.State{
				Action: &CreatingOrderAction{},
				Events: fsm.Events{
					FailOrder:  OrderFailed,
					PlaceOrder: OrderPlaced,
				},
			},
			OrderFailed: fsm.State{
				Action: &OrderFailedAction{},
				Events: fsm.Events{
					CreateOrder: CreatingOrder,
				},
			},
			OrderPlaced: fsm.State{
				Action: &OrderPlacedAction{},
				Events: fsm.Events{
					ChargeCard: ChargingCard,
				},
			},
			ChargingCard: fsm.State{
				Action: &ChargingCardAction{},
				Events: fsm.Events{
					FailTransaction: TransactionFailed,
					ShipOrder:       OrderShipped,
				},
			},
			TransactionFailed: fsm.State{
				Action: &TransactionFailedAction{},
				Events: fsm.Events{
					ChargeCard: ChargingCard,
				},
			},
			OrderShipped: fsm.State{
				Action: &OrderShippedAction{},
			},
		},
	}
}
