package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/ningenme/miiko-api/proto/gen_go/v1/miikov1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"

	miikov1 "github.com/ningenme/miiko-api/proto/gen_go/v1"
)

type MiikoServer struct{}

func (s *MiikoServer) CategoryGet(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[miikov1.CategoryGetResponse], error) {
	var list []*miikov1.Category
	res := &miikov1.CategoryGetResponse{
		CategoryList: list,
	}
	return connect.NewResponse[miikov1.CategoryGetResponse](res), nil
}

func (s *MiikoServer) CategoryPost(
	ctx context.Context,
	req *connect.Request[miikov1.CategoryPostRequest],
) (*connect.Response[emptypb.Empty], error) {
	return connect.NewResponse[emptypb.Empty](
		&emptypb.Empty{},
	), nil
}

func (s *MiikoServer) TopicGet(
	ctx context.Context,
	req *connect.Request[miikov1.TopicGetRequest],
) (*connect.Response[miikov1.TopicGetResponse], error) {
	return connect.NewResponse[miikov1.TopicGetResponse](&miikov1.TopicGetResponse{}), nil
}

func main() {
	miiko := &MiikoServer{}
	mux := http.NewServeMux()
	path, handler := miikov1connect.NewComproCategoryServiceHandler(miiko)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
	//infra.NingenmeMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("ningenme").FormatDSN())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer infra.NingenmeMysql.Close()
	//
	//infra.ComproMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("compro").FormatDSN())
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer infra.NingenmeMysql.Close()
}
