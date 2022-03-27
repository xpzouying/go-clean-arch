package service

import "github.com/xpzouying/go-clean-arch/internal/usecase/twitter"

type TwitterService struct {
	tc *twitter.Twitter
}

func NewTwitterService(tc *twitter.Twitter) *TwitterService {
	return &TwitterService{tc}
}
