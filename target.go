package plumber

import (
	"context"
	"errors"
)

type TargetService struct{}

func (w TargetService) ListTargets(ctx context.Context, opts *ListOptions) (*TargetsListResponse, error) {
	return nil, errors.New("not implemented")
}

func (w TargetService) CreateTarget(ctx context.Context, opts TargetCreateInput) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (w TargetService) GetTarget(ctx context.Context, id int64) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (w TargetService) DeleteTarget(ctx context.Context, id int64) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (w TargetService) UpdateTarget(ctx context.Context, id int64, data TargetUpdateInput) (*TargetResponse, error) {
	return nil, errors.New("not implemented")
}

func (w TargetService) CreateToken(ctx context.Context, opts TargetCreateInput) (*TokenResponse, error) {
	return nil, errors.New("not implemented")
}

type TargetsListResponse struct{}

type TargetCreateResponse struct{}

type TargetCreateInput struct{}

type TargetUpdateInput struct{}

type TargetResponse struct{}

type TokenResponse struct{}

type ListOptions struct{}
