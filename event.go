package xigubo

import (
	"context"
	"errors"
)

var _ Events = (*events)(nil)

type Events interface {
	ListEvents(ctx context.Context, opts *ListOptions) (*EventsListResponse, error)
	CreateEvent(ctx context.Context, opts EventCreateInput) (*EventResponse, error)
	GetEvent(ctx context.Context, id int64) (*EventResponse, error)
}

type EventsListResponse struct{}

type EventCreateResponse struct{}

type EventCreateInput struct{}

type EventUpdateInput struct{}

type EventResponse struct{}

type events struct {
	client *Client
}

func (e events) ListEvents(ctx context.Context, opts *ListOptions) (*EventsListResponse, error) {
	return nil, errors.New("not implemented")
}

func (e events) CreateEvent(ctx context.Context, opts EventCreateInput) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}

func (e events) GetEvent(ctx context.Context, id int64) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}
