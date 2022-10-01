package bizangela

import (
	"context"

	"github.com/tuihub/librarian/app/sephirah/internal/biz/biztiphereth"
	"github.com/tuihub/librarian/internal/lib/libmq"
	mapper "github.com/tuihub/protos/pkg/librarian/mapper/v1"
	porter "github.com/tuihub/protos/pkg/librarian/porter/v1"
	searcher "github.com/tuihub/protos/pkg/librarian/searcher/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAngela, NewAngelaBase, NewPullAccountTopic, NewPullSteamAccountAppRelationTopic)

type Angela struct{}
type AngelaBase struct {
	mapper   mapper.LibrarianMapperServiceClient
	searcher searcher.LibrarianSearcherServiceClient
	porter   porter.LibrarianPorterServiceClient
}

func NewAngelaBase(
	mClient mapper.LibrarianMapperServiceClient,
	pClient porter.LibrarianPorterServiceClient,
	sClient searcher.LibrarianSearcherServiceClient,
) (*AngelaBase, error) {
	return &AngelaBase{
		mapper:   mClient,
		porter:   pClient,
		searcher: sClient,
	}, nil
}

func NewAngela(
	mq *libmq.MQ,
	pullAccount *libmq.TopicImpl[biztiphereth.PullAccountInfo],
	pullSteamAccountAppRelation *libmq.TopicImpl[PullSteamAccountAppRelation],
) (*Angela, error) {
	if err := mq.RegisterTopic(pullAccount); err != nil {
		return nil, err
	}
	if err := mq.RegisterTopic(pullSteamAccountAppRelation); err != nil {
		return nil, err
	}
	return &Angela{}, nil
}

func NewPullAccountTopic(
	a *AngelaBase,
	sr *libmq.TopicImpl[PullSteamAccountAppRelation],
) *libmq.TopicImpl[biztiphereth.PullAccountInfo] {
	return libmq.NewTopic[biztiphereth.PullAccountInfo](
		"PullAccountInfo",
		func() biztiphereth.PullAccountInfo {
			return biztiphereth.PullAccountInfo{}
		},
		func(info biztiphereth.PullAccountInfo) error {
			resp, err := a.porter.PullAccount(context.TODO(), &porter.PullAccountRequest{AccountId: &librarian.AccountID{
				Platform:          biztiphereth.ToLibrarianAccountPlatform(info.Platform),
				PlatformAccountId: info.PlatformAccountID,
			}})
			if err != nil {
				return err
			}
			switch info.Platform {
			case biztiphereth.AccountPlatformSteam:
				// TODO save account data
				resp.GetAccount()
				return sr.
					Publish(PullSteamAccountAppRelation{SteamID: info.PlatformAccountID})
			default:
				return nil
			}
		},
	)
}

func NewPullSteamAccountAppRelationTopic(a *AngelaBase) *libmq.TopicImpl[PullSteamAccountAppRelation] {
	return libmq.NewTopic[PullSteamAccountAppRelation](
		"PullSteamAccountAppRelation",
		func() PullSteamAccountAppRelation {
			return PullSteamAccountAppRelation{}
		},
		func(r PullSteamAccountAppRelation) error {
			panic("impl me")
		},
	)
}
