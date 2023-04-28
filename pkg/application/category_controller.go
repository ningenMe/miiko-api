package application

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/ningenMe/miiko-api/pkg/infra"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MiikoController struct{}

var categoryRepository = infra.CategoryRepository{}

func (s *MiikoController) CategoryGet(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[miikov1.CategoryGetResponse], error) {

	list := categoryRepository.GetList()

	var viewList []*miikov1.Category
	for _, category := range list {

		viewList = append(viewList, &miikov1.Category{
			CategoryId:          category.CategoryId,
			CategoryDisplayName: category.CategoryDisplayName,
			CategorySystemName:  category.CategorySystemName,
			CategoryOrder:       category.CategoryOrder,
		})
	}

	res := &miikov1.CategoryGetResponse{
		CategoryList: viewList,
	}
	return connect.NewResponse[miikov1.CategoryGetResponse](res), nil
}

func (s *MiikoController) CategoryPost(
	ctx context.Context,
	req *connect.Request[miikov1.CategoryPostRequest],
) (*connect.Response[miikov1.CategoryPostResponse], error) {

	categoryId := req.Msg.CategoryId
	if req.Msg.GetCategory() != nil {
		categoryId = req.Msg.GetCategoryId()
		if categoryId == "" {
			categoryId = infra.GetNewCategoryId()
		}
		categoryRepository.Upsert(
			&infra.CategoryDto{
				CategoryId:          categoryId,
				CategoryDisplayName: req.Msg.GetCategory().GetCategoryDisplayName(),
				CategorySystemName:  req.Msg.GetCategory().GetCategorySystemName(),
				CategoryOrder:       req.Msg.GetCategory().GetCategoryOrder(),
			})
	} else {
		categoryRepository.Delete(categoryId)
	}

	return connect.NewResponse[miikov1.CategoryPostResponse](
		&miikov1.CategoryPostResponse{
			CategoryId: categoryId,
		},
	), nil
}

func (s *MiikoController) TopicGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicGetRequest],
) (*connect.Response[miikov1.TopicGetResponse], error) {
	return connect.NewResponse[miikov1.TopicGetResponse](&miikov1.TopicGetResponse{}), nil
}
