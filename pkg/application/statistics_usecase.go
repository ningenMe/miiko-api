package application

import (
	"fmt"
	"github.com/ningenMe/miiko-api/pkg/infra"
	miikov1 "github.com/ningenMe/miiko-api/proto/gen_go/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatisticsUsecase struct{}

var statisticsRepository infra.StatisticsRepository

func (StatisticsUsecase) StatisticsGet() (*miikov1.StatisticsGetResponse, error) {

	//データ取得
	statisticsDto, err := statisticsRepository.Get()
	if err != nil {
		return &miikov1.StatisticsGetResponse{}, fmt.Errorf("statistics not found")
	}

	return &miikov1.StatisticsGetResponse{
		CategorySize:                  statisticsDto.CategorySize,
		TopicSize:                     statisticsDto.TopicSize,
		ProblemSize:                   statisticsDto.ProblemSize,
		TagSize:                       statisticsDto.TagSize,
		ReferenceSize:                 statisticsDto.ReferenceSize,
		LastUpdatedCategoryTimestamp:  timestamppb.New(statisticsDto.LastUpdatedCategoryTimestamp),
		LastUpdatedTopicTimestamp:     timestamppb.New(statisticsDto.LastUpdatedTopicTimestamp),
		LastUpdatedProblemTimestamp:   timestamppb.New(statisticsDto.LastUpdatedProblemTimestamp),
		LastUpdatedTagTimestamp:       timestamppb.New(statisticsDto.LastUpdatedTagTimestamp),
		LastUpdatedReferenceTimestamp: timestamppb.New(statisticsDto.LastUpdatedReferenceTimestamp),
	}, nil
}
