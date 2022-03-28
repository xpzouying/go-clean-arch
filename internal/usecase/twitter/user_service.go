package twitter

import (
	"context"

	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

func (t *Twitter) CreateUser(ctx context.Context, name, avatar string) (uid int, err error) {
	return t.twitterService.CreateUser(ctx, name, avatar)
}

func (t *Twitter) GetUser(ctx context.Context, uid int) (user.User, error) {
	return t.twitterService.GetUser(ctx, uid)
}
