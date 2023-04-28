package application

import (
	"context"
	"github.com/bufbuild/connect-go"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
)

type HealthController struct{}

func (s *HealthController) Check(
	ctx context.Context,
	req *connect.Request[miikov1.HealthServiceCheckRequest],
) (*connect.Response[miikov1.HealthServiceCheckResponse], error) {
	return connect.NewResponse[miikov1.HealthServiceCheckResponse](&miikov1.HealthServiceCheckResponse{
		Status: miikov1.HealthServiceCheckResponse_SERVING_STATUS_SERVING,
	}), nil
}
