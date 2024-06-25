package domain

import "context"

type EventTag struct{}

func (t EventTag) tagEvent() {}

type Event interface {
	EventID() string
	tagEvent()
}

type EventPublisher interface {
	PublishEvent(ctx context.Context, event Event)
}

type EventListener interface {
	OnEvent(ctx context.Context, event Event)
}
