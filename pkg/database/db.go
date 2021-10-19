package database

import (
	"context"
)

type Manager interface {
	InitDataBase()
	EventDao() EventInterface
}

const eventTableName = "t_event"

type EventInterface interface {
	NewEvent(ctx context.Context, message string) error
	GetByEventId(ctx context.Context, eventId string) (*EventModel, error)
	List(ctx context.Context) ([]*EventModel, error)
}
