package main

import (
	"context"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"

	"github.com/go-kratos/kratos/v2/log"
)

func (c *Client) TestGebura(ctx context.Context) { //nolint:funlen,gocognit // no need
	var appID, appID2 *librarian.InternalID
	if resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{App: &librarian.App{
		Id:               nil,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Details:          nil,
		Name:             "test app 1",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description",
		IconImageUrl:     "",
		HeroImageUrl:     "",
		Tags:             nil,
		AltNames:         nil,
	}}); err != nil {
		log.Fatal(err)
	} else {
		appID = resp.GetId()
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
		log.Fatal(err)
	} else if len(resp.GetApps()) != 1 ||
		resp.GetApps()[0].GetId().GetId() != appID.GetId() {
		log.Fatal("inconsistent app id")
	}
	if _, err := c.cli.UpdateApp(ctx, &pb.UpdateAppRequest{App: &librarian.App{
		Id:               appID,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Details:          nil,
		Name:             "test app 1",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description update",
		IconImageUrl:     "",
		HeroImageUrl:     "",
		Tags:             nil,
		AltNames:         nil,
	}}); err != nil {
		log.Fatal(err)
	}
	if resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{App: &librarian.App{
		Id:               nil,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Details:          nil,
		Name:             "test app 2",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description",
		IconImageUrl:     "",
		HeroImageUrl:     "",
		Tags:             nil,
		AltNames:         nil,
	}}); err != nil {
		log.Fatal(err)
	} else {
		appID2 = resp.GetId()
	}
	if _, err := c.cli.SearchApps(ctx, &pb.SearchAppsRequest{
		Paging:   defaultPaging,
		Keywords: "2",
	}); err != nil {
		log.Fatal(err)
	}
	if _, err := c.cli.GetBindApps(ctx, &pb.GetBindAppsRequest{AppId: appID2}); err != nil {
		log.Fatal(err)
	}
	if _, err := c.cli.PurchaseApp(ctx, &pb.PurchaseAppRequest{AppId: appID2}); err != nil {
		log.Fatal(err)
	}
	if resp, err := c.cli.GetPurchasedApps(ctx, &pb.GetPurchasedAppsRequest{}); err != nil {
		log.Fatal(err)
	} else if len(resp.GetApps()) != 1 || resp.GetApps()[0].GetId().GetId() != appID2.GetId() {
		log.Fatal("unexpected search result")
	}
	if _, err := c.cli.MergeApps(ctx, &pb.MergeAppsRequest{
		Base: &librarian.App{
			Id:               appID,
			Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
			SourceAppId:      "",
			SourceUrl:        nil,
			Details:          nil,
			Name:             "test app 1",
			Type:             librarian.AppType_APP_TYPE_GAME,
			ShortDescription: "test app description update",
			IconImageUrl:     "",
			HeroImageUrl:     "",
			Tags:             nil,
			AltNames:         nil,
		},
		Merged: appID2,
	}); err != nil {
		log.Fatal(err)
	}
	if resp, err := c.cli.GetPurchasedApps(ctx, &pb.GetPurchasedAppsRequest{}); err != nil {
		log.Fatal(err)
	} else if len(resp.GetApps()) != 1 || resp.GetApps()[0].GetId().GetId() != appID.GetId() {
		log.Fatal("unexpected search result")
	}
	if resp, err := c.cli.ListAppPackages(ctx, &pb.ListAppPackagesRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		SourceFilter:        nil,
		IdFilter:            nil,
		AssignedAppIdFilter: nil,
	}); err != nil {
		log.Fatal(err)
	} else if resp.GetPaging().GetTotalSize() != 0 {
		log.Fatal("unexpected app package list result")
	}

	resp, err := c.cli.CreateAppPackage(ctx, &pb.CreateAppPackageRequest{
		AppPackage: &librarian.AppPackage{
			Id:          nil,
			Source:      0,
			SourceId:    nil,
			Name:        "test app package",
			Description: "test",
			Binary:      nil,
			Public:      false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	appPackageID := resp.GetId()
	if resp2, err2 := c.cli.ListAppPackages(ctx, &pb.ListAppPackagesRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		SourceFilter:        nil,
		IdFilter:            nil,
		AssignedAppIdFilter: nil,
	}); err2 != nil {
		log.Fatal(err2)
	} else if resp2.GetPaging().GetTotalSize() != 1 || resp2.GetAppPackages()[0].GetDescription() != "test" {
		log.Fatal("unexpected app package list result")
	}
	if _, err2 := c.cli.UpdateAppPackage(ctx, &pb.UpdateAppPackageRequest{
		AppPackage: &librarian.AppPackage{
			Id:          appPackageID,
			Source:      0,
			SourceId:    nil,
			Name:        "test app package",
			Description: "test2",
			Binary:      nil,
			Public:      false,
		},
	}); err2 != nil {
		log.Fatal(err2)
	}
	if resp2, err2 := c.cli.ListAppPackages(ctx, &pb.ListAppPackagesRequest{
		Paging: &librarian.PagingRequest{
			PageNum:  1,
			PageSize: 1,
		},
		SourceFilter:        nil,
		IdFilter:            nil,
		AssignedAppIdFilter: nil,
	}); err2 != nil {
		log.Fatal(err2)
	} else if resp2.GetPaging().GetTotalSize() != 1 || resp2.GetAppPackages()[0].GetDescription() != "test2" {
		log.Fatal("unexpected app package list result")
	}
}
