package twitter

import (
	"context"
)

type FeedInfo struct {
	FeedID int
	Text   string

	AuthorID     int
	AuthorName   string
	AuthorAvatar string
}

func (t *Twitter) CreateFeed(ctx context.Context, uid int, text string) (*FeedInfo, error) {
	u, err := t.userRepo.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	fid, err := t.feedRepo.CreateFeed(ctx, uid, text)
	if err != nil {
		return nil, err
	}

	return &FeedInfo{
		FeedID:       fid,
		Text:         text,
		AuthorID:     uid,
		AuthorName:   u.Name,
		AuthorAvatar: u.Avatar,
	}, nil
}

func (t *Twitter) DeleteFeed(ctx context.Context, uid, feedID int) error {

	return t.feedRepo.DeleteFeed(ctx, uid, feedID)
}
