// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/xpzouying/go-clean-arch/internal/domain/twitter"
	"github.com/xpzouying/go-clean-arch/internal/repo/feed"
	"github.com/xpzouying/go-clean-arch/internal/repo/user"
	"github.com/xpzouying/go-clean-arch/internal/service"
	twitter2 "github.com/xpzouying/go-clean-arch/internal/usecase/twitter"
)

// Injectors from wire.go:

func initTwitterSvc(dsn string) (*service.TwitterService, error) {
	db, err := newGormDB(dsn)
	if err != nil {
		return nil, err
	}
	feedRepo := feed.NewFeedRepo(db)
	userRepo := user.NewUserRepo(db)
	domainService := twitter.NewTwitterService(feedRepo, userRepo)
	twitterTwitter := twitter2.NewTwitter(domainService)
	twitterService := service.NewTwitterService(twitterTwitter)
	return twitterService, nil
}