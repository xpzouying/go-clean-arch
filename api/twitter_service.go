package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TwitterService interface {
	GetUser(ctx context.Context, req *GetUserReq) (*GetUserReply, error)
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserReply, error)
}

func RegisterTwitterHTTPServer(h *http.ServeMux, svc TwitterService) {

	h.HandleFunc("/get-user", makeGetUser(svc))
	h.HandleFunc("/create-user", makeCreateUser(svc))
}

func encodeResponse(w http.ResponseWriter, in interface{}) {
	data, err := json.Marshal(in)
	if err != nil {
		return
	}

	_, _ = w.Write(data)
}

func decodeRequest(req *http.Request, out interface{}) error {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &out)
}
