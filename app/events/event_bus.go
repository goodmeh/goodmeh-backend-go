package events

type EventType int

type Event struct {
	Type    EventType
	Payload any
}

type Subscriber func(Event)

type EventBus struct {
	subscribers map[EventType][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]Subscriber),
	}
}

func (eb *EventBus) Subscribe(eventType EventType, subscriber Subscriber) {
	if _, exists := eb.subscribers[eventType]; !exists {
		eb.subscribers[eventType] = []Subscriber{}
	}
	eb.subscribers[eventType] = append(eb.subscribers[eventType], subscriber)
}

func (eb *EventBus) Publish(event Event) {
	if subscribers, exists := eb.subscribers[event.Type]; exists {
		for _, subscriber := range subscribers {
			subscriber(event)
		}
	}
}
