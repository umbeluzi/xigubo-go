package xigubo

import (
	"context"
	"errors"
)

var _ EventTypes = (*eventTypes)(nil)

type EventTypes interface {
	ListEventTypes(ctx context.Context, opts *ListOptions) (*EventTypesListResponse, error)
	CreateEventType(ctx context.Context, opts EventTypeCreateInput) (*EventTypeResponse, error)
	GetEventType(ctx context.Context, id int64) (*EventTypeResponse, error)
	DeleteEventType(ctx context.Context, id int64) (*EventTypeResponse, error)
	UpdateEventType(ctx context.Context, id int64, data EventTypeUpdateInput) (*EventTypeResponse, error)
}

type EventTypesListResponse struct{}

type EventTypeCreateResponse struct{}

type EventTypeCreateInput struct{}

type EventTypeUpdateInput struct{}

type EventTypeResponse struct{}

type eventTypes struct {
	client *Client
}

func (w eventTypes) ListEventTypes(ctx context.Context, opts *ListOptions) (*EventTypesListResponse, error) {
	return nil, errors.New("not implemented")
}

func (w eventTypes) CreateEventType(ctx context.Context, opts EventTypeCreateInput) (*EventTypeResponse, error) {
	return nil, errors.New("not implemented")
}

func (w eventTypes) GetEventType(ctx context.Context, id int64) (*EventTypeResponse, error) {
	return nil, errors.New("not implemented")
}

func (w eventTypes) DeleteEventType(ctx context.Context, id int64) (*EventTypeResponse, error) {
	return nil, errors.New("not implemented")
}

func (w eventTypes) UpdateEventType(ctx context.Context, id int64, data EventTypeUpdateInput) (*EventTypeResponse, error) {
	return nil, errors.New("not implemented")
}
