package xigubo

import (
	"context"
	"errors"
)

var _ Webhooks = (*webhooks)(nil)

type Webhooks interface {
	ListWebhooks(ctx context.Context, opts *ListOptions) (*WebhooksListResponse, error)
	CreateWebhook(ctx context.Context, opts WebhookCreate) (*WebhookResponse, error)
	GetWebhook(ctx context.Context, id int64) (*WebhookResponse, error)
	DeleteWebhook(ctx context.Context, id int64) (*WebhookResponse, error)
	UpdateWebhook(ctx context.Context, id int64, data WebhookUpdate) (*WebhookResponse, error)
}

type WebhooksListResponse struct{}

type WebhookCreateResponse struct{}

type WebhookCreate struct{}

type WebhookUpdate struct{}

type WebhookResponse struct{}

type webhooks struct {
	client *Client
}

func (w webhooks) ListWebhooks(ctx context.Context, opts *ListOptions) (*WebhooksListResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) CreateWebhook(ctx context.Context, opts WebhookCreate) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) GetWebhook(ctx context.Context, id int64) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) DeleteWebhook(ctx context.Context, id int64) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) UpdateWebhook(ctx context.Context, id int64, data WebhookUpdate) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}
