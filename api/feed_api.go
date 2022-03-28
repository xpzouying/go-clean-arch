package api

import "net/http"

type (
	FeedInfo struct {
		FID  uint32 `json:"fid"`
		Text string `json:"text"`
	}

	AuthorInfo struct {
		UID          uint32 `json:"uid"`
		AuthorName   string `json:"author_name"`
		AuthorAvatar string `json:"author_avatar"`
	}

	CardInfo struct {
		AuthorInfo
		FeedInfo
	}

	CreateFeedReq struct {
		UID  uint32 `json:"uid"`
		Text string `json:"text"`
	}

	CreateFeedReply struct {
		CardInfo
	}

	DeleteFeedReq struct {
		UID    uint32 `json:"uid"`
		FeedID uint32 `json:"fid"`
	}

	ListFeedReply struct {
		Feeds []CardInfo `json:"feeds"`
	}
)

func makeCreateFeed(svc TwitterService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		var r CreateFeedReq
		if err := decodeRequest(req, &r); err != nil {
			return
		}

		result, err := svc.CreateFeed(req.Context(), &r)
		if err != nil {
			return
		}

		encodeResponse(w, result)
	}
}

func makeDeleteFeed(svc TwitterService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var r DeleteFeedReq
		if err := decodeRequest(req, &r); err != nil {
			return
		}

		if err := svc.DeleteFeed(req.Context(), &r); err != nil {
			return
		}

		encodeResponse(w, map[string]interface{}{"result": true})
	}
}

func makeListFeed(svc TwitterService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		result, err := svc.ListFeeds(req.Context())
		if err != nil {
			return
		}

		encodeResponse(w, result)
	}
}
