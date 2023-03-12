package goswitch

import (
	"context"
	"errors"
)

type Webhooks interface {
	ListWebhooks(ctx context.Context, opts *ListOptions) (*WebhooksListResponse, error)
	CreateWebhook(ctx context.Context, opts WebhookCreateInput) (*WebhookResponse, error)
	GetWebhook(ctx context.Context, id int64) (*WebhookResponse, error)
	DeleteWebhook(ctx context.Context, id int64) (*WebhookResponse, error)
	UpdateWebhook(ctx context.Context, id int64, data WebhookUpdateInput) (*WebhookResponse, error)
}

type WebhooksListResponse struct{}

type WebhookCreateResponse struct{}

type WebhookCreateInput struct{}

type WebhookUpdateInput struct{}

type WebhookResponse struct{}

type webhooks struct {
	client *Client
}

func (w webhooks) ListWebhooks(ctx context.Context, opts *ListOptions) (*WebhooksListResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) CreateWebhook(ctx context.Context, opts WebhookCreateInput) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) GetWebhook(ctx context.Context, id int64) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) DeleteWebhook(ctx context.Context, id int64) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w webhooks) UpdateWebhook(ctx context.Context, id int64, data WebhookUpdateInput) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}
