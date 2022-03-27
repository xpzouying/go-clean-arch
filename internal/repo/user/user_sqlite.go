package user

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

type userPO struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement"`

	Name   string    `gorm:"column:name"`
	Avatar string    `gorm:"column:avatar"`
	Crtime time.Time `gorm:"column:crtime"`
	Uptime time.Time `gorm:"column:uptime"`
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user.UserRepo {

	return &userRepo{db}
}

func (rp *userRepo) CreateUser(ctx context.Context, name, avatar string) (int, error) {

	u := userPO{Name: name, Avatar: avatar}

	if err := rp.db.WithContext(ctx).Create(&u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (rp *userRepo) GetUser(ctx context.Context, uid int) (*user.User, error) {
	var u userPO
	err := rp.db.WithContext(ctx).
		Where("id = ?", uid).
		Take(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user.User{
		Uid:    u.ID,
		Name:   u.Name,
		Avatar: u.Avatar,
	}, nil
}
