package xigubo

import (
	"context"
	"errors"
)

var _ Subscribers = (*subscribers)(nil)

type Subscribers interface {
	ListSubscribers(ctx context.Context, opts *ListOptions) (*SubscribersListResponse, error)
	CreateSubscriber(ctx context.Context, opts SubscriberCreateInput) (*SubscriberResponse, error)
	GetSubscriber(ctx context.Context, id int64) (*SubscriberResponse, error)
	DeleteSubscriber(ctx context.Context, id int64) (*SubscriberResponse, error)
	UpdateSubscriber(ctx context.Context, id int64, data SubscriberUpdateInput) (*SubscriberResponse, error)
	CreateToken(ctx context.Context, opts SubscriberCreateInput) (*TokenResponse, error)
	GetTokenInfo(ctx context.Context, id int64) (*TokenInfoResponse, error)
}

type SubscribersListResponse struct{}

type SubscriberCreateResponse struct{}

type SubscriberCreateInput struct{}

type SubscriberUpdateInput struct{}

type SubscriberResponse struct{}

type TokenResponse struct{}

type TokenInfoResponse struct{}

type subscribers struct {
	client *Client
}

func (t subscribers) ListSubscribers(ctx context.Context, opts *ListOptions) (*SubscribersListResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscribers) CreateSubscriber(ctx context.Context, opts SubscriberCreateInput) (*SubscriberResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscribers) GetSubscriber(ctx context.Context, id int64) (*SubscriberResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscribers) DeleteSubscriber(ctx context.Context, id int64) (*SubscriberResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscribers) UpdateSubscriber(ctx context.Context, id int64, data SubscriberUpdateInput) (*SubscriberResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscribers) CreateToken(ctx context.Context, opts SubscriberCreateInput) (*TokenResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscribers) GetTokenInfo(ctx context.Context, id int64) (*TokenInfoResponse, error) {
	return nil, errors.New("not implemented")
}
