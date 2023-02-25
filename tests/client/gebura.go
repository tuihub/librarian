package main

import (
	"context"

	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"
	librarian "github.com/tuihub/protos/pkg/librarian/v1"
)

func (c *Client) TestGebura(ctx context.Context) {
	var appID *librarian.InternalID
	if resp, err := c.cli.CreateApp(ctx, &pb.CreateAppRequest{App: &librarian.App{
		Id:               nil,
		Source:           librarian.AppSource_APP_SOURCE_INTERNAL,
		SourceAppId:      "",
		SourceUrl:        nil,
		Name:             "test app 1",
		Type:             librarian.AppType_APP_TYPE_GAME,
		ShortDescription: "test app description",
		ImageUrl:         "",
		Details:          nil,
	}}); err != nil {
		panic(err)
	} else {
		appID = resp.Id
	}
	if resp, err := c.cli.ListApp(ctx, &pb.ListAppRequest{
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
	} else if len(resp.GetAppList()) != 1 ||
		resp.GetAppList()[0].GetId().GetId() != appID.GetId() {
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
		ImageUrl:         "",
		Details:          nil,
	}}); err != nil {
		panic(err)
	}
}
