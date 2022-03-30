//go:build wireinject
// +build wireinject

package main

import (
	twitterdo "github.com/xpzouying/go-clean-arch/internal/domain/twitter"
	"github.com/xpzouying/go-clean-arch/internal/repo/feed"
	"github.com/xpzouying/go-clean-arch/internal/repo/user"
	"github.com/xpzouying/go-clean-arch/internal/service"
	"github.com/xpzouying/go-clean-arch/internal/usecase/twitter"

	"github.com/google/wire"
)

func initTwitterSvc(dsn string) (*service.TwitterService, error) {

	panic(
		wire.Build(
			newGormDB,
			user.NewUserRepo,
			feed.NewFeedRepo,
			twitterdo.NewTwitterService,
			twitter.NewTwitter,
			service.NewTwitterService,
		),
	)
}
