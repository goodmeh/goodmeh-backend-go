package events

type EventType int

type Subscriber func(any)

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

func (eb *EventBus) Publish(eventType EventType, payload any) {
	if subscribers, exists := eb.subscribers[eventType]; exists {
		for _, subscriber := range subscribers {
			subscriber(payload)
		}
	}
}

func AssertHandler[T any](handler func(T)) func(any) {
	return func(payload any) {
		if p, ok := payload.(T); ok {
			handler(p)
		} else {
			panic("Invalid payload type")
		}
	}
}
