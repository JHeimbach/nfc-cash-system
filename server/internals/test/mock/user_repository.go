package mock

import (
	"context"

	"github.com/JHeimbach/nfc-cash-system/server/api"
)

type UserRepository struct {
	Called           map[string]bool
	AuthenticateFunc func(email, password string) (*api.User, error)
}

func (m *UserRepository) Authenticate(_ context.Context, email, password string) (*api.User, error) {
	m.Called["auth"] = true
	return m.AuthenticateFunc(email, password)
}
