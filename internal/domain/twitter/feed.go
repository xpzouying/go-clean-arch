package twitter

import (
	"context"

	"github.com/xpzouying/go-clean-arch/internal/domain/feed"
	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

func (t *twitterService) CreateFeed(ctx context.Context, uid int, text string) (*CardInfo, error) {
	u, err := t.userRepo.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	fid, err := t.feedRepo.CreateFeed(ctx, uid, text)
	if err != nil {
		return nil, err
	}

	return &CardInfo{
		FID:          fid,
		Text:         text,
		UID:          uid,
		AuthorName:   u.Name,
		AuthorAvatar: u.Avatar,
	}, nil
}

func (t *twitterService) DeleteFeed(ctx context.Context, uid, feedID int) error {

	return t.feedRepo.DeleteFeed(ctx, uid, feedID)
}

func (t *twitterService) ListFeeds(ctx context.Context) ([]CardInfo, error) {
	feeds, err := t.feedRepo.ListFeeds(ctx)
	if err != nil {
		return nil, err
	}

	usersInfo, err := t.getFeedsAuthorInfos(ctx, feeds)
	if err != nil {
		return nil, err
	}

	assember := newCardAssembler(feeds, usersInfo)
	return assember.Assemble()
}

func (t *twitterService) getFeedsAuthorInfos(ctx context.Context, feeds []feed.Feed) (map[int]user.User, error) {
	authorUIDs, err := t.getFeedsUIDs(ctx, feeds)
	if err != nil {
		return nil, err
	}

	return t.userRepo.FindUsers(ctx, authorUIDs)
}

func (t *twitterService) getFeedsUIDs(ctx context.Context, feeds []feed.Feed) ([]int, error) {
	uids := make(map[int]struct{}, len(feeds))

	for _, f := range feeds {
		uid := f.AuthorID

		if _, ok := uids[uid]; ok {
			continue
		}

		uids[uid] = struct{}{}
	}

	result := make([]int, 0, len(uids))
	for uid := range uids {
		result = append(result, uid)
	}
	return result, nil
}
