package xigubo

import (
	"context"
	"errors"
)

var _ Events = (*events)(nil)

type Events interface {
	ListEvents(ctx context.Context, opts *ListOptions) (*EventsListResponse, error)
	CreateEvent(ctx context.Context, opts EventCreate) (*EventResponse, error)
	GetEvent(ctx context.Context, id int64) (*EventResponse, error)
}

type EventsListResponse struct{}

type EventCreateResponse struct{}

type EventCreate struct{}

type EventUpdate struct{}

type EventResponse struct{}

type events struct {
	client *Client
}

func (e events) ListEvents(ctx context.Context, opts *ListOptions) (*EventsListResponse, error) {
	return nil, errors.New("not implemented")
}

func (e events) CreateEvent(ctx context.Context, opts EventCreate) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}

func (e events) GetEvent(ctx context.Context, id int64) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}

func (e events) ResendEvent(ctx context.Context, id int64) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}

func (e events) CancelEvent(ctx context.Context, id int64) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}
