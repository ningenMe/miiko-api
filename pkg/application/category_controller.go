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
var topicRepository = infra.TopicRepository{}

func (s *MiikoController) CategoryListGet(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[miikov1.CategoryListGetResponse], error) {

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

	res := &miikov1.CategoryListGetResponse{
		CategoryList: viewList,
	}
	return connect.NewResponse[miikov1.CategoryListGetResponse](res), nil
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

func (s *MiikoController) TopicListGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicListGetRequest],
) (*connect.Response[miikov1.TopicListGetResponse], error) {

	categorySystemName := req.Msg.CategorySystemName
	categoryDto := categoryRepository.Get(categorySystemName)

	var list []*infra.TopicDto
	var category *miikov1.Category
	if categoryDto != nil {
		category = &miikov1.Category{
			CategoryId:          categoryDto.CategoryId,
			CategorySystemName:  categoryDto.CategorySystemName,
			CategoryDisplayName: categoryDto.CategoryDisplayName,
			CategoryOrder:       categoryDto.CategoryOrder,
		}
		list = topicRepository.GetListByCategoryId(categoryDto.CategoryId)
	}

	var viewTopicList []*miikov1.Topic
	for _, topic := range list {

		var viewProblemList []*miikov1.Problem
		for _, problem := range topic.ProblemList {
			viewProblemList = append(viewProblemList, &miikov1.Problem{
				ProblemId:          problem.ProblemId,
				Url:                problem.Url,
				ProblemDisplayName: problem.ProblemDisplayName,
				Estimation:         problem.Estimation,
			})
		}

		viewTopicList = append(viewTopicList, &miikov1.Topic{
			TopicId:          topic.TopicId,
			TopicDisplayName: topic.TopicDisplayName,
			TopicOrder:       topic.TopicOrder,
			ProblemList:      viewProblemList,
		})
	}

	return connect.NewResponse[miikov1.TopicListGetResponse](&miikov1.TopicListGetResponse{
		Category:  category,
		TopicList: viewTopicList,
	}), nil
}
