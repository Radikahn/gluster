// eventsBus.go is responsible for holding all events that occurr in gluster
// The eventBus internally utilizes the default queue implementation in directive.Queue
package events

import (
	"github.com/Radikahn/gluster/cmd/device"
	"github.com/Radikahn/gluster/cmd/directive"
	"github.com/google/uuid"
)

// type event struct {
// 	//some event data
// 	cmd *directive.CommandDirective
// 	next    *event
// 	prev    *event
// }

type EventBus struct {
	id     string
	device *device.Device
	queue  *directive.Queue
}

func NewEventBus(device *device.Device) *EventBus {
	id := uuid.NewString()
	q := directive.NewQueue()

	return &EventBus{id: id, device: device, queue: q}
}

// Push an event (commandDirective) to the eventBus
func (e *EventBus) Push(cmd *directive.CommandDirective) {
	e.queue.Push(cmd)
}

// Pop an event (commandDirective) from the eventBus
func (e *EventBus) Pop() (*directive.CommandDirective, string, bool) {
	//TODO: event handler needs to execute Popped command, this function can instead return just a bool
	directive, status := e.queue.Pop()

	return directive, e.device.Id, status
}
