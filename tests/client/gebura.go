package main

import (
	"context"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (c *Client) TestGebura(ctx context.Context) { //nolint:gocognit,funlen // no need
	var appID, appID2 *librarian.InternalID
	if resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{App: &librarian.App{
		Id:               nil,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 1",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description",
		IconImageUrl:     "",
		Tags:             nil,
		Details:          nil,
	}}); err != nil {
		panic(err)
	} else {
		appID = resp.Id
	}
	if resp, err := c.cli.ListApps(ctx, &pb.ListAppsRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		SourceFilter:   nil,
		TypeFilter:     nil,
		IdFilter:       nil,
		ContainDetails: false,
	}); err != nil {
		panic(err)
	} else if len(resp.GetApps()) != 1 ||
		resp.GetApps()[0].GetId().GetId() != appID.GetId() {
		panic("inconsistent app id")
	}
	if _, err := c.cli.UpdateApp(ctx, &pb.UpdateAppRequest{App: &librarian.App{
		Id:               appID,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 1",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description update",
		IconImageUrl:     "",
		Tags:             nil,
		Details:          nil,
	}}); err != nil {
		panic(err)
	}
	if resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{App: &librarian.App{
		Id:               nil,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 2",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description",
		IconImageUrl:     "",
		Tags:             nil,
		Details:          nil,
	}}); err != nil {
		panic(err)
	} else {
		appID2 = resp.Id
	}
	if resp, err := c.cli.SearchApps(ctx, &pb.SearchAppsRequest{
		Paging:   defaultPaging,
		Keywords: "2",
	}); err != nil {
		panic(err)
	} else if len(resp.GetApps()) != 1 || resp.GetApps()[0].GetId().GetId() != appID2.GetId() {
		panic("unexpected search result")
	}
	if _, err := c.cli.GetBindApps(ctx, &pb.GetBindAppsRequest{AppId: appID2}); err != nil {
		panic(err)
	}
	if _, err := c.cli.PurchaseApp(ctx, &pb.PurchaseAppRequest{AppId: appID2}); err != nil {
		panic(err)
	}
	if resp, err := c.cli.GetPurchasedApps(ctx, &pb.GetPurchasedAppsRequest{}); err != nil {
		panic(err)
	} else if len(resp.GetApps()) != 1 || resp.GetApps()[0].Id.GetId() != appID2.GetId() {
		panic("unexpected search result")
	}
	if _, err := c.cli.MergeApps(ctx, &pb.MergeAppsRequest{
		Base: &librarian.App{
			Id:               appID,
			Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
			SourceAppId:      "",
			SourceUrl:        nil,
			Name:             "test app 1",
			Type:             librarian.AppType_APP_TYPE_GAME,
			ShortDescription: "test app description update",
			IconImageUrl:     "",
			Tags:             nil,
			Details:          nil,
		},
		Merged: appID2,
	}); err != nil {
		panic(err)
	}
	if resp, err := c.cli.GetPurchasedApps(ctx, &pb.GetPurchasedAppsRequest{}); err != nil {
		panic(err)
	} else if len(resp.GetApps()) != 1 || resp.GetApps()[0].Id.GetId() != appID.GetId() {
		panic("unexpected search result")
	}
}
