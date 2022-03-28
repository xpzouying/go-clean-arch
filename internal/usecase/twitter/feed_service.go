package twitter

import (
	"context"

	"github.com/xpzouying/go-clean-arch/internal/domain/twitter"
)

func (t *Twitter) ListFeeds(ctx context.Context) ([]twitter.CardInfo, error) {

	return t.twitterService.ListFeeds(ctx)
}

func (t *Twitter) CreateFeed(ctx context.Context, uid int, text string) (*twitter.CardInfo, error) {
	return t.twitterService.CreateFeed(ctx, uid, text)
}

func (t *Twitter) DeleteFeed(ctx context.Context, uid, feedID int) error {

	return t.twitterService.DeleteFeed(ctx, uid, feedID)
}
