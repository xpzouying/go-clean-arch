package service

import (
	"context"

	"github.com/xpzouying/go-clean-arch/api"
)

func (ts *TwitterService) GetUser(ctx context.Context, req *api.GetUserReq) (*api.GetUserReply, error) {

	user, err := ts.tc.GetUser(ctx, int(req.UID))
	if err != nil {
		return nil, err
	}

	return &api.GetUserReply{
		UID:    uint32(user.Uid),
		Name:   user.Name,
		Avatar: user.Avatar,
	}, nil
}

func (ts *TwitterService) CreateUser(ctx context.Context, req *api.CreateUserReq) (*api.CreateUserReply, error) {

	uid, err := ts.tc.CreateUser(ctx, req.Username, req.AvatarURL)
	if err != nil {
		return nil, err
	}

	return &api.CreateUserReply{UID: uint32(uid)}, nil
}
