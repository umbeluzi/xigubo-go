package xigubo

import (
	"context"
	"errors"
)

var _ Subscriptions = (*subscriptions)(nil)

type Subscriptions interface {
	ListSubscriptions(ctx context.Context, opts *ListOptions) (*SubscriptionsListResponse, error)
	CreateSubscription(ctx context.Context, opts SubscriptionCreate) (*SubscriptionResponse, error)
	GetSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error)
	DeleteSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error)
	UpdateSubscription(ctx context.Context, id int64, data SubscriptionUpdate) (*SubscriptionResponse, error)
	CreateCredentials(ctx context.Context, opts SubscriptionCreate) (*CredentialsResponse, error)
	GetCredentialsInfo(ctx context.Context, id int64) (*CredentialsInfoResponse, error)
}

type SubscriptionsListResponse struct{}

type SubscriptionCreateResponse struct{}

type SubscriptionCreate struct {
	URL               string                      `json:"url"`
	Filter            *SubscriptionFilter         `json:"filter"`
	ConnectionTimeout *int                        `json:"connection_timeout"`
	ReadTimeout       *int                        `json:"read_timeout"`
	Security          *SubscriptionSecurityConfig `json:"security"`
}

type SubscriptionUpdate struct{}

type SubscriptionResponse struct{}

type CredentialsResponse struct{}

type CredentialsInfoResponse struct{}

type SubscriptionSecurityConfig struct {
	TLSCACert    *string `json:"tls_ca_cert"`
	TLSCientCert *string `json:"tls_client_cert"`
}

type SubscriptionFilter struct {
	Topics      []string `json:"topics"`
	EnventTypes []string `json:"event_types"`
}

type subscriptions struct {
	client *Client
}

func (t subscriptions) ListSubscriptions(ctx context.Context, opts *ListOptions) (*SubscriptionsListResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) CreateSubscription(ctx context.Context, opts SubscriptionCreate) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) GetSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) DeleteSubscription(ctx context.Context, id int64) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) UpdateSubscription(ctx context.Context, id int64, data SubscriptionUpdate) (*SubscriptionResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) CreateCredentials(ctx context.Context, opts SubscriptionCreate) (*CredentialsResponse, error) {
	return nil, errors.New("not implemented")
}

func (t subscriptions) GetCredentialsInfo(ctx context.Context, id int64) (*CredentialsInfoResponse, error) {
	return nil, errors.New("not implemented")
}
