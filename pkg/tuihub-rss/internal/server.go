package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/tuihub/librarian/pkg/tuihub-go"
	"github.com/tuihub/librarian/pkg/tuihub-rss/internal/converter"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1/porter"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/gorilla/feeds"
)

const defaultPageSize = 20
const defaultServerTimeout = 5 * time.Second

type Server struct {
	mux    *http.ServeMux
	porter *tuihub.Porter
}

func NewServer(porter *tuihub.Porter) (*Server, error) {
	mux := http.NewServeMux()
	server := &Server{
		mux:    mux,
		porter: porter,
	}
	server.routes()
	return server, nil
}

func (s *Server) routes() {
	s.mux.HandleFunc("/", s.handleRSS())
}

func (s *Server) handleRSS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		idStr := r.URL.Path[len("/"):]
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		client, err := s.porter.ReverseCall(ctx)
		if err != nil {
			http.Error(w, "Failed to initialize reverse call", http.StatusInternalServerError)
			return
		}
		resp, err := client.GetNotifyTargetItems(client.WithToken(ctx), &pb.GetNotifyTargetItemsRequest{
			Id: &librarian.InternalID{Id: id},
			Paging: &librarian.PagingRequest{
				PageNum:  1,
				PageSize: defaultPageSize,
			},
		})
		if err != nil {
			http.Error(w, "Feed not found", http.StatusBadRequest)
			return
		}
		dest := resp.GetDestination()
		if dest == nil || dest.GetId() != "rss" {
			http.Error(w, "Feed found but unexpected", http.StatusInternalServerError)
			return
		}
		var config ServeRSSConfig
		err = json.Unmarshal([]byte(dest.GetConfigJson()), &config)
		if err != nil {
			http.Error(w, "Failed to parse feed config", http.StatusInternalServerError)
			return
		}
		feed := feeds.Feed{
			Title:       config.Title,
			Link:        nil,
			Description: "",
			Author:      nil,
			Updated:     time.Time{},
			Created:     time.Time{},
			Id:          "",
			Subtitle:    "",
			Items:       converter.FromPBFeedItems(resp.GetItems()),
			Copyright:   "",
			Image:       nil,
		}
		err = feed.WriteRss(w)
		if err != nil {
			http.Error(w, "Failed to convert feed to RSS", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/rss+xml")
	}
}

func (s *Server) Run(addr string) error {
	server := new(http.Server)
	server.Addr = addr
	server.ReadHeaderTimeout = defaultServerTimeout
	server.Handler = s.mux

	return server.ListenAndServe()
}
