package application

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/ningenMe/miiko-api/pkg/domain"
	"github.com/ningenMe/miiko-api/pkg/infra"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
)

type MiikoController struct{}

// TODO 依存の置き場所を考える
var categoryRepository = infra.CategoryRepository{}
var topicRepository = infra.TopicRepository{}
var problemRepository = infra.ProblemRepository{}

var categoryUsecase = CategoryUsecase{}
var topicUsecase = TopicUsecase{}
var problemUsecase = ProblemUsecase{}
var statisticsUsecase = StatisticsUsecase{}
var authorizationService = domain.AuthorizationService{}
var urlUsecase = UrlUsecase{}

func (controller *MiikoController) CategoryListGet(
	ctx context.Context,
	req *connect.Request[miikov1.CategoryListGetRequest],
) (*connect.Response[miikov1.CategoryListGetResponse], error) {

	body, err := categoryUsecase.CategoryListGet(req.Msg.IsRequiredTopic)

	return connect.NewResponse[miikov1.CategoryListGetResponse](body), err
}

func (controller *MiikoController) CategoryPost(
	ctx context.Context,
	req *connect.Request[miikov1.CategoryPostRequest],
) (*connect.Response[miikov1.CategoryPostResponse], error) {

	body, err := categoryUsecase.CategoryPost(req.Header(), req.Msg.CategoryId, req.Msg.GetCategory())

	return connect.NewResponse[miikov1.CategoryPostResponse](
		body,
	), err
}

func (controller *MiikoController) TopicListGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicListGetRequest],
) (*connect.Response[miikov1.TopicListGetResponse], error) {

	body, err := topicUsecase.TopicListGet(req.Msg.CategorySystemName)

	return connect.NewResponse[miikov1.TopicListGetResponse](body), err
}

func (controller *MiikoController) TopicPost(
	ctx context.Context,
	req *connect.Request[miikov1.TopicPostRequest],
) (*connect.Response[miikov1.TopicPostResponse], error) {

	body, err := topicUsecase.TopicPost(req.Header(), req.Msg.TopicId, req.Msg.GetTopic(), req.Msg.CategoryId)

	return connect.NewResponse[miikov1.TopicPostResponse](body), err
}

func (controller *MiikoController) TopicGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicGetRequest],
) (*connect.Response[miikov1.TopicGetResponse], error) {

	body, err := topicUsecase.TopicGet(req.Msg.TopicId)

	return connect.NewResponse[miikov1.TopicGetResponse](body), err
}

func (controller *MiikoController) ProblemListGet(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemListGetRequest],
) (*connect.Response[miikov1.ProblemListGetResponse], error) {

	body, err := problemUsecase.ProblemListGet(req.Msg.Limit, req.Msg.Offset, req.Msg.IsDesc)

	return connect.NewResponse[miikov1.ProblemListGetResponse](body), err
}

func (controller *MiikoController) ProblemGet(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemGetRequest],
) (*connect.Response[miikov1.ProblemGetResponse], error) {

	body, err := problemUsecase.ProblemGet(req.Msg.ProblemId)

	return connect.NewResponse[miikov1.ProblemGetResponse](body), err
}

func (controller *MiikoController) ProblemPost(
	ctx context.Context,
	req *connect.Request[miikov1.ProblemPostRequest],
) (*connect.Response[miikov1.ProblemPostResponse], error) {

	body, err := problemUsecase.ProblemPost(req.Header(), req.Msg.ProblemId, req.Msg.GetProblem())

	return connect.NewResponse[miikov1.ProblemPostResponse](body), err
}

func (controller *MiikoController) StatisticsGet(
	ctx context.Context,
	req *connect.Request[miikov1.StatisticsGetRequest],
) (*connect.Response[miikov1.StatisticsGetResponse], error) {

	body, err := statisticsUsecase.StatisticsGet()

	return connect.NewResponse[miikov1.StatisticsGetResponse](body), err
}

func (controller *MiikoController) UrlGet(
	ctx context.Context,
	req *connect.Request[miikov1.UrlGetRequest],
) (*connect.Response[miikov1.UrlGetResponse], error) {

	body, err := urlUsecase.UrlGet(req.Header(), req.Msg.Url)

	return connect.NewResponse[miikov1.UrlGetResponse](body), err
}
