package twitter

import (
	"context"
	"errors"

	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

var (
	ErrInvalidUid    = errors.New("invalid uid")
	ErrInvalidName   = errors.New("invalid username")
	ErrInvalidAvatar = errors.New("invalid avatar")
)

func (t *twitterService) CreateUser(ctx context.Context, name, avatar string) (uid int, err error) {
	if len(name) == 0 {
		return 0, ErrInvalidName
	}

	if len(avatar) == 0 {
		return 0, ErrInvalidAvatar
	}

	return t.userRepo.CreateUser(ctx, name, avatar)
}

func (t *twitterService) GetUser(ctx context.Context, uid int) (user.User, error) {
	if uid == 0 {
		return user.User{}, ErrInvalidUid
	}

	return t.userRepo.GetUser(ctx, uid)
}
