package authtest

import "context"

type AuthTester interface {
	TestUserAuthAndTokenRefresh(ctx context.Context) error
}
