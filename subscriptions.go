package xigubo

import (
	"context"
	"errors"
)

var _ Subscriptions = (*subscriptions)(nil)

type Subscriptions interface {
	ListSubscriptions(ctx context.Context, opts *ListOptions) (*SubscriptionsListResponse, error)
	CreateSubscription(ctx context.Context, opts SubscriptionCreateInput) (*SubscriptionResponse, error)
	GetSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error)
	DeleteSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, id int64, data SubscriptionUpdateInput) (*SubscriptionResponse, error)
	CreateCredentials(ctx context.Context, opts SubscriptionCreateInput) (*CredentialsResponse, error)
	GetCredentialsInfo(ctx context.Context, id int64) (*CredentialsInfoResponse, error)
}

type SubscriptionsListResponse struct{}

type SubscriptionCreateResponse struct{}

type SubscriptionCreateInput struct{}

type SubscriptionUpdateInput struct{}

type SubscriptionResponse struct{}

type CredentialsResponse struct{}

type CredentialsInfoResponse struct{}

type subscriptions struct {
	client *Client
}

func (t subscriptions) ListSubscriptions(ctx context.Context, opts *ListOptions) (*SubscriptionsListResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) CreateSubscription(ctx context.Context, opts SubscriptionCreateInput) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) GetSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) DeleteSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) UpdateSubscription(ctx context.Context, id int64, data SubscriptionUpdateInput) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) CreateCredentials(ctx context.Context, opts SubscriptionCreateInput) (*CredentialsResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) GetCredentialsInfo(ctx context.Context, id int64) (*CredentialsInfoResponse, error) {
	return nil, errors.New("not implemented")
}
