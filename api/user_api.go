package api

import "net/http"

type (
	GetUserReq struct {
		UID uint32 `json:"uid"`
	}

	GetUserReply struct {
		UID    uint32 `json:"uid"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	}

	CreateUserReq struct {
		Username  string `json:"username"`
		AvatarURL string `json:"avatar"`
	}

	CreateUserReply struct {
		UID uint32 `json:"uid"`
	}
)

func makeGetUser(svc TwitterService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		var r GetUserReq
		if err := decodeRequest(req, &r); err != nil {
			return
		}

		result, err := svc.GetUser(req.Context(), &r)
		if err != nil {
			return
		}

		resp := &GetUserReply{
			UID:    result.UID,
			Name:   result.Name,
			Avatar: result.Avatar,
		}

		encodeResponse(w, resp)
	}
}

func makeCreateUser(svc TwitterService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var r CreateUserReq
		if err := decodeRequest(req, &r); err != nil {
			return
		}

		result, err := svc.CreateUser(req.Context(), &r)
		if err != nil {
			return
		}

		resp := &CreateUserReply{UID: result.UID}

		encodeResponse(w, resp)
	}
}
