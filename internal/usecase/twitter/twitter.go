package twitter

import (
	"errors"

	"github.com/xpzouying/go-clean-arch/internal/domain/feed"
	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

var (
	ErrInvalidUid    = errors.New("invalid uid")
	ErrInvalidName   = errors.New("invalid username")
	ErrInvalidAvatar = errors.New("invalid avatar")
)

type Twitter struct {
	userRepo user.UserRepo
	feedRepo feed.FeedRepo
}

func NewTwitter(userRepo user.UserRepo, feedRepo feed.FeedRepo) *Twitter {

	return &Twitter{
		userRepo: userRepo,
		feedRepo: feedRepo,
	}
}
