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

var categoryUsecase = CategoryUsecase{}
var topicUsecase = TopicUsecase{}
var problemUsecase = ProblemUsecase{}
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

	body, err := topicUsecase.TopicPost(req.Header(), req.Msg.TopicId, req.Msg.GetTopic(), req.Msg.CategoryId)

	return connect.NewResponse[miikov1.TopicPostResponse](body), err
}

func (s *MiikoController) TopicGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicGetRequest],
) (*connect.Response[miikov1.TopicGetResponse], error) {

	body, err := topicUsecase.TopicGet(req.Msg.TopicId)

	return connect.NewResponse[miikov1.TopicGetResponse](body), err
}

func (s *MiikoController) ProblemListGet(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemListGetRequest],
) (*connect.Response[miikov1.ProblemListGetResponse], error) {

	body, err := problemUsecase.ProblemListGet(req.Msg.Offset, req.Msg.Limit)

	return connect.NewResponse[miikov1.ProblemListGetResponse](body), err
}

func (s *MiikoController) ProblemGet(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemGetRequest],
) (*connect.Response[miikov1.ProblemGetResponse], error) {

	body, err := problemUsecase.ProblemGet(req.Msg.ProblemId)

	return connect.NewResponse[miikov1.ProblemGetResponse](body), err
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
