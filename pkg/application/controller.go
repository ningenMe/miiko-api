package application

import (
	"context"
	"github.com/bufbuild/connect-go"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MiikoController struct{}

func (s *MiikoController) CategoryGet(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[miikov1.CategoryGetResponse], error) {
	var list []*miikov1.Category
	list = append(list, &miikov1.Category{
		CategoryId:          "a",
		CategorySystemName:  "b",
		CategoryDisplayName: "c",
		CategoryOrder:       1,
	})
	res := &miikov1.CategoryGetResponse{
		CategoryList: list,
	}
	return connect.NewResponse[miikov1.CategoryGetResponse](res), nil
}

func (s *MiikoController) CategoryPost(
	ctx context.Context,
	req *connect.Request[miikov1.CategoryPostRequest],
) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse[emptypb.Empty](
		&emptypb.Empty{},
	), nil
}

func (s *MiikoController) TopicGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicGetRequest],
) (*connect.Response[miikov1.TopicGetResponse], error) {
	return connect.NewResponse[miikov1.TopicGetResponse](&miikov1.TopicGetResponse{}), nil
}

func (s *MiikoController) HealthGet(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[miikov1.HealthGetResponse], error) {
	return connect.NewResponse[miikov1.HealthGetResponse](&miikov1.HealthGetResponse{
		Message: "ok",
	}), nil
}
