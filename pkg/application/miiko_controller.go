package application

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/ningenMe/miiko-api/pkg/domain"
	"github.com/ningenMe/miiko-api/pkg/infra"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
)

type MiikoController struct{}

var categoryRepository = infra.CategoryRepository{}
var topicRepository = infra.TopicRepository{}
var problemRepository = infra.ProblemRepository{}
var kiwaApiRepository = infra.KiwaApiRepository{}

var categoryUsecase = CategoryUsecase{}
var topicUsecase = TopicUsecase{}
var authorizationService = domain.AuthorizationService{}

func (s *MiikoController) CategoryListGet(
	ctx context.Context,
	req *connect.Request[miikov1.CategoryListGetRequest],
) (*connect.Response[miikov1.CategoryListGetResponse], error) {

	body, err := categoryUsecase.CategoryListGet(req.Msg.IsRequiredTopic)

	return connect.NewResponse[miikov1.CategoryListGetResponse](body), err
}

func (s *MiikoController) CategoryPost(
	ctx context.Context,
	req *connect.Request[miikov1.CategoryPostRequest],
) (*connect.Response[miikov1.CategoryPostResponse], error) {

	body, err := categoryUsecase.CategoryPost(req.Header(), req.Msg.CategoryId, req.Msg.GetCategory())

	return connect.NewResponse[miikov1.CategoryPostResponse](
		body,
	), err
}

func (s *MiikoController) TopicListGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicListGetRequest],
) (*connect.Response[miikov1.TopicListGetResponse], error) {

	body, err := topicUsecase.TopicListGet(req.Msg.CategorySystemName)

	return connect.NewResponse[miikov1.TopicListGetResponse](body), err
}

func (s *MiikoController) TopicPost(
	ctx context.Context,
	req *connect.Request[miikov1.TopicPostRequest],
) (*connect.Response[miikov1.TopicPostResponse], error) {
	if err := authorizationService.CheckLoggedIn(req.Header()); err != nil {
		return connect.NewResponse[miikov1.TopicPostResponse](
			&miikov1.TopicPostResponse{},
		), err
	}

	topicId := req.Msg.TopicId
	//更新
	if topicId == "" {
		topicId = infra.GetNewTopicId()
	}

	topicRepository.Upsert(
		&infra.TopicDto{
			TopicId:          topicId,
			CategoryId:       req.Msg.CategoryId,
			TopicDisplayName: req.Msg.GetTopic().GetTopicDisplayName(),
			TopicOrder:       req.Msg.GetTopic().GetTopicOrder(),
		})

	return connect.NewResponse[miikov1.TopicPostResponse](
		&miikov1.TopicPostResponse{
			TopicId: topicId,
		},
	), nil
}

func (s *MiikoController) ProblemListGet(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemListGetRequest],
) (*connect.Response[miikov1.ProblemListGetResponse], error) {
	list := problemRepository.GetProblemList(
		req.Msg.Offset,
		req.Msg.Limit,
	)

	var viewProblemList []*miikov1.Problem
	for _, problem := range list {

		var viewTagList []*miikov1.Tag
		for _, tag := range problem.TagList {
			viewTagList = append(viewTagList, &miikov1.Tag{
				TopicId:          tag.TopicId,
				CategoryId:       tag.CategoryId,
				TopicDisplayName: tag.TopicDisplayName,
			})
		}

		viewProblemList = append(viewProblemList, &miikov1.Problem{
			ProblemId:          problem.ProblemId,
			Url:                problem.Url,
			ProblemDisplayName: problem.ProblemDisplayName,
			Estimation:         problem.Estimation,
			TagList:            viewTagList,
		})
	}

	return connect.NewResponse[miikov1.ProblemListGetResponse](
		&miikov1.ProblemListGetResponse{
			ProblemList: viewProblemList,
		}), nil
}

func (s *MiikoController) ProblemGet(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemGetRequest],
) (*connect.Response[miikov1.ProblemGetResponse], error) {

	problem := problemRepository.GetProblem(req.Msg.ProblemId)

	var viewTagList []*miikov1.Tag
	for _, tag := range problem.TagList {
		viewTagList = append(viewTagList, &miikov1.Tag{
			TopicId:          tag.TopicId,
			CategoryId:       tag.CategoryId,
			TopicDisplayName: tag.TopicDisplayName,
		})
	}

	return connect.NewResponse[miikov1.ProblemGetResponse](
		&miikov1.ProblemGetResponse{
			Problem: &miikov1.Problem{
				ProblemId:          problem.ProblemId,
				Url:                problem.Url,
				ProblemDisplayName: problem.ProblemDisplayName,
				Estimation:         problem.Estimation,
				TagList:            viewTagList,
			},
		}), nil
}

func (s *MiikoController) ProblemPost(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemPostRequest],
) (*connect.Response[miikov1.ProblemPostResponse], error) {
	if err := authorizationService.CheckLoggedIn(req.Header()); err != nil {
		return connect.NewResponse[miikov1.ProblemPostResponse](
			&miikov1.ProblemPostResponse{},
		), err
	}

	problemId := req.Msg.ProblemId
	//新規でidを払い出す
	if problemId == "" {
		problemId = infra.GetNewProblemId()
		//urlが同じなのに新規作成してたらおかしいのでエラーにする
		problemWithUrl := problemRepository.GetProblemByUrl(req.Msg.GetProblem().GetUrl())
		if problemWithUrl != nil {
			return connect.NewResponse[miikov1.ProblemPostResponse](
				&miikov1.ProblemPostResponse{},
			), fmt.Errorf("url is duplicated")
		}
	}

	var tagList []*infra.TagDto
	for _, tag := range req.Msg.GetProblem().GetTagList() {
		tagList = append(tagList, &infra.TagDto{
			TopicId: tag.GetTopicId(),
		})
	}

	problem := &infra.ProblemDto{
		ProblemId:          problemId,
		Url:                req.Msg.GetProblem().GetUrl(),
		ProblemDisplayName: req.Msg.GetProblem().GetProblemDisplayName(),
		Estimation:         req.Msg.GetProblem().GetEstimation(),
		TagList:            tagList,
	}

	//NOTE タグがない場合は未登録タグに入れる
	if len(problem.TagList) == 0 {
		problem.TagList = append(problem.TagList, &infra.TagDto{
			TopicId: "topic_000266",
		})
	}

	problemRepository.Upsert(problem)
	problemRepository.DeleteTag(problemId)
	for _, tag := range problem.TagList {
		problemRepository.InsertTag(problemId, tag.TopicId)
	}

	return connect.NewResponse[miikov1.ProblemPostResponse](
		&miikov1.ProblemPostResponse{
			ProblemId: problemId,
		},
	), nil
}

func (s *MiikoController) TopicGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicGetRequest],
) (*connect.Response[miikov1.TopicGetResponse], error) {

	topic := topicRepository.Get(req.Msg.TopicId)

	var viewProblemList []*miikov1.Problem
	for _, problem := range topic.ProblemList {

		var viewTagList []*miikov1.Tag
		for _, tag := range problem.TagList {
			viewTagList = append(viewTagList, &miikov1.Tag{
				TopicId:          tag.TopicId,
				CategoryId:       tag.CategoryId,
				TopicDisplayName: tag.TopicDisplayName,
			})
		}

		viewProblemList = append(viewProblemList, &miikov1.Problem{
			ProblemId:          problem.ProblemId,
			Url:                problem.Url,
			ProblemDisplayName: problem.ProblemDisplayName,
			Estimation:         problem.Estimation,
			TagList:            viewTagList,
		})
	}
	categoryList := categoryRepository.GetList()
	var viewCategory *miikov1.Category
	for _, category := range categoryList {
		if category.CategoryId == topic.CategoryId {
			viewCategory = &miikov1.Category{
				CategoryId:          category.CategoryId,
				CategorySystemName:  category.CategorySystemName,
				CategoryDisplayName: category.CategoryDisplayName,
				CategoryOrder:       category.CategoryOrder,
				TopicSize:           category.TopicSize,
				ProblemSize:         category.ProblemSize,
			}
		}
	}

	return connect.NewResponse[miikov1.TopicGetResponse](
		&miikov1.TopicGetResponse{
			Category: viewCategory,
			Topic: &miikov1.Topic{
				TopicId:          topic.TopicId,
				TopicDisplayName: topic.TopicDisplayName,
				TopicOrder:       topic.TopicOrder,
				ProblemList:      viewProblemList,
			},
		}), nil
}
