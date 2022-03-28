package twitter

import (
	"github.com/xpzouying/go-clean-arch/internal/domain/twitter"
)

type Twitter struct {
	twitterService twitter.DomainService
}

func NewTwitter(twitterService twitter.DomainService) *Twitter {

	return &Twitter{twitterService: twitterService}
}
