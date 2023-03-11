package plumber

import (
	"context"
	"errors"
)

type WebhookService struct{}

func (w WebhookService) ListWebhooks(ctx context.Context, opts *ListOptions) (*WebhooksListResponse, error) {
	return nil, errors.New("not implemented")
}

func (w WebhookService) CreateWebhook(ctx context.Context, opts WebhookCreateInput) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w WebhookService) GetWebhook(ctx context.Context, id int64) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w WebhookService) DeleteWebhook(ctx context.Context, id int64) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

func (w WebhookService) UpdateWebhook(ctx context.Context, id int64, data WebhookUpdateInput) (*WebhookResponse, error) {
	return nil, errors.New("not implemented")
}

type WebhooksListResponse struct{}

type WebhookCreateResponse struct{}

type WebhookCreateInput struct{}

type WebhookUpdateInput struct{}

type WebhookResponse struct{}

type ListOptions struct{}
