package authhandler

import (
	"context"

	"cmd/iam/main.go/internal/delivery/http/api"
)

type AuthHandler struct {
}

var _ api.AuthHandler = (*AuthHandler)(nil)

func (h *AuthHandler) LoginPost(ctx context.Context, req *api.LoginRequest) (*api.TokenResponse, error) {
	return nil, nil

}

func (h *AuthHandler) RefreshPost(ctx context.Context, req *api.RefreshPostReq) (*api.TokenResponse, error) {
	return nil, nil

}

func (h *AuthHandler) SignupPost(ctx context.Context, req *api.SignupRequest) (*api.User, error) {
	return nil, nil

}
