package events

import "time"

type (
	EventInterface interface {
		GetName() string
		GetTime() time.Time
		GetPayload() interface{}
	}

	EventHandlerInterface interface {
		Handle(event EventInterface)
	}

	EventDispatcherInterface interface {
		Register(eventName string, handler EventHandlerInterface) error
		Dispatch(event EventInterface) error
		Remove(eventName string, handler EventHandlerInterface) error
		Has(eventName string, handler EventHandlerInterface) bool
		Clear() error
	}
)
