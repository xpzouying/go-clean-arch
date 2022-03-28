package twitter

import (
	"context"

	"github.com/xpzouying/go-clean-arch/internal/domain/feed"
	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

type CardInfo struct {
	FID  int
	Text string

	UID          int
	AuthorName   string
	AuthorAvatar string
}

type DomainService interface {
	ListFeeds(ctx context.Context) ([]CardInfo, error)
	CreateFeed(ctx context.Context, uid int, text string) (*CardInfo, error)
	DeleteFeed(ctx context.Context, uid, feedID int) error

	GetUser(ctx context.Context, uid int) (user.User, error)
	CreateUser(ctx context.Context, name, avatar string) (uid int, err error)
}

type twitterService struct {
	feedRepo feed.FeedRepo
	userRepo user.UserRepo
}

func NewTwitterService(feedRepo feed.FeedRepo, userRepo user.UserRepo) DomainService {
	return &twitterService{
		feedRepo: feedRepo,
		userRepo: userRepo,
	}
}
