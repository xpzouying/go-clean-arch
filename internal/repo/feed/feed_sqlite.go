package feed

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/xpzouying/go-clean-arch/internal/domain/feed"
)

type Status uint8

const (
	StatusDelete Status = iota
	StatusNormal
)

// feedPO is ORM object for feed
type feedPO struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement"`
	UID    int    `gorm:"column:uid;index"`
	Text   string `gorm:"column:text"`
	Status Status `gorm:"column:status"`

	Crtime time.Time `gorm:"column:crtime;autoCreateTime"`
	Uptime time.Time `gorm:"column:uptime;autoUpdateTime"`
}

func (feedPO) TableName() string {
	return "feed"
}

type feedRepo struct {
	db *gorm.DB
}

func NewFeedRepo(db *gorm.DB) feed.FeedRepo {
	_ = db.AutoMigrate(&feedPO{})

	return &feedRepo{db: db}
}

func (rp *feedRepo) CreateFeed(ctx context.Context, uid int, text string) (int, error) {

	r := feedPO{
		UID:  uid,
		Text: text,
	}
	err := rp.db.WithContext(ctx).Create(&r).Error

	return r.ID, err
}

func (rp *feedRepo) GetFeed(ctx context.Context, fid int) (*feed.Feed, error) {

	var r feedPO
	err := rp.db.WithContext(ctx).
		Where("id = ?", fid).
		Where("status = ?", StatusNormal).
		Take(&r).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &feed.Feed{
		FeedID:   r.ID,
		AuthorID: r.UID,
		Text:     r.Text,
	}, nil
}

func (rp *feedRepo) DeleteFeed(ctx context.Context, uid int, fid int) error {

	return rp.db.WithContext(ctx).
		Where("uid = ?", uid).
		Where("id = ?", fid).
		Update("status", StatusDelete).Error
}
