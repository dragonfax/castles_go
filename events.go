package main

type Event interface{}

type EventC chan Event

var eventSendC = make(EventC)

var eventReceivers = make([]EventC, 0)

// to run in a goroutine to make sure all event listeners get all events.
func MuxEvents() {
	for {
		select {
		case event := <-eventSendC:
			for _, ec := range eventReceivers {
				ec <- event
			}
		}
	}
}

// for anyone to get an event queue of their own
func GetEventReceiver() EventC {
	eventReceiver := make(EventC)
	eventReceivers = append(eventReceivers, eventReceiver)
	return eventReceiver
}

func removeEventReceiver(ec EventC) {
	var index int
	for i := 0; i < len(eventReceivers); i++ {
		if eventReceivers[i] == ec {
			index = i
		}
	}
	eventReceivers = append(eventReceivers[:index], eventReceivers[index+1:]...)
}

func CloseEventReceiver(ec EventC) {
	close(ec)
	removeEventReceiver(ec)

	// incase MuxEvents is already writing to it
	var ok = true
	for ok {
		_, ok = <-ec
	}
}
