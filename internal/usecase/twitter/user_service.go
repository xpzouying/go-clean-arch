package twitter

import (
	"context"

	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

func (t *Twitter) CreateUser(ctx context.Context, name, avatar string) (uid int, err error) {
	if len(name) == 0 {
		return 0, ErrInvalidName
	}

	if len(avatar) == 0 {
		return 0, ErrInvalidAvatar
	}

	return t.userRepo.CreateUser(ctx, name, avatar)
}

func (t *Twitter) GetUser(ctx context.Context, uid int) (*user.User, error) {
	if uid == 0 {
		return nil, ErrInvalidUid
	}

	return t.userRepo.GetUser(ctx, uid)
}
