package flowhandler

import (
	"context"

	"cmd/iam/main.go/internal/delivery/http/api"
)

type FlowHandler struct {
}

var _ api.FlowHandler = (*FlowHandler)(nil)

func (h *FlowHandler) LoginPost(ctx context.Context, req *api.LoginRequest) (*api.TokenResponse, error) {
	return nil, nil

}

func (h *FlowHandler) RefreshPost(ctx context.Context, req *api.RefreshPostReq) (*api.TokenResponse, error) {
	return nil, nil

}

func (h *FlowHandler) SignupPost(ctx context.Context, req *api.SignupRequest) (*api.User, error) {
	return nil, nil

}
