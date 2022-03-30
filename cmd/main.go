package main

import (
	"log"
	"net/http"

	"github.com/xpzouying/go-clean-arch/api"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var dsn = ":memory:"
	twitterSvc, err := initTwitterSvc(dsn)
	if err != nil {
		log.Fatalf("open db failed: %v", err)
	}

	mux := http.NewServeMux()
	api.RegisterTwitterHTTPServer(mux, twitterSvc)

	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func newGormDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}
