package user

import "context"

type User struct {
	Uid    int
	Name   string
	Avatar string
}

type UserRepo interface {
	CreateUser(ctx context.Context, name, avatar string) (int, error)
	GetUser(ctx context.Context, uid int) (*User, error)
}
