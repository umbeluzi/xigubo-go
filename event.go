package plumber

import (
	"context"
	"errors"
)

type EventService struct{}

func (w EventService) ListEvents(ctx context.Context, opts *ListOptions) (*EventsListResponse, error) {
	return nil, errors.New("not implemented")
}

func (w EventService) CreateEvent(ctx context.Context, opts EventCreateInput) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}

func (w EventService) GetEvent(ctx context.Context, id int64) (*EventResponse, error) {
	return nil, errors.New("not implemented")
}

type EventsListResponse struct{}

type EventCreateResponse struct{}

type EventCreateInput struct{}

type EventUpdateInput struct{}

type EventResponse struct{}

type ListOptions struct{}
