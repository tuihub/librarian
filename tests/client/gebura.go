package main

import (
	"context"
	"strconv"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/log"
)

func (c *Client) TestGebura(ctx context.Context) { //nolint:funlen,gocognit // no need
	var appID, appID2 *librarian.InternalID
	if resp, err := c.cli.CreateAppInfo(ctx, &pb.CreateAppInfoRequest{AppInfo: &librarian.AppInfo{
		Id:                 nil,
		Internal:           true,
		Source:             "",
		SourceAppId:        "",
		SourceUrl:          nil,
		Details:            nil,
		Name:               "test app 1",
		Type:               librarian.AppType_APP_TYPE_GAME,
		ShortDescription:   "test app description",
		IconImageUrl:       "",
		BackgroundImageUrl: "",
		CoverImageUrl:      "",
		Tags:               nil,
		AltNames:           nil,
	}}); err != nil {
		log.Fatal(err)
	} else {
		appID = resp.GetId()
	}
	if resp, err := c.cli.ListAppInfos(ctx, &pb.ListAppInfosRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		ExcludeInternal: false,
		SourceFilter:    nil,
		TypeFilter:      nil,
		IdFilter:        nil,
		ContainDetails:  false,
	}); err != nil {
		log.Fatal(err)
	} else if len(resp.GetAppInfos()) != 1 ||
		resp.GetAppInfos()[0].GetId().GetId() != appID.GetId() {
		log.Fatal("inconsistent app id")
	}
	if _, err := c.cli.UpdateAppInfo(ctx, &pb.UpdateAppInfoRequest{AppInfo: &librarian.AppInfo{
		Id:                 appID,
		Internal:           true,
		Source:             "",
		SourceAppId:        "",
		SourceUrl:          nil,
		Details:            nil,
		Name:               "test app 1",
		Type:               librarian.AppType_APP_TYPE_GAME,
		ShortDescription:   "test app description update",
		IconImageUrl:       "",
		BackgroundImageUrl: "",
		CoverImageUrl:      "",
		Tags:               nil,
		AltNames:           nil,
	}}); err != nil {
		log.Fatal(err)
	}
	if resp, err := c.cli.CreateAppInfo(ctx, &pb.CreateAppInfoRequest{AppInfo: &librarian.AppInfo{
		Id:                 nil,
		Internal:           true,
		Source:             "",
		SourceAppId:        "",
		SourceUrl:          nil,
		Details:            nil,
		Name:               "test app 2",
		Type:               librarian.AppType_APP_TYPE_GAME,
		ShortDescription:   "test app description",
		IconImageUrl:       "",
		BackgroundImageUrl: "",
		CoverImageUrl:      "",
		Tags:               nil,
		AltNames:           nil,
	}}); err != nil {
		log.Fatal(err)
	} else {
		appID2 = resp.GetId()
	}
	if _, err := c.cli.SearchAppInfos(ctx, &pb.SearchAppInfosRequest{
		Paging: defaultPaging,
		Query:  "2",
	}); err != nil {
		log.Fatal(err)
	}
	if _, err := c.cli.GetBoundAppInfos(ctx, &pb.GetBoundAppInfosRequest{AppInfoId: appID2}); err != nil {
		log.Fatal(err)
	}
	if _, err := c.cli.PurchaseAppInfo(ctx, &pb.PurchaseAppInfoRequest{
		AppInfoId: &librarian.AppInfoID{
			Internal:    true,
			Source:      "",
			SourceAppId: strconv.FormatInt(appID2.GetId(), 10),
		},
	}); err != nil {
		log.Fatal(err)
	}
	if resp, err := c.cli.GetPurchasedAppInfos(ctx, &pb.GetPurchasedAppInfosRequest{Source: nil}); err != nil {
		log.Fatal(err)
	} else if len(resp.GetAppInfos()) != 1 || resp.GetAppInfos()[0].GetId().GetId() != appID2.GetId() {
		log.Fatal("unexpected search result")
	}
	if _, err := c.cli.MergeAppInfos(ctx, &pb.MergeAppInfosRequest{
		Base: &librarian.AppInfo{
			Id:                 appID,
			Internal:           true,
			Source:             "",
			SourceAppId:        "",
			SourceUrl:          nil,
			Details:            nil,
			Name:               "test app 1",
			Type:               librarian.AppType_APP_TYPE_GAME,
			ShortDescription:   "test app description update",
			IconImageUrl:       "",
			BackgroundImageUrl: "",
			CoverImageUrl:      "",
			Tags:               nil,
			AltNames:           nil,
		},
		Merged: appID2,
	}); err != nil {
		log.Fatal(err)
	}
	if resp, err := c.cli.GetPurchasedAppInfos(ctx, &pb.GetPurchasedAppInfosRequest{Source: nil}); err != nil {
		log.Fatal(err)
	} else if len(resp.GetAppInfos()) != 1 || resp.GetAppInfos()[0].GetId().GetId() != appID.GetId() {
		log.Fatal("unexpected search result")
	}
	if resp, err := c.cli.ListApps(ctx, &pb.ListAppsRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		OwnerIdFilter:           nil,
		IdFilter:                nil,
		AssignedAppInfoIdFilter: nil,
	}); err != nil {
		log.Fatal(err)
	} else if resp.GetPaging().GetTotalSize() != 0 {
		log.Fatal("unexpected app package list result")
	}

	resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{
		App: &pb.App{
			Id:                nil,
			Name:              "test app package",
			Description:       "test",
			AssignedAppInfoId: nil,
			Public:            false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	appPackageID := resp.GetId()
	if resp2, err2 := c.cli.ListApps(ctx, &pb.ListAppsRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		OwnerIdFilter:           nil,
		IdFilter:                nil,
		AssignedAppInfoIdFilter: nil,
	}); err2 != nil {
		log.Fatal(err2)
	} else if resp2.GetPaging().GetTotalSize() != 1 || resp2.GetApps()[0].GetDescription() != "test" {
		log.Fatal("unexpected app package list result")
	}
	if _, err2 := c.cli.UpdateApp(ctx, &pb.UpdateAppRequest{
		App: &pb.App{
			Id:                appPackageID,
			Name:              "test app package",
			Description:       "test2",
			AssignedAppInfoId: nil,
			Public:            false,
		},
	}); err2 != nil {
		log.Fatal(err2)
	}
	if resp2, err2 := c.cli.ListApps(ctx, &pb.ListAppsRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		OwnerIdFilter:           nil,
		IdFilter:                nil,
		AssignedAppInfoIdFilter: nil,
	}); err2 != nil {
		log.Fatal(err2)
	} else if resp2.GetPaging().GetTotalSize() != 1 || resp2.GetApps()[0].GetDescription() != "test2" {
		log.Fatal("unexpected app package list result")
	}
}
