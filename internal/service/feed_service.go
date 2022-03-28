package service

import (
	"context"

	"github.com/xpzouying/go-clean-arch/api"
	"github.com/xpzouying/go-clean-arch/internal/domain/twitter"
)

func (ts *TwitterService) CreateFeed(ctx context.Context, req *api.CreateFeedReq) (*api.CreateFeedReply, error) {

	result, err := ts.tc.CreateFeed(ctx, int(req.UID), req.Text)
	if err != nil {
		return nil, err
	}

	return &api.CreateFeedReply{
		CardInfo: api.CardInfo{
			AuthorInfo: api.AuthorInfo{
				UID:          uint32(result.UID),
				AuthorName:   result.AuthorName,
				AuthorAvatar: result.AuthorAvatar,
			},

			FeedInfo: api.FeedInfo{
				FID:  uint32(result.FID),
				Text: result.Text,
			},
		},
	}, nil
}

func (ts *TwitterService) DeleteFeed(ctx context.Context, req *api.DeleteFeedReq) error {

	return ts.tc.DeleteFeed(ctx, int(req.UID), int(req.FeedID))
}

func (ts *TwitterService) ListFeeds(ctx context.Context) (*api.ListFeedReply, error) {

	result, err := ts.tc.ListFeeds(ctx)
	if err != nil {
		return nil, err
	}

	feeds := ts.assembleCardInfo(result)

	return &api.ListFeedReply{
		Feeds: feeds,
	}, nil
}

func (ts *TwitterService) assembleCardInfo(cards []twitter.CardInfo) []api.CardInfo {

	result := make([]api.CardInfo, 0, len(cards))

	for _, c := range cards {

		result = append(result, api.CardInfo{
			AuthorInfo: api.AuthorInfo{
				UID:          uint32(c.UID),
				AuthorName:   c.AuthorName,
				AuthorAvatar: c.AuthorAvatar,
			},

			FeedInfo: api.FeedInfo{
				FID:  uint32(c.FID),
				Text: c.Text,
			},
		})

	}

	return result
}
