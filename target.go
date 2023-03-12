package goswitch

import (
	"context"
	"errors"
)

type Targets interface {
	ListTargets(ctx context.Context, opts *ListOptions) (*TargetsListResponse, error)
	CreateTarget(ctx context.Context, opts TargetCreateInput) (*TargetResponse, error)
	GetTarget(ctx context.Context, id int64) (*TargetResponse, error)
	DeleteTarget(ctx context.Context, id int64) (*TargetResponse, error)
	UpdateTarget(ctx context.Context, id int64, data TargetUpdateInput) (*TargetResponse, error)
	CreateToken(ctx context.Context, opts TargetCreateInput) (*TokenResponse, error)
	GetTokenInfo(ctx context.Context, id int64) (*TokenInfoResponse, error)
}

type TargetsListResponse struct{}

type TargetCreateResponse struct{}

type TargetCreateInput struct{}

type TargetUpdateInput struct{}

type TargetResponse struct{}

type TokenResponse struct{}

type TokenInfoResponse struct{}

type targets struct {
	client *Client
}

func (t targets) ListTargets(ctx context.Context, opts *ListOptions) (*TargetsListResponse, error) {
	return nil, errors.New("not implemented")
}

func (t targets) CreateTarget(ctx context.Context, opts TargetCreateInput) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (t targets) GetTarget(ctx context.Context, id int64) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (t targets) DeleteTarget(ctx context.Context, id int64) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (t targets) UpdateTarget(ctx context.Context, id int64, data TargetUpdateInput) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (t targets) CreateToken(ctx context.Context, opts TargetCreateInput) (*TokenResponse, error) {
	return nil, errors.New("not implemented")
}

func (t targets) GetTokenInfo(ctx context.Context, id int64) (*TokenInfoResponse, error) {
	return nil, errors.New("not implemented")
}
