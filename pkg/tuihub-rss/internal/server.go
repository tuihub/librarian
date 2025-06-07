package internal

import (
	"context"
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
	addr   string
	server *http.Server
}

func NewServer(porter *tuihub.Porter, addr string) (*Server, error) {
	mux := http.NewServeMux()
	server := &Server{
		mux:    mux,
		porter: porter,
		addr:   addr,
		server: nil,
	}
	server.routes()
	return server, nil
}

func (s *Server) SetPorter(porter *tuihub.Porter) {
	s.porter = porter
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

func (s *Server) Start(ctx context.Context) error {
	server := new(http.Server)
	server.Addr = s.addr
	server.ReadHeaderTimeout = defaultServerTimeout
	server.Handler = s.mux
	server.IdleTimeout = defaultServerTimeout
	server.ReadTimeout = defaultServerTimeout
	server.WriteTimeout = defaultServerTimeout

	s.server = server

	return server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	if s.server == nil {
		return nil
	}
	return s.server.Shutdown(ctx)
}
