// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/miiko.proto

package miikov1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// MiikoServiceName is the fully-qualified name of the MiikoService service.
	MiikoServiceName = "miiko.v1.MiikoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MiikoServiceCategoryListGetProcedure is the fully-qualified name of the MiikoService's
	// CategoryListGet RPC.
	MiikoServiceCategoryListGetProcedure = "/miiko.v1.MiikoService/CategoryListGet"
	// MiikoServiceCategoryPostProcedure is the fully-qualified name of the MiikoService's CategoryPost
	// RPC.
	MiikoServiceCategoryPostProcedure = "/miiko.v1.MiikoService/CategoryPost"
	// MiikoServiceTopicListGetProcedure is the fully-qualified name of the MiikoService's TopicListGet
	// RPC.
	MiikoServiceTopicListGetProcedure = "/miiko.v1.MiikoService/TopicListGet"
	// MiikoServiceTopicGetProcedure is the fully-qualified name of the MiikoService's TopicGet RPC.
	MiikoServiceTopicGetProcedure = "/miiko.v1.MiikoService/TopicGet"
	// MiikoServiceTopicPostProcedure is the fully-qualified name of the MiikoService's TopicPost RPC.
	MiikoServiceTopicPostProcedure = "/miiko.v1.MiikoService/TopicPost"
	// MiikoServiceProblemListGetProcedure is the fully-qualified name of the MiikoService's
	// ProblemListGet RPC.
	MiikoServiceProblemListGetProcedure = "/miiko.v1.MiikoService/ProblemListGet"
	// MiikoServiceProblemGetProcedure is the fully-qualified name of the MiikoService's ProblemGet RPC.
	MiikoServiceProblemGetProcedure = "/miiko.v1.MiikoService/ProblemGet"
	// MiikoServiceProblemPostProcedure is the fully-qualified name of the MiikoService's ProblemPost
	// RPC.
	MiikoServiceProblemPostProcedure = "/miiko.v1.MiikoService/ProblemPost"
	// MiikoServiceStatisticsGetProcedure is the fully-qualified name of the MiikoService's
	// StatisticsGet RPC.
	MiikoServiceStatisticsGetProcedure = "/miiko.v1.MiikoService/StatisticsGet"
	// MiikoServiceUrlGetProcedure is the fully-qualified name of the MiikoService's UrlGet RPC.
	MiikoServiceUrlGetProcedure = "/miiko.v1.MiikoService/UrlGet"
)

// MiikoServiceClient is a client for the miiko.v1.MiikoService service.
type MiikoServiceClient interface {
	CategoryListGet(context.Context, *connect_go.Request[v1.CategoryListGetRequest]) (*connect_go.Response[v1.CategoryListGetResponse], error)
	CategoryPost(context.Context, *connect_go.Request[v1.CategoryPostRequest]) (*connect_go.Response[v1.CategoryPostResponse], error)
	TopicListGet(context.Context, *connect_go.Request[v1.TopicListGetRequest]) (*connect_go.Response[v1.TopicListGetResponse], error)
	TopicGet(context.Context, *connect_go.Request[v1.TopicGetRequest]) (*connect_go.Response[v1.TopicGetResponse], error)
	TopicPost(context.Context, *connect_go.Request[v1.TopicPostRequest]) (*connect_go.Response[v1.TopicPostResponse], error)
	ProblemListGet(context.Context, *connect_go.Request[v1.ProblemListGetRequest]) (*connect_go.Response[v1.ProblemListGetResponse], error)
	ProblemGet(context.Context, *connect_go.Request[v1.ProblemGetRequest]) (*connect_go.Response[v1.ProblemGetResponse], error)
	ProblemPost(context.Context, *connect_go.Request[v1.ProblemPostRequest]) (*connect_go.Response[v1.ProblemPostResponse], error)
	StatisticsGet(context.Context, *connect_go.Request[v1.StatisticsGetRequest]) (*connect_go.Response[v1.StatisticsGetResponse], error)
	UrlGet(context.Context, *connect_go.Request[v1.UrlGetRequest]) (*connect_go.Response[v1.UrlGetResponse], error)
}

// NewMiikoServiceClient constructs a client for the miiko.v1.MiikoService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMiikoServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MiikoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &miikoServiceClient{
		categoryListGet: connect_go.NewClient[v1.CategoryListGetRequest, v1.CategoryListGetResponse](
			httpClient,
			baseURL+MiikoServiceCategoryListGetProcedure,
			opts...,
		),
		categoryPost: connect_go.NewClient[v1.CategoryPostRequest, v1.CategoryPostResponse](
			httpClient,
			baseURL+MiikoServiceCategoryPostProcedure,
			opts...,
		),
		topicListGet: connect_go.NewClient[v1.TopicListGetRequest, v1.TopicListGetResponse](
			httpClient,
			baseURL+MiikoServiceTopicListGetProcedure,
			opts...,
		),
		topicGet: connect_go.NewClient[v1.TopicGetRequest, v1.TopicGetResponse](
			httpClient,
			baseURL+MiikoServiceTopicGetProcedure,
			opts...,
		),
		topicPost: connect_go.NewClient[v1.TopicPostRequest, v1.TopicPostResponse](
			httpClient,
			baseURL+MiikoServiceTopicPostProcedure,
			opts...,
		),
		problemListGet: connect_go.NewClient[v1.ProblemListGetRequest, v1.ProblemListGetResponse](
			httpClient,
			baseURL+MiikoServiceProblemListGetProcedure,
			opts...,
		),
		problemGet: connect_go.NewClient[v1.ProblemGetRequest, v1.ProblemGetResponse](
			httpClient,
			baseURL+MiikoServiceProblemGetProcedure,
			opts...,
		),
		problemPost: connect_go.NewClient[v1.ProblemPostRequest, v1.ProblemPostResponse](
			httpClient,
			baseURL+MiikoServiceProblemPostProcedure,
			opts...,
		),
		statisticsGet: connect_go.NewClient[v1.StatisticsGetRequest, v1.StatisticsGetResponse](
			httpClient,
			baseURL+MiikoServiceStatisticsGetProcedure,
			opts...,
		),
		urlGet: connect_go.NewClient[v1.UrlGetRequest, v1.UrlGetResponse](
			httpClient,
			baseURL+MiikoServiceUrlGetProcedure,
			opts...,
		),
	}
}

// miikoServiceClient implements MiikoServiceClient.
type miikoServiceClient struct {
	categoryListGet *connect_go.Client[v1.CategoryListGetRequest, v1.CategoryListGetResponse]
	categoryPost    *connect_go.Client[v1.CategoryPostRequest, v1.CategoryPostResponse]
	topicListGet    *connect_go.Client[v1.TopicListGetRequest, v1.TopicListGetResponse]
	topicGet        *connect_go.Client[v1.TopicGetRequest, v1.TopicGetResponse]
	topicPost       *connect_go.Client[v1.TopicPostRequest, v1.TopicPostResponse]
	problemListGet  *connect_go.Client[v1.ProblemListGetRequest, v1.ProblemListGetResponse]
	problemGet      *connect_go.Client[v1.ProblemGetRequest, v1.ProblemGetResponse]
	problemPost     *connect_go.Client[v1.ProblemPostRequest, v1.ProblemPostResponse]
	statisticsGet   *connect_go.Client[v1.StatisticsGetRequest, v1.StatisticsGetResponse]
	urlGet          *connect_go.Client[v1.UrlGetRequest, v1.UrlGetResponse]
}

// CategoryListGet calls miiko.v1.MiikoService.CategoryListGet.
func (c *miikoServiceClient) CategoryListGet(ctx context.Context, req *connect_go.Request[v1.CategoryListGetRequest]) (*connect_go.Response[v1.CategoryListGetResponse], error) {
	return c.categoryListGet.CallUnary(ctx, req)
}

// CategoryPost calls miiko.v1.MiikoService.CategoryPost.
func (c *miikoServiceClient) CategoryPost(ctx context.Context, req *connect_go.Request[v1.CategoryPostRequest]) (*connect_go.Response[v1.CategoryPostResponse], error) {
	return c.categoryPost.CallUnary(ctx, req)
}

// TopicListGet calls miiko.v1.MiikoService.TopicListGet.
func (c *miikoServiceClient) TopicListGet(ctx context.Context, req *connect_go.Request[v1.TopicListGetRequest]) (*connect_go.Response[v1.TopicListGetResponse], error) {
	return c.topicListGet.CallUnary(ctx, req)
}

// TopicGet calls miiko.v1.MiikoService.TopicGet.
func (c *miikoServiceClient) TopicGet(ctx context.Context, req *connect_go.Request[v1.TopicGetRequest]) (*connect_go.Response[v1.TopicGetResponse], error) {
	return c.topicGet.CallUnary(ctx, req)
}

// TopicPost calls miiko.v1.MiikoService.TopicPost.
func (c *miikoServiceClient) TopicPost(ctx context.Context, req *connect_go.Request[v1.TopicPostRequest]) (*connect_go.Response[v1.TopicPostResponse], error) {
	return c.topicPost.CallUnary(ctx, req)
}

// ProblemListGet calls miiko.v1.MiikoService.ProblemListGet.
func (c *miikoServiceClient) ProblemListGet(ctx context.Context, req *connect_go.Request[v1.ProblemListGetRequest]) (*connect_go.Response[v1.ProblemListGetResponse], error) {
	return c.problemListGet.CallUnary(ctx, req)
}

// ProblemGet calls miiko.v1.MiikoService.ProblemGet.
func (c *miikoServiceClient) ProblemGet(ctx context.Context, req *connect_go.Request[v1.ProblemGetRequest]) (*connect_go.Response[v1.ProblemGetResponse], error) {
	return c.problemGet.CallUnary(ctx, req)
}

// ProblemPost calls miiko.v1.MiikoService.ProblemPost.
func (c *miikoServiceClient) ProblemPost(ctx context.Context, req *connect_go.Request[v1.ProblemPostRequest]) (*connect_go.Response[v1.ProblemPostResponse], error) {
	return c.problemPost.CallUnary(ctx, req)
}

// StatisticsGet calls miiko.v1.MiikoService.StatisticsGet.
func (c *miikoServiceClient) StatisticsGet(ctx context.Context, req *connect_go.Request[v1.StatisticsGetRequest]) (*connect_go.Response[v1.StatisticsGetResponse], error) {
	return c.statisticsGet.CallUnary(ctx, req)
}

// UrlGet calls miiko.v1.MiikoService.UrlGet.
func (c *miikoServiceClient) UrlGet(ctx context.Context, req *connect_go.Request[v1.UrlGetRequest]) (*connect_go.Response[v1.UrlGetResponse], error) {
	return c.urlGet.CallUnary(ctx, req)
}

// MiikoServiceHandler is an implementation of the miiko.v1.MiikoService service.
type MiikoServiceHandler interface {
	CategoryListGet(context.Context, *connect_go.Request[v1.CategoryListGetRequest]) (*connect_go.Response[v1.CategoryListGetResponse], error)
	CategoryPost(context.Context, *connect_go.Request[v1.CategoryPostRequest]) (*connect_go.Response[v1.CategoryPostResponse], error)
	TopicListGet(context.Context, *connect_go.Request[v1.TopicListGetRequest]) (*connect_go.Response[v1.TopicListGetResponse], error)
	TopicGet(context.Context, *connect_go.Request[v1.TopicGetRequest]) (*connect_go.Response[v1.TopicGetResponse], error)
	TopicPost(context.Context, *connect_go.Request[v1.TopicPostRequest]) (*connect_go.Response[v1.TopicPostResponse], error)
	ProblemListGet(context.Context, *connect_go.Request[v1.ProblemListGetRequest]) (*connect_go.Response[v1.ProblemListGetResponse], error)
	ProblemGet(context.Context, *connect_go.Request[v1.ProblemGetRequest]) (*connect_go.Response[v1.ProblemGetResponse], error)
	ProblemPost(context.Context, *connect_go.Request[v1.ProblemPostRequest]) (*connect_go.Response[v1.ProblemPostResponse], error)
	StatisticsGet(context.Context, *connect_go.Request[v1.StatisticsGetRequest]) (*connect_go.Response[v1.StatisticsGetResponse], error)
	UrlGet(context.Context, *connect_go.Request[v1.UrlGetRequest]) (*connect_go.Response[v1.UrlGetResponse], error)
}

// NewMiikoServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMiikoServiceHandler(svc MiikoServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(MiikoServiceCategoryListGetProcedure, connect_go.NewUnaryHandler(
		MiikoServiceCategoryListGetProcedure,
		svc.CategoryListGet,
		opts...,
	))
	mux.Handle(MiikoServiceCategoryPostProcedure, connect_go.NewUnaryHandler(
		MiikoServiceCategoryPostProcedure,
		svc.CategoryPost,
		opts...,
	))
	mux.Handle(MiikoServiceTopicListGetProcedure, connect_go.NewUnaryHandler(
		MiikoServiceTopicListGetProcedure,
		svc.TopicListGet,
		opts...,
	))
	mux.Handle(MiikoServiceTopicGetProcedure, connect_go.NewUnaryHandler(
		MiikoServiceTopicGetProcedure,
		svc.TopicGet,
		opts...,
	))
	mux.Handle(MiikoServiceTopicPostProcedure, connect_go.NewUnaryHandler(
		MiikoServiceTopicPostProcedure,
		svc.TopicPost,
		opts...,
	))
	mux.Handle(MiikoServiceProblemListGetProcedure, connect_go.NewUnaryHandler(
		MiikoServiceProblemListGetProcedure,
		svc.ProblemListGet,
		opts...,
	))
	mux.Handle(MiikoServiceProblemGetProcedure, connect_go.NewUnaryHandler(
		MiikoServiceProblemGetProcedure,
		svc.ProblemGet,
		opts...,
	))
	mux.Handle(MiikoServiceProblemPostProcedure, connect_go.NewUnaryHandler(
		MiikoServiceProblemPostProcedure,
		svc.ProblemPost,
		opts...,
	))
	mux.Handle(MiikoServiceStatisticsGetProcedure, connect_go.NewUnaryHandler(
		MiikoServiceStatisticsGetProcedure,
		svc.StatisticsGet,
		opts...,
	))
	mux.Handle(MiikoServiceUrlGetProcedure, connect_go.NewUnaryHandler(
		MiikoServiceUrlGetProcedure,
		svc.UrlGet,
		opts...,
	))
	return "/miiko.v1.MiikoService/", mux
}

// UnimplementedMiikoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMiikoServiceHandler struct{}

func (UnimplementedMiikoServiceHandler) CategoryListGet(context.Context, *connect_go.Request[v1.CategoryListGetRequest]) (*connect_go.Response[v1.CategoryListGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.CategoryListGet is not implemented"))
}

func (UnimplementedMiikoServiceHandler) CategoryPost(context.Context, *connect_go.Request[v1.CategoryPostRequest]) (*connect_go.Response[v1.CategoryPostResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.CategoryPost is not implemented"))
}

func (UnimplementedMiikoServiceHandler) TopicListGet(context.Context, *connect_go.Request[v1.TopicListGetRequest]) (*connect_go.Response[v1.TopicListGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.TopicListGet is not implemented"))
}

func (UnimplementedMiikoServiceHandler) TopicGet(context.Context, *connect_go.Request[v1.TopicGetRequest]) (*connect_go.Response[v1.TopicGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.TopicGet is not implemented"))
}

func (UnimplementedMiikoServiceHandler) TopicPost(context.Context, *connect_go.Request[v1.TopicPostRequest]) (*connect_go.Response[v1.TopicPostResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.TopicPost is not implemented"))
}

func (UnimplementedMiikoServiceHandler) ProblemListGet(context.Context, *connect_go.Request[v1.ProblemListGetRequest]) (*connect_go.Response[v1.ProblemListGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.ProblemListGet is not implemented"))
}

func (UnimplementedMiikoServiceHandler) ProblemGet(context.Context, *connect_go.Request[v1.ProblemGetRequest]) (*connect_go.Response[v1.ProblemGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.ProblemGet is not implemented"))
}

func (UnimplementedMiikoServiceHandler) ProblemPost(context.Context, *connect_go.Request[v1.ProblemPostRequest]) (*connect_go.Response[v1.ProblemPostResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.ProblemPost is not implemented"))
}

func (UnimplementedMiikoServiceHandler) StatisticsGet(context.Context, *connect_go.Request[v1.StatisticsGetRequest]) (*connect_go.Response[v1.StatisticsGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.StatisticsGet is not implemented"))
}

func (UnimplementedMiikoServiceHandler) UrlGet(context.Context, *connect_go.Request[v1.UrlGetRequest]) (*connect_go.Response[v1.UrlGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("miiko.v1.MiikoService.UrlGet is not implemented"))
}
