package main

import (
	"log"
	"net/http"

	"github.com/xpzouying/go-clean-arch/api"
	"github.com/xpzouying/go-clean-arch/internal/repo/feed"
	"github.com/xpzouying/go-clean-arch/internal/repo/user"
	"github.com/xpzouying/go-clean-arch/internal/service"
	"github.com/xpzouying/go-clean-arch/internal/usecase/twitter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatalf("open db failed: %v", err)
	}

	userRepo := user.NewUserRepo(db)
	feedRepo := feed.NewFeedRepo(db)
	twitterUC := twitter.NewTwitter(userRepo, feedRepo)

	twitterSvc := service.NewTwitterService(twitterUC)

	mux := http.NewServeMux()
	api.RegisterTwitterHTTPServer(mux, twitterSvc)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}
