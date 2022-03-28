package twitter

import (
	"github.com/xpzouying/go-clean-arch/internal/domain/feed"
	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

type cardAssembler struct {
	feeds     []feed.Feed
	usersInfo map[int]user.User
}

func newCardAssembler(feeds []feed.Feed, usersInfo map[int]user.User) cardAssembler {
	return cardAssembler{feeds, usersInfo}
}

func (ca *cardAssembler) Assemble() ([]CardInfo, error) {

	cardsInfos := make([]CardInfo, 0, len(ca.feeds))

	for _, f := range ca.feeds {

		userInfo, ok := ca.usersInfo[f.AuthorID]
		if !ok {
			continue
		}

		cardInfo := CardInfo{
			FID:          f.FeedID,
			Text:         f.Text,
			UID:          f.AuthorID,
			AuthorName:   userInfo.Name,
			AuthorAvatar: userInfo.Avatar,
		}

		cardsInfos = append(cardsInfos, cardInfo)
	}

	return cardsInfos, nil
}
