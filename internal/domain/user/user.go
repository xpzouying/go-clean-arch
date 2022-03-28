package user

import (
	"context"
	"errors"
)

var (
	ErrUserNotExists = errors.New("user not exists")
)

type User struct {
	Uid    int
	Name   string
	Avatar string
}

type UserRepo interface {
	CreateUser(ctx context.Context, name, avatar string) (int, error)
	GetUser(ctx context.Context, uid int) (User, error)
	FindUsers(ctx context.Context, uid []int) (map[int]User, error)
}

func IsErrUserNotExists(err error) bool {
	return errors.Is(err, ErrUserNotExists)
}
