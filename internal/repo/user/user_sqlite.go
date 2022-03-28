package user

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/xpzouying/go-clean-arch/internal/domain/user"
)

// userPO is ORM object for user
type userPO struct {
	ID int `gorm:"column:id;primaryKey;autoIncrement"`

	Name   string    `gorm:"column:name"`
	Avatar string    `gorm:"column:avatar"`
	Crtime time.Time `gorm:"column:crtime;autoCreateTime"`
	Uptime time.Time `gorm:"column:uptime;autoUpdateTime"`
}

func (userPO) TableName() string {

	return "user"
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) user.UserRepo {
	_ = db.AutoMigrate(&userPO{})

	return &userRepo{db}
}

func (rp *userRepo) CreateUser(ctx context.Context, name, avatar string) (int, error) {

	u := userPO{
		Name:   name,
		Avatar: avatar,
	}

	if err := rp.db.WithContext(ctx).Create(&u).Error; err != nil {
		return 0, err
	}
	return u.ID, nil
}

func (rp *userRepo) GetUser(ctx context.Context, uid int) (user.User, error) {

	users, err := rp.FindUsers(ctx, []int{uid})
	if err != nil {
		return user.User{}, err
	}

	u, ok := users[uid]
	if !ok {
		return user.User{}, user.ErrUserNotExists
	}

	return u, nil
}

func (rp *userRepo) FindUsers(ctx context.Context, uid []int) (map[int]user.User, error) {
	var records []userPO
	err := rp.db.WithContext(ctx).
		Where("id IN ?", uid).
		Find(&records).Error
	if err != nil {
		return nil, err
	}

	users := make(map[int]user.User, len(records))
	for _, r := range records {

		users[r.ID] = user.User{
			Uid:    r.ID,
			Name:   r.Name,
			Avatar: r.Avatar,
		}
	}
	return users, nil
}
