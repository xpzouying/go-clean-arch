package feed

import "context"

type Feed struct {
	FeedID   int
	AuthorID int
	Text     string
}

type FeedRepo interface {
	CreateFeed(ctx context.Context, uid int, text string) (int, error)
	GetFeed(ctx context.Context, fid int) (*Feed, error)
	DeleteFeed(ctx context.Context, uid int, feedID int) error
}
